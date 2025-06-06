DROP TABLE IF EXISTS email_verification_tokens;
CREATE TABLE IF NOT EXISTS email_verification_tokens
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    user_id    UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    token      VARCHAR(255)             NOT NULL UNIQUE,
    email      VARCHAR(255)             NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used_at    TIMESTAMP WITH TIME ZONE,
    is_used    BOOLEAN                  NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- インデックスの作成
CREATE INDEX idx_email_verification_tokens_user_id ON email_verification_tokens (user_id);
CREATE INDEX idx_email_verification_tokens_token ON email_verification_tokens (token);
CREATE INDEX idx_email_verification_tokens_expires_at ON email_verification_tokens (expires_at);

-- コメント追加
COMMENT ON TABLE email_verification_tokens IS 'メール認証トークンを格納するテーブル';
COMMENT ON COLUMN email_verification_tokens.id IS 'トークンの一意な識別子';
COMMENT ON COLUMN email_verification_tokens.user_id IS '関連するユーザーのID';
COMMENT ON COLUMN email_verification_tokens.token IS 'メール認証トークン';
COMMENT ON COLUMN email_verification_tokens.email IS '認証対象のメールアドレス';
COMMENT ON COLUMN email_verification_tokens.expires_at IS 'トークンの有効期限';
COMMENT ON COLUMN email_verification_tokens.used_at IS 'トークンが使用された日時';
COMMENT ON COLUMN email_verification_tokens.is_used IS 'トークンが使用されたかどうか';
COMMENT ON COLUMN email_verification_tokens.created_at IS 'トークンの作成日時';
