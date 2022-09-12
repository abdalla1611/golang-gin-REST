package data

type Book struct{
		ID string		`json:"id"`
		Title string	`json:"title" binding:"required"`
		Year string		`json:"year" binding:"required"`
		Author string	`json:"author" binding:"required"`
}