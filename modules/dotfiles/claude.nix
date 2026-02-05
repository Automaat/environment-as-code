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
        "Bash(fc-list:*)"
        "Bash(ghostty +show-config:*)"
        "Bash(ghostty:*)"
        "Bash(cp:*)"
        "Bash(rm:*)"
        "Bash(lsof:*)"
        "Bash(ps:*)"
        "Bash(git checkout:*)"
        "Bash(git add:*)"
        "Bash(git commit:*)"
        "Bash(git push:*)"
        "Bash(gh pr create:*)"
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

  # Copy CLAUDE.md from repo
  home.file.".claude/CLAUDE.md".source = ../../dotfiles/claude/CLAUDE.md;

  # Claude Code commands
  home.file.".claude/commands/cleanup-code.md".source = ../../dotfiles/claude/commands/cleanup-code.md;
  home.file.".claude/commands/docs-review.md".source = ../../dotfiles/claude/commands/docs-review.md;
  home.file.".claude/commands/pr-review.md".source = ../../dotfiles/claude/commands/pr-review.md;
  home.file.".claude/commands/fix-review.md".source = ../../dotfiles/claude/commands/fix-review.md;
  home.file.".claude/commands/go-review.md".source = ../../dotfiles/claude/commands/go-review.md;
  home.file.".claude/commands/merge-renovate.md".source = ../../dotfiles/claude/commands/merge-renovate.md;

  # Claude Code skills
  home.file.".claude/skills/go-code-review/SKILL.md".source = ../../dotfiles/claude/skills/go-code-review/SKILL.md;
  home.file.".claude/skills/go-code-review/knowledge-base.md".source = ../../dotfiles/claude/skills/go-code-review/knowledge-base.md;
  home.file.".claude/skills/go-code-review/real-world-patterns.md".source = ../../dotfiles/claude/skills/go-code-review/real-world-patterns.md;

  # claude-md-gen skill
  home.file.".claude/skills/claude-md-gen/SKILL.md".source = ../../dotfiles/claude/skills/claude-md-gen/SKILL.md;
  home.file.".claude/skills/claude-md-gen/checklist.md".source = ../../dotfiles/claude/skills/claude-md-gen/checklist.md;
  home.file.".claude/skills/claude-md-gen/customization-guide.md".source = ../../dotfiles/claude/skills/claude-md-gen/customization-guide.md;
  home.file.".claude/skills/claude-md-gen/patterns/triage-workflows.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/triage-workflows.md;
  home.file.".claude/skills/claude-md-gen/patterns/decision-matrices.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/decision-matrices.md;
  home.file.".claude/skills/claude-md-gen/patterns/output-templates.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/output-templates.md;
  home.file.".claude/skills/claude-md-gen/patterns/cove-verification.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/cove-verification.md;
  home.file.".claude/skills/claude-md-gen/patterns/few-shot-examples.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/few-shot-examples.md;
  home.file.".claude/skills/claude-md-gen/patterns/citation-systems.md".source = ../../dotfiles/claude/skills/claude-md-gen/patterns/citation-systems.md;
  home.file.".claude/skills/claude-md-gen/templates/research-knowledge.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/research-knowledge.md;
  home.file.".claude/skills/claude-md-gen/templates/mixed.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/mixed.md;
  home.file.".claude/skills/claude-md-gen/templates/web-app.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/web-app.md;
  home.file.".claude/skills/claude-md-gen/templates/cli-tool.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/cli-tool.md;
  home.file.".claude/skills/claude-md-gen/templates/library.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/library.md;
  home.file.".claude/skills/claude-md-gen/templates/python-cli.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/python-cli.md;
  home.file.".claude/skills/claude-md-gen/templates/api-service.md".source = ../../dotfiles/claude/skills/claude-md-gen/templates/api-service.md;
  home.file.".claude/skills/claude-md-gen/questions/library-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/library-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/api-service-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/api-service-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/python-cli-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/python-cli-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/research-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/research-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/mixed-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/mixed-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/web-app-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/web-app-questions.md;
  home.file.".claude/skills/claude-md-gen/questions/cli-tool-questions.md".source = ../../dotfiles/claude/skills/claude-md-gen/questions/cli-tool-questions.md;

  # obsidian-inbox-cleanup skill
  home.file.".claude/skills/obsidian-inbox-cleanup/SKILL.md".source = ../../dotfiles/claude/skills/obsidian-inbox-cleanup/SKILL.md;
}
