git clone https://github.com/zsh-users/zsh-autosuggestions /home/vscode/.oh-my-zsh/custom/plugins/zsh-autosuggestions &&
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git /home/vscode/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting && 
git clone https://github.com/romkatv/powerlevel10k.git /home/vscode/.oh-my-zsh/custom/themes/powerlevel10k && 
cp /workspaces/contracts/.devcontainer/.zshrc /home/vscode/ && 
cp /workspaces/contracts/.devcontainer/.p10k.zsh /home/vscode/ && 
git config --global user.email \"danielmcis@gmail.com\" && 
git config --global user.name \"Daniel Machado\"