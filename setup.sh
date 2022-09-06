#! /bin/bash

print_info() {
    echo "$(tput setaf 2)$1$(tput setaf 7)"
}

print_hint() {
    echo "$(tput setaf 3)$1$(tput setaf 7)"
}

source $(pwd)/scripts/installers
source $(pwd)/scripts/configurators

# install_xcode
# install_brew
# install_zsh
configure_zsh
# install_vim_plug
# configure_vim
# install_gobrew

# brew install --cask visual-studio-code google-chrome

# casks="todoist obsidian bitwarden"
# install_brew_casks $casks