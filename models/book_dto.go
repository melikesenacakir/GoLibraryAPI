package dto

import "rest-api-2/db/dbstore"

type BookDto struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Stock  int32  `json:"stock"`
}
type CreateBookDto struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Stock  int32  `json:"stock"`
}

func DtoToBook(bookdto *BookDto) *dbstore.Book { //gelen dto posta dönüştürülür
	return &dbstore.Book{
		Name:   bookdto.Name,
		Author: bookdto.Author,
		Stock:  bookdto.Stock,
	}
}

func BookToDto(book *dbstore.Book) *BookDto { //veritabanına dışa dönme(dtoya dönüştürme)
	return &BookDto{
		ID:     book.ID,
		Name:   book.Name,
		Author: book.Author,
		Stock:  book.Stock,
	}
}
func BooksToDto(Book *[]dbstore.Book) []BookDto {
	var result []BookDto
	for _, books := range *Book {
		book := BookToDto(&books)
		result = append(result, *book)
	}

	return result
}
func CreateDtoToBook(bookdto *CreateBookDto) *dbstore.CreateBookParams {
	return &dbstore.CreateBookParams{
		Name:   bookdto.Name,
		Author: bookdto.Author,
		Stock:  bookdto.Stock,
	}
}

func CreateBookToDto(book *dbstore.CreateBookParams) *CreateBookDto {
	return &CreateBookDto{
		Name:   book.Name,
		Author: book.Author,
		Stock:  book.Stock,
	}
}
func UpdateDtoToBook(bookdto *BookDto) *dbstore.UpdateBookParams {
	return &dbstore.UpdateBookParams{
		ID:     bookdto.ID,
		Name:   bookdto.Name,
		Author: bookdto.Author,
		Stock:  bookdto.Stock,
	}
}
