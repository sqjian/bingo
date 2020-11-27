#!/usr/bin/env bash

tee -a /root/.bashrc <<-'EOF'
alias sync='rsync -av --progress --delete /vagrant ~'
alias python='python3'
EOF