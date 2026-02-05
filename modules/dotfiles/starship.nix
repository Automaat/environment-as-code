{ config, pkgs, lib, ... }:

{
  programs.starship = {
    enable = true;
  };

  # Link custom starship config
  home.file.".config/starship.toml" = {
    source = ../../dotfiles/starship/starship.toml;
  };
}
