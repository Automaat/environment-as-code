#!/bin/bash
set -e

echo "🚀 Bootstrapping macOS environment..."

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

error() { echo -e "${RED}❌ $1${NC}" >&2; exit 1; }
success() { echo -e "${GREEN}✅ $1${NC}"; }
info() { echo -e "${YELLOW}ℹ️  $1${NC}"; }

# Check not running as root
if [ "$EUID" -eq 0 ]; then
    error "Do not run this script with sudo or as root"
fi

# 1. Install Xcode CLI Tools
if ! xcode-select -p &>/dev/null; then
    info "Installing Xcode Command Line Tools..."
    xcode-select --install
    info "Press any key after installation completes..."
    read -n 1 -s
else
    success "Xcode CLI Tools already installed"
fi

# 2. Install Nix
if ! command -v nix &>/dev/null; then
    info "Installing Nix..."
    curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install

    # Source Nix
    if [ -e '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh' ]; then
        . '/nix/var/nix/profiles/default/etc/profile.d/nix-daemon.sh'
    fi

    # Verify Nix is available
    if ! command -v nix &>/dev/null; then
        error "Nix installation failed. Try restarting terminal and run script again."
    fi
    success "Nix installed"
else
    success "Nix already installed"
fi

# 3. Enable experimental features
info "Enabling Nix experimental features..."
sudo mkdir -p /etc/nix
if ! grep -q "experimental-features = nix-command flakes" /etc/nix/nix.conf 2>/dev/null; then
    echo "experimental-features = nix-command flakes" | sudo tee -a /etc/nix/nix.conf
    sudo launchctl kickstart -k system/org.nixos.nix-daemon
    sleep 2  # Wait for daemon to restart
    success "Experimental features enabled"
else
    success "Experimental features already enabled"
fi

# 4. Clone repo
REPO_DIR="$HOME/sideprojects/environment-as-code"
if [ ! -d "$REPO_DIR" ]; then
    info "Cloning repository..."
    mkdir -p "$HOME/sideprojects"

    # Prompt for GitHub username
    read -p "GitHub username (for git clone): " GITHUB_USER
    git clone "https://github.com/$GITHUB_USER/environment-as-code.git" "$REPO_DIR"
    success "Repository cloned"
else
    success "Repository already exists"
fi

cd "$REPO_DIR"

# 5. Configure personal details
info "Configuring personal details..."
read -p "Full name: " USER_NAME
read -p "Email: " USER_EMAIL

# Get actual username and home directory
ACTUAL_USER=$(whoami)
ACTUAL_HOME="$HOME"

info "Detected user: $ACTUAL_USER"
info "Detected home: $ACTUAL_HOME"

# Update home.nix
sed -i '' "s/fullName = \".*\";/fullName = \"$USER_NAME\";/" modules/home.nix
sed -i '' "s/userEmail = \".*\";/userEmail = \"$USER_EMAIL\";/" modules/home.nix
sed -i '' "s/home.username = \".*\";/home.username = \"$ACTUAL_USER\";/" modules/home.nix
sed -i '' "s|home.homeDirectory = \".*\";|home.homeDirectory = \"$ACTUAL_HOME\";|" modules/home.nix

# Get hostname and architecture
HOSTNAME=$(scutil --get LocalHostName)
ARCH=$(uname -m)
if [ "$ARCH" = "arm64" ]; then
    NIX_ARCH="aarch64-darwin"
else
    NIX_ARCH="x86_64-darwin"
fi

info "Hostname: $HOSTNAME"
info "Architecture: $NIX_ARCH"

# Update flake.nix
sed -i '' "s/\"M-Skalski-MBP\"/\"$HOSTNAME\"/" flake.nix
sed -i '' "s/\"aarch64-darwin\"/\"$NIX_ARCH\"/" flake.nix
sed -i '' "s/\"marcin.skalski@konghq.com\"/\"$ACTUAL_USER\"/" flake.nix

success "Configuration updated"

# 6. Build and activate
info "Building nix-darwin configuration (this may take 10-15 minutes)..."
nix run nix-darwin -- switch --flake ".#$HOSTNAME"

success "🎉 Bootstrap complete!"
echo ""
info "Next steps:"
echo "  1. Restart your terminal"
echo "  2. Verify: darwin-rebuild switch --flake ~/.config/nix-darwin"
echo "  3. Copy SSH key to GitHub: cat ~/.ssh/id_ed25519.pub | pbcopy"
echo "  4. Commit personal changes: cd $REPO_DIR && git add . && git commit -s -S -m 'feat: personalize config'"
