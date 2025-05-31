-- 施設設備マスタテーブル作成

-- 施設設備マスタテーブル
DROP TABLE IF EXISTS facility_equipment CASCADE;
CREATE TABLE IF NOT EXISTS facility_equipment
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
    facility_type_id INT NOT NULL REFERENCES facility_types(id),
    model_number    VARCHAR(100),
    manufacturer    VARCHAR(100),
    installation_date DATE,
    status          VARCHAR(50) NOT NULL DEFAULT '稼働中',
    CHECK (status IN ('稼働中', '点検中', '故障中', '廃止')),
    location_description TEXT,
    location_latitude   DECIMAL(10, 8),
    location_longitude  DECIMAL(11, 8),
    notes           TEXT,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_facility_equipment_name ON facility_equipment (name);
CREATE INDEX IF NOT EXISTS idx_facility_equipment_facility_type_id ON facility_equipment (facility_type_id);
CREATE INDEX IF NOT EXISTS idx_facility_equipment_status ON facility_equipment (status);

COMMENT ON TABLE facility_equipment IS '施設設備マスタテーブル - 施設設備の情報を管理';
COMMENT ON COLUMN facility_equipment.id IS '施設設備ID - 主キー';
COMMENT ON COLUMN facility_equipment.name IS '施設設備名 - 施設設備の名称';
COMMENT ON COLUMN facility_equipment.facility_type_id IS '施設種別ID - 施設設備の種別';
COMMENT ON COLUMN facility_equipment.model_number IS '型番 - 施設設備の型番';
COMMENT ON COLUMN facility_equipment.manufacturer IS 'メーカー - 施設設備のメーカー';
COMMENT ON COLUMN facility_equipment.installation_date IS '設置日 - 施設設備の設置日';
COMMENT ON COLUMN facility_equipment.status IS '状態 - 施設設備の稼働状態';
COMMENT ON COLUMN facility_equipment.location_description IS '設置場所説明 - 施設設備の設置場所の説明';
COMMENT ON COLUMN facility_equipment.location_latitude IS '位置（緯度） - 施設設備の緯度';
COMMENT ON COLUMN facility_equipment.location_longitude IS '位置（経度） - 施設設備の経度';
COMMENT ON COLUMN facility_equipment.notes IS '備考 - 施設設備に関する備考やメモ';
COMMENT ON COLUMN facility_equipment.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN facility_equipment.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN facility_equipment.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 更新日時を自動更新するトリガー
CREATE TRIGGER update_facility_equipment_updated_at
    BEFORE UPDATE ON facility_equipment
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

-- 初期データ投入（サンプル）
INSERT INTO facility_equipment (name, facility_type_id, model_number, manufacturer, installation_date, status, location_description, notes)
VALUES
    ('中央水路ポンプ', 1, 'WP-2000', '水設備株式会社', '2020-04-01', '稼働中', '中央水路北側入口', '年次点検：毎年4月'),
    ('東部ため池水門', 2, 'DG-500', '水門工業', '2019-07-15', '稼働中', '東部ため池南側', '遠隔操作可能'),
    ('西部農道街灯', 3, 'RL-100', '道路設備工業', '2021-02-28', '点検中', '西部農道沿い1km毎', '太陽光パネル搭載'),
    ('ハウスA温度管理システム', 4, 'TMS-X1', '農業設備株式会社', '2022-01-10', '稼働中', 'ビニールハウスA棟', 'スマホアプリ連携');
