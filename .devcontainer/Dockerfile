FROM golang:1.23-alpine

# Install dev dependencies
RUN apk add --update \
    bash zsh zsh-vcs git sudo \
    age htop inotify-tools \
    nodejs npm \
    gcc musl-dev \
    docker-cli docker-cli-buildx postgresql-client curl

# Install sops
RUN wget https://github.com/mozilla/sops/releases/download/v3.7.3/sops-v3.7.3.linux.amd64 -O /usr/local/bin/sops && \
    chmod +x /usr/local/bin/sops

# Create and switch to a jailed admin user
RUN echo "%sudo ALL=(ALL:ALL) NOPASSWD: ALL" > /etc/sudoers.d/sudo && \
    addgroup sudo && addgroup docker && \
    adduser -D -s /bin/zsh coder && \
    addgroup coder sudo && \
    addgroup coder docker
USER coder

ENV EDITOR="code --wait"

RUN git config --global core.editor "$EDITOR" && git config --global init.defaultBranch main

# Install oh my zsh
ARG ZSH_THEME_NAME="agnoster"
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" && \
    sed -i -e "s/ZSH_THEME=.*/ZSH_THEME=\"$ZSH_THEME_NAME\"/" ~/.zshrc
