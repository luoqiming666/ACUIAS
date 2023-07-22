package services

import (
	"test.com/hello/app/models"
	"test.com/hello/app/repositories"
	"test.com/hello/utils"
)

// 验证身份服务（用到区块链）
type AuthService struct {
	blockchainService BlockchainService
	// ...
}

// 用户服务接口
type UserService interface {
	CreateUserService()
	UpdateUserService()
	DeleteUserService()
	GetUserService()
	IsUsernameExist(username string) bool
	UpdateToken()
	GetAllUserInfo()
}

// 数据库资源控制服务
type Userservice struct {
	userRepo repositories.UserRepository
}

// 创建新的用户服务
func NewUserService(userRepo repositories.UserRepository) *Userservice {
	return &Userservice{
		userRepo: userRepo,
	}
}

// 新建用户
func (s *Userservice) CreateUserService(userPost *models.User) (*models.User, error) {

	user := models.User{
		Id:       userPost.Id,
		Username: userPost.Username,
		Password: userPost.Password,
		Phone:    userPost.Phone,
		Email:    userPost.Email,
	}

	userCreate, err := s.userRepo.CreateUser(&user)
	return userCreate, err

}

// 删除用户
func (s *Userservice) DeleteUserService(username string) (*models.User, error) {

	deleteUser, err := s.userRepo.DeleteUser(username)
	return deleteUser, err

}

// 查询用户
func (s *Userservice) GetUserService(userName string) (*models.User, error) {

	getUser, err := s.userRepo.GetUser(userName)
	return getUser, err

}

// 更新用户数据
func (s *Userservice) UpdateUserService(username string, field string, updateDate interface{}) (*models.User, error) {
	updateUser, err := s.userRepo.UpdateUser(username, field, updateDate)
	return updateUser, err
}

// 判断用户是否存在
func (s *Userservice) IsUsernameExist(username string) bool {

	getUser, _ := s.userRepo.GetUser(username)
	if getUser.Username == "" {
		return false
	} else {
		return true
	}

}

// 更新用户token
func (s *Userservice) UpdateToken(user *models.User) (string, bool) {

	//工具控制类
	ut := &utils.Userutils{}

	randomToken, _ := ut.GenerateToken(uint(user.Id))

	_, err := s.userRepo.UpdateUser(user.Username, "token", randomToken)
	if err != nil {

		print("generate token failed")
		return "", false

	}

	return randomToken, true

}

// 获取用户信息
func (s *Userservice) GetAllUserInfoByID(ID int) (*models.User,error) {

	userGet, err := s.userRepo.GetUserByID(ID)
	return userGet,err

}

// 验证身份服务
func NewAuthService(blockchainService BlockchainService) *AuthService {

	// 验证服务
	return &AuthService{
		blockchainService: blockchainService,
	}
}
