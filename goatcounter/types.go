package goatcounter

import (
	"time"
)

type Count struct {
}
type HitsResponse struct {
	Hits []Hit `json:"hits"`
}
type Hit struct {
	Count  int    `json:"count"`
	PathID int    `json:"path_id"`
	Path   string `json:"path"`
	Event  bool   `json:"event"`
	Title  string `json:"title"`
	Max    int    `json:"max"`
	Stats  []Stat `json:"stats"`
}
type PathsResponse struct {
	Paths []Path `json:"paths"`
	More  bool   `json:"more"`
}
type Path struct {
	Event bool   `json:"event"`
	ID    int    `json:"id"`
	Path  string `json:"path"`
	Title string `json:"title"`
}
type SitesResponse struct {
	Sites []Site `json:"sites"`
}
type Site struct {
	CNAME        string       `json:"cname"`
	CNAMESetupAt time.Time    `json:"cname_setup_at"`
	Code         string       `json:"code"`
	CreatedAt    time.Time    `json:"created_at"`
	FirstHitAt   time.Time    `json:"first_hit_at"`
	ID           int          `json:"id"`
	LinkDomain   string       `json:"link_domain"`
	Parent       int          `json:"parent"`
	ReceivedData bool         `json:"received_data"`
	Settings     SiteSettings `json:"settings"`
	State        string       `json:"state"`
	UpdatedAt    time.Time    `json:"updated_at"`
	UserDefaults UserSettings `json:"user_defaults"`
}
type SiteSettings struct {
	AllowBosmang   bool     `json:"allow_bosmang"`
	AllowCounter   bool     `json:"allow_counter"`
	AllowEmbed     []string `json:"allow_embed"`
	Collect        int      `json:"collect"`
	CollectRegions []string `json:"collect_regions"`
	DataRetention  int      `json:"data_retention"`
	IgnoreIPs      []string `json:"ignore_ips"`
	Public         string   `json:"public"`
	Secret         string   `json:"secret"`
}
type Stat struct {
	Day    string `json:"day"` // 2023-08-16
	Hourly []int  `json:"hourly"`
	Daily  int    `json:"daily"`
}
type Token struct {
	Name        string `json:"name"`
	Permissions int    `json:"permissions"`
}
type Total struct {
	Total       int `json:"total"`
	TotalEvents int `json:"total_events"`
	TotalUTC    int `json:"total_utc"`
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
	Settings      UserSettings      `json:"settings"`
	LastReportAt  time.Time         `json:"last_report_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	Token         Token             `json:"token"`
}
type UserSettings struct {
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
