#!/bin/bash
set -euo pipefail

REPO_URL="https://github.com/Automaat/environment-as-code.git"
REPO_DIR="$HOME/sideprojects/environment-as-code"

info() { printf '\033[1;33m==> %s\033[0m\n' "$1"; }

if ! xcode-select -p &>/dev/null; then
    info "Installing Xcode Command Line Tools (finish the dialog, then press any key)"
    xcode-select --install
    read -r -n 1 -s
fi

if [ ! -x /opt/homebrew/bin/brew ]; then
    info "Installing Homebrew"
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi
eval "$(/opt/homebrew/bin/brew shellenv)"

info "Installing mise"
brew install mise

if [ ! -d "$REPO_DIR" ]; then
    info "Cloning $REPO_URL"
    mkdir -p "$(dirname "$REPO_DIR")"
    git clone "$REPO_URL" "$REPO_DIR"
fi
cd "$REPO_DIR"

info "Installing the eac toolchain"
mise trust --yes mise.toml
mise install --yes

info "Creating the SSH key"
mise exec -- go run ./cmd/eac apply --only system

# The managed git config fetches GitHub over SSH, so tool installs need the
# key registered first.
until ssh -T -o StrictHostKeyChecking=accept-new git@github.com 2>&1 | grep -q "successfully authenticated"; do
    pbcopy < "$HOME/.ssh/id_ed25519.pub"
    info "Public key copied. Add it at https://github.com/settings/ssh/new, then press any key"
    open "https://github.com/settings/ssh/new"
    read -r -n 1 -s
done

info "Converging the machine"
mise exec -- go run ./cmd/eac apply

info "Linking agent configs"
./dotfiles/claude/link.sh

info "Done. Open a new terminal."
