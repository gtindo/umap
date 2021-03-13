package app

import "time"

type Url struct {
	Id        int
	Link      string
	LinkId    string
	CreatedAt time.Time
}

type Visitor struct {
	Id               int
	UrlId            int
	Device           string
	Country          string
	ScreenResolution string
	Browser          string
	OsName           string
}
