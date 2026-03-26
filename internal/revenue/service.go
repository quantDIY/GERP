package revenue

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

// Service defines the boundary for customer relationship mapping and sales origination.
type Service interface {
	GetCustomer(ctx context.Context, id uuid.UUID) (*Customer, error)
	GetSalesOrder(ctx context.Context, id uuid.UUID) (*SalesOrder, error)
	InsertSalesOrder(ctx context.Context, order *SalesOrder) error
}

type revenueService struct {
	client *spanner.Client
}

// NewService provisions the Revenue & Sales service with the dedicated Spanner client.
func NewService(client *spanner.Client) Service {
	return &revenueService{client: client}
}

// GetCustomer retrieves the sales-level representation of a buyer or organization.
func (s *revenueService) GetCustomer(ctx context.Context, id uuid.UUID) (*Customer, error) {
	row, err := s.client.Single().ReadRow(ctx, "Customers", spanner.Key{id.String()}, []string{
		"ID", "Name", "MasterDataID", "AccountManagerID", "CreditLimit", "CreatedAt", "UpdatedAt",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read customer: %w", err)
	}

	var customer Customer
	if err := row.ToStruct(&customer); err != nil {
		return nil, fmt.Errorf("failed to decode customer payload: %w", err)
	}

	return &customer, nil
}

// GetSalesOrder unpacks the origin transaction of the fulfillment saga.
func (s *revenueService) GetSalesOrder(ctx context.Context, id uuid.UUID) (*SalesOrder, error) {
	row, err := s.client.Single().ReadRow(ctx, "SalesOrders", spanner.Key{id.String()}, []string{
		"ID", "CustomerID", "TotalValue", "Status", "CreatedAt", "UpdatedAt",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read sales order: %w", err)
	}

	var order SalesOrder
	if err := row.ToStruct(&order); err != nil {
		return nil, fmt.Errorf("failed to decode sales order struct: %w", err)
	}

	return &order, nil
}

// InsertSalesOrder durably writes a new sales origin event to Cloud Spanner,
// placing it in state for the Temporal orchestration layer to pick up.
func (s *revenueService) InsertSalesOrder(ctx context.Context, order *SalesOrder) error {
	mut, err := spanner.InsertStruct("SalesOrders", order)
	if err != nil {
		return err
	}
	
	_, err = s.client.Apply(ctx, []*spanner.Mutation{mut})
	if err != nil {
		return fmt.Errorf("failed to persist sales order ledger: %w", err)
	}
	
	return nil
}
