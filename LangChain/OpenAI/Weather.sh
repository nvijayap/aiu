#!/usr/bin/env bash

uv venv --clear >/dev/null 2>&1

source .venv/bin/activate

uv pip install -q -r Requirements.txt

python Weather.py | jq
