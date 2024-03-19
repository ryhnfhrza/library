package Web

type BookUpdateRequest struct {
	Id       string ` json:"id"`
	Title    string ` json:"title"`
	Author   string ` json:"author"`
	Category string ` json:"category"`
}