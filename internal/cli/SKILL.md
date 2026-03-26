---
name: "internal/cli"
description: "Performs tactical terminal operations for cli subsystem. Use when doing native terminal queries, or when the user mentions cobra subcommands, audit views, or orders creation."
compatibility: ["cobra:latest", "go:1.21"]
metadata:
  gerp-domain: "cli"
  gerp-status: "DRAFT"
  gerp-reviewer: "TODO"
  gerp-id: "skill-cli-sub-01"
---

# CLI Subcommands (internal/cli)

This module holds the isolated tactical operations for the native `gerp` terminal binary. It executes `spf13/cobra` routes to directly manipulate the GraphQL and Spanner environments securely from the terminal.
