FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Switch to root user to install dependencies and set permissions
USER root

# Install Goose command line interface
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Set permissions for Go modules cache directory
RUN chmod -R 777 /go

# Set the owner of the Go module cache to the 'vscode' user
RUN chown -R vscode:vscode /go

# Switch back to the default devcontainer user (vscode)
USER vscode

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
