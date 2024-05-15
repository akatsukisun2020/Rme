package user

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Tokener token生成器，包含refreshtoken 和 普通token。使用AEC加密算法
type Tokener struct {
	UserID       string // 用户id
	UserSecretID string // 用户唯一加密id(勿传递到端上)
	TokenType    int    // 0: access_token, 1: refresh_token, 参见consts中TokenType定义.
	Token        string // token值
	EffectTime   int64  // 生效时间(ms)
}

// GenerateToken 生成accessToken
func (t *Tokener) GenerateToken() error {
	if t.UserID == "" || t.UserSecretID == "" {
		return fmt.Errorf("token gen error, userid or user secretid is empty")
	}

	// ${user_id}_${user_secretid}_${random_str}_${ts} 然后进行对称加密（不用hash）算法，的到token
	plaintext := fmt.Sprintf("%d_%s_%s_%s_%d", t.TokenType, t.UserID,
		t.UserSecretID, uuid.New().String(), t.EffectTime)

	// 从配置中读取key TODO
	appkey := []byte("1234567891234567") // 16 or 32 byte // TODO：从配置中读取
	// 创建AES加密块
	block, err := aes.NewCipher(appkey)
	if err != nil {
		return fmt.Errorf("token gen error, err[%s]", err.Error())
	}
	// 创建GCM模式的加密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("token gen error, err[%s]", err.Error())
	}
	// 获取nonce
	nonceBytes := []byte{}
	for len(nonceBytes) < aesGCM.NonceSize() {
		nonceBytes = append(nonceBytes, []byte(t.UserSecretID)...)
	}
	nonce := nonceBytes[:aesGCM.NonceSize()]

	// 加密
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)
	t.Token = string(ciphertext)
	return nil
}

// ParseToken 解析token
func (t *Tokener) ParseToken() error {
	if t.Token == "" {
		return fmt.Errorf("token parse error, token is empty")
	}

	// 从配置中读取key  TODO
	appkey := []byte("1234567891234567") // 16 or 32 byte // TODO：从配置中读取
	// 创建AES加密块
	block, err := aes.NewCipher(appkey)
	if err != nil {
		return fmt.Errorf("token parse error, err[%s]", err.Error())
	}
	// 创建GCM模式的加密器
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("token parse error, err[%s]", err.Error())
	}
	// 获取nonce
	nonceBytes := []byte{}
	for len(nonceBytes) < aesGCM.NonceSize() {
		nonceBytes = append(nonceBytes, []byte(t.UserSecretID)...)
	}
	nonce := nonceBytes[:aesGCM.NonceSize()]

	// 解密数据
	decrypted, err := aesGCM.Open(nil, nonce, []byte(t.Token), nil)
	if err != nil {
		return fmt.Errorf("token parse error, err[%s]", err.Error())
	}

	arrs := strings.Split(string(decrypted), "_")
	if len(arrs) != 5 {
		return fmt.Errorf("token parse error, token format error, decrypted:%s", string(decrypted))
	}

	t.TokenType, _ = strconv.Atoi(arrs[0])
	t.UserID = arrs[1]
	t.UserSecretID = arrs[2]
	t.EffectTime, _ = strconv.ParseInt(arrs[4], 10, 64)
	return nil
}
