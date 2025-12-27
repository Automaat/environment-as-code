{ ... }:

{
  # Copy klaudiush config from chezmoi source
  home.file.".klaudiush/config.toml".source = /Users/marcin.skalski@konghq.com/.local/share/chezmoi/private_dot_klaudiush/private_config.toml;
}
