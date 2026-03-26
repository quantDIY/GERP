---
name: "internal/revenue"
description: "Performs B2B and B2C sales tracking for revenue subsystem. Use when doing cross-domain order orchestration, or when the user mentions customers, sales orders, or credit limits."
compatibility: ["spanner:latest", "go:1.21"]
metadata:
  gerp-domain: "revenue"
  gerp-status: "DRAFT"
  gerp-reviewer: "TODO"
  gerp-id: "skill-revenue-01"
---

# Revenue & Sales (internal/revenue)

This module powers the top-line engine. It manages customer profiles, credit allocations, and the sales orders that act as the genesis for downstream fulfillment and manufacturing workflows.

## Constraints
- **Zero Foreign Keys:** `MasterDataID` physically connects to the `mdm` universal translator. `AccountManagerID` connects to the `hcm` personnel graph. These are pure UUIDs.
- **Amounts:** Standard `int64` minor unit (cents) structures enforce financial precision.
