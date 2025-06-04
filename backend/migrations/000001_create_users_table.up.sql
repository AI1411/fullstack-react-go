DROP TABLE IF EXISTS organizations CASCADE;
CREATE TABLE organizations
(
    id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name       VARCHAR(64)                           NOT NULL,
    type       VARCHAR(50)                           NOT NULL CHECK (type IN ('MAFF', 'regional', 'prefecture')), -- 組織の種類（農林水産本省、地方農政局、都道府県）
    parent_id  INTEGER REFERENCES organizations (id) NULL,
    sort_order INTEGER                  DEFAULT 0,
    is_active  BOOLEAN                  DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_organizations_name ON organizations (name);
CREATE INDEX IF NOT EXISTS idx_organizations_type ON organizations (type);
CREATE INDEX IF NOT EXISTS idx_organizations_parent_id ON organizations (parent_id);
CREATE INDEX IF NOT EXISTS idx_organizations_sort_order ON organizations (sort_order);
CREATE INDEX IF NOT EXISTS idx_organizations_is_active ON organizations (is_active);
-- テーブルコメント
COMMENT ON TABLE organizations IS '組織マスタ - 農林水産省、地方農政局、都道府県などの組織情報を管理';
COMMENT ON COLUMN organizations.id IS '組織ID（主キー、自動掲番）';
COMMENT ON COLUMN organizations.name IS '組織名';
COMMENT ON COLUMN organizations.type IS '組織の種類（MAFF: 農林水産本省、regional: 地方農政局、prefecture: 都道府県）';
COMMENT ON COLUMN organizations.parent_id IS '親組織ID（階層構造を持つための外部キー）';
COMMENT ON COLUMN organizations.sort_order IS '表示順序';
COMMENT ON COLUMN organizations.is_active IS '有効フラグ（TRUE: 有効、FALSE: 無効）';
-- ダミーデータ挿入
INSERT INTO organizations (name, type, parent_id, sort_order, is_active)
VALUES ('農林水産省', 'MAFF', NULL, 1, true),
       ('北海道農政事務所', 'regional', 1, 1, true),
       ('東北農政局', 'regional', 1, 2, true),
       ('関東農政局', 'regional', 1, 3, true),
       ('中部農政局', 'regional', 1, 4, true),
       ('近畿農政局', 'regional', 1, 5, true),
       ('中国四国農政局', 'regional', 1, 6, true),
       ('九州農政局', 'regional', 1, 7, true),
       ('沖縄総合事務局農林水産部', 'regional', 1, 8, true),
       -- 北海道農政事務所管内
       ('北海道', 'prefecture', 2, 1, true),
       -- 東北農政局管内
       ('青森県', 'prefecture', 3, 1, true),
       ('岩手県', 'prefecture', 3, 2, true),
       ('宮城県', 'prefecture', 3, 3, true),
       ('秋田県', 'prefecture', 3, 4, true),
       ('山形県', 'prefecture', 3, 5, true),
       ('福島県', 'prefecture', 3, 6, true),
       -- 関東農政局管内
       ('茨城県', 'prefecture', 4, 1, true),
       ('栃木県', 'prefecture', 4, 2, true),
       ('群馬県', 'prefecture', 4, 3, true),
       ('埼玉県', 'prefecture', 4, 4, true),
       ('千葉県', 'prefecture', 4, 5, true),
       ('東京都', 'prefecture', 4, 6, true),
       ('神奈川県', 'prefecture', 4, 7, true),
       ('山梨県', 'prefecture', 4, 8, true),
       ('長野県', 'prefecture', 4, 9, true),
       ('静岡県', 'prefecture', 4, 10, true),
       -- 中部農政局管内
       ('新潟県', 'prefecture', 5, 1, true),
       ('富山県', 'prefecture', 5, 2, true),
       ('石川県', 'prefecture', 5, 3, true),
       ('福井県', 'prefecture', 5, 4, true),
       ('岐阜県', 'prefecture', 5, 5, true),
       ('愛知県', 'prefecture', 5, 6, true),
       ('三重県', 'prefecture', 5, 7, true),
       -- 近畿農政局管内
       ('滋賀県', 'prefecture', 6, 1, true),
       ('京都府', 'prefecture', 6, 2, true),
       ('大阪府', 'prefecture', 6, 3, true),
       ('兵庫県', 'prefecture', 6, 4, true),
       ('奈良県', 'prefecture', 6, 5, true),
       ('和歌山県', 'prefecture', 6, 6, true),
       -- 中国四国農政局管内
       ('鳥取県', 'prefecture', 7, 1, true),
       ('島根県', 'prefecture', 7, 2, true),
       ('岡山県', 'prefecture', 7, 3, true),
       ('広島県', 'prefecture', 7, 4, true),
       ('山口県', 'prefecture', 7, 5, true),
       ('徳島県', 'prefecture', 7, 6, true),
       ('香川県', 'prefecture', 7, 7, true),
       ('愛媛県', 'prefecture', 7, 8, true),
       ('高知県', 'prefecture', 7, 9, true),
       -- 九州農政局管内
       ('福岡県', 'prefecture', 8, 1, true),
       ('佐賀県', 'prefecture', 8, 2, true),
       ('長崎県', 'prefecture', 8, 3, true),
       ('熊本県', 'prefecture', 8, 4, true),
       ('大分県', 'prefecture', 8, 5, true),
       ('宮崎県', 'prefecture', 8, 6, true),
       ('鹿児島県', 'prefecture', 8, 7, true),
       -- 沖縄総合事務局農林水産部管内
       ('沖縄県', 'prefecture', 9, 1, true);

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

-- ダミーユーザーデータ挿入
INSERT INTO users (name, email, password)
VALUES ('田中太郎', 'taro.tanaka@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('佐藤花子', 'hanako.sato@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('鈴木一郎', 'ichiro.suzuki@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('高橋幸子', 'sachiko.takahashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('渡辺雄太', 'yuta.watanabe@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('伊藤美咲', 'misaki.ito@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('山本健太', 'kenta.yamamoto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('中村洋子', 'yoko.nakamura@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('小林直人', 'naoto.kobayashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('加藤千尋', 'chihiro.kato@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('松本龍太郎', 'ryutaro.matsumoto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('井上真希', 'maki.inoue@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('木村大輔', 'daisuke.kimura@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('林優子', 'yuko.hayashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('斎藤拓也', 'takuya.saito@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('清水恵子', 'keiko.shimizu@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('山田隆史', 'takashi.yamada@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('中島裕太', 'yuta.nakajima@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('岡田彩香', 'ayaka.okada@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
       ('後藤光希', 'koki.goto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G');
