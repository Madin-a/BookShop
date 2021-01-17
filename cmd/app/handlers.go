package app

import (
	"HumosBooks/DataBase"
	"HumosBooks/cmd/app/tokens"
	"HumosBooks/models"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"time"
)

func (server *MainServer) RegistrationHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.RegisterData
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}

	if len(requestBody.Login) < 6 {
		fmt.Println("Login must have more than 6 symbols!")
		err = json.NewEncoder(writer).Encode("Login must have more than 6 symbols!")
		return
	}
	registerData := requestBody
	ok, err := models.AddUserToDB(server.DB, registerData)
	if !ok {
		fmt.Println("error in insert to db")
		err = json.NewEncoder(writer).Encode("error in insert to db")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
		//fmt.Println("error in insert to db")
	}
	//fmt.Println(registerData)
	fmt.Println("User has been registered. ")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode("insert to db")
}

func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.LoginData
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	loginData := requestBody
	row := server.DB.QueryRow(DataBase.GetUserByLoginAndPass, loginData.Login, loginData.Password)
	fmt.Println(row.Err())
	var user models.RegisterData
	row.Scan(&user.ID,
		&user.Name,
		&user.Surname,
		&user.Age,
		&user.Gender,
		&user.Email,
		&user.Phone,
		&user.Login,
		&user.Password,
		&user.Role,
		&user.Remove)
	fmt.Println(user)

	if user.ID == 0 || user.Remove == true {
		writer.WriteHeader(http.StatusNonAuthoritativeInfo)
		err = json.NewEncoder(writer).Encode("Invalid_data")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}

	token := tokens.CreateToken(user.ID, user.Login, user.Password, user.Role)
	var responseBody models.LoginResponse
	responseBody.Description = "Ok"
	responseBody.Token = token

	err = json.NewEncoder(writer).Encode(responseBody)
	if err != nil {
		fmt.Println("Error while writing responseBody. Error is", err)
	}

}
func (server *MainServer) AddBookHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody models.BookData
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}

	bookData := requestBody
	fmt.Println(requestBody)
	ok, err := models.AddBookToDB(server.DB, bookData)
	if !ok {
		fmt.Println("error in insert to db")
		err = json.NewEncoder(writer).Encode("error in insert to db")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
		//fmt.Println("error in insert to db")
	}
	//fmt.Println(registerData)
	fmt.Println("Book was added to Database")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode("insert to db")
}

func (server *MainServer) SearchBookHandler(writer http.ResponseWriter, reader *http.Request, params httprouter.Params) {

	var requestBody models.BookData
	err := json.NewDecoder(reader.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	var requestString string
	if len(requestBody.Name) > 2 {
		requestString += "name like '%" + requestBody.Name + "%' "
	}
	if len(requestBody.Author) > 2 {
		if len(requestString) > 2 {
			requestString += " or "
		}
		requestString += "author like '%" + requestBody.Author + "%' "
	}
	if len(requestBody.Description) > 2 {
		if len(requestString) > 2 {
			requestString += " or "
		}
		requestString += "description like '%" + requestBody.Description + "%'"
	}
	if len(requestBody.Category) > 2 {
		if len(requestString) > 2 {
			requestString += " or "
		}
		requestString += "category like '%" + requestBody.Category + "%'"
	}

	if len(requestString) != 0 {
		requestString = " where " + requestString
	}
	fmt.Println(requestString)
	rows, err := server.DB.Query(DataBase.SelectBook + requestString)
	if err != nil {
		fmt.Println("ошибка", err)
		return
	}
	books := []models.BookData{}
	for rows.Next() {
		book := models.BookData{}
		err := rows.Scan(
			&book.ID, &book.Name, &book.Author, &book.Category, &book.Price, &book.Description)
		if err != nil {
			fmt.Println("ошибка", err)
			continue
		}
		books = append(books, book)
	}
	err = json.NewEncoder(writer).Encode(books)
}

func (server *MainServer) SearchBookByIDHandler(writer http.ResponseWriter, reader *http.Request, params httprouter.Params) {
	var requestBody struct {
		ID int64 `json:"id"'`
	}
	err := json.NewDecoder(reader.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	/*_, err = models.SearchBookByID(server.DB, requestBody.ID)

	if err != nil {
		fmt.Println("Can't get book:", err)
		return
	}
	err = json.NewEncoder(writer).Encode("results: ")
	if err != nil {
		fmt.Println("Can't find connection: ", err)
		return
	}*/
		rows, err := server.DB.Query(DataBase.SelectBookByID, requestBody.ID)
		if err != nil {
			fmt.Println("ошибка", err)
			return
		}
		books := []models.BookData{}
		for rows.Next() {
			book := models.BookData{}
			err := rows.Scan(
				&book.ID, &book.Name, &book.Author, &book.Category, &book.Price, &book.Description)
			if err != nil {
				fmt.Println("ошибка", err)
				continue
			}
			books = append(books, book)
		}
		err = json.NewEncoder(writer).Encode(books)
	}

func (server *MainServer) DeleteBookByHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody struct {
		ID int64 `json:"id"'`
	}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	_, err = models.DeleteBookFromDB(server.DB, requestBody.ID)

	if err != nil {
		fmt.Println("Can't delete book:", err)
		return
	}
	err = json.NewEncoder(writer).Encode("Book was deleted!")
	if err != nil {
		fmt.Println("Can't find connection: ", err)
		return
	}
}

func (server *MainServer) BooksBuyHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody struct {
		BooksId []string `json:"books_id"`
	}
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil { //проверка правильности json
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode("Invalid_json")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	bearerToken := request.Header.Get("Authorization")
	Token := bearerToken[len("Bearer "):]
	claims := tokens.ParseToken(Token)

	booksArr := requestBody.BooksId
	var requestString string
	for i := 0; i < len(booksArr); i++ {
		if i == 0 {
			requestString += " where id = " + booksArr[i]
		} else {
			requestString += " or id = " + booksArr[i]
		}
	}
	rows, err := server.DB.Query(DataBase.SelectBooks + requestString)
	if err != nil {
		fmt.Println("error to get book's prices", err)
	}
	fmt.Println(DataBase.SelectBooks + requestString)
	fmt.Println(requestBody)
	var sum float64
	sum = float64(0)
	for rows.Next() {
		id := 0
		price := float64(0)
		err := rows.Scan(
			&id, &price)
		if err != nil {
			fmt.Println("ошибка", err)
			continue
		}
		sum += price
	}
	buyList := models.BuyListResponse{
		UserID:        claims.ID,
		Books:         strings.Join(booksArr, ","),
		Total:         sum,
		DateOfShoping: time.Now().Unix(),
	}
	ok, err := models.AddToArchive(server.DB, buyList)
	if !ok {
		fmt.Println("error in insert to db")
		err = json.NewEncoder(writer).Encode("error in insert to db")
		if err != nil {
			fmt.Println("Can't find connection")
			return
		}
		return
	}
	fmt.Println("The purchase was added to Database")
	writer.WriteHeader(http.StatusOK)
	err = json.NewEncoder(writer).Encode("insert to db")
}

