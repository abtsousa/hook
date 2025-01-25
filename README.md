# Hook

## Overview

Hook is a simple, lightning-fast CLI tool that uses the Gemini API to generate shell commands based on a user query.

Written 100% in Go. Should work in Linux, macOS and Windows!

⚠️ AI makes mistakes. Double-check the generated commands before executing them. ⚠️

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
Now just type `? description` and the suggested command will be automatically redirected into your shell input. No more copy-pasting!

## TODO

- Pre-built binaries
- Proper config file support, better (OS-agnostic) API key handling
- Windows command shell and Powershell testing
- Easier ZSH configuration on install and other shell support
- Local (offline) LLM support
- OpenAI, Claude support, more?

## License

GPL v3
