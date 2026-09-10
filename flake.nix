{
  description = "macOS system configuration with nix-darwin + home-manager";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    darwin = {
      url = "github:LnL7/nix-darwin";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    home-manager = {
      url = "github:nix-community/home-manager";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    # No inputs.nixpkgs.follows: the overlay serves this flake's own build,
    # pinned against the nixpkgs, bun2nix, and rust toolchain it needs.
    oh-my-pi.url = "github:can1357/oh-my-pi";
  };

  outputs = { self, nixpkgs, darwin, home-manager, oh-my-pi }: {
    darwinConfigurations = {
      # Replace with your hostname or use: $(scutil --get LocalHostName)
      "JJ4M9J6X2M" = darwin.lib.darwinSystem {
        system = "aarch64-darwin"; # or "aarch64-darwin"

        modules = [
          ./modules/darwin.nix

          # Allow unfree packages
          { nixpkgs.config.allowUnfree = true; }

          {
            nixpkgs.overlays = [
              oh-my-pi.overlays.default
              (final: prev: {
                pre-commit = prev.pre-commit.overrideAttrs (old: {
                  doCheck = false;
                });
              })
            ];
          }

          home-manager.darwinModules.home-manager
          {
            home-manager.useGlobalPkgs = true;
            home-manager.useUserPackages = true;
            home-manager.backupFileExtension = "hm-backup";
            home-manager.users."marcin.skalski" = import ./modules/home.nix;
          }
        ];
      };
    };
  };
}
