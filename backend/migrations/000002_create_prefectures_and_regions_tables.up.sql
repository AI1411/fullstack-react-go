-- 地域マスタテーブル
DROP TABLE IF EXISTS regions CASCADE;
CREATE TABLE IF NOT EXISTS regions
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_regions_name ON regions (name);

COMMENT ON TABLE regions IS '地域マスタテーブル';
COMMENT ON COLUMN regions.id IS '地域ID';
COMMENT ON COLUMN regions.name IS '地域名';
-- 都道府県と地域の初期データを挿入
INSERT INTO regions (name)
VALUES ('北海道・東北'),
       ('関東'),
       ('中部'),
       ('近畿'),
       ('中国'),
       ('四国'),
       ('九州・沖縄');

-- 都道府県マスタテーブル
DROP TABLE IF EXISTS prefectures CASCADE;
CREATE TABLE IF NOT EXISTS prefectures
(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(10) NOT NULL UNIQUE,
    region_id INT         NOT NULL REFERENCES regions (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_prefectures_name ON prefectures (name);
CREATE INDEX IF NOT EXISTS idx_prefectures_region_id ON prefectures (region_id);
COMMENT ON COLUMN prefectures.id IS '都道府県名';
COMMENT ON COLUMN prefectures.name IS '都道府県名';
COMMENT ON COLUMN prefectures.region_id IS '地域ID';

-- 都道府県マスタデータのINSERT
-- region_id: 1=北海道・東北, 2=関東, 3=中部, 4=近畿, 5=中国, 6=四国, 7=九州・沖縄

INSERT INTO prefectures (name, region_id)
VALUES
-- 北海道・東北地方 (region_id: 1)
('北海道', 1),
('青森県', 1),
('岩手県', 1),
('宮城県', 1),
('秋田県', 1),
('山形県', 1),
('福島県', 1),

-- 関東地方 (region_id: 2)
('茨城県', 2),
('栃木県', 2),
('群馬県', 2),
('埼玉県', 2),
('千葉県', 2),
('東京都', 2),
('神奈川県', 2),

-- 中部地方 (region_id: 3)
('新潟県', 3),
('富山県', 3),
('石川県', 3),
('福井県', 3),
('山梨県', 3),
('長野県', 3),
('岐阜県', 3),
('静岡県', 3),
('愛知県', 3),

-- 近畿地方 (region_id: 4)
('三重県', 4),
('滋賀県', 4),
('京都府', 4),
('大阪府', 4),
('兵庫県', 4),
('奈良県', 4),
('和歌山県', 4),

-- 中国地方 (region_id: 5)
('鳥取県', 5),
('島根県', 5),
('岡山県', 5),
('広島県', 5),
('山口県', 5),

-- 四国地方 (region_id: 6)
('徳島県', 6),
('香川県', 6),
('愛媛県', 6),
('高知県', 6),

-- 九州・沖縄地方 (region_id: 7)
('福岡県', 7),
('佐賀県', 7),
('長崎県', 7),
('熊本県', 7),
('大分県', 7),
('宮崎県', 7),
('鹿児島県', 7),
('沖縄県', 7)

ON CONFLICT (name) DO NOTHING;