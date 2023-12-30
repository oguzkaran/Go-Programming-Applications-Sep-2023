package dal

import (
	"PaymentServiceApp/data/entity"
	"PaymentServiceApp/data/repository"
	"database/sql"
)

type PaymentAppHelper struct {
	userRepository    *repository.UserRepository
	productRepository *repository.ProductRepository
	paymentRepository *repository.PaymentRepository
}

func NewPaymentAppHelper(db *sql.DB) *PaymentAppHelper {
	return &PaymentAppHelper{repository.NewUserRepository(db), repository.NewProductRepository(db), repository.NewPaymentRepository(db)}
}

func (ph *PaymentAppHelper) SaveUser(user *entity.User) (*entity.User, error) {
	return ph.userRepository.Save(user)
}

func (ph *PaymentAppHelper) FindUserByUsername(username string) (*entity.User, error) {
	return ph.userRepository.FindById(username)
}

//...
