#!/bin/bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
ROOT_DIR=$(dirname "$SCRIPT_DIR")

# rm -rf "$ROOT_DIR/screenshots"
mkdir -p "$ROOT_DIR/screenshots"

if [[ "$1" == "-f" ]]; then
	while read -r tapefile; do
		cd "$(dirname "$ROOT_DIR/$tapefile")" && vhs <"$ROOT_DIR/$tapefile"
	done < <(find . -name screenshot.tape)
else
	while read -r tapefile; do
		cd "$(dirname "$ROOT_DIR/$tapefile")" && vhs <"$ROOT_DIR/$tapefile"
	done < <(cd "$ROOT_DIR" && git status -u -s | awk '{print $2}' | grep screenshot.tape)
fi
