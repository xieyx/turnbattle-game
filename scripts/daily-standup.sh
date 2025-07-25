#!/bin/bash
# scripts/daily-standup.sh - 生成每日站会报告

echo "=== 昨日完成的PR ==="
gh pr list --state merged --limit 10 --json number,title,author

echo "=== 进行中的Issues ==="
gh issue list --label "status:in-progress" --json number,title,assignees

echo "=== 待Review的PR ==="
gh pr list --state open --json number,title,author,reviewRequests

# 使用方式: ./scripts/daily-standup.sh
