#!/bin/bash
set -e

echo "🔧 Fixing configuration for current user..."

# Find repo directory
if [ -f "flake.nix" ] && [ -d "modules" ]; then
    REPO_DIR="$(pwd)"
elif [ -d "$HOME/sideprojects/environment-as-code" ]; then
    REPO_DIR="$HOME/sideprojects/environment-as-code"
else
    echo "❌ Error: Cannot find environment-as-code repository"
    echo "   Run this script from the repo directory or ensure it's in ~/sideprojects/environment-as-code"
    exit 1
fi

cd "$REPO_DIR"
echo "Working in: $REPO_DIR"
echo ""

# Get actual values
ACTUAL_USER=$(whoami)
ACTUAL_HOME="$HOME"
HOSTNAME=$(scutil --get LocalHostName)
ARCH=$(uname -m)
if [ "$ARCH" = "arm64" ]; then
    NIX_ARCH="aarch64-darwin"
else
    NIX_ARCH="x86_64-darwin"
fi

echo "Detected configuration:"
echo "  User: $ACTUAL_USER"
echo "  Home: $ACTUAL_HOME"
echo "  Hostname: $HOSTNAME"
echo "  Architecture: $NIX_ARCH"
echo ""

# Fix home.nix
echo "Updating modules/home.nix..."
sed -i '' "s/^  home.username = \".*\";$/  home.username = \"$ACTUAL_USER\";/" modules/home.nix
sed -i '' "s|^  home.homeDirectory = \".*\";$|  home.homeDirectory = \"$ACTUAL_HOME\";|" modules/home.nix

# Fix flake.nix
echo "Updating flake.nix..."
sed -i '' "s/\"M-Skalski-MBP\"/\"$HOSTNAME\"/g" flake.nix
sed -i '' "s/\"aarch64-darwin\"/\"$NIX_ARCH\"/g" flake.nix
sed -i '' "s/\"x86_64-darwin\"/\"$NIX_ARCH\"/g" flake.nix
sed -i '' "s/\"marcin.skalski@konghq.com\"/\"$ACTUAL_USER\"/g" flake.nix

echo ""
echo "✅ Configuration updated!"
echo ""
echo "Verification:"
echo "─────────────────────────────────────"
grep -E "home.username|home.homeDirectory" modules/home.nix | sed 's/^/  /'
echo ""
grep -E "darwinConfigurations|system =|home-manager.users" flake.nix | sed 's/^/  /'
echo "─────────────────────────────────────"
echo ""

read -p "Run nix-darwin build now? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "Building nix-darwin configuration..."
    sudo --preserve-env=HOME nix --extra-experimental-features "nix-command flakes" run nix-darwin -- switch --flake ".#$HOSTNAME"
else
    echo "Skipped. Run manually with:"
    echo "  sudo --preserve-env=HOME nix --extra-experimental-features \"nix-command flakes\" run nix-darwin -- switch --flake \".#$HOSTNAME\""
fi
