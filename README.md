# Hook

## Overview

Hook is a lightning-fast CLI tool that uses AI to generate shell commands based on a user query. Simply write what you want to do in the terminal, and let AI do the work!

Written 100% in Go. Should work in Linux, macOS and Windows!

**⚠️ AI makes mistakes. Please double-check the generated commands before executing them. ⚠️**

## Prerequisites

- Go
- Gemini API Key ([get a free one from Google AI Studio](https://aistudio.google.com/app/apikey))

## Setup

Set up your API key:
- Option 1: Set `GEMINI_API_KEY` environment variable
- Option 2 (deprecated): Create a `.api-key` file with your API key

## Installation

1. Clone the repository

2. Build the project:

```bash
go build .
```

3. Move the executable file somewhere in your PATH, for instance `/usr/local/bin`:

```bash
sudo mv ./hook /usr/local/bin
```

## Usage

```bash
./hook "describe what you want to do"
```

The program outputs the command that should be typed in the console.

## Zsh Configuration

Add to your `.zshrc`:
```bash
function __hook() {print -z $(hook ${@:2})}
alias '?'=__hook
```
Now just type `? description` and the shell prompt will be auto-filled with the suggested command. No more copy-pasting!

## TODO

- Pre-built binaries
- Easier setup and configuration (setup script)
- Proper config file support, better (OS-agnostic) API key handling
- Windows command shell and Powershell testing
- Easier ZSH configuration on install and other shell support
- Local (offline) LLM support
- OpenAI, Claude support, more?

## Acknowledgements

Inspired by [Warp terminal](https://www.warp.dev/)'s AI prompt.

## License

GPL v3
