-- 外部キー制約の追加

-- assessment_itemsテーブルの外部キー制約を追加
ALTER TABLE assessment_items
    ADD CONSTRAINT fk_assessment_items_facility_type
        FOREIGN KEY (facility_type_id) REFERENCES facility_types(id);

ALTER TABLE assessment_items
    ADD CONSTRAINT fk_assessment_items_damage_level
        FOREIGN KEY (damage_level_id) REFERENCES damage_levels(id);

-- 外部キー制約のインデックスを追加（パフォーマンス向上のため）
CREATE INDEX IF NOT EXISTS idx_assessment_items_facility_type_id ON assessment_items (facility_type_id);
CREATE INDEX IF NOT EXISTS idx_assessment_items_damage_level_id ON assessment_items (damage_level_id);

-- 必要に応じて他の外部キー制約も追加
