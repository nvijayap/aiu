#!/usr/bin/env bash

cd `dirname $0`

if [ $# -ne 1 ]; then
  printf "\nNeed city name as arg\n\n"; exit 1
fi

nc -zv 127.0.0.1 11434 >/dev/null 2>&1

[ $? -ne 0 ] && brew services start ollama && sleep 3

ollama list | grep -q ^llama3.1

if [ $? -ne 0 ]; then
  ollama pull llama3.1
fi

go build && ./a2a "$1"
