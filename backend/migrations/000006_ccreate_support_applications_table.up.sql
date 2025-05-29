-- 支援申請テーブルの作成
CREATE TABLE support_applications
(
    application_id   VARCHAR(10) PRIMARY KEY,
    application_date DATE         NOT NULL,
    applicant_name   VARCHAR(100) NOT NULL,
    disaster_name    VARCHAR(200) NOT NULL,
    requested_amount BIGINT       NOT NULL,
    CHECK (requested_amount >= 0),
    status           VARCHAR(20)  NOT NULL DEFAULT '審査中',
    CHECK (status IN ('審査中', '書類確認中', '承認済', '完了', '支払処理中', '却下')),
    reviewed_at      TIMESTAMP,
    approved_at      TIMESTAMP,
    completed_at     TIMESTAMP,
    notes            TEXT,
    created_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- インデックスの作成
CREATE INDEX IF NOT EXISTS idx_support_applications_status ON support_applications (status);
CREATE INDEX IF NOT EXISTS idx_support_applications_application_date ON support_applications (application_date);
CREATE INDEX IF NOT EXISTS idx_support_applications_disaster_name ON support_applications (disaster_name);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE support_applications IS '支援申請管理テーブル - 災害時の農業被害に対する支援申請を格納';
COMMENT ON COLUMN support_applications.application_id IS '申請ID - 主キー（例：A001, A002...）';
COMMENT ON COLUMN support_applications.application_date IS '申請日 - 申請が提出された日付';
COMMENT ON COLUMN support_applications.applicant_name IS '申請者名 - 個人名または法人名';
COMMENT ON COLUMN support_applications.disaster_name IS '災害名 - 対象となる災害の名称';
COMMENT ON COLUMN support_applications.requested_amount IS '申請金額 - 申請する支援金額（円）';
COMMENT ON COLUMN support_applications.status IS 'ステータス - 申請の処理状況（審査中、書類確認中、承認済、完了、支払処理中、却下）';
COMMENT ON COLUMN support_applications.reviewed_at IS '審査完了日時 - 申請の審査が完了した日時';
COMMENT ON COLUMN support_applications.approved_at IS '承認日時 - 申請が承認された日時';
COMMENT ON COLUMN support_applications.completed_at IS '処理完了日時 - 支援金の支払いなど全ての処理が完了した日時';
COMMENT ON COLUMN support_applications.notes IS '備考 - 申請に関する備考やメモ';
COMMENT ON COLUMN support_applications.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN support_applications.updated_at IS '更新日時 - レコード最終更新日時';

-- 更新日時を自動更新するトリガー関数
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 更新トリガーの作成
CREATE TRIGGER update_support_applications_updated_at
    BEFORE UPDATE
    ON support_applications
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- ダミーデータの投入
INSERT INTO support_applications (application_id,
                                  application_date,
                                  applicant_name,
                                  disaster_name,
                                  requested_amount,
                                  status,
                                  reviewed_at,
                                  approved_at,
                                  completed_at,
                                  notes)
VALUES
-- 審査中のケース
('A001', '2024-07-18', '山田農園', '京都府洪水被害', 2500000, '審査中', NULL, NULL, NULL, '水稲の冠水被害による損失'),

-- 書類確認中のケース
('A002', '2024-07-16', '大阪みかん生産組合', '大阪府地滑り被害', 1800000, '書類確認中', '2024-07-20 10:30:00', NULL,
 NULL, '追加書類の提出待ち'),

-- 承認済のケース
('A003', '2024-07-12', '兵庫果樹園', '兵庫県霜害', 3200000, '承認済', '2024-07-15 14:20:00', '2024-07-16 09:15:00',
 NULL, '果樹の枝折れ被害'),

-- 審査中のケース
('A004', '2024-07-10', '奈良田んぼ組合', '奈良県干ばつ被害', 1500000, '審査中', NULL, NULL, NULL,
 '水不足による収量減少'),

-- 支払処理中のケース
('A005', '2024-07-05', '滋賀グリーンハウス', '滋賀県風害', 4200000, '支払処理中', '2024-07-08 11:45:00',
 '2024-07-09 16:30:00', NULL, 'ビニールハウス倒壊被害'),

-- 完了のケース
('A006', '2024-07-01', '和歌山みかん農園', '和歌山県水害', 1700000, '完了', '2024-07-03 13:20:00',
 '2024-07-04 10:00:00', '2024-07-10 15:45:00', '支援金支払い完了'),

-- 審査中のケース
('A007', '2024-06-28', '三重農業協同組合', '三重県地滑り被害', 2900000, '審査中', NULL, NULL, NULL,
 '棚田の土砂流出被害'),

-- 書類確認中のケース
('A008', '2024-06-22', '愛知野菜生産者', '愛知県霜害', 1200000, '書類確認中', '2024-06-25 09:30:00', NULL, NULL,
 '露地野菜の冷害'),

-- 完了のケース
('A009', '2024-06-18', '岐阜畜産農家', '岐阜県干ばつ被害', 3500000, '完了', '2024-06-20 14:15:00',
 '2024-06-21 11:20:00', '2024-06-28 16:00:00', '飼料作物の被害補償'),

-- 承認済のケース
('A010', '2024-06-15', '静岡果樹園', '静岡県風害', 1900000, '承認済', '2024-06-18 10:45:00', '2024-06-19 13:30:00',
 NULL, 'みかん樹の倒木被害');

-- データ確認用のクエリ
SELECT application_id                                           AS "申請ID",
       application_date                                         AS "申請日",
       applicant_name                                           AS "申請者",
       disaster_name                                            AS "災害名",
       CONCAT(TO_CHAR(requested_amount, 'FM999,999,999'), '円') AS "申請金額",
       status                                                   AS "ステータス",
       notes                                                    AS "備考"
FROM support_applications
ORDER BY application_date DESC;

-- ステータス別集計
SELECT status                                                            AS "ステータス",
       COUNT(*)                                                          AS "件数",
       CONCAT(TO_CHAR(SUM(requested_amount), 'FM999,999,999,999'), '円') AS "総申請金額"
FROM support_applications
GROUP BY status
ORDER BY CASE status
             WHEN '審査中' THEN 1
             WHEN '書類確認中' THEN 2
             WHEN '承認済' THEN 3
             WHEN '支払処理中' THEN 4
             WHEN '完了' THEN 5
             WHEN '却下' THEN 6
             END;

-- 月別申請件数
SELECT TO_CHAR(application_date, 'YYYY-MM')                              AS "申請月",
       COUNT(*)                                                          AS "申請件数",
       CONCAT(TO_CHAR(SUM(requested_amount), 'FM999,999,999,999'), '円') AS "申請金額合計"
FROM support_applications
GROUP BY TO_CHAR(application_date, 'YYYY-MM')
ORDER BY "申請月" DESC;
