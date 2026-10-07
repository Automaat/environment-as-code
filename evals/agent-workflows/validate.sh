#!/usr/bin/env bash

set -euo pipefail

corpus_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cases_file="${1:-$corpus_dir/cases.json}"
schema_file="$corpus_dir/schema.json"

if [[ ! -f "$cases_file" ]]; then
  printf 'cases file is not a regular file: %s\n' "$cases_file" >&2
  exit 1
fi

uvx --from check-jsonschema==0.38.2 check-jsonschema \
  --schemafile "$schema_file" \
  "$cases_file"

duplicate_ids="$(jq -r '[.cases[].id] | group_by(.)[] | select(length > 1)[0]' "$cases_file")"
if [[ -n "$duplicate_ids" ]]; then
  printf 'duplicate case ids:\n%s\n' "$duplicate_ids" >&2
  exit 1
fi
