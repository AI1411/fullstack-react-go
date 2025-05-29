-- 既存のENUM型を削除して再作成
DROP TYPE IF EXISTS disaster_status CASCADE;
CREATE TYPE disaster_status AS ENUM ('pending', 'under_review', 'in_progress', 'completed');

-- 既存テーブルを削除して再作成
DROP TABLE IF EXISTS disasters CASCADE;
CREATE TABLE IF NOT EXISTS disasters
(
    id                      UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    disaster_code           VARCHAR(20)              NOT NULL UNIQUE,
    name                    VARCHAR(100)             NOT NULL,
    prefecture_id           INT                      NOT NULL REFERENCES prefectures (id),
    occurred_at             TIMESTAMP WITH TIME ZONE NOT NULL,
    summary                 TEXT                     NOT NULL,
    disaster_type           VARCHAR(30)              NOT NULL,
    status                  disaster_status          NOT NULL DEFAULT 'pending',
    impact_level            VARCHAR(20)              NOT NULL,
    affected_area_size      DECIMAL(10, 2),
    estimated_damage_amount DECIMAL(15, 2),
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_disasters_prefecture_id ON disasters (prefecture_id);
CREATE INDEX idx_disasters_disaster_type ON disasters (disaster_type);
CREATE INDEX idx_disasters_status ON disasters (status);
CREATE INDEX idx_disasters_occurred_at ON disasters (occurred_at);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE disasters IS '農業災害情報管理テーブル - 各種災害の詳細情報を格納';
COMMENT ON COLUMN disasters.id IS '災害ID - 主キー';
COMMENT ON COLUMN disasters.disaster_code IS '災害コード - 管理用の一意識別子 (例: D2024-001)';
COMMENT ON COLUMN disasters.name IS '災害名 - 災害の名称';
COMMENT ON COLUMN disasters.prefecture_id IS '都道府県ID - 災害が発生した都道府県のID';
COMMENT ON COLUMN disasters.occurred_at IS '発生日時 - 災害が発生した日時';
COMMENT ON COLUMN disasters.summary IS '被害概要 - 災害による被害の詳細説明';
COMMENT ON COLUMN disasters.disaster_type IS '災害種別 - 洪水, 地滑り, 雹害, 干ばつ, 風害, 地震, 霜害, 病害虫など';
COMMENT ON COLUMN disasters.status IS '状態 - pending(未着手), under_review(審査中), in_progress(対応中), completed(完了)のいずれか';
COMMENT ON COLUMN disasters.impact_level IS '被害レベル - 軽微, 中程度, 深刻, 甚大などの被害度合い';
COMMENT ON COLUMN disasters.affected_area_size IS '被害面積 - ヘクタール (ha) 単位での被害エリアの広さ';
COMMENT ON COLUMN disasters.estimated_damage_amount IS '被害推定金額 - 円単位での被害総額';
COMMENT ON COLUMN disasters.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN disasters.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN disasters.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';
