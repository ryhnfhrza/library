package Web

type BookAddRequest struct {
	Id       string `validate:"required,max=10,min=1" json:"id"`
	Title    string `validate:"required,max=50,min=1" json:"title"`
	Author   string `validate:"required,max=30,min=1" json:"author"`
	Category string `validate:"required,max=20,min=1" json:"category"`
}