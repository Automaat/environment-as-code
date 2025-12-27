{ pkgs, ... }:

{
  # Enable nix-darwin
  system.stateVersion = 5;

  # Set primary user for homebrew and system defaults
  system.primaryUser = "marcin.skalski@konghq.com";

  # Nix configuration
  nix.settings = {
    experimental-features = "nix-command flakes";
    trusted-users = [ "@admin" ];
  };

  # System packages available to all users
  environment.systemPackages = with pkgs; [
    vim
    git
  ];

  # Homebrew integration for casks and PHP
  homebrew = {
    enable = true;

    onActivation = {
      cleanup = "zap";
      autoUpdate = true;
      upgrade = true;
    };

    taps = [
      "shivammathur/php"
    ];

    # GUI applications
    casks = import ./casks.nix;

    # Keep PHP in homebrew (not migrated to nix)
    brews = [
      # Add PHP formula here if needed
    ];
  };

  # System activation scripts
  system.activationScripts.postActivation.text = ''
    echo "Checking Xcode Command Line Tools..."
    if ! xcode-select -p &> /dev/null; then
      echo "Installing Xcode Command Line Tools..."
      xcode-select --install
    fi

    echo "Creating PARA folders..."
    mkdir -p "$HOME/Documents/1_Projects"
    mkdir -p "$HOME/Documents/2_Areas"
    mkdir -p "$HOME/Documents/3_Resources"
    mkdir -p "$HOME/Documents/4_Archive"
  '';

  # Enable Touch ID for sudo
  security.pam.services.sudo_local.touchIdAuth = true;

  # System defaults
  system.defaults = {
    dock = {
      autohide = true;
      show-recents = false;
    };

    finder = {
      AppleShowAllExtensions = true;
      FXEnableExtensionChangeWarning = false;
    };

    NSGlobalDomain = {
      AppleShowAllExtensions = true;
      InitialKeyRepeat = 15;
      KeyRepeat = 2;
    };
  };

  # Configure users
  users.users."marcin.skalski@konghq.com" = {
    home = "/Users/marcin.skalski@konghq.com";
    description = "Marcin Skalski";
  };
}
