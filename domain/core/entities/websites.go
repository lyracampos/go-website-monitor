package entities

import "time"

type WebSite struct {
	Id        int
	Name      string
	Url       string
	Status    int
	LastCheck time.Time
}

type WebSites []WebSite
