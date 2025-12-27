{ config, ... }:

{
  # Claude Code settings
  home.file.".claude/settings.json".text = builtins.toJSON {
    env = {
      SHELL = "/bin/zsh";
    };
    includeCoAuthoredBy = false;
    permissions = {
      allow = [
        "WebSearch"
        "WebFetch"
        "Read"
        "Glob"
        "Grep"
        "TodoWrite"
        "Bash(actionlint:*)"
        "Bash(base64:*)"
        "Bash(cat:*)"
        "Bash(checkmake:*)"
        "Bash(command:*)"
        "Bash(diff:*)"
        "Bash(docker pull:*)"
        "Bash(echo:*)"
        "Bash(gh pr checks:*)"
        "Bash(gh pr view:*)"
        "Bash(gh repo list:*)"
        "Bash(gh repo view:*)"
        "Bash(gh run list:*)"
        "Bash(gh run view:*)"
        "Bash(gh run watch:*)"
        "Bash(gh variable list:*)"
        "Bash(git branch:*)"
        "Bash(git cat-file:*)"
        "Bash(git cat:*)"
        "Bash(git config --get:*)"
        "Bash(git describe:*)"
        "Bash(git diff:*)"
        "Bash(git fetch:*)"
        "Bash(git log:*)"
        "Bash(git ls-files:*)"
        "Bash(git ls-tree:*)"
        "Bash(git remote:*)"
        "Bash(git rev-list:*)"
        "Bash(git rev-parse:*)"
        "Bash(git show:*)"
        "Bash(git status:*)"
        "Bash(git worktree list:*)"
        "Bash(go fmt:*)"
        "Bash(go mod:*)"
        "Bash(golangci-lint:*)"
        "Bash(grep:*)"
        "Bash(hadolint:*)"
        "Bash(head:*)"
        "Bash(jq:*)"
        "Bash(ls:*)"
        "Bash(make check:*)"
        "Bash(make fmt:*)"
        "Bash(make format:*)"
        "Bash(make lint:*)"
        "Bash(make test:*)"
        "Bash(markdownlint:*)"
        "Bash(mise:*)"
        "Bash(mkdir:*)"
        "Bash(printenv:*)"
        "Bash(printf:*)"
        "Bash(shellspec:*)"
        "Bash(task check:*)"
        "Bash(task fmt:*)"
        "Bash(task format:*)"
        "Bash(task lint:*)"
        "Bash(task test:*)"
        "Bash(tree:*)"
        "Bash(which:*)"
        "Bash(yamllint:*)"
        "WebFetch"
      ];
    };
    hooks = {
      PreToolUse = [
        {
          matcher = "Bash";
          hooks = [
            {
              type = "command";
              command = "klaudiush --hook-type PreToolUse --trace";
            }
          ];
        }
        {
          matcher = "Write|Edit|MultiEdit";
          hooks = [
            {
              type = "command";
              command = "klaudiush --hook-type PreToolUse --trace";
              timeout = 30;
            }
          ];
        }
      ];
      Notification = [
        {
          hooks = [
            {
              type = "command";
              command = "klaudiush --hook-type Notification --trace";
            }
          ];
        }
      ];
    };
    statusLine = {
      type = "command";
      command = "bash ${config.home.homeDirectory}/.claude/statusline-command.sh";
    };
    alwaysThinkingEnabled = true;
    model = "opus";
  };

  # Status line script
  home.file.".claude/statusline-command.sh".text = ''
    #!/usr/bin/env bash

    # Read JSON input from stdin
    input=$(cat)

    # Extract current working directory
    cwd=$(echo "$input" | jq -r '.workspace.current_dir')

    # Extract project name (github.com/org/repo format if applicable)
    if echo "$cwd" | grep -q 'Projects/github.com/'; then
        project=$(echo "$cwd" | sed -E 's|.*/Projects/github.com/([^/]+/[^/]+).*|\1|')
    else
        project=$(basename "$cwd")
    fi

    # Check if we're in a git repository and get the branch name
    if git -C "$cwd" rev-parse --git-dir > /dev/null 2>&1; then
        branch=$(git -C "$cwd" rev-parse --abbrev-ref HEAD 2>/dev/null)
        project_branch="\033[36m''${project}\033[0m:\033[33m''${branch}\033[0m"
    else
        # Just show project in cyan if not in git repo
        project_branch="\033[36m''${project}\033[0m"
    fi

    # Extract file context from the JSON input if available
    # The input may contain a 'file' field with the current file being edited
    file=$(echo "$input" | jq -r '.file // empty')

    # Build the status line
    if [[ -n "$file" ]]; then
        # Display: project:branch | file: filename
        printf "%b | file: \033[32m%s\033[0m\n\n" "$project_branch" "$file"
    else
        # Just show project:branch if no file context
        printf "%b\n\n" "$project_branch"
    fi
  '';

  # Copy CLAUDE.md from chezmoi source
  home.file.".claude/CLAUDE.md".source = "${config.home.homeDirectory}/.local/share/chezmoi/dot_claude/CLAUDE.md";

  # Claude Code commands
  home.file.".claude/commands/cleanup-code.md".source = "${config.home.homeDirectory}/.local/share/chezmoi/dot_claude/commands/cleanup-code.md";
  home.file.".claude/commands/docs-review.md".source = "${config.home.homeDirectory}/.local/share/chezmoi/dot_claude/commands/docs-review.md";
  home.file.".claude/commands/pr-review.md".source = "${config.home.homeDirectory}/.local/share/chezmoi/dot_claude/commands/pr-review.md";

  # Claude Code skills
  home.file.".claude/skills/go-code-review/SKILL.md".source = "${config.home.homeDirectory}/.claude/skills/go-code-review/SKILL.md";
  home.file.".claude/skills/go-code-review/knowledge-base.md".source = "${config.home.homeDirectory}/.claude/skills/go-code-review/knowledge-base.md";
  home.file.".claude/skills/go-code-review/real-world-patterns.md".source = "${config.home.homeDirectory}/.claude/skills/go-code-review/real-world-patterns.md";
}
