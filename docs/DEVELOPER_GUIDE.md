# GERP Developer Onboarding Guide

Welcome to GERP (Google ERP). This system is a FAANG-grade, multi-domain Enterprise Resource Planning matrix structurally isolated by Domain-Driven Design constraints.

## 1. The Golden Thread (UUID Stitching)
To guarantee distributed, isolated micro-database performance at scale, GERP completely abandons standard SQL foreign keys. No domain is permitted to link its physical state via SQL layer to another domain. 

Instead, GERP employs the **Golden Thread**: strict UUID pointers managed entirely in application space. A `SalesOrder` in `internal/revenue` physically stores a `CustomerID` (UUID), but it never explicitly joins that to the `mdm.GlobalEntities` table at the Spanner level. The graph stitching occurs strictly in memory at the BFF level.

## 2. The 7 Tier-1 Domains
GERP separates its global state into 7 isolated execution environments:
1. **Finance (`internal/finance`):** The immutable double-entry ledger (`Accounts`, `LedgerEntries`, `LineItems`).
2. **Human Capital (`internal/hcm`):** The employee and compensation engine (`Employees`, `PayrollRuns`).
3. **Supply Chain (`internal/scm`):** The physical tracking bounds (`Products`, `InventoryLots`).
4. **Enterprise Asset Management (`internal/eam`):** The infrastructure core (`Assets`, `MaintenanceLogs`).
5. **Legal & Compliance (`internal/legal`):** The sovereign audit logging domain (`Contracts`, `ComplianceAudits`).
6. **Revenue & Sales (`internal/revenue`):** The top-line commercial router (`Customers`, `SalesOrders`).
7. **Learning Management (`internal/lms`):** Educational compliance bounds (`Courses`, `Enrollments`, `Certifications`).

**The Universal Translator:** `internal/mdm` bridges these environments by tracking a `GlobalEntity` mapped to downstream `LocalID` pointers.

## 3. The Temporal Saga Orchestrator (`internal/pipeline`)
To execute multi-domain mutations (e.g., selling a product requires adjusting inventory and charging the ledger), you cannot wrap distinct Spanner databases in a single ACID transaction. 

GERP solves this with **Temporal Workflows**. The `pipeline` handles transactions asynchronously. If it subtracts 10 servers from SCM but the Finance domain rejects the ledger charge due to limits, the Temporal saga automatically executes a **Compensating Rollback** (via `ReverseInventoryActivity`) to guarantee eventual consistency and mathematically eliminate phantom locks.

## 4. The GraphQL BFF (`cmd/gateway`)
The Backend-For-Frontend receives a universal, deeply nested query from our UIs. When a client requests an `InventoryLot` and its associated `Warehouse`, the Gateway fetches the `scm.InventoryLot`, notices the `WarehouseID` pointer, mathematically fans out the query to the isolated `eam.Service`, and stitches the object graph back together dynamically using `gqlgen` resolvers.
