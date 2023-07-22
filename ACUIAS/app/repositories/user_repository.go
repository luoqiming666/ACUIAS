package repositories

import (
	"test.com/hello/app/models"
	"test.com/hello/database"
)

// 数据库repository
type UserRepository struct {
	// UsersRepo *models.User
}

// 获取用户
func (ur *UserRepository) GetUser(username string) (*models.User, error) {

	user := &models.User{}
	db := database.DB
	err := db.Where("username=?", username).First(user).Error

	return user, err

}

// 用id获取用户
func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {

	user := &models.User{}
	db := database.DB
	err := db.Where("id=?", id).First(user).Error

	return user, err

}

// 新增用户
func (ur *UserRepository) CreateUser(userPost *models.User) (*models.User, error) {

	db := database.DB
	err := db.Create(userPost).Error

	return userPost, err

}

// 删除用户
func (ur *UserRepository) DeleteUser(username string) (*models.User, error) {
	db := database.DB

	ur1 := &UserRepository{}
	user, _ := ur1.GetUser(username)

	err := db.Debug().Where("username=?", username).Delete(user).Error

	return user, err

}

// 更新用户数据
func (ur *UserRepository) UpdateUser(username string, field string, updateData interface{}) (*models.User, error) {

	db := database.DB

	ur1 := &UserRepository{}
	user, _ := ur1.GetUser(username)

	err := db.Debug().Model(user).Update(field, updateData).Error

	return user, err
}
