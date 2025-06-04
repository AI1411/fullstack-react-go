-- マスターデータテーブル作成

-- 災害種別マスタテーブル
DROP TABLE IF EXISTS disaster_types CASCADE;
CREATE TABLE IF NOT EXISTS disaster_types
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50)              NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_disaster_types_name ON disaster_types (name);

COMMENT ON TABLE disaster_types IS '災害種別マスタテーブル - 災害の種類を管理';
COMMENT ON COLUMN disaster_types.id IS '災害種別ID - 主キー';
COMMENT ON COLUMN disaster_types.name IS '災害種別名 - 洪水, 地滑り, 雹害, 干ばつ, 風害, 地震, 霜害, 病害虫など';
COMMENT ON COLUMN disaster_types.description IS '説明 - 災害種別の詳細説明';
COMMENT ON COLUMN disaster_types.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN disaster_types.updated_at IS '更新日時 - レコード最終更新日時';

-- 初期データ投入
INSERT INTO disaster_types (name, description)
VALUES ('洪水', '河川の氾濫や大雨による浸水被害'),
       ('地滑り', '土砂崩れや地滑りによる被害'),
       ('雹害', '雹（ひょう）の降下による農作物への物理的被害'),
       ('干ばつ', '長期間の少雨による水不足被害'),
       ('風害', '強風や台風による被害'),
       ('地震', '地震による農地・施設への被害'),
       ('霜害', '霜による農作物への被害'),
       ('病害虫', '病気や害虫による農作物への被害'),
       ('長雨', '長期間の降雨による被害'),
       ('低温', '異常低温による農作物への被害'),
       ('高温', '異常高温による農作物への被害'),
       ('塩害', '海水の浸入による塩害'),
       ('雪害', '大雪や雪崩による被害'),
       ('その他', 'その他の自然災害による被害');

-- 被害程度マスタテーブル
DROP TABLE IF EXISTS damage_levels CASCADE;
CREATE TABLE IF NOT EXISTS damage_levels
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50)              NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_damage_levels_name ON damage_levels (name);

COMMENT ON TABLE damage_levels IS '被害程度マスタテーブル - 被害の程度を管理';
COMMENT ON COLUMN damage_levels.id IS '被害程度ID - 主キー';
COMMENT ON COLUMN damage_levels.name IS '被害程度名 - 軽微, 中程度, 深刻, 甚大など';
COMMENT ON COLUMN damage_levels.description IS '説明 - 被害程度の詳細説明';
COMMENT ON COLUMN damage_levels.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN damage_levels.updated_at IS '更新日時 - レコード最終更新日時';

-- 初期データ投入
INSERT INTO damage_levels (name, description)
VALUES ('軽微', '被害は小規模で、通常の営農活動への影響は限定的'),
       ('中程度', '被害は一定規模あり、営農活動に部分的な影響がある'),
       ('深刻', '被害規模が大きく、営農活動に大きな影響がある'),
       ('甚大', '被害が極めて大きく、営農活動の継続が困難な状態'),
       ('不明', '被害程度が調査中または不明な状態');

-- 施設種別マスタテーブル
DROP TABLE IF EXISTS facility_types CASCADE;
CREATE TABLE IF NOT EXISTS facility_types
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50)              NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_facility_types_name ON facility_types (name);

COMMENT ON TABLE facility_types IS '施設種別マスタテーブル - 農業用施設の種類を管理';
COMMENT ON COLUMN facility_types.id IS '施設種別ID - 主キー';
COMMENT ON COLUMN facility_types.name IS '施設種別名 - 水路, ため池, 農道, ビニールハウスなど';
COMMENT ON COLUMN facility_types.description IS '説明 - 施設種別の詳細説明';
COMMENT ON COLUMN facility_types.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN facility_types.updated_at IS '更新日時 - レコード最終更新日時';

-- 初期データ投入
INSERT INTO facility_types (name, description)
VALUES ('水路', '農業用水を供給するための水路施設'),
       ('ため池', '農業用水を貯水するための池'),
       ('農道', '農地へのアクセスや農作業のための道路'),
       ('ビニールハウス', '作物栽培用のビニールハウス施設'),
       ('畜舎', '家畜を飼育するための施設'),
       ('育苗施設', '苗を育てるための施設'),
       ('集出荷施設', '農産物の集荷・出荷のための施設'),
       ('農業倉庫', '農機具や資材を保管するための倉庫'),
       ('灌漑施設', '農地に水を供給するためのポンプや設備'),
       ('暗渠排水', '農地の地下排水施設'),
       ('その他', 'その他の農業関連施設');

-- 役割マスタテーブル
DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE IF NOT EXISTS roles
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50)              NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_roles_name ON roles (name);

COMMENT ON TABLE roles IS '役割マスタテーブル - ユーザーの役割を管理';
COMMENT ON COLUMN roles.id IS '役割ID - 主キー';
COMMENT ON COLUMN roles.name IS '役割名 - 管理者, 一般ユーザー, 査定員など';
COMMENT ON COLUMN roles.description IS '説明 - 役割の詳細説明と権限';
COMMENT ON COLUMN roles.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN roles.updated_at IS '更新日時 - レコード最終更新日時';

-- 初期データ投入
INSERT INTO roles (name, description)
VALUES ('システム管理者', 'システム全体の管理権限を持つ最高権限者'),
       ('組織管理者', '特定の組織内でのユーザー管理や設定管理を行う権限者'),
       ('査定員', '災害被害の査定を行う権限を持つユーザー'),
       ('申請処理担当者', '支援申請の処理を行う権限を持つユーザー'),
       ('データ入力担当者', '災害情報や被害情報の入力を行うユーザー'),
       ('閲覧専用ユーザー', '情報の閲覧のみ可能なユーザー');

-- ユーザーロールテーブル（ユーザーと役割の多対多関連）
DROP TABLE IF EXISTS user_roles CASCADE;
CREATE TABLE IF NOT EXISTS user_roles
(
    user_id    UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role_id    INT                      NOT NULL REFERENCES roles (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id)
);

CREATE INDEX IF NOT EXISTS idx_user_roles_user_id ON user_roles (user_id);
CREATE INDEX IF NOT EXISTS idx_user_roles_role_id ON user_roles (role_id);

COMMENT ON TABLE user_roles IS 'ユーザーロール関連テーブル - ユーザーと役割の関連を管理';
COMMENT ON COLUMN user_roles.user_id IS 'ユーザーID - ユーザーテーブルの外部キー';
COMMENT ON COLUMN user_roles.role_id IS '役割ID - 役割テーブルの外部キー';
COMMENT ON COLUMN user_roles.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN user_roles.updated_at IS '更新日時 - レコード最終更新日時';

-- ユーザー組織テーブル（ユーザーと組織の多対多関連）
DROP TABLE IF EXISTS user_organizations CASCADE;
CREATE TABLE IF NOT EXISTS user_organizations
(
    user_id         UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    organization_id INT                      NOT NULL REFERENCES organizations (id) ON DELETE CASCADE,
    is_primary      BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, organization_id)
);

CREATE INDEX IF NOT EXISTS idx_user_organizations_user_id ON user_organizations (user_id);
CREATE INDEX IF NOT EXISTS idx_user_organizations_organization_id ON user_organizations (organization_id);

COMMENT ON TABLE user_organizations IS 'ユーザー組織関連テーブル - ユーザーと組織の関連を管理';
COMMENT ON COLUMN user_organizations.user_id IS 'ユーザーID - ユーザーテーブルの外部キー';
COMMENT ON COLUMN user_organizations.organization_id IS '組織ID - 組織テーブルの外部キー';
COMMENT ON COLUMN user_organizations.is_primary IS '主所属フラグ - ユーザーの主所属組織かどうか';
COMMENT ON COLUMN user_organizations.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN user_organizations.updated_at IS '更新日時 - レコード最終更新日時';

-- 更新日時を自動更新するトリガー関数
CREATE OR REPLACE FUNCTION update_master_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 更新トリガーの作成
CREATE TRIGGER update_disaster_types_updated_at
    BEFORE UPDATE
    ON disaster_types
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_damage_levels_updated_at
    BEFORE UPDATE
    ON damage_levels
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_facility_types_updated_at
    BEFORE UPDATE
    ON facility_types
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_roles_updated_at
    BEFORE UPDATE
    ON roles
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_user_roles_updated_at
    BEFORE UPDATE
    ON user_roles
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_organizations_updated_at
    BEFORE UPDATE
    ON organizations
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_user_organizations_updated_at
    BEFORE UPDATE
    ON user_organizations
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

-- ユーザーと組織の関連付けのダミーデータ挿入
INSERT INTO user_organizations (user_id, organization_id, is_primary)
VALUES
-- 農林水産省所属
('1cccd405-4150-4cf4-a534-7461f6aca1b3', 1, true),  -- 田中太郎: 農林水産省（主所属）
('186e8952-f98f-4b45-8ca0-736047faa401', 1, true),  -- 佐藤花子: 農林水産省（主所属）
('12d8bfe1-96ff-491a-b36a-166d98f3b270', 1, false), -- 鈴木一郎: 農林水産省（副所属）

-- 関東農政局所属
('4362c0ab-cd9b-4694-be22-503e0f6a524c', 2, true),  -- 鈴木一郎: 関東農政局（主所属）
('def28980-a550-49d7-a73e-1d09961f5296', 2, true),  -- 高橋幸子: 関東農政局（主所属）
('39c8ccf9-ab70-4dc5-8cfc-b92337aaa175', 2, true),  -- 渡辺雄太: 関東農政局（主所属）

-- 東京都農林水産部所属
('2977a69b-3baf-4df2-9894-76fb78773cb5', 3, true),  -- 伊藤美咲: 東京都農林水産部（主所属）
('da0b56a2-03b0-41b5-96ec-c68fb6fedde9', 3, true),  -- 山本健太: 東京都農林水産部（主所属）
('2ee198e4-666f-4614-8da5-834dc4c7139d', 3, false), -- 中村洋子: 東京都農林水産部（副所属）

-- 青森県農林水産部所属
('31b92890-45fc-4cb0-8344-6dedf2f8446b', 4, true),  -- 中村洋子: 青森県農林水産部（主所属）
('a3f6cb3b-3bf4-4162-9004-ffd8fe81b33a', 4, true),  -- 小林直人: 青森県農林水産部（主所属）
('577f5a91-6814-4bf8-a55b-792402d17153', 4, true),  -- 加藤千尋: 青森県農林水産部（主所属）

-- JA全農所属
('f61c1d9c-f098-4c74-9faf-9a49b4c19781', 5, true),  -- 松本龍太郎: JA全農（主所属）
('87fbf2c6-341e-46c5-a2fc-1d9834c55a7e', 5, true),  -- 井上真希: JA全農（主所属）
('4b9d8271-b8a1-4e24-bab2-62ac44ec73a8', 5, true),  -- 木村大輔: JA全農（主所属）

-- JA東京所属
('e3dbf9e4-9182-4c4e-b927-6c0153bc07e4', 6, true),  -- 林優子: JA東京（主所属）
('f08bcf78-248b-42ec-9df1-143b21e4b253', 6, true),  -- 斎藤拓也: JA東京（主所属）
('558b9c39-f56d-4e6f-844c-28b2424cc4a7', 6, true),  -- 清水恵子: JA東京（主所属）

-- 複数組織に所属する例
('d0db121b-77a5-48cb-a4ed-25d085c363d7', 1, false), -- 山田隆史: 農林水産省（副所属）
('dcb243f6-4d77-4cb4-af9c-9d5f645580c8', 5, true),  -- 山田隆史: JA全農（主所属）
('d0db121b-77a5-48cb-a4ed-25d085c363d7', 3, false), -- 中島裕太: 東京都農林水産部（副所属）
('d0db121b-77a5-48cb-a4ed-25d085c363d7', 6, true),  -- 中島裕太: JA東京（主所属）
('d0db121b-77a5-48cb-a4ed-25d085c363d7', 2, false), -- 岡田彩香: 関東農政局（副所属）
('d0db121b-77a5-48cb-a4ed-25d085c363d7', 4, true); -- 岡田彩香: 青森県農林水産部（主所属）
