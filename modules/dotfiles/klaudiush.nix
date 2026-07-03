{ ... }:

{
  # klaudiush validator config, single source of truth at the XDG location.
  # Legacy ~/.klaudiush/config.toml is retired in favour of this file.
  xdg.configFile."klaudiush/config.toml".source = ../../dotfiles/klaudiush/config.toml;
}
