{ config, ... }:

{

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

  # Claude Code hooks
  home.file.".claude/hooks/mark-session-end.sh" = {
    source = ../../dotfiles/claude/hooks/mark-session-end.sh;
    executable = true;
  };

  # Claude Code commands
  home.file.".claude/commands/cleanup-code.md".source = ../../dotfiles/claude/commands/cleanup-code.md;
  home.file.".claude/commands/docs-review.md".source = ../../dotfiles/claude/commands/docs-review.md;
  home.file.".claude/commands/pr-review.md".source = ../../dotfiles/claude/commands/pr-review.md;
  home.file.".claude/skills/fix-review/SKILL.md".source = ../../dotfiles/claude/skills/fix-review/SKILL.md;
  home.file.".claude/skills/fix-review/references/examples.md".source = ../../dotfiles/claude/skills/fix-review/references/examples.md;
  home.file.".claude/commands/fix-review-auto.md".source = ../../dotfiles/claude/commands/fix-review-auto.md;
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

  # notion-inbox-cleanup skill
  home.file.".claude/skills/notion-inbox-cleanup/SKILL.md".source = ../../dotfiles/claude/skills/notion-inbox-cleanup/SKILL.md;

  # docx skill (document creation/editing)
  home.file.".claude/skills/docx".source = ../../dotfiles/claude/skills/docx;
}
