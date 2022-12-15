package repository

import "backend/domain/model"

type TagRepository interface {
	ListTags() ([]*model.Tag, error)
}
