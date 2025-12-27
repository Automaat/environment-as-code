# environment-as-code

Automated macOS environment setup using nix-darwin and home-manager.

## Features

- **Declarative package management** - All packages defined in Nix
- **Dotfiles management** - Managed via home-manager
- **Reproducible setup** - Identical environments across machines
- **Homebrew integration** - GUI apps installed via Homebrew casks
- **Development tools** - Git, SSH, oh-my-zsh, vim, mise, Claude Code
- **System configuration** - macOS defaults and preferences

## Quick Setup

### Prerequisites

1. Xcode Command Line Tools:
   ```bash
   xcode-select --install
   ```

2. Install Nix (using Determinate Systems installer):
   ```bash
   curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
   ```

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/environment-as-code.git
   cd environment-as-code
   ```

2. Set your user info (optional):
   ```bash
   export USER_NAME="Your Name"
   export USER_EMAIL="your.email@example.com"
   ```

3. Get your hostname:
   ```bash
   scutil --get LocalHostName
   ```

4. Update `flake.nix` with your hostname and system architecture:
   - Replace `"default"` with your hostname from step 3
   - Change `"aarch64-darwin"` to `"x86_64-darwin"` if on Intel Mac

5. Build and activate:
   ```bash
   nix run nix-darwin -- switch --flake .#$(scutil --get LocalHostName)
   ```

6. Restart your terminal and verify:
   ```bash
   darwin-rebuild switch --flake ~/.config/nix-darwin
   ```

### Post-Setup

```bash
# SSH key will be auto-generated if not present
# Copy it to GitHub:
cat ~/.ssh/id_ed25519.pub | pbcopy

# PARA folders created automatically in ~/Documents/
ls ~/Documents/
```

## What Gets Installed

### CLI Tools (via Nix)

**Development:**
- go, node, python (3.10, 3.13, 3.14), rust, dart, deno
- pyenv, rbenv, mise
- gh, git, docker, docker-compose, lima, orbstack

**Kubernetes/Cloud:**
- kubectl, kubectx, kubie, helm, k9s, kind, minikube, skaffold, eksctl
- consul, kumactl
- awscli, aws-iam-authenticator, aws-sso-cli, saml2aws

**CLI Utilities:**
- bat, fzf, jq, tree, autojump, starship
- pre-commit, shellcheck, actionlint, tflint, vale
- ruff, uv, oxlint, yamlfmt, yamllint

**Build Tools:**
- cmake, ninja, autoconf, hugo, buf

**Security:**
- bitwarden-cli, gnupg, pinentry_mac, yubikey-manager

### Applications (via Homebrew Cask)

**Productivity:**
- 1Password, Bitwarden, Todoist, Obsidian, Espanso, Contexts, Freedom

**Development:**
- Visual Studio Code, Ghostty, Insomnia, pgAdmin4, OrbStack

**Browsers:**
- Vivaldi

**Cloud:**
- Google Drive, Google Cloud SDK, Home Assistant

**Communication:**
- Discord, ChatGPT

**Media:**
- XnViewMP, Adobe Acrobat Reader, Elgato Stream Deck, Whisper Hotkey

**Fonts:**
- Fira Code Nerd Font, Iosevka, Iosevka Nerd Font

**LaTeX:**
- BasicTeX

### Configurations

**Shell (zsh):**
- oh-my-zsh with plugins: git, brew, docker, golang, helm, kubectl, terraform, fzf
- zsh-autosuggestions, zsh-syntax-highlighting
- ZPlug plugins: you-should-use, clean-history
- autojump, starship (disabled by default)

**Git:**
- User name/email (from env or defaults)
- GPG signing enabled (key: 683BF754A0B005CD)
- Pull rebase, sign-off, LFS support

**Terminal (Ghostty):**
- Iosevka 23pt font with Nerd Font symbols
- Everforest Light theme
- Performance optimized for Claude Code (200M scrollback)

**Other:**
- Vim with Darcula theme
- K9s with Everforest skin
- Mise config (Go 1.24.9, uv latest)
- Claude Code settings, commands, and skills
- Klaudiush validators

**macOS Defaults:**
- Dock autohide, no recents
- Finder show all extensions
- Fast key repeat
- Touch ID for sudo

## Structure

```
.
├── flake.nix                    # Main entrypoint
├── modules/
│   ├── darwin.nix              # macOS system configuration
│   ├── home.nix                # Home-manager configuration
│   ├── packages.nix            # CLI packages list
│   ├── casks.nix               # GUI apps list
│   ├── programs/
│   │   ├── zsh.nix            # Zsh + oh-my-zsh config
│   │   ├── git.nix            # Git configuration
│   │   ├── vim.nix            # Vim configuration
│   │   └── vscode.nix         # VSCode (placeholder)
│   └── dotfiles/
│       ├── ghostty.nix        # Ghostty terminal config
│       ├── k9s.nix            # K9s config + skin
│       ├── mise.nix           # Mise tool versions
│       ├── claude.nix         # Claude Code setup
│       └── klaudiush.nix      # Klaudiush validators
├── README.md                    # This file
└── MIGRATION.md                 # Chezmoi → Nix guide
```

## Updating Configuration

### Add/Remove Packages

1. Edit package list:
   ```bash
   vim modules/packages.nix      # CLI tools
   vim modules/casks.nix         # GUI apps
   ```

2. Apply changes:
   ```bash
   darwin-rebuild switch --flake ~/.config/nix-darwin
   ```

### Update Dotfiles

1. Edit the relevant module:
   ```bash
   vim modules/programs/zsh.nix   # Shell config
   vim modules/dotfiles/ghostty.nix  # Terminal config
   ```

2. Apply:
   ```bash
   darwin-rebuild switch --flake ~/.config/nix-darwin
   ```

### Update All Packages

```bash
# Update flake inputs (nixpkgs, nix-darwin, home-manager)
nix flake update

# Rebuild with new versions
darwin-rebuild switch --flake ~/.config/nix-darwin
```

### Sync Across Machines

1. Commit changes:
   ```bash
   cd ~/sideprojects/environment-as-code
   git add .
   git commit -s -S -m "feat: update configuration"
   git push
   ```

2. On other machine:
   ```bash
   cd ~/sideprojects/environment-as-code
   git pull
   darwin-rebuild switch --flake .#$(scutil --get LocalHostName)
   ```

## Customization

### User Info

Set via environment variables before first build:

```bash
export USER_NAME="Your Name"
export USER_EMAIL="your.email@example.com"
```

Or edit `modules/home.nix` to change defaults.

### Machine-Specific Config

Create per-host configs in `flake.nix`:

```nix
darwinConfigurations = {
  "work-laptop" = darwin.lib.darwinSystem {
    # work-specific config
  };
  "personal-mac" = darwin.lib.darwinSystem {
    # personal config
  };
};
```

### Work-Specific Aliases

The zsh config includes conditional aliases for Kong/Kuma work if `~/kong` directory exists.

## Migration from Chezmoi

See [MIGRATION.md](MIGRATION.md) for detailed guide on migrating from the previous chezmoi setup.

## Requirements

- macOS (tested on Sequoia)
- Apple Silicon (aarch64) or Intel (x86_64)
- Admin access
- Internet connection

## Troubleshooting

### Command Not Found After Install

Restart terminal or reload PATH:

```bash
exec zsh
```

### Homebrew Casks Not Installing

Check homebrew installation:

```bash
which brew
brew --version
```

### Nix Daemon Issues

Restart nix daemon:

```bash
sudo launchctl kickstart -k system/org.nixos.nix-daemon
```

### Rebuild From Scratch

```bash
# Remove nix-darwin state
sudo rm -rf /etc/nix/nix.conf /run/current-system

# Rebuild
darwin-rebuild switch --flake ~/.config/nix-darwin
```

### Check What Changed

```bash
# See what would change (dry run)
darwin-rebuild build --flake ~/.config/nix-darwin
nix store diff-closures /run/current-system ./result
```

## Benefits Over Homebrew/Chezmoi

- **Atomic updates**: Rollback if something breaks
- **Reproducible**: Exact same environment on any machine
- **Declarative**: Single source of truth for all config
- **Faster**: Binary cache, parallel builds
- **Version pinning**: Lock exact package versions
- **System config**: Manage macOS settings declaratively

## License

MIT
