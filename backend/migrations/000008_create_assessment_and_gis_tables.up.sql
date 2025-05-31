-- 査定・GIS関連テーブル作成

-- 更新日時を自動更新するトリガー関数（前のマイグレーションで定義されている可能性があるが、念のため再定義）
CREATE OR REPLACE FUNCTION update_master_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 査定テーブル
DROP TABLE IF EXISTS assessments CASCADE;
CREATE TABLE IF NOT EXISTS assessments
(
    id                  SERIAL PRIMARY KEY,
    disaster_id         UUID NOT NULL REFERENCES disasters(id) ON DELETE CASCADE,
    assessor_id         INT NOT NULL REFERENCES users(id),
    assessment_date     DATE NOT NULL,
    assessment_type     VARCHAR(50) NOT NULL,
    status              VARCHAR(30) NOT NULL DEFAULT '進行中',
    CHECK (status IN ('準備中', '進行中', '完了', '差戻し', '承認済')),
    assessment_method   VARCHAR(50) NOT NULL,
    CHECK (assessment_method IN ('現地査定', 'リモート査定', '書類査定', '緊急査定')),
    assessment_summary  TEXT,
    damage_amount       DECIMAL(15, 2),
    approved_amount     DECIMAL(15, 2),
    approval_date       TIMESTAMP WITH TIME ZONE,
    approved_by         INT REFERENCES users(id),
    notes               TEXT,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_assessments_disaster_id ON assessments (disaster_id);
CREATE INDEX IF NOT EXISTS idx_assessments_assessor_id ON assessments (assessor_id);
CREATE INDEX IF NOT EXISTS idx_assessments_assessment_date ON assessments (assessment_date);
CREATE INDEX IF NOT EXISTS idx_assessments_status ON assessments (status);

COMMENT ON TABLE assessments IS '査定テーブル - 災害被害の査定情報を管理';
COMMENT ON COLUMN assessments.id IS '査定ID - 主キー';
COMMENT ON COLUMN assessments.disaster_id IS '災害ID - 査定対象の災害ID';
COMMENT ON COLUMN assessments.assessor_id IS '査定者ID - 査定を行ったユーザーのID';
COMMENT ON COLUMN assessments.assessment_date IS '査定日 - 査定が行われた日付';
COMMENT ON COLUMN assessments.assessment_type IS '査定種別 - 現地査定、リモート査定など';
COMMENT ON COLUMN assessments.status IS '状態 - 査定の進行状況';
COMMENT ON COLUMN assessments.assessment_method IS '査定方法 - 査定の実施方法';
COMMENT ON COLUMN assessments.assessment_summary IS '査定概要 - 査定結果の概要';
COMMENT ON COLUMN assessments.damage_amount IS '被害金額 - 査定された被害金額';
COMMENT ON COLUMN assessments.approved_amount IS '承認金額 - 承認された支援金額';
COMMENT ON COLUMN assessments.approval_date IS '承認日時 - 査定が承認された日時';
COMMENT ON COLUMN assessments.approved_by IS '承認者ID - 査定を承認したユーザーのID';
COMMENT ON COLUMN assessments.notes IS '備考 - 査定に関する備考やメモ';
COMMENT ON COLUMN assessments.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN assessments.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN assessments.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 査定項目テーブル
DROP TABLE IF EXISTS assessment_items CASCADE;
CREATE TABLE IF NOT EXISTS assessment_items
(
    id                  SERIAL PRIMARY KEY,
    assessment_id       INT NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
    item_name           VARCHAR(100) NOT NULL,
    facility_type_id    INT, -- 外部キー制約は後で追加（facility_typesテーブルへの参照）
    damage_level_id     INT, -- 外部キー制約は後で追加（damage_levelsテーブルへの参照）
    damage_description  TEXT NOT NULL,
    damage_amount       DECIMAL(15, 2) NOT NULL,
    approved_amount     DECIMAL(15, 2),
    location_latitude   DECIMAL(10, 8),
    location_longitude  DECIMAL(11, 8),
    notes               TEXT,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_assessment_items_assessment_id ON assessment_items (assessment_id);
CREATE INDEX IF NOT EXISTS idx_assessment_items_facility_type_id ON assessment_items (facility_type_id);
CREATE INDEX IF NOT EXISTS idx_assessment_items_damage_level_id ON assessment_items (damage_level_id);

COMMENT ON TABLE assessment_items IS '査定項目テーブル - 査定の詳細項目を管理';
COMMENT ON COLUMN assessment_items.id IS '査定項目ID - 主キー';
COMMENT ON COLUMN assessment_items.assessment_id IS '査定ID - 関連する査定のID';
COMMENT ON COLUMN assessment_items.item_name IS '項目名 - 査定項目の名称';
COMMENT ON COLUMN assessment_items.facility_type_id IS '施設種別ID - 被害を受けた施設の種別';
COMMENT ON COLUMN assessment_items.damage_level_id IS '被害程度ID - 被害の程度';
COMMENT ON COLUMN assessment_items.damage_description IS '被害説明 - 被害状況の詳細説明';
COMMENT ON COLUMN assessment_items.damage_amount IS '被害金額 - 項目ごとの被害金額';
COMMENT ON COLUMN assessment_items.approved_amount IS '承認金額 - 項目ごとの承認金額';
COMMENT ON COLUMN assessment_items.location_latitude IS '位置（緯度） - 被害箇所の緯度';
COMMENT ON COLUMN assessment_items.location_longitude IS '位置（経度） - 被害箇所の経度';
COMMENT ON COLUMN assessment_items.notes IS '備考 - 査定項目に関する備考やメモ';
COMMENT ON COLUMN assessment_items.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN assessment_items.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN assessment_items.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 査定コメントテーブル
DROP TABLE IF EXISTS assessment_comments CASCADE;
CREATE TABLE IF NOT EXISTS assessment_comments
(
    id                  SERIAL PRIMARY KEY,
    assessment_id       INT NOT NULL REFERENCES assessments(id) ON DELETE CASCADE,
    user_id             INT NOT NULL REFERENCES users(id),
    comment_text        TEXT NOT NULL,
    comment_time        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    parent_comment_id   INT REFERENCES assessment_comments(id),
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_assessment_comments_assessment_id ON assessment_comments (assessment_id);
CREATE INDEX IF NOT EXISTS idx_assessment_comments_user_id ON assessment_comments (user_id);
CREATE INDEX IF NOT EXISTS idx_assessment_comments_parent_comment_id ON assessment_comments (parent_comment_id);

COMMENT ON TABLE assessment_comments IS '査定コメントテーブル - 査定に関するコメントを管理';
COMMENT ON COLUMN assessment_comments.id IS 'コメントID - 主キー';
COMMENT ON COLUMN assessment_comments.assessment_id IS '査定ID - 関連する査定のID';
COMMENT ON COLUMN assessment_comments.user_id IS 'ユーザーID - コメントを投稿したユーザーのID';
COMMENT ON COLUMN assessment_comments.comment_text IS 'コメント本文 - コメントの内容';
COMMENT ON COLUMN assessment_comments.comment_time IS 'コメント時間 - コメントが投稿された時間';
COMMENT ON COLUMN assessment_comments.parent_comment_id IS '親コメントID - 返信先のコメントID（スレッド構造用）';
COMMENT ON COLUMN assessment_comments.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN assessment_comments.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN assessment_comments.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- GIS情報テーブル
DROP TABLE IF EXISTS gis_data CASCADE;
CREATE TABLE IF NOT EXISTS gis_data
(
    id                  SERIAL PRIMARY KEY,
    disaster_id         UUID NOT NULL REFERENCES disasters(id) ON DELETE CASCADE,
    data_type           VARCHAR(50) NOT NULL,
    CHECK (data_type IN ('被害エリア', '避難経路', '施設位置', 'リソース配置', 'その他')),
    name                VARCHAR(100) NOT NULL,
    description         TEXT,
    geometry_type       VARCHAR(20) NOT NULL,
    CHECK (geometry_type IN ('POINT', 'LINESTRING', 'POLYGON', 'MULTIPOINT', 'MULTILINESTRING', 'MULTIPOLYGON')),
    geometry_data       TEXT NOT NULL, -- GeoJSON形式のデータを格納
    properties          JSONB, -- 追加のプロパティをJSON形式で格納
    created_by          INT NOT NULL REFERENCES users(id),
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_gis_data_disaster_id ON gis_data (disaster_id);
CREATE INDEX IF NOT EXISTS idx_gis_data_data_type ON gis_data (data_type);
CREATE INDEX IF NOT EXISTS idx_gis_data_created_by ON gis_data (created_by);

COMMENT ON TABLE gis_data IS 'GIS情報テーブル - 地理空間情報を管理';
COMMENT ON COLUMN gis_data.id IS 'GISデータID - 主キー';
COMMENT ON COLUMN gis_data.disaster_id IS '災害ID - 関連する災害のID';
COMMENT ON COLUMN gis_data.data_type IS 'データ種別 - GISデータの種類';
COMMENT ON COLUMN gis_data.name IS '名称 - GISデータの名称';
COMMENT ON COLUMN gis_data.description IS '説明 - GISデータの説明';
COMMENT ON COLUMN gis_data.geometry_type IS 'ジオメトリ種別 - 地理データの形状タイプ';
COMMENT ON COLUMN gis_data.geometry_data IS 'ジオメトリデータ - GeoJSON形式の地理データ';
COMMENT ON COLUMN gis_data.properties IS 'プロパティ - 追加のプロパティ情報（JSON形式）';
COMMENT ON COLUMN gis_data.created_by IS '作成者ID - データを作成したユーザーのID';
COMMENT ON COLUMN gis_data.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN gis_data.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN gis_data.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 通知テーブル
DROP TABLE IF EXISTS notifications CASCADE;
CREATE TABLE IF NOT EXISTS notifications
(
    id                  SERIAL PRIMARY KEY,
    user_id             INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title               VARCHAR(200) NOT NULL,
    message             TEXT NOT NULL,
    notification_type   VARCHAR(50) NOT NULL,
    CHECK (notification_type IN ('システム', '災害情報', '査定', '申請', 'リマインダー', 'その他')),
    related_entity_type VARCHAR(50),
    related_entity_id   VARCHAR(100),
    is_read             BOOLEAN NOT NULL DEFAULT FALSE,
    read_at             TIMESTAMP WITH TIME ZONE,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications (user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_is_read ON notifications (is_read);
CREATE INDEX IF NOT EXISTS idx_notifications_notification_type ON notifications (notification_type);
CREATE INDEX IF NOT EXISTS idx_notifications_created_at ON notifications (created_at);

COMMENT ON TABLE notifications IS '通知テーブル - ユーザーへの通知を管理';
COMMENT ON COLUMN notifications.id IS '通知ID - 主キー';
COMMENT ON COLUMN notifications.user_id IS 'ユーザーID - 通知の宛先ユーザーID';
COMMENT ON COLUMN notifications.title IS 'タイトル - 通知のタイトル';
COMMENT ON COLUMN notifications.message IS 'メッセージ - 通知の本文';
COMMENT ON COLUMN notifications.notification_type IS '通知種別 - 通知の種類';
COMMENT ON COLUMN notifications.related_entity_type IS '関連エンティティ種別 - 通知に関連するエンティティの種類';
COMMENT ON COLUMN notifications.related_entity_id IS '関連エンティティID - 通知に関連するエンティティのID';
COMMENT ON COLUMN notifications.is_read IS '既読フラグ - 通知が既読かどうか';
COMMENT ON COLUMN notifications.read_at IS '既読日時 - 通知が既読になった日時';
COMMENT ON COLUMN notifications.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN notifications.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN notifications.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 更新日時を自動更新するトリガー
CREATE TRIGGER update_assessments_updated_at
    BEFORE UPDATE ON assessments
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_assessment_items_updated_at
    BEFORE UPDATE ON assessment_items
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_assessment_comments_updated_at
    BEFORE UPDATE ON assessment_comments
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_gis_data_updated_at
    BEFORE UPDATE ON gis_data
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_notifications_updated_at
    BEFORE UPDATE ON notifications
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_notifications_updated_at
    BEFORE UPDATE ON notifications
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

-- 通知テーブルへのダミーデータ挿入
INSERT INTO notifications (user_id, title, message, notification_type, related_entity_type, related_entity_id, is_read, read_at)
VALUES
-- システム通知
(1, 'システムメンテナンスのお知らせ', '明日2025年6月1日午前2時から4時までシステムメンテナンスを実施します。この間サービスはご利用いただけません。', 'システム', NULL, NULL, true, '2025-05-31 09:15:00+09'),
(2, 'システムアップデート完了', 'システムのアップデートが完了しました。新機能の詳細はお知らせページをご確認ください。', 'システム', NULL, NULL, false, NULL),

-- 災害情報通知
(3, '新規災害情報登録', '新たな災害情報「2025年青森県豪雨」が登録されました。', '災害情報', '災害', '1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p', true, '2025-05-30 14:25:10+09'),
(4, '災害ステータス更新', '災害「2025年関東地方大雨被害」のステータスが「対応中」に更新されました。', '災害情報', '災害', '2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7q', false, NULL),
(5, '被害報告が提出されました', '青森県りんご被害の報告が提出されました。確認をお願いします。', '災害情報', '災害', '3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8r', false, NULL),

-- 査定関連通知
(6, '査定が割り当てられました', '2025年東京都風害の査定担当者として割り当てられました。詳細を確認してください。', '査定', '査定', '1', true, '2025-05-28 10:45:22+09'),
(7, '査定結果が提出されました', '査定ID: 2の結果が提出されました。承認作業をお願いします。', '査定', '査定', '2', false, NULL),
(8, '査定コメントが追加されました', '担当査定に新しいコメントが追加されました。', '査定', 'コメント', '3', false, NULL),

-- 申請関連通知
(9, '支援申請が提出されました', '新しい支援申請が提出されました。審査をお願いします。', '申請', '支援申請', '1', true, '2025-05-25 16:30:00+09'),
(10, '申請ステータス更新', 'あなたの支援申請（ID: 2）のステータスが「審査中」に更新されました。', '申請', '支援申請', '2', false, NULL),
(11, '申請が承認されました', 'あなたの支援申請（ID: 3）が承認されました。詳細は申請履歴をご確認ください。', '申請', '支援申請', '3', true, '2025-05-23 09:10:45+09'),

-- リマインダー通知
(12, '査定期限のリマインダー', '担当の査定（ID: 4）の提出期限が3日後に迫っています。', 'リマインダー', '査定', '4', false, NULL),
(13, 'ドキュメント提出リマインダー', '被害状況の追加ドキュメントの提出期限は明日までです。', 'リマインダー', 'ドキュメント', NULL, true, '2025-05-20 13:40:30+09'),

-- その他の通知
(14, '研修会のお知らせ', '次回の査定員研修会は2025年6月15日に開催されます。参加希望者は登録をお願いします。', 'その他', 'イベント', 'training-2025-06', false, NULL),
(15, '新しい施設登録マニュアルの公開', '施設登録に関する新しいマニュアルが公開されました。マニュアルページからご確認ください。', 'その他', 'マニュアル', 'facility-registration-v2', true, '2025-05-18 11:20:15+09'),

-- 複数ユーザーへの同一通知
(16, '緊急：システムセキュリティアップデート', 'セキュリティ上の理由により、全ユーザーのパスワード変更をお願いします。', 'システム', NULL, NULL, false, NULL),
(17, '緊急：システムセキュリティアップデート', 'セキュリティ上の理由により、全ユーザーのパスワード変更をお願いします。', 'システム', NULL, NULL, true, '2025-05-15 15:30:00+09'),
(18, '緊急：システムセキュリティアップデート', 'セキュリティ上の理由により、全ユーザーのパスワード変更をお願いします。', 'システム', NULL, NULL, false, NULL),

-- 最近の通知
(19, '災害報告書テンプレート更新', '災害報告書の新しいテンプレートが公開されました。今後の報告にはこちらをご使用ください。', 'その他', 'テンプレート', 'disaster-report-v3', false, NULL),
(20, '祝：システム利用者1000人達成', '本システムの利用者が1000人を達成しました。ご利用ありがとうございます。', 'システム', NULL, NULL, false, NULL);
