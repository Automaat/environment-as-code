# Pre-Migration Audit Results

## Summary
Comprehensive audit completed. Added 7 high-priority configs to nix-darwin setup.

## Implemented (Phase 1)

### Added Configs
1. **Starship** - `modules/dotfiles/starship.nix` + `dotfiles/starship/starship.toml`
   - Custom prompt format for Go/Python/Rust/K8s
   - Language-specific symbols, performance tuning

2. **SSH** - `modules/dotfiles/ssh.nix` + `dotfiles/ssh/config`
   - OrbStack/Colima integration
   - Homelab host (192.168.0.241)
   - GitHub key selection
   - Bitwarden SSH agent fallback
   - **NOTE:** Edit homelab IP/user for your setup

3. **Global Gitignore** - `dotfiles/git/.gitignore_global`
   - macOS, editors, Python, Node, build artifacts
   - Linked via `modules/programs/git.nix`

4. **VSCode** - `modules/programs/vscode.nix` + `dotfiles/vscode/{settings,keybindings}.json`
   - Theme: Nord Light / Everforest Light
   - Font: IosevkaTerm Nerd Font Mono
   - Formatters: ruff (Python), YAML, markdown
   - Git: autoStage, smartCommit, alwaysSignOff
   - Mise: explicit bin path

5. **Docker** - `modules/dotfiles/docker.nix` + `dotfiles/docker/config.json`
   - ECR auth (709329596171.dkr.ecr.us-west-2.amazonaws.com)
   - osxkeychain credsStore
   - orbstack context

### Updated Files
- `modules/home.nix` - imported starship, ssh, docker modules
- `modules/programs/git.nix` - added core.excludesfile
- `modules/casks.nix` - fixed todoist-app → todoist

## Verification
```bash
# Run rebuild
sudo darwin-rebuild switch --flake .

# Check symlinks
ls -la ~/.config/starship.toml
ls -la ~/.ssh/config
ls -la ~/.gitignore_global
ls -la ~/Library/Application\ Support/Code/User/settings.json

# Original files backed up with .backup suffix
ls -la ~/.config/starship.toml.backup
ls -la ~/.ssh/config.backup
```

## Not Tracked (Sensitive/Ephemeral)
- `~/.ssh/id_*` - SSH private keys (backup separately via encrypted storage)
- `~/.aws/credentials` - Use aws-sso-cli
- `~/.npmrc` - Regenerate auth tokens
- `~/.config/raycast/config.json` - API tokens
- `~/.gnupg/private-keys-v1.d/` - GPG private keys (backup separately)

## TODO (Low Priority)
1. **Espanso** - `~/Library/Application Support/espanso/`
   - Minimal customization (mostly defaults)
   - Add if text expansion actively used

2. **Local binaries** - `~/.local/bin/`
   - `ddev`, `ddev-starship` - Document in README as manual install
   - `claude` - Claude CLI (separate install)
   - `klaudiush` - Already configured via nix

3. **gcloud-cli cask** - Already have google-cloud-sdk CLI in packages.nix
   - Verify if cask redundant, remove if needed

## Migration Checklist
When setting up new machine:

1. Clone repo, run darwin-rebuild
2. **Edit machine-specific values:**
   - `dotfiles/ssh/config` - homelab IP, Bitwarden agent path
   - `modules/home.nix` - fullName, userEmail, username, homeDirectory
3. **Restore secrets:**
   - GPG key: `gpg --export-secret-keys 683BF754A0B005CD > private.gpg`
   - SSH keys: Copy `~/.ssh/id_ed25519`, `~/.ssh/id_ed25519_homelab`
   - Import GPG: `gpg --import private.gpg`
4. **VSCode extensions:**
   - Extensions not managed by nix
   - List current: `code --list-extensions`
   - Install script or manual
5. **Verify:**
   - Git signing works: `git commit -S -s -m "test"`
   - Starship prompt displays
   - SSH to homelab/GitHub
   - VSCode settings applied

## Coverage
- **Already tracked:** 120+ packages, 58+ casks, 10+ dotfile configs
- **Added:** 7 configs (starship, SSH, gitignore, VSCode x2, docker)
- **Sensitive (skip):** 7 categories (keys, credentials, tokens)
- **Optional:** 3 items (espanso, local bins, gcloud-cli cask)

## Notes
- All backups created automatically with `.backup` suffix
- Symlinks managed by home-manager
- SSH config uses `~` instead of full paths for portability
- VSCode settings include mise bin path (needs username in path)
