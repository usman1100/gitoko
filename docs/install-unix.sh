#!/bin/bash
# install.sh - Installation script for Gitoko

set -e

# Define colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Installing Gitoko - Git Cherry-Picking Tool${NC}"

# Detect OS and architecture
OS="$(uname -s)"
ARCH="$(uname -m)"

# Set download URL based on latest release
REPO="github.com/usman1100/gitoko"
API_URL="https://api.github.com/repos/usman1100/gitoko/releases/latest"
DOWNLOAD_URL=""

# Get the latest release info
if command -v curl &> /dev/null; then
    RELEASE_INFO=$(curl -s $API_URL)
elif command -v wget &> /dev/null; then
    RELEASE_INFO=$(wget -q -O- $API_URL)
else
    echo -e "${RED}Error: Neither curl nor wget found. Please install one of them.${NC}"
    exit 1
fi

# Determine the correct binary based on OS and architecture
if [[ "$OS" == "Darwin" ]]; then
    if [[ "$ARCH" == "arm64" ]]; then
        BINARY_NAME="gitoko-mac-arm64"
    else
        BINARY_NAME="gitoko-mac-amd64"
    fi
elif [[ "$OS" == "Linux" ]]; then
    if [[ "$ARCH" == "aarch64" || "$ARCH" == "arm64" ]]; then
        BINARY_NAME="gitoko-linux-arm64"
    else
        BINARY_NAME="gitoko-linux-amd64"
    fi
else
    echo -e "${RED}Unsupported OS: $OS. Please download the appropriate binary manually.${NC}"
    exit 1
fi

# Extract download URL
if command -v jq &> /dev/null; then
    DOWNLOAD_URL=$(echo "$RELEASE_INFO" | jq -r ".assets[] | select(.name == \"$BINARY_NAME\") | .browser_download_url")
else
    # Fallback to grep and sed if jq is not available
    DOWNLOAD_URL=$(echo "$RELEASE_INFO" | grep -o "\"browser_download_url\":\"[^\"]*$BINARY_NAME\"" | sed 's/"browser_download_url":"\(.*\)"/\1/')
fi

if [[ -z "$DOWNLOAD_URL" ]]; then
    echo -e "${RED}Error: Could not find download URL for $BINARY_NAME${NC}"
    exit 1
fi

# Create bin directory if it doesn't exist
mkdir -p ~/bin

# Download the binary
echo -e "${BLUE}Downloading Gitoko...${NC}"
if command -v curl &> /dev/null; then
    curl -sL "$DOWNLOAD_URL" -o ~/bin/gitoko
elif command -v wget &> /dev/null; then
    wget -q "$DOWNLOAD_URL" -O ~/bin/gitoko
fi

# Make it executable
chmod +x ~/bin/gitoko

# Add ~/bin to PATH if it's not already there
if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    SHELL_NAME=$(basename "$SHELL")
    if [[ "$SHELL_NAME" == "zsh" ]]; then
        PROFILE_FILE="$HOME/.zshrc"
    elif [[ "$SHELL_NAME" == "bash" ]]; then
        if [[ "$OS" == "Darwin" ]]; then
            PROFILE_FILE="$HOME/.bash_profile"
        else
            PROFILE_FILE="$HOME/.bashrc"
        fi
    else
        PROFILE_FILE="$HOME/.profile"
    fi
    
    echo -e "${BLUE}Adding ~/bin to your PATH in $PROFILE_FILE${NC}"
    echo 'export PATH="$HOME/bin:$PATH"' >> "$PROFILE_FILE"
    echo -e "${GREEN}Added ~/bin to PATH. Please restart your terminal or run: source $PROFILE_FILE${NC}"
else
    echo -e "${GREEN}~/bin is already in your PATH${NC}"
fi

echo -e "${GREEN}Gitoko has been successfully installed!${NC}"
echo -e "${BLUE}You can now use it by running 'gitoko' in any git repository.${NC}"