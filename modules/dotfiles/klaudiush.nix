{ ... }:

{
  # Copy klaudiush config from repo
  home.file.".klaudiush/config.toml".source = ../../dotfiles/klaudiush/config.toml;
}
