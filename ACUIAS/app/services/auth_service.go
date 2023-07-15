package services

// 验证身份逻辑

// 验证身份服务（用到区块链）
type AuthService struct {
	blockchainService BlockchainService
	// ...
}

// 验证身份服务
func NewAuthService(blockchainService BlockchainService) *AuthService {

	// 验证服务
	return &AuthService{
		blockchainService: blockchainService,
	}
}

// 实现其他方法，如登录、注册等
