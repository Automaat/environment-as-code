{ pkgs }:

with pkgs; [
  # Development Tools
  autoconf
  cmake
  coreutils
  diffutils
  gh
  git
  go
  gofumpt
  # golang-migrate  # Not in nixpkgs - use `migrate` or install via go
  hugo
  ninja
  nodejs
  opam
  # opencode  # Not in nixpkgs - install via npm if needed
  pyenv
  rbenv
  # ruby-build  # Included with rbenv
  dart-sass  # Was: sass
  deno
  dart
  rustup

  # Python Tools
  pipx
  python310
  python313
  python314
  ruff
  uv
  virtualenv

  # CLI Tools
  actionlint
  autojump
  bat
  buf
  chezmoi
  # create-dmg  # Not in nixpkgs - keep in homebrew if needed
  firebase-tools
  fzf
  gcalcli
  jd-diff-patch
  jq
  jwt-cli
  oniguruma
  # oxlint  # May not be in stable nixpkgs yet
  ponysay
  pre-commit
  rsync
  shellcheck
  starship
  tflint
  tree
  vale
  yamlfmt
  yamllint
  yubikey-manager
  yt-dlp
  zplug
  zsh-autosuggestions
  zsh-syntax-highlighting

  # AWS Tools
  aws-iam-authenticator
  aws-sso-cli
  awscli2
  # gimme-aws-creds  # Python tool - install via pipx if needed
  saml2aws

  # Kubernetes Tools
  consul
  k9s
  eksctl
  kubernetes-helm
  kind
  kubectl
  kubectx
  kubie
  # kumactl  # Not in nixpkgs - install from Kong/Kuma releases
  minikube
  skaffold

  # Container Tools
  docker
  docker-buildx
  docker-compose
  docker-credential-helpers
  lima

  # Database Tools
  postgresql_14

  # Security Tools
  bitwarden-cli
  gnupg
  pinentry_mac

  # Media/FFmpeg
  ffmpeg

  # Version managers
  mise

  # CLI utilities
  # gemini-cli  # Not in nixpkgs - install via npm/pip if needed
]
