package user

import (
	"backend/domain/model"
	"sort"
)

func (r *userRepository) ListUsers() (users []*model.User, err error) {
	err = r.DB.Preload("Articles").Find(&users).Error

	sort.Slice(users, func(i, j int) bool { return len(users[i].Articles) > len(users[j].Articles) })
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
