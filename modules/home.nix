{ config, pkgs, lib, ... }:

let
  # User-specific variables - edit these values
  fullName = "Marcin Skalski";
  userEmail = "skalskimarcin33@gmail.com";

in {
  # Home Manager settings
  home.stateVersion = "24.11";
  home.username = "marcin.skalski@konghq.com";
  home.homeDirectory = "/Users/marcin.skalski@konghq.com";

  # Install CLI packages
  home.packages = import ./packages.nix { inherit pkgs; };

  # Import program configurations
  imports = [
    ./programs/zsh.nix
    ./programs/git.nix
    ./programs/vim.nix
    ./programs/vscode.nix
    ./dotfiles/ghostty.nix
    ./dotfiles/hammerspoon.nix
    ./dotfiles/k9s.nix
    ./dotfiles/mise.nix
    ./dotfiles/claude.nix
    ./dotfiles/klaudiush.nix
    ./dotfiles/starship.nix
    ./dotfiles/ssh.nix
    ./dotfiles/docker.nix
  ];

  # Pass variables to imported modules
  _module.args = { inherit fullName userEmail; };

  # Activation scripts
  home.activation = {
    # Generate SSH key if it doesn't exist
    generateSshKey = lib.hm.dag.entryAfter ["writeBoundary"] ''
      if [ ! -f "$HOME/.ssh/id_ed25519" ]; then
        $DRY_RUN_CMD mkdir -p "$HOME/.ssh"
        $DRY_RUN_CMD ssh-keygen -t ed25519 -C "${userEmail}" -f "$HOME/.ssh/id_ed25519" -N ""
        echo "SSH key generated. Add it to GitHub:"
        echo "  cat ~/.ssh/id_ed25519.pub | pbcopy"
      fi
    '';
  };

  # Environment variables
  home.sessionVariables = {
    EDITOR = "vim";
    GOPRIVATE = "github.com/Kong/*";
    MISE_AUTO_INSTALL = "1";
  };

  # Let Home Manager install and manage itself
  programs.home-manager.enable = true;
}
