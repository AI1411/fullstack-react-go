-- 市町村マスタテーブル
-- 全国の都道府県・市区町村の基本情報を管理する
DROP TABLE IF EXISTS municipalities CASCADE;
CREATE TABLE IF NOT EXISTS municipalities
(
    id                      SERIAL PRIMARY KEY,                                  -- 内部ID（主キー）
    prefecture_code         VARCHAR(2)   NOT NULL REFERENCES prefectures (code), -- 都道府県ID（外部キー）
    organization_code       VARCHAR(6)   NOT NULL UNIQUE,                        -- 団体コード（総務省地方公共団体コード）
    prefecture_name_kanji   VARCHAR(10)  NOT NULL,                               -- 都道府県名（漢字表記）
    municipality_name_kanji VARCHAR(50)  NOT NULL,                               -- 市区町村名（漢字表記）
    prefecture_name_kana    VARCHAR(20)  NOT NULL,                               -- 都道府県名（カタカナ表記）
    municipality_name_kana  VARCHAR(100) NOT NULL,                               -- 市区町村名（カタカナ表記）
    is_active               BOOLEAN      NOT NULL DEFAULT TRUE                   -- 有効フラグ（デフォルトは有効）合併などで無効化する場合に使用
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_municipalities_prefecture_id ON municipalities (prefecture_code);
CREATE INDEX IF NOT EXISTS idx_municipalities_organization_code ON municipalities (organization_code);
CREATE INDEX IF NOT EXISTS idx_municipalities_prefecture_kanji ON municipalities (prefecture_name_kanji);
CREATE INDEX IF NOT EXISTS idx_municipalities_municipality_kanji ON municipalities (municipality_name_kanji);
CREATE INDEX IF NOT EXISTS idx_municipalities_pref_muni_kanji ON municipalities (prefecture_name_kanji, municipality_name_kanji);
CREATE INDEX IF NOT EXISTS idx_municipalities_prefecture_kana ON municipalities (prefecture_name_kana);
CREATE INDEX IF NOT EXISTS idx_municipalities_municipality_kana ON municipalities (municipality_name_kana);
CREATE INDEX IF NOT EXISTS idx_municipalities_is_active ON municipalities (is_active);

-- テーブルコメント
COMMENT ON TABLE municipalities IS '市町村マスタテーブル - 全国の都道府県・市区町村の基本情報を管理';

-- カラムコメント
COMMENT ON COLUMN municipalities.id IS '内部ID（主キー、自動採番）';
COMMENT ON COLUMN municipalities.prefecture_code IS '都道府県ID（外部キー、都道府県マスタのID）';
COMMENT ON COLUMN municipalities.organization_code IS '団体コード（総務省地方公共団体コード、6桁）';
COMMENT ON COLUMN municipalities.prefecture_name_kanji IS '都道府県名（漢字表記）';
COMMENT ON COLUMN municipalities.municipality_name_kanji IS '市区町村名（漢字表記）';
COMMENT ON COLUMN municipalities.prefecture_name_kana IS '都道府県名（カタカナ表記）';
COMMENT ON COLUMN municipalities.municipality_name_kana IS '市区町村名（カタカナ表記）';
COMMENT ON COLUMN municipalities.is_active IS '有効フラグ（TRUE: 有効、FALSE: 無効）';

-- ダミーデータを挿入
INSERT INTO municipalities (prefecture_code, organization_code, prefecture_name_kanji, municipality_name_kanji,
                            prefecture_name_kana, municipality_name_kana)
VALUES ('01', '01001', '北海道', '札幌市', 'ホッカイドウ', 'サッポロシ'),
       ('02', '02001', '青森県', '青森市', 'アオモリケン', 'アオモリシ'),
       ('03', '03001', '岩手県', '盛岡市', 'イワテケン', 'モリオカシ'),
       ('04', '04001', '宮城県', '仙台市', 'ミヤギケン', 'センダイシ'),
       ('05', '05001', '秋田県', '秋田市', 'アキタケン', 'アキタシ'),
       ('06', '06001', '山形県', '山形市', 'ヤマガタケン', 'ヤマガタシ'),
       ('07', '07001', '福島県', '福島市', 'フクシマケン', 'フクシマシ'),
       ('08', '08001', '茨城県', '水戸市', 'イバラキケン', 'ミトシ'),
       ('09', '09001', '栃木県', '宇都宮市', 'トチギケン', 'ウツノミヤシ'),
       ('10', '10001', '群馬県', '前橋市', 'グンマケン', 'マエバシシ');

-- 工種区分マスタ
DROP TABLE IF EXISTS work_categories CASCADE;
CREATE TABLE work_categories
(
    id            SERIAL PRIMARY KEY,
    category_name VARCHAR(20) NOT NULL,
    icon_name     VARCHAR(50) NOT NULL, -- アイコンファイル名
    sort_order    INTEGER     NOT NULL DEFAULT 0,
    is_active     BOOLEAN     NOT NULL DEFAULT true
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_work_categories_category_name ON work_categories (category_name);
CREATE INDEX IF NOT EXISTS idx_work_categories_icon_name ON work_categories (icon_name);
CREATE INDEX IF NOT EXISTS idx_work_categories_sort_order ON work_categories (sort_order);
CREATE INDEX IF NOT EXISTS idx_work_categories_is_active ON work_categories (is_active);

-- テーブルコメント
COMMENT ON TABLE work_categories IS '工種区分マスタ - 工種区分の基本情報を管理';

-- カラムコメント
COMMENT ON COLUMN work_categories.id IS '工種区分ID（主キー、自動掲番）';
COMMENT ON COLUMN work_categories.category_name IS '工種区分名（漢字表記）';
COMMENT ON COLUMN work_categories.icon_name IS 'アイコンファイル名';
COMMENT ON COLUMN work_categories.sort_order IS '表示順序';
COMMENT ON COLUMN work_categories.is_active IS '有効フラグ（TRUE: 有効、FALSE: 無効）';

-- 工種区分マスタに初期データを挿入
INSERT INTO work_categories (category_name, icon_name, sort_order)
VALUES ('農地', '', 10),
       ('水路', '', 20),
       ('農道', '', 30),
       ('ため池', '', 40),
       ('頭首工', '', 50),
       ('揚水機', '', 60),
       ('堤防', '', 70),
       ('橋梁', '', 80),
       ('農地保全施設', '', 90);

-- 単価マスタ
DROP TABLE IF EXISTS unit_prices CASCADE;
CREATE TABLE unit_prices
(
    id              BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,       -- 単価ID（主キー、自動掲番）
    category_id     INTEGER        NOT NULL REFERENCES work_categories (id),
    prefecture_code VARCHAR(2)     NOT NULL REFERENCES prefectures (code), -- 都道府県コード（外部キー、prefecturesテーブルのコード）
    unit_price      DECIMAL(12, 2) NOT NULL,                               -- 単価（円）
    unit_type       VARCHAR(20)    NOT NULL,                               -- 'per_meter', 'per_sqm', 'per_unit'
    valid_from      DATE           NOT NULL DEFAULT CURRENT_DATE,
    valid_to        DATE,
    notes           TEXT,
    created_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_unit_prices_category_id ON unit_prices (category_id);
CREATE INDEX IF NOT EXISTS idx_unit_prices_prefecture_code ON unit_prices (prefecture_code);
CREATE INDEX IF NOT EXISTS idx_unit_prices_unit_type ON unit_prices (unit_type);
CREATE INDEX IF NOT EXISTS idx_unit_prices_valid_from ON unit_prices (valid_from);
CREATE INDEX IF NOT EXISTS idx_unit_prices_valid_to ON unit_prices (valid_to);

-- テーブルコメント
COMMENT ON TABLE unit_prices IS '単価マスタ - 工種区分ごとの単価を管理';

-- カラムコメント
COMMENT ON COLUMN unit_prices.id IS '単価ID（主キー、自動掲番）';
COMMENT ON COLUMN unit_prices.category_id IS '工種区分ID（外部キー、work_categoriesテーブルのID）';
COMMENT ON COLUMN unit_prices.prefecture_code IS '都道府県コード（外部キー、prefecturesテーブルのコード）';
COMMENT ON COLUMN unit_prices.unit_price IS '単価（円）';
COMMENT ON COLUMN unit_prices.unit_type IS '単位タイプ（"per_meter", "per_sqm", "per_unit"）';
COMMENT ON COLUMN unit_prices.valid_from IS '有効開始日';
COMMENT ON COLUMN unit_prices.valid_to IS '有効終了日';
COMMENT ON COLUMN unit_prices.notes IS '備考';
COMMENT ON COLUMN unit_prices.created_at IS '作成日時';
COMMENT ON COLUMN unit_prices.updated_at IS '更新日時';
