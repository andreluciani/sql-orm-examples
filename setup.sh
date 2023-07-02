#!/bin/bash
sudo apt-get update
sudo rm -f google-chrome-stable_current_amd64.deb
wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
sudo dpkg -i google-chrome-stable_current_amd64.deb
sudo apt-get -f install -y
sudo rm -f google-chrome-stable_current_amd64.deb

# unset PGHOSTADDR