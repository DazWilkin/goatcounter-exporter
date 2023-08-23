package goatcounter

import (
	"time"
)

type Hit struct {
	Count  int    `json:"count"`
	PathID int    `json:"path_id"`
	Path   string `json:"path"`
	Event  bool   `json:"event"`
	Title  string `json:"title"`
	Max    int    `json:"max"`
	Stats  []Stat `json:"stats"`
}
type Settings struct {
	TwentyFourHours       bool      `json:"twenty_four_hours"`
	SundayStartsWeek      bool      `json:"sunday_starts_week"`
	Language              string    `json:"language"`
	DateFormat            string    `json:"date_format"`
	NumberFormat          int       `json:"number_format"`
	TimeZone              string    `json:"timezone"`
	Widgets               []Widget  `json:"widgets"`
	Views                 []View    `json:"views"`
	EmailReports          int       `json:"email_reports"`
	FewerNumbers          bool      `json:"fewer_numbers"`
	FewerNumbersLockUntil time.Time `json:"fewer_numbers_lock_until"`
}
type Stat struct {
	Day    string `json:"day"` // 2023-08-16
	Hourly []int  `json:"hourly"`
	Daily  int    `json:"daily"`
}
type StatsHits struct {
	Hits []Hit `json:"hits"`
}
type StatsTotal struct {
	Total       int `json:"total"`
	TotalEvents int `json:"total_events"`
	TotalUTC    int `json:"total_utc"`
}
type Token struct {
	Name        string `json:"name"`
	Permissions int    `json:"permissions"`
}
type User struct {
	ID            int               `json:"id"`
	Site          int               `json:"site"`
	Email         string            `json:"email"`
	EmailVerified bool              `json:"email_verified"`
	TOTPVerified  bool              `json:"totp_verified"`
	Access        map[string]string `json:"access"`
	LoginAt       time.Time         `json:"login_at"`
	OpenAt        time.Time         `json:"open_at"`
	ResetAt       time.Time         `json:"reset_at,omitempty"`
	Settings      Settings          `json:"settings"`
	LastReportAt  time.Time         `json:"last_report_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Token         Token             `json:"token"`
}
type View struct {
	Name   string `json:"name"`
	Filter string `json:"filter"`
	Daily  bool   `json:"daily"`
	Period string `json:"period"`
}
type Widget struct {
	N string                 `json:"n"`
	S map[string]interface{} `json:"s"`
}
