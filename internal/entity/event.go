package entity

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/photoprism/photoprism/pkg/rnd"
)

// Events
type Event struct {
	ID         string       `gorm:"type:varbinary(36);primary_key;auto_increment:false;" json:"id"`
	Location   string       `gorm:"type:varbinary(16);unique_index;" json:"location"`
	EventName  string       `gorm:"type:varchar(200);" json:"name"`
	EventType  string       `gorm:"type:varbinary(50);" json:"type"`
	EventInfo  string       `gorm:"type:varbinary(2000);" json:"info"`
	EventUrl   string       `gorm:"type:varbinary(200);" json:"url"`
	EventStart time.Time    `gorm:"type:datetime;index;" json:"start"`
	EventEnd   sql.NullTime `gorm:"type:datetime;" json:"end"`
	UpdatedAt  time.Time    `json:"updated"`
}

func (Event) TableName() string {
	return "events"
}

func (e *Event) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", rnd.PPID('e'))
}
