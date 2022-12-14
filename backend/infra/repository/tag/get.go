package tag

import (
	"backend/domain/model"
)

func (r *tagRepository) ListTags() (tags []*model.Tag, err error) {
	err = r.DB.Table("tags").Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}
