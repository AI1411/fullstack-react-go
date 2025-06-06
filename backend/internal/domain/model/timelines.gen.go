// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTimeline = "timelines"

// Timeline mapped from table <timelines>
type Timeline struct {
	ID          int32          `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;comment:タイムラインID - 主キー" json:"id"`                                                                  // タイムラインID - 主キー
	DisasterID  string         `gorm:"column:disaster_id;type:uuid;not null;index:idx_timelines_disaster_id,priority:1;comment:災害ID - 関連する災害のID" json:"disaster_id"`                           // 災害ID - 関連する災害のID
	EventName   string         `gorm:"column:event_name;type:character varying(255);not null;comment:イベント名 - 発生したイベントの名称" json:"event_name"`                                                   // イベント名 - 発生したイベントの名称
	EventTime   time.Time      `gorm:"column:event_time;type:timestamp without time zone;not null;index:idx_timelines_event_time,priority:1;comment:イベント発生日時 - イベントが発生した日時" json:"event_time"` // イベント発生日時 - イベントが発生した日時
	Description string         `gorm:"column:description;type:text;not null;comment:イベント説明 - イベントの詳細な説明" json:"description"`                                                                   // イベント説明 - イベントの詳細な説明
	Severity    *string        `gorm:"column:severity;type:character varying(50);comment:イベントの深刻度 - 低, 中, 高などの深刻度" json:"severity"`                                                            // イベントの深刻度 - 低, 中, 高などの深刻度
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP;comment:作成日時 - レコード作成日時" json:"created_at"`                        // 作成日時 - レコード作成日時
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp without time zone;not null;default:CURRENT_TIMESTAMP;comment:更新日時 - レコード最終更新日時" json:"updated_at"`                      // 更新日時 - レコード最終更新日時
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone;comment:削除日時 - 論理削除用のタイムスタンプ" json:"deleted_at"`                                                      // 削除日時 - 論理削除用のタイムスタンプ
	Disaster    Disaster       `json:"disaster"`
}

// TableName Timeline's table name
func (*Timeline) TableName() string {
	return TableNameTimeline
}
