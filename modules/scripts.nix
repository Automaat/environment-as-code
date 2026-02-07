{ config, pkgs, ... }:

{
  # Custom scripts installed to ~/.local/bin/

  # gh-add-subissue: Add GitHub issue as subissue via GraphQL API
  home.file.".local/bin/gh-add-subissue" = {
    text = ''
      #!/usr/bin/env bash
      # Add a GitHub issue as a subissue to a parent issue
      # Usage: gh-add-subissue <parent-number> <child-number> [owner/repo]

      set -e

      if [ "$#" -lt 2 ]; then
          echo "Usage: gh-add-subissue <parent-number> <child-number> [owner/repo]"
          echo "Example: gh-add-subissue 123 456"
          echo "Example: gh-add-subissue 123 456 owner/repo"
          exit 1
      fi

      PARENT_NUM=$1
      CHILD_NUM=$2
      REPO=''${3:-$(gh repo view --json nameWithOwner -q .nameWithOwner)}

      if [ -z "$REPO" ]; then
          echo "Error: Not in a git repository and no owner/repo specified"
          exit 1
      fi

      OWNER=$(echo "$REPO" | cut -d'/' -f1)
      REPO_NAME=$(echo "$REPO" | cut -d'/' -f2)

      echo "Getting issue IDs for $REPO..."

      # Get parent issue ID
      PARENT_ID=$(gh api graphql -f query="
        query {
          repository(owner: \"$OWNER\", name: \"$REPO_NAME\") {
            issue(number: $PARENT_NUM) {
              id
              title
            }
          }
        }
      " -q '.data.repository.issue.id')

      PARENT_TITLE=$(gh api graphql -f query="
        query {
          repository(owner: \"$OWNER\", name: \"$REPO_NAME\") {
            issue(number: $PARENT_NUM) {
              title
            }
          }
        }
      " -q '.data.repository.issue.title')

      # Get child issue ID
      CHILD_ID=$(gh api graphql -f query="
        query {
          repository(owner: \"$OWNER\", name: \"$REPO_NAME\") {
            issue(number: $CHILD_NUM) {
              id
              title
            }
          }
        }
      " -q '.data.repository.issue.id')

      CHILD_TITLE=$(gh api graphql -f query="
        query {
          repository(owner: \"$OWNER\", name: \"$REPO_NAME\") {
            issue(number: $CHILD_NUM) {
              title
            }
          }
        }
      " -q '.data.repository.issue.title')

      if [ -z "$PARENT_ID" ] || [ -z "$CHILD_ID" ]; then
          echo "Error: Could not find issue IDs"
          exit 1
      fi

      echo "Parent: #$PARENT_NUM - $PARENT_TITLE"
      echo "Child:  #$CHILD_NUM - $CHILD_TITLE"
      echo ""
      echo "Adding #$CHILD_NUM as subissue of #$PARENT_NUM..."

      # Add subissue with required header
      gh api graphql \
        -H "GraphQL-Features: sub_issues" \
        -f query="
          mutation {
            addSubIssue(input: {
              issueId: \"$PARENT_ID\"
              subIssueId: \"$CHILD_ID\"
            }) {
              issue {
                number
                title
                subIssuesSummary {
                  total
                  completed
                }
              }
              subIssue {
                number
                title
              }
            }
          }
        " -q '.data.addSubIssue | "✓ Added #\(.subIssue.number) as subissue of #\(.issue.number)\nParent now has \(.issue.subIssuesSummary.total) subissue(s) (\(.issue.subIssuesSummary.completed) completed)"'

      echo ""
      echo "Done!"
    '';
    executable = true;
  };
}
