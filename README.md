# 🌐 GERP (Google-Enterprise Resource Planning)

![Go Version](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go)
![Cloud Spanner](https://img.shields.io/badge/Google_Cloud-Spanner-4285F4?style=for-the-badge&logo=googlecloud)
![Temporal](https://img.shields.io/badge/Temporal-Sagas-161616?style=for-the-badge&logo=temporal)
![GraphQL](https://img.shields.io/badge/GraphQL-BFF-E10098?style=for-the-badge&logo=graphql)
![MCP](https://img.shields.io/badge/MCP-AI_Ready-FF7C00?style=for-the-badge)

**GERP** is a FAANG-grade, multi-domain Enterprise Resource Planning (ERP) matrix. GERP demonstrates how to scale massive distributed data systems using absolute Domain-Driven Design (DDD), cross-domain Temporal Sagas, and zero SQL foreign keys.

This is not just an ERP. It is a self-aware corporate engine exposed natively to AI agents via the Model Context Protocol (MCP).

---

## 🏛️ Core Architectural Pillars

### 1. The Golden Thread (Zero SQL Foreign Keys)
To achieve infinite horizontal scale, GERP completely abandons database-layer foreign keys between domains. A `SalesOrder` in Revenue does not use an SQL `JOIN` to find a `Customer` in Master Data. Instead, GERP employs the **Golden Thread**: strict `uuid.UUID` pointers managed in application space. Each domain owns its Spanner tables exclusively. 

### 2. Distributed Temporal Sagas
Because domains are isolated, GERP cannot rely on single-database ACID transactions. Instead, it uses **Temporal Workflows** as its nervous system. If an order allocates physical inventory in the SCM domain but fails to lock the ledger in the Finance domain, Temporal automatically executes a mathematical **Compensating Rollback** (`ReverseInventoryActivity`) to guarantee eventual consistency and eliminate phantom locks.

### 3. The GraphQL Backend-For-Frontend (BFF)
Clients never see the distributed complexity. The Go-based GraphQL Gateway (`cmd/gateway`) receives a unified query and fans out requests across the isolated micro-domains in memory, resolving the Golden Thread UUIDs instantly into beautiful, deeply nested JSON graphs.

### 4. The MCP Brain Interface
GERP is designed to be operated by AI. The built-in Model Context Protocol server (`cmd/mcp`) exposes the Spanner audit logs, system status, and Temporal Saga triggers over standard JSON-RPC STDIO. Point Claude Desktop or Cursor at this repository, and the AI can run the company.

---

## 🏗️ The 8 Tier-1 Domains

GERP separates its global state into perfectly isolated execution environments:
* 💰 **Finance (`internal/finance`):** The immutable double-entry ledger.
* 👥 **Human Capital (`internal/hcm`):** The employee and payroll engine.
* 📦 **Supply Chain (`internal/scm`):** Physical inventory and SKU tracking.
* 🏭 **Enterprise Asset (`internal/eam`):** Infrastructure and warehouse management.
* ⚖️ **Legal (`internal/legal`):** The append-only SOC2/SOX compliance audit log.
* 📈 **Revenue (`internal/revenue`):** Top-line sales and customer relationship mapping.
* 🎓 **Learning (`internal/lms`):** Educational compliance and safety certifications.
* 🌐 **Master Data (`internal/mdm`):** The Universal Translator connecting localized IDs to a single "Golden Record".

---

## 🚀 Getting Started (Local Matrix)

GERP includes a massive infrastructure control plane designed for local Docker execution.

**1. Boot the Infrastructure (Spanner, Temporal, Redis)**
```bash
make up
make init-db
```

**2. Inject the Genesis State (Seed the Matrix)**
```bash
go run ./cmd/seed/main.go
```

**3. Start the Execution Engines**
```bash
# Terminal 1: Boot the Temporal Orchestrator
make run-worker

# Terminal 2: Boot the GraphQL Gateway
make run-gateway
```

**4. Command the Matrix**
You can fire cross-domain Sagas using the native CLI Operator:
```bash
go build -o gerp ./cmd/gerp
./gerp orders create
./gerp audit view 99999999-9999-9999-9999-999999999999
```

---

## 🧠 Hooking up your AI (MCP Server)

To allow your AI IDE to read GERP's physical state and trigger workflows, add the following to your `.cursor/mcp.json` or Claude Desktop configuration:

```json
{
  "mcpServers": {
    "gerp-matrix": {
      "command": "go",
      "args": ["run", "./cmd/mcp/main.go"],
      "env": {
        "GERP_GRAPHQL_ENDPOINT": "http://localhost:8080/query",
        "GERP_TEMPORAL_HOST": "localhost:7233",
        "GERP_SPANNER_DB": "projects/gerp-local-dev/instances/gerp-instance/databases/gerp-db"
      }
    }
  }
}
```

---
*Built with precision by the Architect-in-the-Loop Swarm.*
