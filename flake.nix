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
      "default" = darwin.lib.darwinSystem {
        system = "aarch64-darwin"; # or "x86_64-darwin"

        modules = [
          ./modules/darwin.nix

          home-manager.darwinModules.home-manager
          {
            home-manager.useGlobalPkgs = true;
            home-manager.useUserPackages = true;
            home-manager.users."marcin.skalski@konghq.com" = import ./modules/home.nix;
          }
        ];
      };
    };
  };
}
