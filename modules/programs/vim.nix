{ pkgs, ... }:

{
  programs.vim = {
    enable = true;

    plugins = with pkgs.vimPlugins; [
      # Darcula color scheme
      vim-darcula
    ];

    extraConfig = ''
      colorscheme darcula
      set number
      set tabstop=4
    '';
  };
}
