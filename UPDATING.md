# Updating Your Environment Configuration

## When You Change Dotfiles

### Method 1: Edit Through chezmoi (Recommended)
```bash
# Edit source file directly
chezmoi edit ~/.zshrc

# Preview changes
chezmoi diff

# Apply changes
chezmoi apply
```

### Method 2: Add Changed Files Back
```bash
# Edit file directly in home directory
vim ~/.zshrc

# Pull changes back into chezmoi
chezmoi add ~/.zshrc

# Commit and push
chezmoi cd
git add .
git commit -m "updated zshrc"
git push
```

## When You Install New Brew Packages

```bash
# Install package normally
brew install new-package

# Update Brewfile script
chezmoi edit ~/.local/share/chezmoi/run_onchange_install-packages.sh.tmpl
# Add: brew "new-package" to heredoc

# Apply (triggers re-run because content changed)
chezmoi apply

# Commit and push
chezmoi cd
git add run_onchange_install-packages.sh.tmpl
git commit -m "added new-package"
git push
```

## When You Install New Brew Casks

```bash
# Install cask normally
brew install --cask spotify

# Update Brewfile script
chezmoi edit ~/.local/share/chezmoi/run_onchange_install-packages.sh.tmpl
# Add: cask "spotify" to heredoc

# Apply (triggers re-run because content changed)
chezmoi apply

# Commit and push
chezmoi cd
git add run_onchange_install-packages.sh.tmpl
git commit -m "added spotify cask"
git push
```

## Generate Brewfile From Current System

```bash
# Dump all installed packages
brew bundle dump --file=-

# Copy relevant entries to chezmoi
chezmoi edit ~/.local/share/chezmoi/run_onchange_install-packages.sh.tmpl
```

## Update All Tracked Files At Once

```bash
# Re-add all tracked files from home directory
chezmoi re-add

# Review changes
chezmoi diff

# Apply if needed
chezmoi apply

# Commit
chezmoi cd
git add .
git commit -m "updated all dotfiles"
git push
```

## Update Configuration on Another Machine

```bash
# Pull latest changes
chezmoi update

# Or manually
chezmoi cd
git pull
exit
chezmoi apply
```

## Useful Commands

```bash
# See what files chezmoi manages
chezmoi managed

# See what would change
chezmoi diff

# Apply changes with verbose output
chezmoi apply -v

# Dry run (don't actually make changes)
chezmoi apply --dry-run

# Navigate to chezmoi source directory
chezmoi cd

# Check status
chezmoi status

# Forget a file (stop managing it)
chezmoi forget ~/.zshrc
```

## Key Points

- `run_onchange_` scripts re-execute when content changes (auto-detects Brewfile updates)
- chezmoi tracks source state in `~/.local/share/chezmoi` (git repo)
- Changes pushed to GitHub, pulled on new machine with `chezmoi update`
- Templates with `.tmpl` extension support variables and conditionals
- `run_once_` scripts only run first time (tracked in `~/.local/share/chezmoi/.chezmoistate.boltdb`)
