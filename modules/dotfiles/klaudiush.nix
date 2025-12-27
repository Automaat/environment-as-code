{ config, ... }:

{
  # Copy klaudiush config from chezmoi source
  home.file.".klaudiush/config.toml".source = "${config.home.homeDirectory}/.local/share/chezmoi/private_dot_klaudiush/private_config.toml";
}
