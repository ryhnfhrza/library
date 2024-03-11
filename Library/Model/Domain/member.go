package domain

import "time"

type Member struct {
	Id        string
	Name      string
	Email     string
	BirthDate time.Time
	Address string
}