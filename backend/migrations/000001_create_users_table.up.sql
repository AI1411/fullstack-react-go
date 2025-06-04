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

DROP TABLE IF EXISTS roles CASCADE;
CREATE TABLE roles
(
    id              SMALLINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(100) NOT NULL,
    description     TEXT,
    organization_id SMALLINT     NOT NULL REFERENCES organizations (id), -- 組織の権限レベル（例: MAFF, regional, prefecture）,
    is_active       BOOLEAN                  DEFAULT true,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_roles_name ON roles (name);
CREATE INDEX IF NOT EXISTS idx_roles_organization_id ON roles (organization_id);
CREATE INDEX IF NOT EXISTS idx_roles_is_active ON roles (is_active);

COMMENT ON TABLE roles IS '役割マスタ - ユーザーの役割を管理';
COMMENT ON COLUMN roles.id IS '役割ID（主キー、自動掲番）';
COMMENT ON COLUMN roles.name IS '役割名（例: システム管理者、組織管理者、査定員など）';
COMMENT ON COLUMN roles.description IS '役割の説明';
COMMENT ON COLUMN roles.organization_id IS '組織ID（外部キー、組織マスタのID）';
COMMENT ON COLUMN roles.is_active IS '有効フラグ（TRUE: 有効、FALSE: 無効）';
COMMENT ON COLUMN roles.created_at IS '作成日時（レコード作成日時）';
COMMENT ON COLUMN roles.updated_at IS '更新日時（レコード最終更新日時）';

-- 初期データ投入
INSERT INTO roles (name, description, organization_id, is_active)
VALUES ('システム管理者', 'システム全体の管理権限を持つ最高権限者', 1, true),
       ('組織管理者', '特定の組織内でのユーザー管理や設定管理を行う権限者', 2, true),
       ('査定員', '災害被害の査定を行う権限を持つユーザー', 30, true),
       ('申請処理担当者', '支援申請の処理を行う権限を持つユーザー', 31, true),
       ('データ入力担当者', '災害情報や被害情報の入力を行うユーザー', 32, true),
       ('閲覧専用ユーザー', '情報の閲覧のみ可能なユーザー', 33, true);

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users
(
    id                     UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    name                   VARCHAR(255) NOT NULL,
    email                  VARCHAR(255) NOT NULL UNIQUE,
    password               VARCHAR(255) NOT NULL,
    organization_id        INTEGER      REFERENCES organizations (id) ON DELETE SET NULL,
    role_id                SMALLINT     NOT NULL REFERENCES roles (id) ON DELETE SET NULL,
    is_active              BOOLEAN      NOT NULL    DEFAULT true,
    email_verified         BOOLEAN      NOT NULL    DEFAULT false,
    mfa_enabled            BOOLEAN                  DEFAULT false,
    mfa_secret             VARCHAR(255),
    password_reset_token   VARCHAR(255),
    password_reset_expires TIMESTAMP WITH TIME ZONE,
    last_password_change   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at             TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at             TIMESTAMP WITH TIME ZONE
);

CREATE INDEX CONCURRENTLY idx_users_id_hash ON users USING hash (id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_organization_id ON users (organization_id);
CREATE INDEX IF NOT EXISTS idx_users_role_id ON users (role_id);
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users (is_active);

COMMENT ON TABLE users IS 'ユーザーテーブル - システムのユーザー情報を管理';

COMMENT ON COLUMN users.id IS 'ユーザーID（主キー、自動掲番）';
COMMENT ON COLUMN users.name IS 'ユーザー名';
COMMENT ON COLUMN users.email IS 'メールアドレス（ユニーク制約あり）';
COMMENT ON COLUMN users.password IS 'ハッシュ化されたパスワード';
COMMENT ON COLUMN users.organization_id IS '組織ID（外部キー、組織マスタのID）';
COMMENT ON COLUMN users.role_id IS '役割ID（外部キー、役割マスタのID）';
COMMENT ON COLUMN users.is_active IS '有効フラグ（TRUE: 有効、FALSE: 無効）';
COMMENT ON COLUMN users.email_verified IS 'メールアドレスの確認済みフラグ（TRUE: 確認済み、FALSE: 未確認）';
COMMENT ON COLUMN users.mfa_enabled IS '多要素認証の有効フラグ（TRUE: 有効、FALSE: 無効）';
COMMENT ON COLUMN users.mfa_secret IS '多要素認証のシークレットキー（TOTP用）';
COMMENT ON COLUMN users.password_reset_token IS 'パスワードリセット用のトークン';
COMMENT ON COLUMN users.password_reset_expires IS 'パスワードリセットトークンの有効期限';
COMMENT ON COLUMN users.last_password_change IS '最後のパスワード変更日時';
COMMENT ON COLUMN users.created_at IS 'ユーザー作成日時';
COMMENT ON COLUMN users.updated_at IS 'ユーザー情報の最終更新日時';
COMMENT ON COLUMN users.deleted_at IS 'ユーザー削除日時（論理削除用、NULLの場合は削除されていない）';


-- ダミーユーザーデータ挿入
INSERT INTO users (id, name, email, password, role_id, is_active, email_verified)
VALUES ('1cccd405-4150-4cf4-a534-7461f6aca1b3', '田中太郎', 'taro.tanaka@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 1, true, true), -- 農林水産省・システム管理者
       ('186e8952-f98f-4b45-8ca0-736047faa401', '佐藤花子', 'hanako.sato@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 2, true, true), -- 農林水産省・組織管理者
       ('12d8bfe1-96ff-491a-b36a-166d98f3b270', '鈴木一郎', 'ichiro.suzuki@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 東北農政局・査定員
       ('4362c0ab-cd9b-4694-be22-503e0f6a524c', '高橋幸子', 'sachiko.takahashi@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 関東農政局・査定員
       ('def28980-a550-49d7-a73e-1d09961f5296', '渡辺雄太', 'yuta.watanabe@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 中部農政局・査定員
       ('39c8ccf9-ab70-4dc5-8cfc-b92337aaa175', '伊藤美咲', 'misaki.ito@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 近畿農政局・査定員
       ('2977a69b-3baf-4df2-9894-76fb78773cb5', '山本健太', 'kenta.yamamoto@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 中国四国農政局・査定員
       ('da0b56a2-03b0-41b5-96ec-c68fb6fedde9', '中村洋子', 'yoko.nakamura@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 九州農政局・査定員
       ('2ee198e4-666f-4614-8da5-834dc4c7139d', '小林直人', 'naoto.kobayashi@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 3, true, true), -- 沖縄総合事務局・査定員
       ('31b92890-45fc-4cb0-8344-6dedf2f8446b', '加藤千尋', 'chihiro.kato@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 4, true, true), -- 北海道・申請処理担当者
       ('a3f6cb3b-3bf4-4162-9004-ffd8fe81b33a', '松本龍太郎', 'ryutaro.matsumoto@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 4, true, true), -- 宮城県・申請処理担当者
       ('577f5a91-6814-4bf8-a55b-792402d17153', '井上真希', 'maki.inoue@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 4, true, true), -- 東京都・申請処理担当者
       ('f61c1d9c-f098-4c74-9faf-9a49b4c19781', '木村大輔', 'daisuke.kimura@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 4, true, true), -- 愛知県・申請処理担当者
       ('87fbf2c6-341e-46c5-a2fc-1d9834c55a7e', '林優子', 'yuko.hayashi@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 5, true, true), -- 大阪府・データ入力担当者
       ('4b9d8271-b8a1-4e24-bab2-62ac44ec73a8', '斎藤拓也', 'takuya.saito@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 5, true, true), -- 岡山県・データ入力担当者
       ('e3dbf9e4-9182-4c4e-b927-6c0153bc07e4', '清水恵子', 'keiko.shimizu@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 5, true, true), -- 福岡県・データ入力担当者
       ('f08bcf78-248b-42ec-9df1-143b21e4b253', '山田隆史', 'takashi.yamada@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 6, true, true), -- 農林水産省・閲覧専用ユーザー
       ('558b9c39-f56d-4e6f-844c-28b2424cc4a7', '中島裕太', 'yuta.nakajima@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 6, true, true), -- 東北農政局・閲覧専用ユーザー
       ('d0db121b-77a5-48cb-a4ed-25d085c363d7', '岡田彩香', 'ayaka.okada@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 6, true, true), -- 沖縄県・閲覧専用ユーザー
       ('dcb243f6-4d77-4cb4-af9c-9d5f645580c8', '後藤光希', 'koki.goto@example.com',
        '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G', 2, false, false); -- 農林水産省・組織管理者（非アクティブ）

DROP TABLE IF EXISTS user_sessions CASCADE;
CREATE TABLE user_sessions
(
    id            BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    session_id    VARCHAR(255) UNIQUE      NOT NULL,
    user_id       UUID                     NOT NULL REFERENCES users (id),
    ip_address    INET,
    user_agent    TEXT,
    is_active     BOOLEAN                  DEFAULT true,
    expires_at    TIMESTAMP WITH TIME ZONE NOT NULL,
    last_activity TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON user_sessions (user_id);
CREATE INDEX IF NOT EXISTS idx_user_sessions_session_id ON user_sessions (session_id);
COMMENT ON TABLE user_sessions IS 'ユーザーセッションテーブル - ユーザーのログインセッション情報を管理';
COMMENT ON COLUMN user_sessions.id IS 'セッションID（主キー、自動掲番）';
COMMENT ON COLUMN user_sessions.session_id IS 'セッション識別子（ユニーク制約あり）';
COMMENT ON COLUMN user_sessions.user_id IS 'ユーザーID（外部キー、ユーザーテーブルのID）';
COMMENT ON COLUMN user_sessions.ip_address IS 'ユーザーのIPアドレス';
COMMENT ON COLUMN user_sessions.user_agent IS 'ユーザーエージェント情報（ブラウザやデバイス情報）';
COMMENT ON COLUMN user_sessions.is_active IS 'セッションの有効フラグ（TRUE: 有効、FALSE: 無効）';
COMMENT ON COLUMN user_sessions.expires_at IS 'セッションの有効期限';
COMMENT ON COLUMN user_sessions.last_activity IS '最後のアクティビティ日時（セッションの最終アクティビティ）';
COMMENT ON COLUMN user_sessions.created_at IS 'セッション作成日時';

DROP TABLE IF EXISTS login_history CASCADE;
CREATE TABLE login_history
(
    id             BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id        UUID REFERENCES users (id),
    username       VARCHAR(100),         -- ユーザーが削除されても履歴を残すため
    login_type     VARCHAR(50) NOT NULL, -- 'success', 'failed', 'locked', 'mfa_required'
    ip_address     INET,
    user_agent     TEXT,
    failure_reason VARCHAR(255),         -- 失敗理由
    created_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_login_history_user_id ON login_history (user_id);
CREATE INDEX IF NOT EXISTS idx_login_history_username ON login_history (username);
COMMENT ON TABLE login_history IS 'ログイン履歴テーブル - ユーザーのログイン履歴を管理';
COMMENT ON COLUMN login_history.id IS 'ログイン履歴ID（主キー、自動掲番）';
COMMENT ON COLUMN login_history.user_id IS 'ユーザーID（外部キー、ユーザーテーブルのID）';
COMMENT ON COLUMN login_history.username IS 'ユーザー名（ユーザーが削除されても履歴を残すため）';
COMMENT ON COLUMN login_history.login_type IS 'ログインタイプ（成功: success、失敗: failed、ロック: locked、多要素認証必要: mfa_required）';
COMMENT ON COLUMN login_history.ip_address IS 'ユーザーのIPアドレス';
COMMENT ON COLUMN login_history.user_agent IS 'ユーザーエージェント情報（ブラウザやデバイス情報）';
COMMENT ON COLUMN login_history.failure_reason IS 'ログイン失敗理由（失敗時のみ）';
COMMENT ON COLUMN login_history.created_at IS 'ログイン履歴の作成日時';

DROP TABLE IF EXISTS operation_logs CASCADE;
CREATE TABLE operation_logs
(
    id                BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id           UUID REFERENCES users (id),
    action            VARCHAR(100) NOT NULL,
    resource_type     VARCHAR(100),
    resource_id       VARCHAR(100),
    endpoint          VARCHAR(500),
    method            VARCHAR(10),
    ip_address        INET,
    user_agent        TEXT,
    request_data      JSONB,
    response_status   INTEGER,
    execution_time_ms INTEGER,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_operation_logs_user_id ON operation_logs (user_id);
CREATE INDEX IF NOT EXISTS idx_operation_logs_action ON operation_logs (action);
CREATE INDEX IF NOT EXISTS idx_operation_logs_resource_type ON operation_logs (resource_type);
CREATE INDEX IF NOT EXISTS idx_operation_logs_resource_id ON operation_logs (resource_id);
CREATE INDEX IF NOT EXISTS idx_operation_logs_endpoint ON operation_logs (endpoint);
CREATE INDEX IF NOT EXISTS idx_operation_logs_method ON operation_logs (method);
CREATE INDEX IF NOT EXISTS idx_operation_logs_request_data ON operation_logs (request_data);
CREATE INDEX IF NOT EXISTS idx_operation_logs_response_status ON operation_logs (response_status);
CREATE INDEX IF NOT EXISTS idx_operation_logs_execution_time_ms ON operation_logs (execution_time_ms);
CREATE INDEX IF NOT EXISTS idx_operation_logs_created_at ON operation_logs (created_at);

COMMENT ON TABLE operation_logs IS '監査ログテーブル - システムの操作履歴を管理';
COMMENT ON COLUMN operation_logs.id IS '監査ログID（主キー、自動掲番）';
COMMENT ON COLUMN operation_logs.user_id IS 'ユーザーID（外部キー、ユーザーテーブルのID）';
COMMENT ON COLUMN operation_logs.action IS '実行されたアクション（例: create, update, delete）';
COMMENT ON COLUMN operation_logs.resource_type IS 'リソースの種類（例: user, organization, role）';
COMMENT ON COLUMN operation_logs.resource_id IS 'リソースID（アクションが適用されたリソースのID）';
COMMENT ON COLUMN operation_logs.endpoint IS 'APIエンドポイント（例: /api/users, /api/organizations）';
COMMENT ON COLUMN operation_logs.method IS 'HTTPメソッド（例: GET, POST, PUT, DELETE）';
COMMENT ON COLUMN operation_logs.ip_address IS 'ユーザーのIPアドレス';
COMMENT ON COLUMN operation_logs.user_agent IS 'ユーザーエージェント情報（ブラウザやデバイス情報）';
COMMENT ON COLUMN operation_logs.request_data IS 'リクエストデータ（JSON形式）';
COMMENT ON COLUMN operation_logs.response_status IS 'レスポンスステータスコード（例: 200, 404, 500）';
COMMENT ON COLUMN operation_logs.execution_time_ms IS 'リクエストの実行時間（ミリ秒単位）';
COMMENT ON COLUMN operation_logs.created_at IS '監査ログの作成日時';
