#!/usr/bin/env bash

# run.sh

if [ $# -lt 2 ]; then
  printf "\nNeed Args: <Model> <Prompt>\n\n"; exit 1
fi

args="$@"; model=$1

curl -s http://localhost:11434/api/tags | \
  jq -r .models.[].model >/tmp/models

grep -q $model /tmp/models

if [ $? -ne 0 ]; then
  printf "\n$model is not present in https://ollama.com/library\n"
  printf "\nThese are the models that exist in the library ...\n\n"
  cat -n /tmp/models; echo; exit
fi

echo; nc -zv 127.0.0.1 11434 || \
  brew services run ollama

sleep 1; nc -zv 127.0.0.1 11434 >/dev/null 2>&1

ollama ls | grep -q ^$1 || \
  ollama pull $1 >/dev/null 2>&1

go run main.go $args; echo; echo
