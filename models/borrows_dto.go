package dto

import (
	"rest-api-2/db/dbstore"
	"time"
)

type CreateBookborrowDto struct {
	Borrowdate time.Time `json:"borrowdate"`
	Returndate time.Time `json:"returndate"`
	UserID     int64     `json:"user_id"`
	BookID     int64     `json:"book_id"`
	Status     string    `json:"status"`
}

type ReturnbookDto struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type BookborrowDto struct {
	BorrowID   int64     `json:"borrow_id"`
	Borrowdate time.Time `json:"borrowdate"`
	Returndate time.Time `json:"returndate"`
	Status     string    `json:"status"`
	Username   string    `json:"username"`
	BookName   string    `json:"book_name"`
}

func CreateDtoToBorrow(borrow *CreateBookborrowDto) *dbstore.BookBorrowParams {
	return &dbstore.BookBorrowParams{
		Borrowdate: borrow.Borrowdate,
		Returndate: borrow.Returndate,
		Status:     borrow.Status,
		UserID:     borrow.UserID,
		BookID:     borrow.BookID,
	}
}

func CreateBorrowToDto(borrow *dbstore.BookBorrowParams) *CreateBookborrowDto {
	return &CreateBookborrowDto{
		Borrowdate: borrow.Borrowdate,
		Returndate: borrow.Returndate,
		Status:     borrow.Status,
	}
}
func BorrowToDto(borrow *[]dbstore.GetBorrowsRow) []BookborrowDto {
	var result []BookborrowDto

	for _, borrows := range *borrow {
		temp := BookborrowDto{
			BorrowID:   borrows.BorrowID,
			Borrowdate: borrows.Borrowdate,
			Returndate: borrows.Returndate,
			Status:     borrows.Status,
			BookName:   borrows.BookName,
			Username:   borrows.UserUsername,
		}
		result = append(result, temp)
	}
	return result
}

func DtoToBorrow(borrow *BookborrowDto) *dbstore.GetBorrowsRow {
	return &dbstore.GetBorrowsRow{
		BorrowID:     borrow.BorrowID,
		Borrowdate:   borrow.Borrowdate,
		Returndate:   borrow.Returndate,
		Status:       borrow.Status,
		UserUsername: borrow.Username,
		BookName:     borrow.BookName,
	}
}

func DtoToReturnBorrow(borrow *ReturnbookDto) *dbstore.ReturnbookParams {
	return &dbstore.ReturnbookParams{
		ID:     borrow.ID,
		Status: borrow.Status,
	}
}
func BorrowHistoryToDto(borrow *[]dbstore.GetBorrowHistoryRow) []BookborrowDto {
	var result []BookborrowDto

	for _, borrows := range *borrow {
		temp := BookborrowDto{
			BorrowID:   borrows.BorrowID,
			Borrowdate: borrows.Borrowdate,
			Returndate: borrows.Returndate,
			Status:     borrows.Status,
			BookName:   borrows.BookName,
			Username:   borrows.UserUsername,
		}
		result = append(result, temp)
	}
	return result
}
