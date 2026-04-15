#!/usr/bin/env bash

# run.sh

if [ $# -lt 2 ]; then
  printf "\nNeed Args: <Model> <Prompt>\n\n"; exit 1
fi

echo; nc -zv 127.0.0.1 11434 || brew services run ollama
sleep 1; nc -zv 127.0.0.1 11434

ollama ls | grep -q ^$1 || ollama pull $1 >/dev/null 2>&1
go run main.go "$@"; echo; echo
