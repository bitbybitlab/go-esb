package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type BaseDirectoryModel struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
	UpdateTime time.Time `json:"updateTime" db:"update_time"`
	Version    int       `json:"version" db:"version"`
}
