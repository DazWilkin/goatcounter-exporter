package goatcounter

type Endpoint uint8

func (m Endpoint) String() string {
	switch m {
	case Count:
		return "count"
	case Export:
		return "export"
	case Me:
		return "me"
	case Paths:
		return "paths"
	case Sites:
		return "sites"
	case Stats:
		return "stats"
	default:
		return "ERROR"
	}
}

const (
	Count  Endpoint = 0
	Export Endpoint = 1
	Me     Endpoint = 2
	Paths  Endpoint = 3
	Sites  Endpoint = 4
	Stats  Endpoint = 5
)
