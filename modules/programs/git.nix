{ pkgs, fullName, userEmail, ... }:

{
  programs.git = {
    enable = true;

    userName = fullName;
    userEmail = userEmail;

    signing = {
      key = "683BF754A0B005CD";
      signByDefault = true;
    };

    lfs.enable = true;

    extraConfig = {
      pull = {
        rebase = true;
      };

      core = {
        autocrlf = "input";
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
    pinentryPackage = pkgs.pinentry_mac;
  };
}
