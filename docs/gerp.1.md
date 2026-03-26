GERP(1)                                    GERP Matrix Man Page                                    GERP(1)

NAME
       gerp - Google ERP Master Control Plane

SYNOPSIS
       make [COMMAND]

DESCRIPTION
       The GERP control plane utilizes Makefile targets to orchestrate a distributed, 7-domain ERP 
       matrix structurally bound by Temporal Sagas and isolated Spanner data layers.

ARCHITECTURE
       The physical runtimes consist of:
       - Cloud Spanner Emulator (Local instance binding all 7 domains natively).
       - Temporal Server (Local instance managing `GERP_GLOBAL_QUEUE`).
       - Go GraphQL Gateway (Port 8080: Binds the BFF execution).
       - Go Temporal Worker (Background: Listens for Saga dispatches).

COMMANDS
       make up
           Boot the local Docker infrastructure encompassing Spanner, Temporal, and associated matrix layers.

       make init-db
           Create the Spanner emulator instance and execute the physical DDL files for all 7 downstream domains, permanently interleaving the schema state.

       make run-gateway
           Boot the GraphQL Backend-For-Frontend (BFF). Exposes the GraphQL Playground and routing execution path on http://localhost:8080.

       make run-worker
           Boot the Temporal Saga orchestrator. Listens dynamically to the Temporal Server queue and executes multi-domain asynchronous mutations.

       make generate
           Regenerate the physical Go bindings for the unified GraphQL API.

SEE ALSO
       docker-compose(1), temporal(1), gcloud-spanner(1)
