#!/usr/bin/env bash

nc -zv 127.0.0.1 11434 >/dev/null 2>&1

[ $? -ne 0 ] && brew services start ollama

if [ $# -ne 1 ]; then
  printf "\nNeed city name as arg\n\n"; exit 1
fi

sleep 3; go run main.go "$@"
