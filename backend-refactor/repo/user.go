package repo

import "boilerplate/backend-refactor/model"

func (r *Repo) GetUserByName(name string) (model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users"từ request
	var user model.User
	r.db.Where("username = ?", name).First(&user)
	return user, nil
}
