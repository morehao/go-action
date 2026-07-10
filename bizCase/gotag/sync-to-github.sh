#!/bin/bash
set -euo pipefail

# sync-to-github.sh
# 从 GitLab 仓库同步仅开源代码到 GitHub
# 用法: ./sync-to-github.sh <github-remote>

GITHUB_REMOTE="${1:-git@github.com:morehao/go-action.git}"
TEMP_BRANCH="sync-to-github/$(date +%Y%m%d-%H%M%S)"
WORK_DIR="/tmp/gotag-sync"

echo "=== Step 1: Create temp branch ==="
git checkout -b "${TEMP_BRANCH}"

echo "=== Step 2: Remove enterprise-only files ==="
# 删除含 //go:build enterprise 的文件
find . -name "*.go" -exec grep -l "//go:build enterprise" {} \; | while read -r f; do
    echo "  [REMOVE] $f"
    git rm "$f"
done

echo "=== Step 3: Remove enterprise-only directories ==="
# 如果整个目录都没文件了，删除空目录（git 会自动处理）

echo "=== Step 4: Rename _opensource.go files ==="
find . -name "*_opensource.go" | while read -r f; do
    new_name=$(echo "$f" | sed 's/_opensource//')
    echo "  [RENAME] $f -> $new_name"
    git mv "$f" "$new_name"
done

echo "=== Step 5: Clean go.mod enterprise-only dependencies ==="
if [ -f go.mod ]; then
    sed -i '' '/enterprise-only/d' go.mod
    go mod tidy
    git add go.mod go.sum
fi

echo "=== Step 6: Commit and push to GitHub ==="
git commit -m "chore: sync opensource code $(date +%Y-%m-%d)"

echo "=== Step 7: Push to GitHub ==="
git push "${GITHUB_REMOTE}" "${TEMP_BRANCH}:main"

echo "=== Done ==="
echo "Pushed ${TEMP_BRANCH} to ${GITHUB_REMOTE} as main"
echo "Switch back to original branch with: git checkout -"
