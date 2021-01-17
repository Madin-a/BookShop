package DataBase

const AddNewUser = `insert into users(name, surname, age, gender, email, phone, login, password, role) VALUES (($1),($2),($3),($4),($5), ($6), ($7), ($8), ($9))`
const AddNewBook = `insert into books(name, author, category, price, description) VALUES (($1),($2),($3),($4),($5))`
const DeleteBook = `delete from books where id = ($1)`
const GetUserByLoginAndPass = `select * from users where login=($1) and password=($2)`

const SelectBook = `select * from books `
const SelectBooks = `select id, price from books `
const SelectBookByID = `select * from books where id = ($1)`

const AddToArchive = `insert into archive(userID, books, totalPrice, dateOfShoping) VALUES (($1),($2),($3),($4))`
