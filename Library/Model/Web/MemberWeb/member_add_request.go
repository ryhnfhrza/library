package Web

type MemberAddRequest struct {
	Name      string `validate:"required,max=30,min=1" json:"name"`
	Email     string `validate:"required,max=50,min=1,email" json:"email"`
	BirthDate string `validate:"required" json:"birth-date"`
	Address   string `validate:"required,max=20,min=1" json:"address"`
}