package service

import (
	"PaymentServiceApp/data/dal"
	"PaymentServiceApp/service/dto"
)

type PaymentAppDataService struct {
	helper *dal.PaymentAppHelper
}

func NewPaymentAppDataService(url string) *PaymentAppDataService {
	panic("todo")
}

func (ps *PaymentAppDataService) FindUserByUsername(username string) (*dto.UserDTO, error) {
	user, _ := ps.helper.FindUserByUsername(username)

	//...
	return &dto.UserDTO{Username: user.Username, Name: user.Name, Phone: user.Phone}, nil
}
