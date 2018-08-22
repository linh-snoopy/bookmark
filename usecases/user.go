package usecases

import (

)

type UserUsecase struct {}

type UserInteractor interface {
	SignUp(email, password string) error
	Login(email, password string) error
	ChangePassword(old, new string) error
}

func (u UserUsecase) SignUp(email, password string) error {
	return nil
}

func (u UserUsecase) Login(email, password string) error {
	return nil
}

func (u UserUsecase) ChangePassword(old, new string) error {
	return nil
}