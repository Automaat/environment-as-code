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

info "Converging the machine"
mise exec -- go run ./cmd/eac apply

info "Linking agent configs"
./dotfiles/claude/link.sh

info "Done. Open a new terminal, then add the SSH key to GitHub: pbcopy < ~/.ssh/id_ed25519.pub"
