package repository

import (
	"PaymentServiceApp/data/entity"
	"database/sql"
)

const saveCmd = `insert into users (username, password, name, phone) values ($1, $2, $3, $4)`
const findAllCmd = "select * from users"
const findByIdCmd = `select * from users where username=$1`
const countCmd = "select count(*) from users"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Count() (int, error) {
	rows, err := ur.db.Query(countCmd)

	if err != nil {
		return 0, err
	}

	var count int

	err = rows.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
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

func (ur *UserRepository) FindById(username string) (*entity.User, error) {
	rows, err := ur.db.Query(findByIdCmd, username)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	var password, name, phone string

	err = rows.Scan(&username, &password, &name, &phone)

	if err != nil {
		return nil, err
	}

	return entity.NewUser(username, password, name, phone), nil
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

func (ur *UserRepository) Save(user *entity.User) (*entity.User, error) {
	_, err := ur.db.Exec(saveCmd, user.Username, user.Password, user.Name, user.Phone)

	if err != nil {
		return nil, err
	}

	return user, nil
}

///////////////////////////////////////////////////////////////////////////////////////

func (ur *UserRepository) DeleteById(id string) error {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) Delete(entity *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) DeleteAll() error {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) DeleteAllById(ids []string) error {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) ExistsById(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) SaveAll(entities ...entity.User) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
