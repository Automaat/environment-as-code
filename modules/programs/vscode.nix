{ config, pkgs, lib, ... }:

{
  # VSCode is installed via homebrew cask (modules/casks.nix)
  # This module links settings and keybindings

  # Link VSCode settings
  home.file."Library/Application Support/Code/User/settings.json" = {
    source = ../../dotfiles/vscode/settings.json;
  };

  # Link VSCode keybindings
  home.file."Library/Application Support/Code/User/keybindings.json" = {
    source = ../../dotfiles/vscode/keybindings.json;
  };

  # Note: VSCode extensions are not managed by nix
  # To list current extensions: code --list-extensions
  # Extensions will need to be reinstalled manually or via script
}
