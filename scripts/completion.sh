#!/usr/bin/env bash

rm -rf completions
mkdir completions
go run . completion bash > completions/heycart-cli.bash
go run . completion zsh > completions/heycart-cli.zsh
go run . completion fish > completions/heycart-cli.fish