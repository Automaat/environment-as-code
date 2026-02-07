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
  };

  outputs = { self, nixpkgs, darwin, home-manager }: {
    darwinConfigurations = {
      # Replace with your hostname or use: $(scutil --get LocalHostName)
      "JJ4M9J6X2M" = darwin.lib.darwinSystem {
        system = "aarch64-darwin"; # or "aarch64-darwin"

        modules = [
          ./modules/darwin.nix

          # Allow unfree packages
          { nixpkgs.config.allowUnfree = true; }

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
