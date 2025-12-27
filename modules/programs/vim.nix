{ pkgs, ... }:

{
  programs.vim = {
    enable = true;

    plugins = with pkgs.vimPlugins; [
      # Note: vim-darcula not in nixpkgs
      # Install manually via vim-plug if needed
    ];

    extraConfig = ''
      set number
      set tabstop=4
    '';
  };
}
