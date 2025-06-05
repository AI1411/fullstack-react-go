DROP TABLE IF EXISTS email_histories;
CREATE TABLE email_histories
(
    id            BIGSERIAL PRIMARY KEY,
    user_id       UUID                                               NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    email         VARCHAR(255)                                       NOT NULL,
    subject       VARCHAR(500)                                       NOT NULL,
    email_type    VARCHAR(100)                                       NOT NULL, -- register, verification, password_reset等
    provider      VARCHAR(50)                                        NOT NULL, -- smtp, sendgrid等
    error_message TEXT,
    sent_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status        VARCHAR(50)                                        NOT NULL, -- success, failed, pending等
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_email_histories_user_id ON email_histories (user_id);
CREATE INDEX IF NOT EXISTS idx_email_histories_email ON email_histories (email);
CREATE INDEX IF NOT EXISTS idx_email_histories_status ON email_histories (status);
CREATE INDEX IF NOT EXISTS idx_email_histories_email_type ON email_histories (email_type);
CREATE INDEX IF NOT EXISTS idx_email_histories_sent_at ON email_histories (sent_at);

-- table comments
COMMENT ON TABLE email_histories IS 'ユーザーのメール送信履歴を保存するテーブル';
-- column comments
COMMENT ON COLUMN email_histories.id IS 'メール履歴ID';
COMMENT ON COLUMN email_histories.user_id IS 'ユーザーID';
COMMENT ON COLUMN email_histories.email IS '送信先メールアドレス';
COMMENT ON COLUMN email_histories.subject IS 'メール件名';
COMMENT ON COLUMN email_histories.email_type IS 'メール種別（welcome, verification, password_reset等）';
COMMENT ON COLUMN email_histories.provider IS 'メール送信プロバイダー（smtp, sendgrid等）';
COMMENT ON COLUMN email_histories.error_message IS 'エラーメッセージ（送信失敗時）';
COMMENT ON COLUMN email_histories.sent_at IS 'メール送信日時';
COMMENT ON COLUMN email_histories.status IS 'メール送信ステータス（sent, failed, pending等）';
COMMENT ON COLUMN email_histories.created_at IS '作成日時';
COMMENT ON COLUMN email_histories.updated_at IS '更新日時';
