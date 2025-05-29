-- マスターデータテーブル作成

-- 災害種別マスタテーブル
DROP TABLE IF EXISTS disaster_types CASCADE;
CREATE TABLE IF NOT EXISTS disaster_types
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
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
VALUES
    ('洪水', '河川の氾濫や大雨による浸水被害'),
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
    name        VARCHAR(50) NOT NULL UNIQUE,
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
VALUES
    ('軽微', '被害は小規模で、通常の営農活動への影響は限定的'),
    ('中程度', '被害は一定規模あり、営農活動に部分的な影響がある'),
    ('深刻', '被害規模が大きく、営農活動に大きな影響がある'),
    ('甚大', '被害が極めて大きく、営農活動の継続が困難な状態'),
    ('不明', '被害程度が調査中または不明な状態');

-- 施設種別マスタテーブル
DROP TABLE IF EXISTS facility_types CASCADE;
CREATE TABLE IF NOT EXISTS facility_types
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(50) NOT NULL UNIQUE,
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
VALUES
    ('水路', '農業用水を供給するための水路施設'),
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
    name        VARCHAR(50) NOT NULL UNIQUE,
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
VALUES
    ('システム管理者', 'システム全体の管理権限を持つ最高権限者'),
    ('組織管理者', '特定の組織内でのユーザー管理や設定管理を行う権限者'),
    ('査定員', '災害被害の査定を行う権限を持つユーザー'),
    ('申請処理担当者', '支援申請の処理を行う権限を持つユーザー'),
    ('データ入力担当者', '災害情報や被害情報の入力を行うユーザー'),
    ('閲覧専用ユーザー', '情報の閲覧のみ可能なユーザー');

-- ユーザーロールテーブル（ユーザーと役割の多対多関連）
DROP TABLE IF EXISTS user_roles CASCADE;
CREATE TABLE IF NOT EXISTS user_roles
(
    user_id     INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id     INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, role_id)
);

CREATE INDEX IF NOT EXISTS idx_user_roles_user_id ON user_roles (user_id);
CREATE INDEX IF NOT EXISTS idx_user_roles_role_id ON user_roles (role_id);

COMMENT ON TABLE user_roles IS 'ユーザーロール関連テーブル - ユーザーと役割の関連を管理';
COMMENT ON COLUMN user_roles.user_id IS 'ユーザーID - ユーザーテーブルの外部キー';
COMMENT ON COLUMN user_roles.role_id IS '役割ID - 役割テーブルの外部キー';
COMMENT ON COLUMN user_roles.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN user_roles.updated_at IS '更新日時 - レコード最終更新日時';

-- 組織テーブル
DROP TABLE IF EXISTS organizations CASCADE;
CREATE TABLE IF NOT EXISTS organizations
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    type        VARCHAR(50) NOT NULL,
    prefecture_id INT REFERENCES prefectures(id),
    parent_id   INT REFERENCES organizations(id),
    description TEXT,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_organizations_name ON organizations (name);
CREATE INDEX IF NOT EXISTS idx_organizations_type ON organizations (type);
CREATE INDEX IF NOT EXISTS idx_organizations_prefecture_id ON organizations (prefecture_id);
CREATE INDEX IF NOT EXISTS idx_organizations_parent_id ON organizations (parent_id);

COMMENT ON TABLE organizations IS '組織テーブル - 行政機関や団体の組織情報を管理';
COMMENT ON COLUMN organizations.id IS '組織ID - 主キー';
COMMENT ON COLUMN organizations.name IS '組織名 - 組織の名称';
COMMENT ON COLUMN organizations.type IS '組織種別 - 国, 都道府県, 市町村, JAなど';
COMMENT ON COLUMN organizations.prefecture_id IS '都道府県ID - 組織が所属する都道府県';
COMMENT ON COLUMN organizations.parent_id IS '親組織ID - 階層構造を表現するための親組織ID';
COMMENT ON COLUMN organizations.description IS '説明 - 組織の詳細説明';
COMMENT ON COLUMN organizations.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN organizations.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN organizations.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 初期データ投入（サンプル）
INSERT INTO organizations (name, type, prefecture_id, parent_id, description)
VALUES
    ('農林水産省', '国', NULL, NULL, '農林水産行政を担当する国の行政機関'),
    ('関東農政局', '国', 13, 1, '関東地方の農政を担当する農林水産省の地方機関'),
    ('東京都農林水産部', '都道府県', 13, NULL, '東京都の農林水産行政を担当する部署'),
    ('青森県農林水産部', '都道府県', 2, NULL, '青森県の農林水産行政を担当する部署'),
    ('JA全農', '団体', NULL, NULL, '全国農業協同組合連合会'),
    ('JA東京', '団体', 13, 5, '東京都の農業協同組合');

-- ユーザー組織テーブル（ユーザーと組織の多対多関連）
DROP TABLE IF EXISTS user_organizations CASCADE;
CREATE TABLE IF NOT EXISTS user_organizations
(
    user_id         INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    organization_id INT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    is_primary      BOOLEAN NOT NULL DEFAULT FALSE,
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
    BEFORE UPDATE ON disaster_types
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_damage_levels_updated_at
    BEFORE UPDATE ON damage_levels
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_facility_types_updated_at
    BEFORE UPDATE ON facility_types
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_roles_updated_at
    BEFORE UPDATE ON roles
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_user_roles_updated_at
    BEFORE UPDATE ON user_roles
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_organizations_updated_at
    BEFORE UPDATE ON organizations
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();

CREATE TRIGGER update_user_organizations_updated_at
    BEFORE UPDATE ON user_organizations
    FOR EACH ROW
EXECUTE FUNCTION update_master_updated_at_column();
