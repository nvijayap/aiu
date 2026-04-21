#!/usr/bin/env bash

trap "brew services stop ollama" EXIT SIGINT SIGTERM

nc -zv 127.0.0.1 11434 || brew services start ollama

sleep 2; go run main.go "$@"
