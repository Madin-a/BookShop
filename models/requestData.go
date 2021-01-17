package models

import (
	"HumosBooks/DataBase"
	"database/sql"
	"fmt"
)

type LoginData struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Description string `json:"description"`
	Token string `json:"token"`
	Role string `json:"role"`
}
type RegisterData struct {
	ID       int64 	`json:"id"`
	Name     string `json:"name, omitempty"`
	Surname  string `json:"surname, omitempty"`
	Age      int64 	`json:"age, omitempty"`
	Gender   string `json:"gender, omitempty"`
	Email 	string 	`json:"email, omitempty"`
	Phone 	int64 	`json:"phone, omitempty"`
	Login    string `json:"login, omitempty"`
	Password string `json:"password, omitempty"`
	Role   string `json:"role, omitempty"`
	Remove   bool 	`json:"remove, omitempty"`
}

func AddUserToDB(database *sql.DB, data RegisterData) (ok bool, err error) {
	_, err = database.Exec(DataBase.AddNewUser, data.Name, data.Surname, data.Age, data.Gender, data.Email, data.Phone, data.Login, data.Password, data.Role)
	if err != nil {
		fmt.Println(`Can't add new user`, err)
		return false, err
	}
	return true, nil
}

type BookData struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}

func AddBookToDB(database *sql.DB, data BookData) (ok bool, err error) {
	_, err = database.Exec(DataBase.AddNewBook, data.Name, data.Author, data.Category, data.Price, data.Description)
	if err != nil {
		fmt.Println(`Can't add new book`, err)
		return false, err
	}
	return true, nil
}

func DeleteBookFromDB(database *sql.DB, bookID int64) (ok bool, err error)  {
	_, err = database.Exec(DataBase.DeleteBook, bookID)
	if err != nil {
		fmt.Println(`Can't delete book`, err)
		return false, err
	}
	return true, nil
}
/*func SearchBookByID(database *sql.DB, bookID int64)(ok bool, err error)  {
	_, err = database.Exec(DataBase.SelectBookByID, bookID)
	if err != nil {
		fmt.Println(`Can't get book`, err)
		return false, err
	}
	return true, nil
}*/

type BuyListResponse struct {
	ID int64
	UserID int64
	Books string
	Total float64
	DateOfShoping int64
}

func AddToArchive(database *sql.DB, buyList BuyListResponse) (ok bool, err error) {
	_, err = database.Exec(DataBase.AddToArchive, buyList.UserID, buyList.Books, buyList.Total, buyList.DateOfShoping)
	if err != nil {
		fmt.Println(`Can't add to archive`, err)
		return false, err
	}
	return true, nil
}