CREATE TABLE IF NOT EXISTS disasters
(
    id                      UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    disaster_code           VARCHAR(20)  NOT NULL UNIQUE,
    name                    VARCHAR(100) NOT NULL,
    region_id               UUID         NOT NULL REFERENCES regions (id),
    occurred_at             TIMESTAMP    NOT NULL,
    summary                 TEXT         NOT NULL,
    disaster_type           VARCHAR(30)  NOT NULL,
    status                  VARCHAR(20)  NOT NULL,
    impact_level            VARCHAR(20)  NOT NULL,
    affected_area_size      DECIMAL(10, 2),
    estimated_damage_amount DECIMAL(15, 2),
    created_at              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_disasters_region_id ON disasters (region_id);
CREATE INDEX IF NOT EXISTS idx_disasters_disaster_type ON disasters (disaster_type);
CREATE INDEX IF NOT EXISTS idx_disasters_status ON disasters (status);
CREATE INDEX IF NOT EXISTS idx_disasters_occurred_at ON disasters (occurred_at);
