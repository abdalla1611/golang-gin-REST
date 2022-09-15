package services

import (
	"main/repository"
	"main/Data"
	"github.com/google/uuid"
)

type Services struct{
	Repo *repository.Repository
}

func NewServices() *Services {
	return &Services{
		Repo : repository.NewRepository() ,
	}
}
func (s *Services) GetAllBooks() ([] Data.Book, error) {
	return s.Repo.GetAllBooks()
}

func (s *Services) GetBook(id string) (Data.Book ,error) {
	return s.Repo.GetBook(id)
}

func (s *Services) DeleteBook(id string)  error {
	return s.Repo.DeleteBook(id)
}

func (s *Services) UpdateBook(book Data.Book) (Data.Book, error){
	if err :=s.Repo.UpdateBook(book) ; err != nil{
		return Data.Book{}, err
	}
	return book, nil
}

func (s *Services) AddBook(book Data.Book)  (Data.Book, int, error) {
	book.ID = uuid.New().String()
	book,err := s.Repo.AddBook(book)
	if err != nil {
		return Data.Book{}, 400, err
	}
	return book, 201, nil
}
