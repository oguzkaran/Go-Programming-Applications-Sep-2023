package service

import (
	"PaymentServiceApp/data/dal"
	"PaymentServiceApp/service/dto"
)

type PaymentAppDataService struct {
	helper *dal.PaymentAppHelper
}

func (ps *PaymentAppDataService) FindUserByUsername(username string) (*dto.UserDTO, error) {
	user, _ := ps.helper.FindUserByUsername(username)

	//...
	return &dto.UserDTO{user.Username, user.Name, user.Phone}, nil
}
