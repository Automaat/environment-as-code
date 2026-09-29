#!/bin/bash
set -euo pipefail

REPO_URL="https://github.com/Automaat/environment-as-code.git"
REPO_DIR="$HOME/sideprojects/environment-as-code"
INSTALLER="https://raw.githubusercontent.com/Automaat/zakwas/main/install.sh"

info() { printf '\033[1;33m==> %s\033[0m\n' "$1"; }

# ssh -T exits 1 even when authenticated (GitHub offers no shell), so the
# exit status can't be trusted; match the greeting instead.
github_ssh_ok() {
    local out
    out=$(ssh -T -o StrictHostKeyChecking=accept-new git@github.com 2>&1 || true)
    [[ "$out" == *"successfully authenticated"* ]]
}

main() {
    # Piped as `curl … | bash`, stdin is the script itself: prompts would
    # swallow script text. The script lives in main, called on the last
    # line, so bash has parsed all of it before stdin moves to the terminal.
    if [ ! -t 0 ] && [ -r /dev/tty ]; then
        exec </dev/tty
    fi

    # The pinned release, so a fresh Mac runs what this config was tested with.
    local version
    version=$(curl -fsSL "https://raw.githubusercontent.com/Automaat/environment-as-code/main/dotfiles/mise/config.toml" |
        sed -n 's|^"github:Automaat/zakwas" = "\(.*\)"$|\1|p')

    curl -fsSL "$INSTALLER" | bash -s -- --repo "$REPO_URL" --dir "$REPO_DIR" --version "$version" --no-apply
    zakwas() { "$HOME/.local/bin/zakwas" "$@"; }
    cd "$REPO_DIR"

    info "Creating the SSH key"
    zakwas apply -y --only system

    # The managed git config fetches GitHub over SSH, so tool installs need
    # the key registered first.
    until github_ssh_ok; do
        pbcopy < "$HOME/.ssh/id_ed25519.pub"
        info "Public key copied. Add it at https://github.com/settings/ssh/new, then press any key"
        open "https://github.com/settings/ssh/new"
        read -r -n 1 -s
    done

    # Around 60 pinned tools resolve through the GitHub API; without a token
    # a fresh install can hit the unauthenticated rate limit.
    if [ -z "${MISE_GITHUB_TOKEN:-}" ]; then
        info "GitHub token for tool downloads (any token without scopes works; Enter to skip)"
        read -r -s token
        if [ -n "$token" ]; then
            export MISE_GITHUB_TOKEN="$token"
        fi
    fi

    # Installs Homebrew and mise first, then everything else.
    info "Converging the machine"
    zakwas apply

    info "Linking agent configs"
    ./dotfiles/claude/link.sh

    info "Done. Open a new terminal."
}

main "$@"; exit
