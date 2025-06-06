// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFacilityType = "facility_types"

// FacilityType mapped from table <facility_types>
type FacilityType struct {
	ID          int32     `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:施設種別ID - 主キー" json:"id"`                                                         // 施設種別ID - 主キー
	Name        string    `gorm:"column:name;type:character varying(50);not null;index:idx_facility_types_name,priority:1;comment:施設種別名 - 水路, ため池, 農道, ビニールハウスなど" json:"name"` // 施設種別名 - 水路, ため池, 農道, ビニールハウスなど
	Description *string   `gorm:"column:description;type:text;comment:説明 - 施設種別の詳細説明" json:"description"`                                                                      // 説明 - 施設種別の詳細説明
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:作成日時 - レコード作成日時" json:"created_at"`                // 作成日時 - レコード作成日時
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:更新日時 - レコード最終更新日時" json:"updated_at"`              // 更新日時 - レコード最終更新日時
}

// TableName FacilityType's table name
func (*FacilityType) TableName() string {
	return TableNameFacilityType
}
