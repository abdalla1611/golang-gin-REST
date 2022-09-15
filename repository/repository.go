package repository

import(
	"main/Data"
	"errors"
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

func (repo *Repository) GetAllBooks() ([]Data.Book, error){
	var books []Data.Book

	for _,book :=range repo.DB {
		books = append(books, book)
	}
	return books ,nil
}

func (repo *Repository) GetBook(id string) (Data.Book ,error){
	if book, exist := repo.DB[id] ; exist{
		return book ,nil
	}
	return Data.Book{} , errors.New("book not found !")
}

func (repo *Repository) DeleteBook(id string) error {
	if _, exist := repo.DB[id] ; !exist{
		return errors.New("book not found !")
	}
	delete(repo.DB, id)

	return nil
}

func (repo *Repository) UpdateBook(book Data.Book) error{
	if _, exist := repo.DB[book.ID] ; !exist{
		return errors.New("book not found !")
	}
	
	repo.DB[book.ID] = book
	return nil
}

func (repo *Repository) AddBook(book Data.Book) (Data.Book, error){
	if _, exist := repo.DB[book.ID] ; exist{
		return Data.Book{},errors.New("already exist !")
	}
	repo.DB[book.ID] = book
	return book , nil
}



