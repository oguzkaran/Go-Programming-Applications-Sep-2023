package service

import (
	"PaymentServiceApp/data/dal"
	"PaymentServiceApp/data/entity"
	"PaymentServiceApp/service/dto"
	"database/sql"
)

type PaymentAppDataService struct {
	helper *dal.PaymentAppHelper
}

func NewPaymentAppDataService(connStr string) (*PaymentAppDataService, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &PaymentAppDataService{dal.NewPaymentAppHelper(db)}, nil
}

func (ps *PaymentAppDataService) FindUserByUsername(username string) (*dto.UserDTO, error) {
	user, _ := ps.helper.FindUserByUsername(username)

	//...

	return dto.NewUserDTO(user.Username, user.Name, user.Phone), nil
}

func (ps *PaymentAppDataService) SaveUser(userSaveDTO *dto.UserSaveDTO) error {
	user := entity.NewUser(userSaveDTO.Username, userSaveDTO.Password, userSaveDTO.Name, userSaveDTO.Phone)
	_, err := ps.helper.SaveUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (ps *PaymentAppDataService) Close() error {
	return ps.helper.Close()
}
