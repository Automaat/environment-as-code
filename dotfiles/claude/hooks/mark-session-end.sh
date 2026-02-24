#!/bin/bash

INPUT=$(cat)
SESSION_ID=$(echo "$INPUT" | jq -r '.session_id')
CWD=$(echo "$INPUT" | jq -r '.cwd')
STOP_HOOK_ACTIVE=$(echo "$INPUT" | jq -r '.stop_hook_active')

# Prevent infinite loops if hook triggers continuation
if [ "$STOP_HOOK_ACTIVE" = "true" ]; then
  exit 0
fi

# Extract repo and PR from workdir path
# Path format: /tmp/auto-claude-dev/worktrees/REPO/pr-XXX
if [[ "$CWD" =~ /worktrees/([^/]+)/(pr-[0-9]+) ]]; then
  REPO="${BASH_REMATCH[1]}"
  PR="${BASH_REMATCH[2]}"
  MARKER_NAME="${REPO}-${PR}"
else
  # Fallback for non-PR paths (shouldn't happen in daemon context)
  MARKER_NAME="unknown"
fi

# Debug: log hook firing
DEBUG_LOG="$HOME/.auto-claude/hook-debug.log"
{
  echo "$(date): Stop hook fired"
  echo "  CWD=$CWD"
  echo "  Marker: $MARKER_NAME"
  echo "  SESSION_ID=$SESSION_ID"
} >> "$DEBUG_LOG"

# Write marker to central location
MARKER_DIR="$HOME/.auto-claude/markers"
mkdir -p "$MARKER_DIR"
MARKER_FILE="$MARKER_DIR/${MARKER_NAME}.marker"

echo "$SESSION_ID" > "$MARKER_FILE"

exit 0
