#!/bin/bash

set -e

echo "=================================="
echo "Environment Setup via chezmoi"
echo "=================================="
echo ""

# Check if chezmoi is installed
if ! command -v chezmoi &> /dev/null; then
    echo "chezmoi not found. Installing via Homebrew..."

    # Check if Homebrew is installed
    if ! command -v brew &> /dev/null; then
        echo "Homebrew not found. Installing Homebrew first..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

        # Add Homebrew to PATH for Apple Silicon Macs
        if [[ $(uname -m) == "arm64" ]]; then
            eval "$(/opt/homebrew/bin/brew shellenv)"
        fi
    fi

    brew install chezmoi
fi

echo ""
echo "Running chezmoi to set up your environment..."
echo ""

# Initialize and apply chezmoi configuration
chezmoi init --apply

echo ""
echo "=================================="
echo "Setup complete!"
echo "=================================="
echo ""
echo "Next steps:"
echo "1. Restart your terminal or run: source ~/.zshrc"
echo "2. Run 'vim +PlugInstall' to install Vim plugins"
echo "3. If you generated a new SSH key, add it to GitHub"
echo ""
