package legal

import (
	"time"

	"github.com/google/uuid"
)

// Contract holds the legal execution status of B2B or B2C agreements.
type Contract struct {
	ID             uuid.UUID `json:"id" spanner:"ID"`
	CounterpartyID uuid.UUID `json:"counterparty_id" spanner:"CounterpartyID"` // Golden Thread: Links to MDM
	Type           string    `json:"type" spanner:"Type"`                      // e.g., "NDA", "MSA", "SLA"
	ValidFrom      time.Time `json:"valid_from" spanner:"ValidFrom"`
	ValidTo        time.Time `json:"valid_to" spanner:"ValidTo"`
	CreatedAt      time.Time `json:"created_at" spanner:"CreatedAt"`
	UpdatedAt      time.Time `json:"updated_at" spanner:"UpdatedAt"`
}

// ComplianceAudit acts as the append-only SOX/SOC2 ledger for the absolute truth of system changes.
type ComplianceAudit struct {
	ID             uuid.UUID `json:"id" spanner:"ID"`
	TargetRecordID uuid.UUID `json:"target_record_id" spanner:"TargetRecordID"` // Golden Thread: Links to ANY modified row across all domains
	ActorID        uuid.UUID `json:"actor_id" spanner:"ActorID"`                // Golden Thread: HCM Employee who made the change
	Action         string    `json:"action" spanner:"Action"`                   // e.g., "UPDATE_SALARY", "VOID_INVOICE"
	AuditTimestamp time.Time `json:"audit_timestamp" spanner:"AuditTimestamp"`  // Exact execution boundary
}
