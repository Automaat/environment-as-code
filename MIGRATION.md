# Migration Guide: Chezmoi → Nix-Darwin

Complete guide for migrating from chezmoi + Homebrew to nix-darwin + home-manager.

## Overview

**Before (Chezmoi):**
- Homebrew for package management
- Chezmoi for dotfiles
- Manual setup scripts
- Template-based config

**After (Nix-Darwin):**
- Nix for CLI packages
- Homebrew for GUI apps only
- Home-manager for dotfiles
- Declarative everything

## Pre-Migration Checklist

### 1. Backup Current State

```bash
# Export current Homebrew state
brew bundle dump --file=~/Brewfile.backup

# Backup dotfiles
chezmoi archive > ~/chezmoi-backup.tar.gz

# List installed packages
brew list > ~/brew-packages.txt
brew list --cask > ~/brew-casks.txt

# Backup any manual configs
cp ~/.zshrc ~/.zshrc.backup
cp ~/.gitconfig ~/.gitconfig.backup
```

### 2. Document Custom Changes

Review and note:
- Custom aliases in `.zshrc`
- Work-specific paths
- Machine-specific settings
- SSH keys location
- GPG key ID

## Migration Steps

### Step 1: Install Nix

```bash
# Install Nix (Determinate Systems installer)
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install

# Restart terminal
exec zsh
```

### Step 2: Clone Nix Config

```bash
# Clone the migrated config
git clone https://github.com/yourusername/environment-as-code.git
cd environment-as-code
```

### Step 3: Configure User Info

```bash
# Set your info (or edit modules/home.nix directly)
export USER_NAME="Your Name"
export USER_EMAIL="your.email@example.com"

# Get hostname
scutil --get LocalHostName
```

### Step 4: Update Flake

Edit `flake.nix`:

```nix
# Replace "default" with your hostname
darwinConfigurations = {
  "your-hostname" = darwin.lib.darwinSystem {
    # ...
  };
};
```

Update architecture if Intel Mac:
```nix
system = "x86_64-darwin";  # instead of "aarch64-darwin"
```

### Step 5: First Build

```bash
# Build and activate nix-darwin
nix run nix-darwin -- switch --flake .#$(scutil --get LocalHostName)

# Restart terminal
exec zsh
```

### Step 6: Verify Installation

```bash
# Check nix-darwin is active
darwin-rebuild --version

# Verify packages
which go
which kubectl
which gh

# Check dotfiles
cat ~/.zshrc | head
cat ~/.gitconfig
cat ~/.config/ghostty/config
```

### Step 7: Manual Migrations

#### SSH Keys

If you have existing SSH keys:

```bash
# Keys are preserved at ~/.ssh/
# Nix will not overwrite existing keys
ls ~/.ssh/

# If you want new keys:
mv ~/.ssh/id_ed25519 ~/.ssh/id_ed25519.old
darwin-rebuild switch --flake ~/.config/nix-darwin
```

#### GPG Keys

If you use different GPG key:

```bash
# List keys
gpg --list-secret-keys

# Update modules/programs/git.nix
# Change signing.key = "YOUR_KEY_ID";

darwin-rebuild switch --flake ~/.config/nix-darwin
```

#### Work-Specific Paths

If you have custom work paths, edit `modules/programs/zsh.nix`:

```nix
# Add to initExtra:
if [ -d "$HOME/your-work-dir" ]; then
  alias your-alias="your command"
  export YOUR_VAR="value"
fi
```

### Step 8: Cleanup Chezmoi (Optional)

```bash
# Stop chezmoi
chezmoi purge

# Remove chezmoi dotfiles source
rm -rf ~/.local/share/chezmoi

# Uninstall chezmoi (now managed by nix)
# brew uninstall chezmoi  # Not needed, nix takes over
```

### Step 9: Remove Homebrew Formulae

Nix now manages CLI tools. Keep homebrew for casks only:

```bash
# List what's installed via brew (not cask)
brew list --formula

# These are now in nix, can uninstall:
brew uninstall go node python@3.13 kubectl helm k9s docker

# Or clean all formulae:
brew list --formula | xargs brew uninstall

# Keep homebrew itself for casks
# DON'T run: /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/uninstall.sh)"
```

## Configuration Comparison

### Dotfile Locations

| Chezmoi | Nix-Darwin |
|---------|------------|
| `~/.local/share/chezmoi/dot_zshrc.tmpl` | `modules/programs/zsh.nix` |
| `~/.local/share/chezmoi/dot_gitconfig.tmpl` | `modules/programs/git.nix` |
| `~/.local/share/chezmoi/dot_config/ghostty/config` | `modules/dotfiles/ghostty.nix` |
| `~/.local/share/chezmoi/dot_config/k9s/*` | `modules/dotfiles/k9s.nix` |
| `~/.local/share/chezmoi/run_onchange_install-packages.sh.tmpl` | `modules/packages.nix` + `modules/casks.nix` |

### Update Workflow

| Task | Chezmoi | Nix-Darwin |
|------|---------|------------|
| Add package | Edit Brewfile, `brew bundle` | Edit `packages.nix`, `darwin-rebuild switch` |
| Update dotfile | `chezmoi edit ~/.zshrc` | `vim modules/programs/zsh.nix` |
| Apply changes | `chezmoi apply` | `darwin-rebuild switch --flake ~/.config/nix-darwin` |
| Sync machines | `chezmoi init`, `chezmoi apply` | `git pull`, `darwin-rebuild switch` |

### Templating

**Chezmoi:**
```zsh
{{- if eq .chezmoi.os "darwin" }}
alias sb="cd {{ .chezmoi.homeDir }}/Documents"
{{- end }}
```

**Nix:**
```nix
initExtra = ''
  alias sb="cd $HOME/Documents"
'';
```

## Common Issues

### "Command not found" after migration

```bash
# Reload shell
exec zsh

# Or check PATH
echo $PATH | tr ':' '\n'
```

### Packages not available in nixpkgs

Some packages might not exist in nixpkgs. Options:

1. **Keep in homebrew:**
   ```nix
   # modules/darwin.nix
   homebrew.brews = [
     "some-rare-package"
   ];
   ```

2. **Search for alternative name:**
   ```bash
   # Search nixpkgs
   nix search nixpkgs package-name
   ```

### Dotfiles not updating

```bash
# Check what would change
darwin-rebuild build --flake ~/.config/nix-darwin
nix store diff-closures /run/current-system ./result

# Force rebuild
darwin-rebuild switch --flake ~/.config/nix-darwin --recreate-lock-file
```

### Rollback if something breaks

```bash
# List generations
darwin-rebuild --list-generations

# Rollback to previous
darwin-rebuild rollback

# Or specific generation
darwin-rebuild switch --rollback-to 42
```

## Package Name Mappings

Common homebrew → nixpkgs name differences:

| Homebrew | Nixpkgs |
|----------|---------|
| `awscli` | `awscli2` |
| `helm` | `kubernetes-helm` |
| `jd` | `jd-diff-patch` |
| `postgresql@14` | `postgresql_14` |
| `python@3.13` | `python313` |
| `ykman` | `yubikey-manager` |
| `docker-credential-helper` | `docker-credential-helpers` |

## Benefits After Migration

### Reproducibility

```bash
# Exact same environment on new machine:
git clone repo
nix run nix-darwin -- switch --flake .
```

### Atomic Updates

```bash
# Update safely, rollback if breaks
darwin-rebuild switch
# If broken:
darwin-rebuild rollback
```

### Version Pinning

```bash
# Lock exact versions
nix flake lock

# Update specific input
nix flake update nixpkgs
```

### Faster Rebuilds

- Binary cache (no compilation)
- Parallel builds
- Only changed packages rebuild

## Next Steps

1. **Commit initial config:**
   ```bash
   cd ~/sideprojects/environment-as-code
   git add .
   git commit -s -S -m "feat: migrate to nix-darwin"
   git push
   ```

2. **Test on another machine** (if available)

3. **Document custom changes** in CLAUDE.md

4. **Set up automatic updates** (optional):
   ```bash
   # Add to cron or launchd
   darwin-rebuild switch --flake ~/.config/nix-darwin && nix-collect-garbage -d
   ```

## Troubleshooting

### Homebrew conflicts with Nix

```bash
# Ensure nix binaries take precedence
echo $PATH
# Should show /run/current-system/sw/bin before /opt/homebrew/bin
```

### Oh-my-zsh not loading

```bash
# Check installation
ls ~/.oh-my-zsh

# Reinstall if needed (managed by home-manager)
darwin-rebuild switch --flake ~/.config/nix-darwin --recreate-lock-file
```

### GPG signing not working

```bash
# Check pinentry
which pinentry-mac
ls -la ~/.gnupg/gpg-agent.conf

# Restart gpg-agent
gpgconf --kill gpg-agent
darwin-rebuild switch --flake ~/.config/nix-darwin
```

## Getting Help

- [Nix-Darwin Manual](https://daiderd.com/nix-darwin/)
- [Home-Manager Manual](https://nix-community.github.io/home-manager/)
- [NixOS Discourse](https://discourse.nixos.org/)
- [Nix Pills](https://nixos.org/guides/nix-pills/)

## Reverting to Chezmoi (Emergency)

If you need to revert:

```bash
# Uninstall nix-darwin
darwin-uninstaller

# Remove nix
/nix/nix-installer uninstall

# Restore chezmoi
tar -xzf ~/chezmoi-backup.tar.gz
chezmoi init
chezmoi apply

# Reinstall homebrew packages
brew bundle --file=~/Brewfile.backup
```
