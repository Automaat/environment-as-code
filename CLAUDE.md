# environment-as-code

Declarative macOS system configuration using nix-darwin and home-manager.

## Project Structure

```
.
├── flake.nix                    # Flake entrypoint, system definition
├── flake.lock                   # Dependency lock file
├── modules/
│   ├── darwin.nix              # macOS system settings (Dock, Finder, Touch ID)
│   ├── home.nix                # Home-manager config, user environment
│   ├── packages.nix            # CLI tools (kubectl, gh, aws, etc.)
│   ├── casks.nix               # GUI apps via Homebrew (VSCode, Ghostty, etc.)
│   ├── scripts.nix             # Custom scripts (gh-add-subissue, etc.)
│   ├── programs/               # Program-specific configs
│   │   ├── git.nix            # Git config (GPG signing, aliases)
│   │   ├── vim.nix            # Vim config + Darcula theme
│   │   └── zsh.nix            # oh-my-zsh, plugins, aliases
│   └── dotfiles/               # Dotfile configurations
│       ├── claude.nix         # Claude Code settings + commands
│       ├── ghostty.nix        # Ghostty terminal config
│       ├── k9s.nix            # K9s + Everforest skin
│       ├── mise.nix           # Mise tool versions
│       ├── ssh.nix            # SSH config
│       └── starship.nix       # Starship prompt
├── dotfiles/                    # Source dotfiles (referenced by modules)
├── bootstrap.sh                 # Fresh Mac setup script
└── setup.sh                     # Manual installation script
```

## Tech Stack

**Language:** Nix
**Framework:** nix-darwin (system) + home-manager (user env)
**Package Manager:** Nix (CLI tools) + Homebrew (GUI apps)
**Dependency Management:** Nix flakes with lock file
**Version Control:** Git with GPG signing

## Development Workflow

### Adding/Updating Dotfiles or Program Configs

**Most common workflow** - modifying program configurations:

1. Edit relevant module:
   ```bash
   vim modules/programs/zsh.nix      # Shell config
   vim modules/dotfiles/ghostty.nix  # Terminal config
   vim modules/programs/git.nix      # Git settings
   ```

2. Validate syntax:
   ```bash
   nix flake check
   ```

3. Build configuration (test without applying):
   ```bash
   darwin-rebuild build --flake ~/.config/nix-darwin
   ```

4. Inspect what will change:
   ```bash
   nix store diff-closures /run/current-system ./result
   ```

5. Apply if satisfied:
   ```bash
   darwin-rebuild switch --flake ~/.config/nix-darwin
   ```

6. Commit changes:
   ```bash
   git add .
   git commit -s -S -m "feat: update <program> config"
   git push
   ```

### Adding/Removing Packages

1. Edit package list:
   ```bash
   vim modules/packages.nix    # CLI tools
   vim modules/casks.nix       # GUI apps
   ```

2. Follow build-then-switch workflow above (steps 2-6)

### Updating Dependencies

```bash
# Update flake inputs (nixpkgs, nix-darwin, home-manager)
nix flake update

# Review what will change
nix flake lock --update-input nixpkgs  # Update specific input only

# Rebuild with new versions
darwin-rebuild build --flake ~/.config/nix-darwin
nix store diff-closures /run/current-system ./result
darwin-rebuild switch --flake ~/.config/nix-darwin
```

### Syncing Across Machines

1. On machine A (after making changes):
   ```bash
   git add .
   git commit -s -S -m "feat: description"
   git push
   ```

2. On machine B:
   ```bash
   cd ~/sideprojects/environment-as-code
   git pull
   darwin-rebuild switch --flake .#$(scutil --get LocalHostName)
   ```

## Module Organization

### When to Create New Module

**Create in `modules/programs/`:**
- Program with complex config requiring multiple options
- Examples: git.nix (GPG, aliases), zsh.nix (oh-my-zsh, plugins), vim.nix (theme, settings)

**Create in `modules/dotfiles/`:**
- Simple dotfile deployment via home.file
- Examples: ghostty.nix (terminal config), k9s.nix (k8s UI config)

**Add to existing module:**
- Single package addition → `modules/packages.nix`
- Single cask addition → `modules/casks.nix`
- System setting → `modules/darwin.nix`

### Module Structure

**For programs/ (complex config):**

```nix
{ config, pkgs, ... }:

{
  programs.<program-name> = {
    enable = true;
    # Program-specific options
  };
}
```

**For dotfiles/ (simple file deployment):**

```nix
{ config, pkgs, ... }:

{
  home.file.".<config-path>" = {
    source = ../../dotfiles/<program>/<file>;
    # OR for templated content:
    text = ''
      config content here
    '';
  };
}
```

## Quality Gates

### Pre-Rebuild Checklist

Run before darwin-rebuild switch:

- [ ] **Syntax validation:** `nix flake check`
- [ ] **Build test:** `darwin-rebuild build --flake ~/.config/nix-darwin`
- [ ] **Inspect diff:** `nix store diff-closures /run/current-system ./result`
- [ ] **Git committed:** Changes committed to enable rollback
- [ ] **Git pushed:** Synced to remote (prevents loss across machines)

### Commands

```bash
# Validate syntax
nix flake check

# Test build without applying
darwin-rebuild build --flake ~/.config/nix-darwin

# See what will change
nix store diff-closures /run/current-system ./result

# Apply changes
darwin-rebuild switch --flake ~/.config/nix-darwin

# Commit and push
git add . && git commit -s -S -m "feat: description" && git push
```

## Build-Then-Switch Verification

**Two-phase safety workflow** to prevent breaking system:

### Phase 1: Build & Inspect

```bash
# 1. Validate configuration syntax
nix flake check
# ✓ Pass → continue
# ✗ Fail → fix syntax errors in modules/*.nix

# 2. Build without switching
darwin-rebuild build --flake ~/.config/nix-darwin
# ✓ Pass → ./result symlink created
# ✗ Fail → fix module errors (missing packages, invalid options)

# 3. Inspect what will change
nix store diff-closures /run/current-system ./result
# Review output for unexpected changes
# Check package versions, new packages, removed packages
```

### Phase 2: Apply

```bash
# 4. Switch to new configuration
darwin-rebuild switch --flake ~/.config/nix-darwin
# ✓ Pass → new config active
# ✗ Fail → rollback available (see Troubleshooting)

# 5. Verify in new terminal
exec zsh
# Test changed programs/configs

# 6. Commit if satisfied
git add . && git commit -s -S -m "feat: description" && git push
```

**Critical:** Always commit after successful switch. Uncommitted changes risk loss when syncing across machines.

## Pre-Rebuild Validation

**Quick triage before rebuilding** - catch errors early:

### 1-Minute Validation

```bash
# Check for uncommitted changes (reminder to commit later)
git status

# Validate Nix syntax
nix flake check

# Quick build test
darwin-rebuild build --flake ~/.config/nix-darwin --fast
```

### 3-Minute Validation

Add if making significant changes:

```bash
# Full build test
darwin-rebuild build --flake ~/.config/nix-darwin

# Review diff
nix store diff-closures /run/current-system ./result | head -50

# Check if critical packages affected
nix store diff-closures /run/current-system ./result | grep -E "(git|zsh|nix|darwin)"
```

### 5-Minute Validation

Add if updating dependencies:

```bash
# See what inputs changed
nix flake metadata

# Review flake.lock diff
git diff flake.lock

# Full diff review
nix store diff-closures /run/current-system ./result
```

**Workflow Decision:**
- Minor config tweak → 1-min validation → switch
- New module or package → 3-min validation → switch
- Flake update → 5-min validation → switch

## Common Commands

```bash
# Apply configuration changes
darwin-rebuild switch --flake ~/.config/nix-darwin

# Test build without applying
darwin-rebuild build --flake ~/.config/nix-darwin

# Validate flake syntax
nix flake check

# See what will change
nix store diff-closures /run/current-system ./result

# Update all dependencies
nix flake update

# Update specific input
nix flake lock --update-input nixpkgs

# Rollback to previous generation
darwin-rebuild switch --rollback

# List generations
darwin-rebuild --list-generations

# Garbage collect old generations
nix-collect-garbage -d

# Get current hostname
scutil --get LocalHostName
```

## Anti-Patterns

**AVOID:**

- ❌ **Skipping `nix flake check`** - Syntax errors break rebuild
- ❌ **Direct switch without build test** - Risk breaking system, hard to diagnose
- ❌ **Not committing changes before rebuild** - Can't rollback to working state easily
- ❌ **Forgetting to push after successful rebuild** - Changes lost when syncing across machines
- ❌ **Editing files in ~/.config/nix-darwin directly** - Changes overwritten by git pull; edit in repo
- ❌ **Adding machine-specific paths** - Use home.homeDirectory variable, not hardcoded paths
- ❌ **Not reviewing diff** - Unexpected package updates can introduce issues

**REASON:**
- Syntax validation catches errors before system state affected
- Build test ensures config valid without disrupting running system
- Git discipline prevents loss of working configurations
- Machine portability requires variable substitution

## Module Templates

### New Program Config Template

When adding complex program configuration in `modules/programs/<name>.nix`:

```nix
{ config, pkgs, ... }:

{
  programs.<program-name> = {
    enable = true;

    # Package override if needed
    package = pkgs.<program-name>;

    # Program-specific options
    # Consult: https://home-manager-options.extranix.com

    # Example: aliases
    shellAliases = {
      "shortcut" = "full command";
    };

    # Example: environment variables
    sessionVariables = {
      VAR_NAME = "value";
    };

    # Example: config file content
    extraConfig = ''
      config line 1
      config line 2
    '';
  };
}
```

### New Dotfile Config Template

When adding simple dotfile in `modules/dotfiles/<name>.nix`:

```nix
{ config, pkgs, ... }:

{
  # Deploy existing dotfile
  home.file.".<config-path>/<filename>" = {
    source = ../../dotfiles/<program>/<filename>;
  };

  # OR generate config inline
  home.file.".<config-path>/<filename>".text = ''
    # Config content
    setting = value
  '';

  # For XDG config dirs
  xdg.configFile."<program>/<filename>" = {
    source = ../../dotfiles/<program>/<filename>;
  };
}
```

### Adding to flake.nix

After creating module, import in `modules/home.nix` or `modules/darwin.nix`:

```nix
# In modules/home.nix
imports = [
  ./programs/new-program.nix
  ./dotfiles/new-dotfile.nix
];

# OR in flake.nix darwinSystem modules list
modules = [
  ./modules/darwin.nix
  ./modules/new-module.nix  # System-level module
];
```

## Troubleshooting

### Rollback After Bad Switch

```bash
# Immediate rollback to previous generation
darwin-rebuild switch --rollback

# OR select specific generation
darwin-rebuild --list-generations
darwin-rebuild switch --switch-generation <number>
```

### Build Failures

```bash
# Check syntax first
nix flake check

# Detailed error output
darwin-rebuild build --flake ~/.config/nix-darwin --show-trace

# Check if package exists
nix search nixpkgs <package-name>

# Check home-manager options
# https://home-manager-options.extranix.com
```

### Changes Not Applied

```bash
# Verify symlink
ls -la ~/.config/nix-darwin
# Should point to ~/sideprojects/environment-as-code

# Restart services
launchctl kickstart -k user/$UID/org.nixos.nix-daemon

# Full rebuild
darwin-rebuild switch --flake ~/.config/nix-darwin --recreate-lock-file
```

### Nix Daemon Issues

```bash
# Restart daemon
sudo launchctl kickstart -k system/org.nixos.nix-daemon

# Check status
launchctl list | grep nix
```

### Git Conflicts After Pull

```bash
# Stash local changes
git stash

# Pull remote
git pull

# Reapply local changes
git stash pop

# Resolve conflicts manually, then rebuild
```

## Bootstrap New Machine

From fresh macOS install:

```bash
# 1. Run bootstrap script (automated)
curl -fsSL https://raw.githubusercontent.com/yourusername/environment-as-code/main/bootstrap.sh | bash

# OR manual setup:

# 2. Install Xcode CLI Tools
xcode-select --install

# 3. Install Nix
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install

# 4. Clone repo
git clone https://github.com/yourusername/environment-as-code.git ~/sideprojects/environment-as-code
cd ~/sideprojects/environment-as-code

# 5. Update flake.nix with hostname
vim flake.nix
# Replace "JJ4M9J6X2M" with output of: scutil --get LocalHostName

# 6. Update user info in modules/home.nix
vim modules/home.nix
# Set name and email

# 7. Build and activate
nix run nix-darwin -- switch --flake .#$(scutil --get LocalHostName)

# 8. Restart terminal
exec zsh

# 9. Verify
darwin-rebuild switch --flake ~/.config/nix-darwin
```

## Quality Gate Matrix

Use this checklist before rebuilding:

| Gate | Command | Pass Criteria | Fail Action |
|------|---------|---------------|-------------|
| **Syntax** | `nix flake check` | Exit 0 | Fix .nix syntax errors |
| **Build** | `darwin-rebuild build` | `./result` exists | Fix module errors, check package names |
| **Diff Review** | `nix store diff-closures` | Expected changes only | Investigate unexpected packages |
| **Git Clean** | `git status` | No uncommitted critical changes | Commit or stash |
| **Git Synced** | `git push --dry-run` | Nothing to push OR push succeeds | Push to prevent loss |

**Thresholds:**
- Minor change: Syntax + Build required
- New module/package: Syntax + Build + Diff Review required
- Dependency update: ALL gates required

**Scoring:**
- All pass → Safe to switch
- 1 fail → Fix and retry
- Multiple fails → Review change scope, consider breaking into smaller steps

## Extensibility

To add sections as project evolves:

1. Add heading in appropriate location (e.g., after Common Commands)
2. Follow structure: description → concrete examples → commands
3. Include anti-patterns if applicable
4. Add templates if task is repetitive

**Useful additions:**
- Machine-specific configurations (work vs personal)
- Custom scripts documentation (modules/scripts.nix)
- Integration with external tools (Homebrew, mise, etc.)
- Backup/restore procedures

See `.claude/skills/claude-md-gen/customization-guide.md` for:
- Adding custom sections
- Creating reusable patterns
- Adapting as requirements change
