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
