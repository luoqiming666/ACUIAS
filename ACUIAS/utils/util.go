package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 用户工具
type Userutils struct {
}

// JWT签名密钥
var SigningKey = []byte("d378y473dudhwqi2903093281sygsyuqq878rkltrug8943")

// md5加密秘钥
var Md5PrivateKey = "d378y473dudhwqi2903093281sygsyvwdqb7y4782bwdqkjdq"

// 生成Token
func (ut *Userutils) GenerateToken(userID uint) (string, error) {
	// 创建一个新的Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 设置Token有效期为24小时
	})

	// 使用密钥进行签名并获取完整的Token字符串
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		print("\n\ncan't generate random token\n\n")
		return "", err
	}

	return tokenString, nil
}

// 解析并验证Token
func ParseAndValidateToken(tokenString string) (bool, error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 指定使用相同的密钥进行签名验证
		return SigningKey, nil
	})

	if err != nil {
		return false, err
	}

	// 检查Token是否有效
	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	// 检查Token是否过期
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, fmt.Errorf("invalid token claims")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if expirationTime.Before(time.Now()) {
		return false, fmt.Errorf("token has expired")
	}

	return true, nil
}

// 加密函数
func (ut *Userutils) EncryptWithMD5(password string) string {
	// 将密码与密钥拼接
	data := []byte(password + Md5PrivateKey)

	// 使用MD5进行加密
	hasher := md5.New()
	hasher.Write(data)

	// 获取加密后的字节数组
	encryptedData := hasher.Sum(nil)

	// 将字节数组转换为字符串
	encryptedString := hex.EncodeToString(encryptedData)

	return encryptedString
}

// 验证密码
func (ut *Userutils) VerifyPassword(password, hashedPassword string) bool {

	// 校验密码
	passwordVerify := ut.EncryptWithMD5(password)

	// 密码正确
	if hashedPassword == passwordVerify {
		return true
	} else {
		return false
	}

}

// 写入文件
func (ut *Userutils) WriteFile(filepath string, data interface{}) {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("err occur", err.Error())
		return
	}
	defer f.Close()

	byteData, ok := data.([]byte)
	if !ok {
		fmt.Println("can't convert to byte")
		return
	}

	_, err = f.Write(byteData)
	if err != nil {
		fmt.Println("can't write file")
		return

	}

}
