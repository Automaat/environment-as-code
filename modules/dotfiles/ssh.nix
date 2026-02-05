{ config, pkgs, lib, ... }:

{
  # Link SSH config
  # NOTE: Homelab IP and Bitwarden agent path are specific to this machine
  # Edit dotfiles/ssh/config for your setup
  home.file.".ssh/config" = {
    source = ../../dotfiles/ssh/config;
  };

  # Ensure SSH directory exists with correct permissions
  home.activation.ensureSshDir = lib.hm.dag.entryBefore ["checkLinkTargets"] ''
    $DRY_RUN_CMD mkdir -p "$HOME/.ssh"
    $DRY_RUN_CMD chmod 700 "$HOME/.ssh"
  '';
}
