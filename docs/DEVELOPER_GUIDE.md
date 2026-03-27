# GERP Developer Onboarding Guide

Welcome to GERP (Google ERP). This system is a FAANG-grade, multi-domain Enterprise Resource Planning matrix structurally isolated by Domain-Driven Design constraints.

## 1. The Golden Thread (UUID Stitching)
To guarantee distributed, isolated micro-database performance at scale, GERP completely abandons standard SQL foreign keys. No domain is permitted to link its physical state via SQL layer to another domain. 

Instead, GERP employs the **Golden Thread**: strict UUID pointers managed entirely in application space. A `SalesOrder` in `internal/revenue` physically stores a `CustomerID` (UUID), but it never explicitly joins that to the `mdm.GlobalEntities` table at the Spanner level. The graph stitching occurs strictly in memory at the BFF level.

## 2. The 8 Tier-1 Domains
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

## 5. The CLI Control Plane (`cmd/gerp`)
The `gerp` binary is the native operator terminal giving system administrators direct, low-latency control and query capabilities into the core Spanner domains and Temporal execution queues using `spf13/cobra` (routing) and `spf13/viper` (configuration).

- **Config Bound:** The binary dynamically binds to `.gerp.yaml` in the user's home directory across different shell environments or relies on `GERP_` prefixed environment variables.
- By bypassing the BFF using commands like `gerp audit view`, the CLI securely isolates direct native Spanner telemetry for compliance operations.

## 6. Infrastructure & Deployment (`build/`, `deploy/`)
GERP utilizes mathematically reproducible cloud infrastructure and secure container limits:
- **Docker Packaging (`build/docker/`):** Both `cmd/gateway` and `cmd/worker` exist natively as multi-stage Alpine images optimized for scalable deployment pods.
- **Terraform (`deploy/terraform/`):** GCP-specific HashiCorp constraints physicalizing the raw Cloud Spanner limits (`google_spanner_instance`, `google_spanner_database`) absent rigid DDL schema mappings.
- **Ansible (`deploy/ansible/`):** Orchestrates the deployment configuration variables on the executing hardware targeting Docker container lifecycles.

## 7. The AI Brain Interface (MCP)
GERP exposes FAANG-grade capabilities autonomously to outside LLMs via standard JSON-RPC STDIO built off the Model Context Protocol.

The `cmd/mcp` Server binds to AI agents running on Cursor IDE, Claude Desktop, or centralized orchestration engines via `.cursor/mcp.json`. Agents inherently use the physical tools (`gerp_status`, `gerp_create_order`, `gerp_audit_view`) to mechanically command the Spanner databases and trigger automated Temporal subroutines strictly adhering to the `internal/cli` environmental scope bindings.

## 8. COAMS (Content Operating and Management System)

COAMS represents the eighth primary domain of GERP, specifically architected as a Markdown-Native, AI-First content platform. Traditional headless CMS patterns relying on complex graphical JSON/ASTs have been discarded in favor of raw Markdown, natively optimizing it for AI integration.

### Core Architecture & The "Ignorant Engine"
- **Isolated AlloyDB Sharding:** Unlike 
the core Spanner nodes, COAMS stores vector embeddings via AlloyDB using PostgreSQL's `pgvector` paired with Declarative Partitioning (`PARTITION BY LIST`). Data is physically separated by `channel_id` (e.g., engineering, hr) ensuring Zero-Leak Architecture.
- **Microsecond Graph Linking:** The `doc:uuid` pointers are validated mathematically before writing by the **Agent-Index**. The system outright rejects writes if an internal link points to a non-existent or unauthorized document.

### System Integration & The Publish Saga
The COAMS ingestion logic lives entirely within the temporal queues (`internal/pipeline/coams_publish_saga.go`). Executing `gerp coams sync ./docs` triggers:
1. Pure Go Goldmark AST markdown parsing.
2. Link Integrity Verification.
3. Vertex AI Vector Embeddings generation.
4. Partitioned AlloyDB storage.
5. Dynamic broadcast update of the GraphQL schema ensuring immediate downstream synchronization.

### Bootstrapping the QuanuX Knowledge Vector
By design, COAMS is self-documenting. To understand the operational footprint of the system, any AI agent or Operator can directly consult the embedded `SKILL.md` file:
[Consult the COAMS SKILL.md Here](file:///Users/Duncan/GERP/internal/coams/docs/SKILL.md)

Additionally, COAMS autonomously generates manual pages for CLI execution utilizing the `gerp coams gen-man` control hook. This fundamentally ensures that agents never suffer from blank-page syndrome when learning how to navigate the repository layer.
