package legal

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

// Service defines the boundary for corporate governance, contracts, and immutable tracking.
type Service interface {
	GetContract(ctx context.Context, id uuid.UUID) (*Contract, error)
	LogAudit(ctx context.Context, audit *ComplianceAudit) error
}

type legalService struct {
	client *spanner.Client
}

// NewService provisions the Legal & Compliance service with the dedicated Spanner client.
func NewService(client *spanner.Client) Service {
	return &legalService{client: client}
}

// GetContract cleanly fetches a single agreement from the database without requiring cross-table joints.
func (s *legalService) GetContract(ctx context.Context, id uuid.UUID) (*Contract, error) {
	row, err := s.client.Single().ReadRow(ctx, "Contracts", spanner.Key{id.String()}, []string{
		"ID", "CounterpartyID", "Type", "ValidFrom", "ValidTo", "CreatedAt", "UpdatedAt",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read contract: %w", err)
	}

	var contract Contract
	if err := row.ToStruct(&contract); err != nil {
		return nil, fmt.Errorf("failed to decode contract payload: %w", err)
	}

	return &contract, nil
}

// LogAudit enforces the append-only SOX/SOC2 isolation log. 
// It guarantees that cross-domain mutations are permanently recorded.
func (s *legalService) LogAudit(ctx context.Context, audit *ComplianceAudit) error {
	mut, err := spanner.InsertStruct("ComplianceAudits", audit)
	if err != nil {
		return err
	}
	
	_, err = s.client.Apply(ctx, []*spanner.Mutation{mut})
	if err != nil {
		return fmt.Errorf("failed to persist immutable audit log: %w", err)
	}
	
	return nil
}
