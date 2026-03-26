-- internal/legal/schema.ddl
-- Cloud Spanner DDL (Strict DDD Isolation)

CREATE TABLE Contracts (
    ID STRING(36) NOT NULL,
    CounterpartyID STRING(36) NOT NULL, -- Soft Link to MDM
    Type STRING(50) NOT NULL,
    ValidFrom TIMESTAMP NOT NULL,
    ValidTo TIMESTAMP NOT NULL,
    CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
    UpdatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY (ID);

CREATE TABLE ComplianceAudits (
    ID STRING(36) NOT NULL,
    TargetRecordID STRING(36) NOT NULL, -- Soft Link to ANY domain's UUID
    ActorID STRING(36) NOT NULL,        -- Soft Link to HCM Employee
    Action STRING(100) NOT NULL,
    AuditTimestamp TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)
) PRIMARY KEY (ID);

-- Rapid audit retrieval for SOC2 compliance verification on any specific record
CREATE INDEX AuditsByTarget ON ComplianceAudits(TargetRecordID);
