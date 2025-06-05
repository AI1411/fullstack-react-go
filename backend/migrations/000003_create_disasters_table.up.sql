-- 既存のENUM型を削除して再作成
DROP TYPE IF EXISTS disaster_status CASCADE;
CREATE TYPE disaster_status AS ENUM ('pending', 'under_review', 'in_progress', 'completed');

-- 既存テーブルを削除して再作成
DROP TABLE IF EXISTS disasters CASCADE;
CREATE TABLE IF NOT EXISTS disasters
(
    id                      UUID PRIMARY KEY                  DEFAULT gen_random_uuid(),
    name                    VARCHAR(100)             NOT NULL,
    municipality_id         INT                      NOT NULL REFERENCES municipalities (id),
    occurred_at             TIMESTAMP WITH TIME ZONE NOT NULL,
    summary                 TEXT                     NOT NULL,
    work_category_id        BIGINT                   NOT NULL,
    status                  disaster_status          NOT NULL DEFAULT 'pending',
    affected_area_size      DECIMAL(10, 2),
    estimated_damage_amount DECIMAL(15, 2),
    latitude                DECIMAL(10, 8)           NULL,
    longitude               DECIMAL(11, 8)           NULL,
    address                 TEXT                     NULL,
    place_id                VARCHAR(255)             NULL,
    created_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at              TIMESTAMP WITH TIME ZONE
);

-- インデックス作成
CREATE INDEX IF NOT EXISTS idx_disasters_municipality_id ON disasters (municipality_id);
CREATE INDEX IF NOT EXISTS idx_disasters_work_category_id ON disasters (work_category_id);
CREATE INDEX IF NOT EXISTS idx_disasters_status ON disasters (status);
CREATE INDEX IF NOT EXISTS idx_disasters_occurred_at ON disasters (occurred_at);
CREATE INDEX IF NOT EXISTS idx_disasters_latitude_longitude ON disasters (latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_disasters_place_id ON disasters (place_id);

-- pg_bigm インデックス作成
CREATE INDEX IF NOT EXISTS idx_disaster_reports_title_bigm ON disasters USING gin (name gin_bigm_ops);
CREATE INDEX IF NOT EXISTS idx_disaster_reports_description_bigm ON disasters USING gin (summary gin_bigm_ops);

-- テーブルとカラムにコメントを追加
COMMENT ON TABLE disasters IS '農業災害情報管理テーブル - 各種災害の詳細情報を格納';
COMMENT ON COLUMN disasters.id IS '災害ID - 主キー';
COMMENT ON COLUMN disasters.name IS '災害名 - 災害の名称';
COMMENT ON COLUMN disasters.municipality_id IS '自治体ID - 関連する自治体のID';
COMMENT ON COLUMN disasters.occurred_at IS '発生日時 - 災害が発生した日時';
COMMENT ON COLUMN disasters.summary IS '概要 - 災害の簡単な説明';
COMMENT ON COLUMN disasters.work_category_id IS '工種区分ID - 関連する作業カテゴリのID';
COMMENT ON COLUMN disasters.summary IS '被害概要 - 災害による被害の詳細説明';
COMMENT ON COLUMN disasters.status IS '状態 - pending(未着手), under_review(審査中), in_progress(対応中), completed(完了)のいずれか';
COMMENT ON COLUMN disasters.affected_area_size IS '被害面積 - ヘクタール (ha) 単位での被害エリアの広さ';
COMMENT ON COLUMN disasters.estimated_damage_amount IS '被害推定金額 - 円単位での被害総額';
COMMENT ON COLUMN disasters.latitude IS '緯度 - 災害発生地点の緯度座標';
COMMENT ON COLUMN disasters.longitude IS '経度 - 災害発生地点の経度座標';
COMMENT ON COLUMN disasters.address IS '住所 - Google Maps APIから取得した住所情報';
COMMENT ON COLUMN disasters.place_id IS 'Google Place ID - Google Maps APIの場所識別子';
COMMENT ON COLUMN disasters.created_at IS '作成日時 - レコード作成日時';
COMMENT ON COLUMN disasters.updated_at IS '更新日時 - レコード最終更新日時';
COMMENT ON COLUMN disasters.deleted_at IS '削除日時 - 論理削除用のタイムスタンプ';

INSERT INTO disasters (name, municipality_id, occurred_at, summary, work_category_id, status, affected_area_size, estimated_damage_amount, latitude, longitude, address, place_id)
VALUES
    ('台風による河川氾濫', 1, '2023-08-15 09:30:00+09', '台風15号により住宅25棟が浸水被害を受けました。避難者数は約150名です。', 1, 'pending', 120.50, 15000000.00, 35.6762, 139.6503, '東京都渋谷区道玄坂1-2-3', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('集中豪雨による土砂災害', 2, '2023-07-22 14:15:00+09', '集中豪雨の影響で道路が通行止めとなり、復旧作業が進められています。', 2, 'in_progress', 85.30, 8500000.00, 34.6937, 135.5023, '大阪府大阪市北区梅田2-4-9', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('地震による建物倒壊', 3, '2023-09-10 03:45:00+09', '震度6弱の地震により、電力供給が停止し、約800世帯が影響を受けています。', 3, 'completed', 200.00, 25000000.00, 35.1815, 136.9066, '愛知県名古屋市中村区名駅1-1-1', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('大雪による道路封鎖', 4, '2024-01-08 06:20:00+09', '記録的大雪のため、120名の住民が避難所に避難しています。', 1, 'under_review', 150.75, 12000000.00, 43.0642, 141.3469, '北海道札幌市中央区北1条西2-1', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('竜巻による被害', 5, '2023-06-18 16:50:00+09', '竜巻により農地約300㎡が被害を受け、農作物に甚大な被害が発生しました。', 2, 'pending', 300.25, 5000000.00, 35.4437, 139.6380, '神奈川県横浜市西区高島2-19-12', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('雷雨による停電', 6, '2023-08-03 19:30:00+09', '落雷の影響で、公共施設5箇所が使用不能となっています。', 3, 'in_progress', 75.00, 3500000.00, 34.3853, 132.4553, '広島県広島市中区基町9-42', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('洪水による浸水', 7, '2023-07-30 11:45:00+09', '豪雨により、上下水道施設に損傷が発生し、復旧作業を実施中です。', 1, 'completed', 180.60, 18000000.00, 34.2259, 131.6106, '山口県下関市豊前田町3-3-1', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('地滑りによる道路寸断', 8, '2023-10-15 08:15:00+09', '地滑りのため、学校8校が臨時休校となりました。', 2, 'pending', 95.40, 7500000.00, 36.2048, 138.2529, '長野県長野市大字長野元善町462', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('豪雨による橋梁損傷', 9, '2023-09-25 13:20:00+09', '台風により住宅45棟が被害を受けました。避難者数は約280名です。', 3, 'under_review', 220.80, 22000000.00, 33.5904, 130.4017, '福岡県福岡市博多区博多駅前2-1-1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('強風による倒木', 10, '2023-05-12 07:40:00+09', '強風の影響で道路が通行止めとなり、復旧作業が進められています。', 1, 'in_progress', 65.20, 4200000.00, 38.2682, 140.8694, '宮城県仙台市青葉区中央1-3-1', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('雹による農作物被害', 11, '2023-06-05 15:30:00+09', '雹により、電力供給が停止し、約450世帯が影響を受けています。', 2, 'completed', 110.30, 9800000.00, 36.5658, 136.6586, '石川県金沢市広坂1-1-1', 'ChIJJWpMaHmmGGARxQdRpIRpnNI'),
    ('津波による沿岸被害', 12, '2023-11-08 04:25:00+09', '津波のため、85名の住民が避難所に避難しています。', 3, 'pending', 350.00, 35000000.00, 39.7036, 141.1527, '岩手県盛岡市内丸10-1', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('火山灰による交通麻痺', 13, '2023-04-20 12:10:00+09', '火山灰により農地約180㎡が被害を受け、農作物に甚大な被害が発生しました。', 1, 'under_review', 180.45, 13500000.00, 31.5966, 130.5571, '鹿児島県鹿児島市山下町11-1', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('凍結による水道管破裂', 14, '2024-01-20 05:55:00+09', '凍結の影響で、公共施設12箇所が使用不能となっています。', 2, 'in_progress', 90.70, 6800000.00, 40.8244, 140.7400, '秋田県秋田市山王4-1-1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('高波による堤防決壊', 15, '2023-10-30 17:35:00+09', '高波により、上下水道施設に損傷が発生し、復旧作業を実施中です。', 3, 'completed', 280.20, 28000000.00, 26.2124, 127.6792, '沖縄県那覇市泉崎1-2-2', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('落雷による火災', 16, '2023-07-14 20:15:00+09', '落雷のため、学校6校が臨時休校となりました。', 1, 'pending', 55.80, 3800000.00, 35.4437, 133.0504, '鳥取県鳥取市東町1-220', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('ゲリラ豪雨による冠水', 17, '2023-08-28 14:45:00+09', 'ゲリラ豪雨により住宅38棟が被害を受けました。避難者数は約200名です。', 2, 'under_review', 160.90, 16000000.00, 34.6565, 133.9195, '岡山県岡山市北区内山下2-4-6', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('暴風雨による屋根損傷', 18, '2023-09-12 10:20:00+09', '暴風雨の影響で道路が通行止めとなり、復旧作業が進められています。', 3, 'in_progress', 125.40, 11200000.00, 35.6586, 139.7454, '東京都台東区上野公園5-20', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('大雨による河川決壊', 19, '2023-06-25 09:30:00+09', '大雨により、電力供給が停止し、約680世帯が影響を受けています。', 1, 'completed', 240.60, 24000000.00, 33.2382, 130.2988, '熊本県熊本市中央区水前寺6-18-1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('積雪による建物倒壊', 20, '2024-02-05 06:40:00+09', '大雪のため、95名の住民が避難所に避難しています。', 2, 'pending', 140.20, 14500000.00, 36.6953, 137.2113, '富山県富山市新桜町7-38', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('台風による農地被害', 21, '2023-09-18 11:25:00+09', '台風により農地約420㎡が冠水し、農作物に甚大な被害が発生しました。', 3, 'under_review', 420.80, 18500000.00, 34.2259, 134.2491, '徳島県徳島市万代町1-1', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('地震による道路陥没', 22, '2023-08-07 02:15:00+09', '地震の影響で、公共施設9箇所が使用不能となっています。', 1, 'in_progress', 105.30, 8900000.00, 33.8416, 132.7656, '愛媛県松山市一番町4-4-2', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('豪雨による土石流', 23, '2023-07-09 16:50:00+09', '土石流により、上下水道施設に損傷が発生し、復旧作業を実施中です。', 2, 'completed', 195.70, 19800000.00, 33.5597, 133.5311, '高知県高知市丸ノ内1-2-20', 'ChIJXWfMj-eMGGARlp2RGYpjnNI'),
    ('強風による電線切断', 24, '2023-05-30 18:40:00+09', '強風のため、学校11校が臨時休校となりました。', 3, 'pending', 78.90, 5600000.00, 34.3402, 132.4594, '広島県呉市中央4-1-6', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('雹嵐による車両損傷', 25, '2023-06-12 13:15:00+09', '雹嵐により住宅52棟が被害を受けました。避難者数は約320名です。', 1, 'under_review', 185.40, 17800000.00, 36.0614, 140.1239, '茨城県水戸市笠原町978-6', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('洪水による橋脚損傷', 26, '2023-08-20 08:05:00+09', '洪水の影響で道路が通行止めとなり、復旧作業が進められています。', 2, 'in_progress', 260.50, 26500000.00, 36.5658, 138.1927, '群馬県前橋市大手町1-1-1', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('竜巻による施設破損', 27, '2023-07-28 15:20:00+09', '竜巻により、電力供給が停止し、約520世帯が影響を受けています。', 3, 'completed', 135.80, 12800000.00, 36.3944, 139.0608, '埼玉県さいたま市浦和区高砂3-15-1', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('大雨による地盤沈下', 28, '2023-09-05 12:35:00+09', '大雨のため、72名の住民が避難所に避難しています。', 1, 'pending', 98.60, 7200000.00, 35.6051, 140.1233, '千葉県千葉市中央区中央3-10-8', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('台風による看板倒壊', 29, '2023-10-02 19:50:00+09', '台風により農地約280㎡が被害を受け、農作物に甚大な被害が発生しました。', 2, 'under_review', 280.30, 15200000.00, 35.4478, 139.6425, '神奈川県川崎市川崎区宮本町1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('雷雨による通信障害', 30, '2023-06-08 21:10:00+09', '雷雨の影響で、公共施設7箇所が使用不能となっています。', 3, 'in_progress', 88.20, 6100000.00, 35.1815, 136.9066, '愛知県豊田市西町3-60', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('強風による屋根瓦飛散', 31, '2023-04-15 14:30:00+09', '強風により、上下水道施設に損傷が発生し、復旧作業を実施中です。', 1, 'completed', 112.90, 9600000.00, 34.7369, 135.5023, '大阪府堺市北区百舌鳥赤畑町1-1-1', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('集中豪雨による河川増水', 32, '2023-08-11 07:45:00+09', '集中豪雨のため、学校14校が臨時休校となりました。', 2, 'pending', 205.70, 20900000.00, 34.6851, 135.8048, '奈良県奈良市登大路町30', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('地震による液状化', 33, '2023-11-14 04:55:00+09', '地震により住宅61棟が被害を受けました。避難者数は約380名です。', 3, 'under_review', 310.40, 31500000.00, 34.2261, 135.1675, '和歌山県和歌山市小松原通1-1', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('豪雪による交通麻痺', 34, '2024-01-12 09:15:00+09', '豪雪の影響で道路が通行止めとなり、復旧作業が進められています。', 1, 'in_progress', 175.60, 16800000.00, 37.9026, 139.0235, '新潟県新潟市中央区学校町通1番町602-1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('台風による電柱倒壊', 35, '2023-09-28 16:25:00+09', '台風により、電力供給が停止し、約750世帯が影響を受けています。', 2, 'completed', 155.80, 14200000.00, 35.2689, 139.1072, '静岡県静岡市葵区追手町9-6', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('雹による温室破損', 36, '2023-05-25 13:40:00+09', '雹のため、128名の住民が避難所に避難しています。', 3, 'pending', 142.30, 11900000.00, 35.3912, 136.7223, '岐阜県岐阜市薮田南2-1-1', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('洪水による農機具損失', 37, '2023-07-18 10:30:00+09', '洪水により農地約350㎡が冠水し、農作物に甚大な被害が発生しました。', 1, 'under_review', 350.90, 17600000.00, 35.0116, 135.7681, '京都府京都市上京区下立売通新町西入薮ノ内町', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('竜巻による工場損傷', 38, '2023-08-16 17:55:00+09', '竜巻の影響で、公共施設13箇所が使用不能となっています。', 2, 'in_progress', 225.70, 22800000.00, 34.6851, 133.9195, '三重県津市広明町13', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('地滑りによる家屋倒壊', 39, '2023-10-08 06:20:00+09', '地滑りにより、上下水道施設に損傷が発生し、復旧作業を実施中です。', 3, 'completed', 190.40, 19200000.00, 35.5036, 134.2384, '滋賀県大津市京町3-1-1', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('強風による広告塔転倒', 40, '2023-06-20 19:15:00+09', '強風のため、学校9校が臨時休校となりました。', 1, 'pending', 102.50, 8400000.00, 35.5031, 134.2384, '兵庫県神戸市中央区加納町6-5-1', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('豪雨による用水路氾濫', 41, '2023-07-05 11:50:00+09', '豪雨により住宅74棟が被害を受けました。避難者数は約460名です。', 2, 'under_review', 268.80, 26200000.00, 35.1815, 139.7673, '山梨県甲府市丸の内1-6-1', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('台風による港湾施設破損', 42, '2023-09-22 08:35:00+09', '台風の影響で道路が通行止めとなり、復旧作業が進められています。', 3, 'in_progress', 320.60, 32500000.00, 35.1815, 139.7673, '静岡県浜松市中区元城町103-2', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('雷雨による変電所故障', 43, '2023-08-24 20:25:00+09', '雷雨により、電力供給が停止し、約920世帯が影響を受けています。', 1, 'completed', 145.30, 13800000.00, 37.7608, 138.9470, '福井県福井市大手3-17-1', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('地震による石垣崩落', 44, '2023-11-01 03:10:00+09', '地震のため、156名の住民が避難所に避難しています。', 2, 'pending', 178.90, 16500000.00, 35.6762, 139.6503, '長野県松本市丸の内3-7', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('豪雨による下水道逆流', 45, '2023-06-30 15:05:00+09', '豪雨により農地約290㎡が被害を受け、農作物に甚大な被害が発生しました。', 3, 'under_review', 290.70, 15800000.00, 36.2048, 138.2529, '山形県山形市松波2-8-1', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4'),
    ('強風による送電線断線', 46, '2023-05-18 12:45:00+09', '強風の影響で、公共施設16箇所が使用不能となっています。', 1, 'in_progress', 98.40, 7800000.00, 38.2682, 140.8694, '福島県福島市杉妻町2-16', 'ChIJJWpMaHmmGGARxQdffIRpnNI'),
    ('雹による太陽光パネル破損', 47, '2023-07-12 14:20:00+09', '雹により、上下水道施設に損傷が発生し、復旧作業を実施中です。', 2, 'completed', 167.50, 14900000.00, 40.8244, 140.7400, '青森県青森市長島1-1-1', 'ChIJf7NMtNJ8GGARSw1lp2XjnNI'),
    ('洪水による河川護岸崩壊', 48, '2023-08-29 09:55:00+09', '洪水のため、学校18校が臨時休校となりました。', 3, 'pending', 385.20, 38800000.00, 33.8416, 130.4017, '佐賀県佐賀市城内1-1-59', 'ChIJXWfMj-eMGGARlp2FGYpjnNI'),
    ('台風による海岸侵食', 49, '2023-10-12 18:30:00+09', '台風により住宅89棟が被害を受けました。避難者数は約510名です。', 1, 'under_review', 425.60, 42800000.00, 32.7503, 129.8777, '長崎県長崎市江戸町2-13', 'ChIJoRyG2-eMGGARkVcKPYc7ZU4'),
    ('地震による斜面崩壊', 50, '2023-11-20 05:40:00+09', '地震の影響で道路が通行止めとなり、復旧作業が進められています。', 2, 'in_progress', 295.80, 29200000.00, 33.2382, 131.6128, '大分県大分市大手町3-1-1', 'ChIJLfyY2x1XAzsR6Z8g9k4Qy4');
