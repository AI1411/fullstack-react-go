-- 既存のENUM型を削除して再作成
DROP TYPE IF EXISTS disaster_status CASCADE;
CREATE TYPE disaster_status AS ENUM ('pending', 'under_review', 'in_progress', 'completed');

-- 既存テーブルを削除して再作成
DROP TABLE IF EXISTS disasters CASCADE;
CREATE TABLE IF NOT EXISTS disasters
(
    id                      UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    disaster_code           VARCHAR(20)              NOT NULL UNIQUE,
    name                    VARCHAR(100)             NOT NULL,
    prefecture_id           INT                      NOT NULL REFERENCES prefectures (id),
    occurred_at             TIMESTAMP WITH TIME ZONE NOT NULL,
    summary                 TEXT                     NOT NULL,
    disaster_type           VARCHAR(30)              NOT NULL,
    status                  disaster_status          NOT NULL DEFAULT 'pending',
    impact_level            VARCHAR(20)              NOT NULL,
    affected_area_size      DECIMAL(10, 2),
    estimated_damage_amount DECIMAL(15, 2),
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_disasters_prefecture_id ON disasters (prefecture_id);
CREATE INDEX idx_disasters_disaster_type ON disasters (disaster_type);
CREATE INDEX idx_disasters_status ON disasters (status);
CREATE INDEX idx_disasters_occurred_at ON disasters (occurred_at);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE disasters IS '農業災害情報管理テーブル - 各種災害の詳細情報を格納';
COMMENT ON COLUMN disasters.id IS '災害ID - 主キー';
COMMENT ON COLUMN disasters.disaster_code IS '災害コード - 管理用の一意識別子 (例: D2024-001)';
COMMENT ON COLUMN disasters.name IS '災害名 - 災害の名称';
COMMENT ON COLUMN disasters.prefecture_id IS '都道府県ID - 災害が発生した都道府県のID';
COMMENT ON COLUMN disasters.occurred_at IS '発生日時 - 災害が発生した日時';
COMMENT ON COLUMN disasters.summary IS '被害概要 - 災害による被害の詳細説明';
COMMENT ON COLUMN disasters.disaster_type IS '災害種別 - 洪水, 地滑り, 雹害, 干ばつ, 風害, 地震, 霜害, 病害虫など';
COMMENT ON COLUMN disasters.status IS '状態 - pending(未着手), under_review(審査中), in_progress(対応中), completed(完了)のいずれか';
COMMENT ON COLUMN disasters.impact_level IS '被害レベル - 軽微, 中程度, 深刻, 甚大などの被害度合い';
COMMENT ON COLUMN disasters.affected_area_size IS '被害面積 - ヘクタール (ha) 単位での被害エリアの広さ';
COMMENT ON COLUMN disasters.estimated_damage_amount IS '被害推定金額 - 円単位での被害総額';
COMMENT ON COLUMN disasters.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN disasters.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN disasters.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

-- 農業災害テストデータの挿入
-- 前提: prefecturesテーブルに都道府県データが存在することを想定

INSERT INTO disasters (disaster_code,
                       name,
                       prefecture_id,
                       occurred_at,
                       summary,
                       disaster_type,
                       status,
                       impact_level,
                       affected_area_size,
                       estimated_damage_amount,
                       created_at,
                       updated_at)
VALUES
-- 2024年の災害データ
('D2024-001',
 '令和6年7月豪雨による水稲被害',
 13, -- 東京都
 '2024-07-10 14:30:00+09',
 '集中豪雨により水田が冠水し、出穂期の水稲に甚大な被害が発生。約200ヘクタールの水田で倒伏・流失が確認された。',
 '洪水',
 'completed',
 '甚大',
 200.50,
 150000000.00,
 '2024-07-11 09:00:00+09',
 '2024-08-15 16:30:00+09'),
('D2024-002',
 '春季霜害によるりんご園被害',
 2, -- 青森県
 '2024-04-18 03:45:00+09',
 '開花期の異常低温により、県内主要産地のりんご園で花芽・幼果に霜害が発生。収穫量の大幅な減少が予想される。',
 '霜害',
 'in_progress',
 '深刻',
 350.75,
 280000000.00,
 '2024-04-18 08:00:00+09',
 '2024-05-20 14:20:00+09'),
('D2024-003',
 '梅雨期長雨による野菜価格高騰',
 14, -- 神奈川県
 '2024-06-15 00:00:00+09',
 '梅雨前線の停滞により長期間の降雨が続き、露地野菜の生育不良と病害が多発。特にキャベツ、レタスの出荷量が大幅減少。',
 '長雨',
 'under_review',
 '中程度',
 120.30,
 45000000.00,
 '2024-06-20 10:15:00+09',
 '2024-06-25 11:45:00+09'),
('D2024-004',
 'ひょう害によるぶどう栽培被害',
 19, -- 山梨県
 '2024-05-28 16:20:00+09',
 '局地的な雹の襲来により、ぶどう栽培地域で葉や幼果に物理的損傷が発生。特に甲州種の被害が深刻。',
 '雹害',
 'in_progress',
 '深刻',
 85.20,
 120000000.00,
 '2024-05-28 18:30:00+09',
 '2024-06-10 09:20:00+09'),
('D2024-005',
 '台風15号による農業施設損壊',
 12, -- 千葉県
 '2024-09-08 22:15:00+09',
 '台風15号の強風によりビニールハウスや農業倉庫が損壊。施設園芸作物への直接被害も確認されている。',
 '風害',
 'pending',
 '甚大',
 45.80,
 200000000.00,
 '2024-09-09 07:00:00+09',
 '2024-09-09 07:00:00+09'),
-- 2023年の災害データ
('D2023-078',
 '夏季干ばつによる畑作物被害',
 1, -- 北海道
 '2023-08-05 12:00:00+09',
 '記録的な少雨により土壌水分が不足し、じゃがいも、とうもろこし等の畑作物に深刻な生育障害が発生。',
 '干ばつ',
 'completed',
 '深刻',
 1250.00,
 350000000.00,
 '2023-08-10 14:30:00+09',
 '2023-10-31 17:00:00+09'),
('D2023-089',
 'いもち病大発生による稲作被害',
 15, -- 新潟県
 '2023-07-22 00:00:00+09',
 '高温多湿な気象条件下でいもち病が大発生し、県内の水稲作付面積の約15%に被害が拡大。',
 '病害虫',
 'completed',
 '中程度',
 450.25,
 95000000.00,
 '2023-07-25 08:45:00+09',
 '2023-09-15 16:20:00+09'),
('D2023-091',
 '地滑りによる棚田被害',
 17, -- 石川県
 '2023-06-30 04:20:00+09',
 '豪雨による地盤の緩みで山間部の棚田において地滑りが発生。貴重な文化的景観と農地が失われた。',
 '地滑り',
 'completed',
 '深刻',
 12.50,
 25000000.00,
 '2023-06-30 09:00:00+09',
 '2023-08-20 15:30:00+09'),
('D2023-045',
 '春季低温による茶葉品質低下',
 22, -- 静岡県
 '2023-04-12 06:00:00+09',
 '新茶の摘採期に異常低温が継続し、茶葉の生育遅延と品質低下が発生。一番茶の収量・品質ともに例年を大きく下回る。',
 '低温',
 'completed',
 '中程度',
 180.40,
 80000000.00,
 '2023-04-15 11:20:00+09',
 '2023-06-10 14:50:00+09'),
-- 軽微な被害の例
('D2024-006',
 '局地的豪雨による小規模冠水',
 11, -- 埼玉県
 '2024-08-20 15:45:00+09',
 '短時間の集中豪雨により一部農地で冠水が発生したが、排水対応により被害は最小限に抑制された。',
 '洪水',
 'completed',
 '軽微',
 8.30,
 2500000.00,
 '2024-08-20 18:00:00+09',
 '2024-08-22 10:30:00+09'),
('D2024-007',
 'コナガ発生による白菜軽微被害',
 27, -- 大阪府
 '2024-10-05 00:00:00+09',
 '秋冬野菜の生育期にコナガの発生が確認されたが、早期防除により被害は軽微に留まった。',
 '病害虫',
 'in_progress',
 '軽微',
 25.60,
 8000000.00,
 '2024-10-08 09:30:00+09',
 '2024-10-15 14:20:00+09'),
-- 審査中・未着手の災害
('D2024-008',
 '秋季台風による果樹園被害報告',
 34, -- 広島県
 '2024-10-12 08:30:00+09',
 '台風による強風で柿、みかん等の果樹に落果被害が発生した可能性があり、現在被害状況を調査中。',
 '風害',
 'under_review',
 '中程度',
 95.20,
 NULL, -- 査定中のため金額未確定
 '2024-10-15 10:00:00+09',
 '2024-10-20 16:45:00+09'),
('D2024-009',
 '病害虫発生報告（詳細調査待ち）',
 46, -- 鹿児島県
 '2024-11-01 00:00:00+09',
 'さつまいも栽培地域で新たな病害の発生が疑われており、専門機関による詳細な調査が必要な状況。',
 '病害虫',
 'pending',
 '不明',
 NULL, -- 調査前のため面積未確定
 NULL, -- 調査前のため金額未確定
 '2024-11-05 13:20:00+09',
 '2024-11-05 13:20:00+09');