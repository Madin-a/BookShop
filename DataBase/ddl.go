package DataBase

const CreateUsersAccount = `Create table if not exists users(
	id integer primary key autoincrement,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	email text not null,
	phone integer not null,
	login text not null unique,
	password text not null,
	role text not null,
	remove boolean not null default false
);`

const CreateTableBooks = `Create table if not exists books(
	id integer primary key autoincrement,
	name text not null,
	author text not null,
	category text not null,
	price integer not null,
	description text not null
);`
const CreateTableArchive = `Create table if not exists archive(
	id integer primary key autoincrement,
	userID integer not null,
	books text not null,
	totalPrice float not null,
	dateOfShoping integer not null
);`