package models

import "time"

type Task struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`

	IsCompleted bool       `json:"is_completed"`
	CompletedAt *time.Time `json:"completed_at"`

	ReloadType  string `json:"reload_type"`  // daily | weekly | custom
	ReloadEvery int    `json:"reload_every"` // для custom (в днях)

	ResetTime    string `json:"reset_time"`    // "03:00", "12:00"
	ResetWeekday *int   `json:"reset_weekday"` // 0–6 (опционально, для weekly)

	LastDoneAt  *time.Time `json:"last_done_at"`
	NextResetAt time.Time  `json:"next_reset_at"`

	Order int `json:"order" gorm:"default:-1"`

	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ExeName  string `json:"exe_name"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
