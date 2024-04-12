package models

import (
	"encoding/json"
	"time"
)

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index;type:datetime;default:null" json:"deleted_at"`
}

func (m Model) MarshalJSON() ([]byte, error) {
	type Alias Model
	if m.DeletedAt.IsZero() {
		return json.Marshal(&struct {
			Alias
			DeletedAt interface{} `json:"deleted_at"`
		}{
			Alias:     (Alias)(m),
			DeletedAt: nil,
		})
	}
	return json.Marshal(&struct {
		Alias
		DeletedAt time.Time `json:"deleted_at,omitempty"`
	}{
		Alias:     (Alias)(m),
		DeletedAt: m.DeletedAt,
	})
}
