// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameFacilityEquipment = "facility_equipment"

// FacilityEquipment mapped from table <facility_equipment>
type FacilityEquipment struct {
	ID                  int32          `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:施設設備ID - 主キー" json:"id"`                                                                     // 施設設備ID - 主キー
	Name                string         `gorm:"column:name;type:character varying(100);not null;index:idx_facility_equipment_name,priority:1;comment:施設設備名 - 施設設備の名称" json:"name"`                       // 施設設備名 - 施設設備の名称
	FacilityTypeID      int32          `gorm:"column:facility_type_id;type:integer;not null;index:idx_facility_equipment_facility_type_id,priority:1;comment:施設種別ID - 施設設備の種別" json:"facility_type_id"` // 施設種別ID - 施設設備の種別
	ModelNumber         *string        `gorm:"column:model_number;type:character varying(100);comment:型番 - 施設設備の型番" json:"model_number"`                                                                // 型番 - 施設設備の型番
	Manufacturer        *string        `gorm:"column:manufacturer;type:character varying(100);comment:メーカー - 施設設備のメーカー" json:"manufacturer"`                                                            // メーカー - 施設設備のメーカー
	InstallationDate    *time.Time     `gorm:"column:installation_date;type:date;comment:設置日 - 施設設備の設置日" json:"installation_date"`                                                                      // 設置日 - 施設設備の設置日
	Status              string         `gorm:"column:status;type:character varying(50);not null;index:idx_facility_equipment_status,priority:1;default:稼働中;comment:状態 - 施設設備の稼働状態" json:"status"`       // 状態 - 施設設備の稼働状態
	LocationDescription *string        `gorm:"column:location_description;type:text;comment:設置場所説明 - 施設設備の設置場所の説明" json:"location_description"`                                                         // 設置場所説明 - 施設設備の設置場所の説明
	LocationLatitude    *float64       `gorm:"column:location_latitude;type:numeric(10,8);comment:位置（緯度） - 施設設備の緯度" json:"location_latitude"`                                                           // 位置（緯度） - 施設設備の緯度
	LocationLongitude   *float64       `gorm:"column:location_longitude;type:numeric(11,8);comment:位置（経度） - 施設設備の経度" json:"location_longitude"`                                                         // 位置（経度） - 施設設備の経度
	Notes               *string        `gorm:"column:notes;type:text;comment:備考 - 施設設備に関する備考やメモ" json:"notes"`                                                                                          // 備考 - 施設設備に関する備考やメモ
	CreatedAt           time.Time      `gorm:"column:created_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:作成日時 - レコード作成日時" json:"created_at"`                            // 作成日時 - レコード作成日時
	UpdatedAt           time.Time      `gorm:"column:updated_at;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP;comment:更新日時 - レコード最終更新日時" json:"updated_at"`                          // 更新日時 - レコード最終更新日時
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp with time zone;comment:削除日時 - 論理削除用のタイムスタンプ" json:"deleted_at"`                                                          // 削除日時 - 論理削除用のタイムスタンプ
	FacilityType        FacilityType   `json:"facility_type"`
}

// TableName FacilityEquipment's table name
func (*FacilityEquipment) TableName() string {
	return TableNameFacilityEquipment
}
