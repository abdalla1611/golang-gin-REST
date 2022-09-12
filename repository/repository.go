package repository
import(
	"main/Data"
)
type Repository struct{
	DB map[string]Data.Book
}

func NewRepository() *Repository {
	example := Data.Book{
		ID : "1",
		Title : "Data Structure",
		Year : "2022" ,
		Author: "Abdalla Sleman" ,
	}

	return &Repository{
		DB : map[string]Data.Book{example.ID : example} ,
	}
}

