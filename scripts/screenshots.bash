#!/bin/bash

SCRIPT_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
ROOT_DIR=$(dirname "$SCRIPT_DIR")

rm -rf "$ROOT_DIR/screenshots"
mkdir "$ROOT_DIR/screenshots"

while read -r tapefile; do
	cd "$(dirname "$ROOT_DIR/$tapefile")" && vhs <"$ROOT_DIR/$tapefile"
done < <(cd "$ROOT_DIR" && git status -u | grep screenshot.tape)
