{ pkgs, ... }:

{
  # VSCode is installed via homebrew cask (modules/casks.nix)
  # This module configures extensions and settings

  # programs.vscode = {
  #   enable = false;  # Using homebrew cask instead
  #
  #   # Add extensions here:
  #   # extensions = with pkgs.vscode-extensions; [
  #   #   # Example:
  #   #   # golang.go
  #   #   # ms-python.python
  #   # ];
  #
  #   # User settings
  #   # userSettings = {
  #   #   "editor.fontSize" = 14;
  #   #   "editor.fontFamily" = "Iosevka";
  #   # };
  # };

  # Note: To migrate VSCode extensions, you can:
  # 1. List current extensions: code --list-extensions
  # 2. Map them to nixpkgs vscode-extensions
  # 3. Add them to the extensions list above
  # 4. Enable programs.vscode and remove vscode from casks.nix
}
