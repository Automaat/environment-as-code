{ config, pkgs, lib, ... }:

{
  programs.zsh = {
    enable = true;

    enableCompletion = true;
    autosuggestion.enable = true;
    syntaxHighlighting.enable = true;

    # History settings
    history = {
      size = 100000000000000;
      save = 100000000000000;
      extended = true;
      ignoreAllDups = true;
      ignoreDups = true;
      share = true;
      path = "${config.home.homeDirectory}/.zsh_history";
    };

    # oh-my-zsh configuration
    oh-my-zsh = {
      enable = true;
      theme = "robbyrussell";
      plugins = [
        "git"
        "brew"
        "docker"
        "docker-compose"
        "golang"
        "helm"
        "httpie"
        "kubectl"
        "macos"
        "pip"
        "python"
        "terraform"
        "fzf"
      ];
    };

    # ZPlug plugins
    zplug = {
      enable = true;
      plugins = [
        { name = "MichaelAquilina/zsh-you-should-use"; }
        { name = "automaat/zsh-clean-history"; tags = [ from:github at:main ]; }
      ];
    };

    # Additional shell initialization
    initContent = ''
      # XDG Base Directory - ensures k9s uses ~/.config instead of Application Support
      export XDG_CONFIG_HOME="$HOME/.config"

      # FZF integration
      source ${pkgs.fzf}/share/fzf/key-bindings.zsh
      source ${pkgs.fzf}/share/fzf/completion.zsh

      # autojump
      source ${pkgs.autojump}/share/autojump/autojump.zsh

      # pyenv
      export PYENV_ROOT="$HOME/.pyenv"
      command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"
      eval "$(pyenv init -)"

      # rbenv
      eval "$(rbenv init - zsh)"

      # mise
      eval "$(mise activate zsh)"

      # saml2aws completion
      eval "$(saml2aws --completion-script-zsh)"

      # opam configuration
      [[ ! -r ~/.opam/opam-init/init.zsh ]] || source ~/.opam/opam-init/init.zsh > /dev/null 2> /dev/null

      # Aliases
      alias sb="cd $HOME/Library/Mobile\ Documents/iCloud~md~obsidian/Documents"

      # Work-specific aliases (conditional on ~/kong directory existing)
      if [ -d "$HOME/kong" ]; then
        alias sync-kuma="gh repo sync automaat/kuma -b master"
        alias sync-kuma-website="gh repo sync automaat/kuma-website -b master"
        alias awsk="aws --profile saml --region us-west-1"
        alias kol="bash $HOME/kong/konnect-on-call-tool-kit/cli/kolsch.sh"
        alias kctl="$HOME/kong/mesh/kuma/build/artifacts-darwin-arm64/kumactl/kumactl"
      fi

      # Kuma dev bin (conditional)
      if [ -d "$HOME/.kuma-dev/kuma-master/bin" ]; then
        export PATH="$HOME/.kuma-dev/kuma-master/bin:$PATH"
      fi

      # Antigravity (conditional)
      if [ -d "$HOME/.antigravity/antigravity/bin" ]; then
        export PATH="$HOME/.antigravity/antigravity/bin:$PATH"
      fi

      # Additional PATH entries
      export PATH="$HOME/bin:$PATH"
      export PATH="/opt/homebrew/opt/ruby/bin:$PATH"
      export PATH="$PATH:$HOME/go/bin"
      export PATH="$PATH:$HOME/.local/bin"
      export PATH="$PATH:/Applications/ToolHive.app/Contents/Resources/bin/darwin-arm64"

      # Custom aliases
      alias cs="carousel-splitter"
    '';
  };

  # FZF configuration
  programs.fzf = {
    enable = true;
    enableZshIntegration = true;
  };

  # Direnv for per-project environments
  programs.direnv = {
    enable = true;
    enableZshIntegration = true;
    nix-direnv.enable = true;
  };

  # Starship prompt
  programs.starship = {
    enable = true;
    settings = {
      # Add custom starship config here if needed
    };
  };
}
