#!/usr/bin/env bash
set -euo pipefail


if [ $# -lt 1 ]; then
echo "Usage: ./init.sh github.com/username/project"; exit 1; fi
MOD=$1


# ganti placeholder di semua file
rg -l "MODULE_PATH" | xargs sed -i.bak "s#MODULE_PATH#${MOD}#g" || true
rm -f **/*.bak


# set module di go.mod
sed -i.bak "s#module .*#module ${MOD}#" go.mod && rm -f go.mod.bak


go mod tidy


echo "Initialized module to ${MOD}"