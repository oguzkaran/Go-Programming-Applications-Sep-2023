package repository

import (
	"PaymentServiceApp/data/entity"
	"database/sql"
)

const saveCmd = `call sp_insert_user($1, $2, $3, $4)`
const findAllCmd = "select * from get_all_users()"
const findByIdCmd = `select * from get_user_by_username($1)`
const countCmd = "select * from get_user_count()"
const existsByUsernameCmd = `select * from user_exists($1)`

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

	rows.Next()

	var count int

	err = rows.Scan(&count)

	if err != nil {
		return 0, err
	}

	err = rows.Close()

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *UserRepository) ExistsById(username string) (bool, error) {
	rows, err := ur.db.Query(existsByUsernameCmd, username)

	if err != nil {
		return false, err
	}

	rows.Next()

	var exists bool

	err = rows.Scan(&exists)

	if err != nil {
		return false, err
	}

	err = rows.Close()

	if err != nil {
		return false, err
	}

	return exists, nil
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

	err = rows.Close()

	if err != nil {
		return nil, err
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

	err = rows.Close()

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

	err = rows.Close()

	if err != nil {
		return err
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
	panic("not implemented yet!...")
}

func (ur *UserRepository) Delete(entity *entity.User) error {
	panic("not implemented yet!...")
}

func (ur *UserRepository) DeleteAll() error {
	panic("not implemented yet!...")
}

func (ur *UserRepository) DeleteAllById(ids []string) error {
	panic("not implemented yet!...")
}

func (ur *UserRepository) SaveAll(entities ...entity.User) ([]entity.User, error) {
	panic("not implemented yet!...")
}
