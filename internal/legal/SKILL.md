---
name: "internal/legal"
description: "Performs corporate compliance and audit logging for legal subsystem. Use when doing cross-domain compliance tracking, or when the user mentions contracts, SOX, or immutable audit trails."
compatibility: ["spanner:latest", "go:1.21"]
metadata:
  gerp-domain: "legal"
  gerp-status: "DRAFT"
  gerp-reviewer: "TODO"
  gerp-id: "skill-legal-01"
---

# Legal & Compliance (internal/legal)

This module enforces immutable corporate governance. It tracks active B2B/B2C contracts and maintains a strict SOX/SOC2 compliance audit log across the entire matrix.

## Constraints
- **Zero Foreign Keys:** `CounterpartyID` tracks back to MDM. `TargetRecordID` tracks any UUID across all Domains. `ActorID` tracks to HCM.
- **Immutability:** `ComplianceAudits` are append-only. They mathematically prove physical state changes across the system.
