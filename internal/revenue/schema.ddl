-- internal/revenue/schema.ddl
-- Cloud Spanner DDL (Strict DDD Isolation)

CREATE TABLE Customers (
    ID STRING(36) NOT NULL,
    Name STRING(255) NOT NULL,
    MasterDataID STRING(36) NOT NULL, -- Soft Link to MDM 
    AccountManagerID STRING(36),      -- Soft Link to HCM Employee
    CreditLimit INT64 NOT NULL,
    CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    UpdatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY (ID);

CREATE TABLE SalesOrders (
    ID STRING(36) NOT NULL,
    CustomerID STRING(36) NOT NULL,
    TotalValue INT64 NOT NULL,
    Status STRING(50) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    UpdatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY (ID);
