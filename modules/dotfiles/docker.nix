{ config, pkgs, lib, ... }:

{
  # Link Docker config
  home.file.".docker/config.json" = {
    source = ../../dotfiles/docker/config.json;
  };
}
