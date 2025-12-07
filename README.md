# environment-as-code

Automated macOS environment setup using Homebrew Bundle and chezmoi.

## Features

- **Homebrew packages & casks** - Declarative package management
- **Dotfiles management** - .zshrc, .vimrc, iTerm2 config
- **Development tools** - Git, SSH, oh-my-zsh, vim-plug, mise
- **PARA folder structure** - Organized documents hierarchy
- **Machine-specific configs** - Templates for different setups

## Quick Setup

### New Machine

```bash
# Clone and run setup
git clone https://github.com/yourusername/environment-as-code.git
cd environment-as-code
./setup.sh
```

This will:

1. Install Homebrew (if needed)
2. Install chezmoi
3. Run chezmoi to set up everything

### Post-Setup

```bash
# Restart terminal or reload shell
source ~/.zshrc

# Install Vim plugins
vim +PlugInstall

# Add SSH key to GitHub (if generated)
cat ~/.ssh/id_ed25519.pub
```

## What Gets Installed

### CLI Tools (via Homebrew)

- docker, k9s, pyenv, gh, kubectl, kubectx, helm
- autojump, zsh-autosuggestions, zsh-syntax-highlighting, starship
- mise, zplug, fzf, saml2aws
- jq, tree, pre-commit, awscli

### Applications (via Homebrew Cask)

- todoist, obsidian, bitwarden, 1password
- visual-studio-code, google-chrome, google-drive
- contexts, ghostty, docker
- adobe-acrobat-reader, espanso

### Configurations

- oh-my-zsh with custom plugins
- Vim with vim-plug
- Ghostty terminal configuration
- Git config (name, email, signing)
- SSH key generation
- PARA folder structure (Projects/Areas/Resources/Archive)
- Claude Code settings and custom commands
- mise config (tool versions)

## Updating Configuration

See [UPDATING.md](UPDATING.md) for detailed instructions on:

- Editing dotfiles
- Adding new packages
- Syncing across machines
- Managing chezmoi

### Quick Reference

```bash
# Edit dotfile
chezmoi edit ~/.zshrc

# Add new package (after installing)
chezmoi edit ~/.local/share/chezmoi/run_onchange_install-packages.sh.tmpl

# Apply changes
chezmoi apply

# Commit and push
chezmoi cd
git add .
git commit -m "description"
git push
```

## Structure

```text
~/.local/share/chezmoi/          # chezmoi source directory
├── .chezmoi.toml.tmpl           # Config with prompts
├── dot_zshrc.tmpl               # .zshrc template
├── dot_vimrc                    # .vimrc
├── dot_config/iterm/            # iTerm2 config
├── run_once_before_*.sh         # Pre-setup scripts
├── run_once_after_*.sh.tmpl     # Post-setup scripts
├── run_onchange_install-packages.sh.tmpl  # Brewfile
├── UPDATING.md                  # Update instructions
└── README.md                    # This file
```

## Customization

### Machine-Specific Settings

Edit templates to add conditionals:

```bash
chezmoi edit ~/.local/share/chezmoi/dot_zshrc.tmpl
```

Use conditionals:

```zsh
{{- if eq .chezmoi.hostname "work-laptop" }}
# Work-specific aliases
{{- end }}
```

### Email and Name

First run will prompt for:

- Email address
- Full name

These are templated into:

- Git config
- SSH key generation

## Requirements

- macOS (tested on Sequoia)
- Internet connection
- Admin access for Xcode Command Line Tools

## Troubleshooting

### Xcode Not Installed

```bash
xcode-select --install
```

### Homebrew Path Issues (Apple Silicon)

```bash
eval "$(/opt/homebrew/bin/brew shellenv)"
```

### chezmoi Changes Not Applying

```bash
chezmoi diff    # See what would change
chezmoi apply -v  # Verbose output
```

### Reset chezmoi State

```bash
rm -rf ~/.local/share/chezmoi/.chezmoistate.boltdb
chezmoi apply
```

## License

MIT
