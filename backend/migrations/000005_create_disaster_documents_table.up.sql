DROP TABLE IF EXISTS disaster_documents CASCADE;
CREATE TABLE IF NOT EXISTS disaster_documents
(
    id            SERIAL PRIMARY KEY,
    disaster_id   UUID         NOT NULL REFERENCES disasters (id) ON DELETE CASCADE,
    title         VARCHAR(255) NOT NULL,
    document_type VARCHAR(50)  NOT NULL,
    file_path     VARCHAR(500) NOT NULL,
    mime_type     VARCHAR(100) NOT NULL,
    description   TEXT,
    uploaded_by   VARCHAR(255) NOT NULL,
    is_public     BOOLEAN      NOT NULL DEFAULT FALSE,
    upload_date   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_disaster_documents_disaster_id ON disaster_documents (disaster_id);
CREATE INDEX IF NOT EXISTS idx_disaster_documents_document_type ON disaster_documents (document_type);
CREATE INDEX IF NOT EXISTS idx_disaster_documents_upload_date ON disaster_documents (upload_date);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE disaster_documents IS '災害関連書類管理テーブル - 各災害に関連する文書や画像ファイルを格納';
COMMENT ON COLUMN disaster_documents.id IS '書類ID - 主キー';
COMMENT ON COLUMN disaster_documents.disaster_id IS '災害ID - 関連する災害のID';
COMMENT ON COLUMN disaster_documents.title IS '書類タイトル - 文書の名称';
COMMENT ON COLUMN disaster_documents.document_type IS '書類種別 - 報告書, 写真, 申請書, 証明書など';
COMMENT ON COLUMN disaster_documents.file_path IS 'ファイルパス - ファイルの保存場所';
COMMENT ON COLUMN disaster_documents.mime_type IS 'MIMEタイプ - ファイルの形式を示すMIMEタイプ';
COMMENT ON COLUMN disaster_documents.description IS '説明 - ファイルの説明や備考';
COMMENT ON COLUMN disaster_documents.uploaded_by IS 'アップロード者 - ファイルをアップロードしたユーザー名';
COMMENT ON COLUMN disaster_documents.is_public IS '公開フラグ - 一般公開するかどうか';
COMMENT ON COLUMN disaster_documents.upload_date IS 'アップロード日時 - ファイルがアップロードされた日時';
COMMENT ON COLUMN disaster_documents.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN disaster_documents.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN disaster_documents.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- サンプルデータの挿入
DO
$$
    DECLARE
        disaster_id1 UUID;
        disaster_id2 UUID;
        disaster_id3 UUID;
        disaster_id4 UUID;
        disaster_id5 UUID;
    BEGIN
        -- 対象となる災害IDを取得（存在するデータから）
        SELECT id INTO disaster_id1 FROM disasters WHERE disaster_code = 'D2024-001' LIMIT 1;
        SELECT id INTO disaster_id2 FROM disasters WHERE disaster_code = 'D2024-002' LIMIT 1;
        SELECT id INTO disaster_id3 FROM disasters WHERE disaster_code = 'D2023-078' LIMIT 1;
        SELECT id INTO disaster_id4 FROM disasters WHERE disaster_code = 'D2024-005' LIMIT 1;
        SELECT id INTO disaster_id5 FROM disasters WHERE disaster_code = 'D2023-091' LIMIT 1;

        -- 令和6年7月豪雨に関する書類
        INSERT INTO disaster_documents (disaster_id, title, document_type, file_path, mime_type, description,
                                        uploaded_by, is_public, upload_date)
        VALUES (disaster_id1, '被害状況調査報告書', '報告書', '/documents/2024/07/D2024-001/damage_report_20240712.pdf',
                 'application/pdf', '農林水産省と東京都の合同調査チームによる被害状況の詳細報告',
                '農林水産課 田中', TRUE, '2024-07-12 17:30:00'),
               (disaster_id1, '冠水地域航空写真', '写真', '/documents/2024/07/D2024-001/aerial_photo_20240711.jpg',
                 'image/jpeg', 'ドローンで撮影した冠水した水田の航空写真', '災害対策室 佐藤', TRUE,
                '2024-07-11 10:15:00'),
               (disaster_id1, '復旧計画書', '計画書', '/documents/2024/07/D2024-001/recovery_plan_20240715.pdf',
                 'application/pdf', '被災地域の農地・施設の復旧に関する計画書', '復旧対策部 鈴木', FALSE,
                '2024-07-15 14:45:00'),
               (disaster_id1, '被災農家リスト', '台帳', '/documents/2024/07/D2024-001/affected_farmers_20240713.xlsx',
                 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
                '被災した農家のリストと被害状況の詳細', '農林水産課 山田', FALSE, '2024-07-13 16:20:00');

        -- 春季霜害に関する書類
        INSERT INTO disaster_documents (disaster_id, title, document_type, file_path,  mime_type, description,
                                        uploaded_by, is_public, upload_date)
        VALUES (disaster_id2, '霜害被害写真集', '写真', '/documents/2024/04/D2024-002/frost_damage_photos_20240419.zip',
                 'application/zip', '被害を受けたりんご園の写真集（圧縮ファイル）', '調査員 高橋', TRUE,
                '2024-04-19 16:30:00'),
               (disaster_id2, '専門家評価レポート', '報告書',
                '/documents/2024/04/D2024-002/expert_assessment_20240425.pdf',  'application/pdf',
                '農業気象学専門家による霜害被害の評価と今後の見通し', '青森県農業研究所 伊藤', TRUE,
                '2024-04-25 15:00:00'),
               (disaster_id2, '支援申請書類テンプレート', '申請書',
                '/documents/2024/04/D2024-002/support_application_template.docx',
                'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
                '被災農家が使用する支援申請書類のテンプレート', '農業支援課 中村', TRUE, '2024-04-26 09:15:00');

        -- 夏季干ばつに関する書類
        INSERT INTO disaster_documents (disaster_id, title, document_type, file_path,  mime_type, description,
                                        uploaded_by, is_public, upload_date)
        VALUES (disaster_id3, '干ばつ影響調査書', '報告書',
                '/documents/2023/08/D2023-078/drought_impact_assessment.pdf',  'application/pdf',
                '干ばつが農作物に与えた影響の詳細調査', '北海道農政部 小林', TRUE, '2023-08-20 11:30:00'),
               (disaster_id3, '地下水位データ', 'データ', '/documents/2023/08/D2023-078/groundwater_levels_202308.csv',
                 'text/csv', '干ばつ期間中の地域別地下水位の観測データ', '水資源管理課 加藤', FALSE,
                '2023-08-15 10:45:00'),
               (disaster_id3, '緊急灌水作業記録', '作業記録',
                '/documents/2023/08/D2023-078/emergency_irrigation_log.pdf',  'application/pdf',
                '自衛隊と地元農家による緊急灌水作業の実施記録', '災害対策本部 斎藤', TRUE, '2023-08-16 17:00:00');

        -- 台風による農業施設損壊に関する書類
        INSERT INTO disaster_documents (disaster_id, title, document_type, file_path,  mime_type, description,
                                        uploaded_by, is_public, upload_date)
        VALUES (disaster_id4, '台風被害状況写真', '写真', '/documents/2024/09/D2024-005/typhoon_damage_photos.zip',
                 'application/zip', '台風15号による農業施設の被害状況写真一式', '調査班 松本', TRUE,
                '2024-09-09 13:20:00'),
               (disaster_id4, '被害施設一覧', '台帳', '/documents/2024/09/D2024-005/damaged_facilities_list.xlsx',
                 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
                '被害を受けた農業施設の一覧と被害状況', '農業施設課 井上', FALSE, '2024-09-10 10:30:00'),
               (disaster_id4, '強風災害対策マニュアル', 'マニュアル',
                '/documents/2024/09/D2024-005/wind_damage_prevention_manual.pdf',  'application/pdf',
                '今後の台風に備えた農業施設の強風対策マニュアル', '防災対策課 清水', TRUE, '2024-09-15 15:45:00');

        -- 地滑りによる棚田被害に関する書類
        INSERT INTO disaster_documents (disaster_id, title, document_type, file_path,  mime_type, description,
                                        uploaded_by, is_public, upload_date)
        VALUES (disaster_id5, '地滑り発生前後の比較写真', '写真',
                '/documents/2023/07/D2023-091/landslide_before_after.pdf',  'application/pdf',
                '地滑り発生前と発生後の棚田状況の比較写真集', '地質調査班 林', TRUE, '2023-07-05 14:30:00'),
               (disaster_id5, '地質調査報告書', '報告書', '/documents/2023/07/D2023-091/geological_survey_report.pdf',
                 'application/pdf', '被災地域の地質状況と今後の地滑りリスク評価', '地質専門家 渡辺', FALSE,
                '2023-07-10 16:15:00'),
               (disaster_id5, '復旧工事計画書', '計画書', '/documents/2023/07/D2023-091/restoration_plan.pdf',
                'application/pdf', '棚田の復旧工事に関する詳細計画と工程表', '土木復旧班 中島', TRUE,
                '2023-07-20 11:45:00'),
               (disaster_id5, '文化的景観保全計画', '計画書',
                '/documents/2023/07/D2023-091/cultural_landscape_conservation.pdf',  'application/pdf',
                '棚田の文化的景観としての価値を保全するための復旧方針', '文化財保護課 木村', TRUE,
                '2023-07-25 13:30:00');
    END
$$;
