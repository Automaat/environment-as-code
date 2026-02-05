{ pkgs, fullName, userEmail, ... }:

{
  programs.git = {
    enable = true;
    lfs.enable = true;

    signing = {
      key = "683BF754A0B005CD";
      signByDefault = true;
    };

    settings = {
      user = {
        name = fullName;
        email = userEmail;
      };

      pull = {
        rebase = true;
      };

      core = {
        autocrlf = "input";
        excludesfile = "~/.gitignore_global";
      };

      url."ssh://git@github.com/" = {
        insteadOf = "https://github.com/";
      };

      format = {
        signOff = true;
      };

      commit = {
        gpgsign = true;
      };
    };
  };

  # GPG configuration for git signing
  programs.gpg = {
    enable = true;
  };

  services.gpg-agent = {
    enable = true;
    pinentry.package = pkgs.pinentry_mac;
  };

  # Link global gitignore
  home.file.".gitignore_global" = {
    source = ../../dotfiles/git/.gitignore_global;
  };
}
