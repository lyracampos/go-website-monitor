package entities

import "time"

type WebSite struct {
	Id        int
	Name      string //obrigatóio, 3+
	Url       string //obrigatório, válido, 10+
	Status    int //obrigatório
	LastCheck time.Time
}

type WebSites []WebSite
