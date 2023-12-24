package repository

import (
	"PaymentServiceApp/data/entity"
	"database/sql"
)

const saveCmd = `insert into users (username, password, name, phone) values ($1, $2, $3, $4)`
const findAllCmd = "select * from users"
const findByIdCmd = `select * from users where username=$1`

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Save(user *entity.User) (*entity.User, error) {
	_, err := ur.db.Exec(saveCmd, user.Username, user.Password, user.Name, user.Phone)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindAll() ([]entity.User, error) {
	rows, err := ur.db.Query(findAllCmd)

	if err != nil {
		return nil, err
	}

	var users []entity.User

	for rows.Next() {
		var username, password, name, phone string

		err = rows.Scan(&username, &password, &name, &phone)

		if err != nil {
			return nil, err
		}

		users = append(users, *entity.NewUser(username, password, name, phone))
	}

	return users, nil
}

func (ur *UserRepository) FindAllFunc(f func(*entity.User)) error {
	rows, err := ur.db.Query(findAllCmd)

	if err != nil {
		return err
	}

	for rows.Next() {
		var username, password, name, phone string

		err = rows.Scan(&username, &password, &name, &phone)

		if err != nil {
			return err
		}
		f(entity.NewUser(username, password, name, phone))
	}

	return nil
}

func (ur *UserRepository) FindById(username string) (*entity.User, error) {
	panic("Not implemented yet!...")
}

//...
