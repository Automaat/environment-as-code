#! /bin/bash

print_info() {
    echo "$(tput setaf 2)$1$(tput setaf 7)"
}

print_hint() {
    echo "$(tput setaf 3)$1$(tput setaf 7)"
}

source $(pwd)/scripts/installers
source $(pwd)/scripts/configurators
source $(pwd)/configurations/brew_casks

## basics
install_xcode
install_brew
configure_github_ssh_key "fill.your@email.com"

## zsh with plugins
install_zsh
install_zplug
install_fzf
configure_zsh

## VIM
install_vim_plug
configure_vim

## Golang
install_gobrew

## Working directories
create_para_folders

## Apps
install_brew_casks $basic_casks
install_brew_casks $company_specific_casks
# brew install docker