package Web

import "time"

type MemberResponse struct {
	Id        string `json:"id"` 
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate time.Time `json:"birth-date"`
	Address   string `json:"address"`
}