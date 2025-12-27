{ ... }:

{
  xdg.configFile."mise/config.toml".text = ''
    [settings]
    idiomatic_version_file_enable_tools = ["ruby"]

    [tools]
    go = "1.24.9"
    uv = "latest"
  '';
}
