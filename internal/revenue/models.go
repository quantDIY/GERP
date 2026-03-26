package revenue

import (
	"time"

	"github.com/google/uuid"
)

// Customer acts as the sales-domain representation of an entity, holding revenue-specific data.
type Customer struct {
	ID               uuid.UUID `json:"id" spanner:"ID"`
	Name             string    `json:"name" spanner:"Name"`
	MasterDataID     uuid.UUID `json:"master_data_id" spanner:"MasterDataID"`         // Golden Thread: Maps to MDM's universal canonical truth
	AccountManagerID uuid.UUID `json:"account_manager_id" spanner:"AccountManagerID"` // Golden Thread: Maps to HCM's internal Employee graph
	CreditLimit      int64     `json:"credit_limit" spanner:"CreditLimit"`            // Minor units (cents)
	CreatedAt        time.Time `json:"created_at" spanner:"CreatedAt"`
	UpdatedAt        time.Time `json:"updated_at" spanner:"UpdatedAt"`
}

// SalesOrder acts as the origin boundary for cross-domain operational cascades.
type SalesOrder struct {
	ID         uuid.UUID `json:"id" spanner:"ID"`
	CustomerID uuid.UUID `json:"customer_id" spanner:"CustomerID"` // Parent Sales mapping
	TotalValue int64     `json:"total_value" spanner:"TotalValue"` // Minor units (cents)
	Status     string    `json:"status" spanner:"Status"`          // e.g., "DRAFT", "PENDING_CREDIT_APPROVAL", "FULFILLED"
	CreatedAt  time.Time `json:"created_at" spanner:"CreatedAt"`
	UpdatedAt  time.Time `json:"updated_at" spanner:"UpdatedAt"`
}
