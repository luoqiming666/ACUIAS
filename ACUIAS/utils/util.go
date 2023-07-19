package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// JWT签名密钥
var signingKey = []byte("d378y473dudhwqi2903093281sygsyuqq878rkltrug8943")

// 生成Token
func generateToken(userID uint) (string, error) {
	// 创建一个新的Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 设置Token有效期为24小时
	})

	// 使用密钥进行签名并获取完整的Token字符串
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 加密密码
func encryptPassword(password string) (string, error) {
	// 将密码转换为字节数组
	passwordBytes := []byte(password)

	// 使用bcrypt进行密码哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 将哈希处理后的密码转换为字符串
	hashedPasswordString := string(hashedPassword)

	return hashedPasswordString, nil
}

// 验证密码
func verifyPassword(password, hashedPassword string) error {
	// 将哈希处理后的密码转换为字节数组
	hashedPasswordBytes := []byte(hashedPassword)

	// 验证密码是否正确
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(password))
	if err != nil {
		return err
	}

	return nil
}
