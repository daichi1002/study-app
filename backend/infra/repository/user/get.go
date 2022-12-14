package user

import (
	"backend/domain/model"
)

func (r *userRepository) ListUsers() (users []*model.User, err error) {
	err = r.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetLoginUser(email string) (user *model.User, err error) {
	err = r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
