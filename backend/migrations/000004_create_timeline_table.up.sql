DROP TABLE IF EXISTS timelines CASCADE;
CREATE TABLE IF NOT EXISTS timelines
(
    id          SERIAL PRIMARY KEY,
    disaster_id UUID         NOT NULL REFERENCES disasters (id) ON DELETE CASCADE,
    event_name  VARCHAR(255) NOT NULL,
    event_time  TIMESTAMP    NOT NULL,
    description TEXT         NOT NULL,
    severity    VARCHAR(50),
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_timelines_disaster_id ON timelines (disaster_id);
CREATE INDEX IF NOT EXISTS idx_timelines_event_time ON timelines (event_time);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE timelines IS '災害タイムライン管理テーブル - 各災害のイベント履歴を格納';
COMMENT ON COLUMN timelines.id IS 'タイムラインID - 主キー';
COMMENT ON COLUMN timelines.disaster_id IS '災害ID - 関連する災害のID';
COMMENT ON COLUMN timelines.event_name IS 'イベント名 - 発生したイベントの名称';
COMMENT ON COLUMN timelines.event_time IS 'イベント発生日時 - イベントが発生した日時';
COMMENT ON COLUMN timelines.description IS 'イベント説明 - イベントの詳細な説明';
COMMENT ON COLUMN timelines.severity IS 'イベントの深刻度 - 低, 中, 高などの深刻度';
COMMENT ON COLUMN timelines.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN timelines.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN timelines.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';
