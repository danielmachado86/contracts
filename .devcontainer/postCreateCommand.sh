#!/bin/bash

# clone terminal plugins
git clone https://github.com/zsh-users/zsh-autosuggestions $HOME/.oh-my-zsh/custom/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git $HOME/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting
git clone https://github.com/romkatv/powerlevel10k.git $HOME/.oh-my-zsh/custom/themes/powerlevel10k

# load plugins config files
cp $PWD/.devcontainer/.zshrc $HOME
cp $PWD/.devcontainer/.p10k.zsh $HOME

# configure git user info
git config --global user.email \"danielmcis@gmail.com\"
git config --global user.name \"Daniel Machado\"