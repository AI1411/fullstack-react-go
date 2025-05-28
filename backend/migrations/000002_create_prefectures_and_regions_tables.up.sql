-- 都道府県マスタテーブル
CREATE TABLE IF NOT EXISTS prefectures
(
    id         UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    code       VARCHAR(2)  NOT NULL UNIQUE,
    name       VARCHAR(10) NOT NULL UNIQUE,
    created_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_prefectures_code ON prefectures (code);

-- 地域マスタテーブル
CREATE TABLE IF NOT EXISTS regions
(
    id            UUID PRIMARY KEY     DEFAULT gen_random_uuid(),
    prefecture_id UUID        NOT NULL REFERENCES prefectures (id),
    name          VARCHAR(50) NOT NULL,
    code          VARCHAR(10) NOT NULL UNIQUE,
    created_at    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP,
    UNIQUE (prefecture_id, name)
);

CREATE INDEX IF NOT EXISTS idx_regions_prefecture_id ON regions (prefecture_id);
CREATE INDEX IF NOT EXISTS idx_regions_code ON regions (code);
