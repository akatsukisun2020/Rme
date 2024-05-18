package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// Tokener token生成器，包含refreshtoken 和 普通token。使用AEC加密算法
type Tokener struct {
	UserID       string // 用户id
	UserSecretID string // 用户唯一加密id(勿传递到端上)
	TokenType    int    // 0: access_token, 1: refresh_token, 参见consts中TokenType定义.
	Token        string // token值
	ExpireTime   int64  // 过期时间(ms)token过期时间，端上根据这个时间，决定何时进行刷新accessToken.
}

// GenerateToken 生成accessToken
func (t *Tokener) GenerateToken(ctx context.Context) error {
	if t.UserID == "" || t.UserSecretID == "" {
		return fmt.Errorf("token gen error, userid or user secretid is empty")
	}

	// ${user_id}_${user_secretid}_${random_str}_${ts} 然后进行对称加密（不用hash）算法，的到token
	plaintext := fmt.Sprintf("%d_%s_%s_%s_%d", t.TokenType, t.UserID,
		t.UserSecretID, uuid.New().String(), t.ExpireTime)

	appkey, _ := g.Cfg().Get(ctx, "custom.token.appkey")
	// 创建AES加密块
	block, err := aes.NewCipher([]byte(appkey.String()))
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
	t.Token = base64.StdEncoding.EncodeToString(ciphertext)
	return nil
}

// ParseToken 解析token
func (t *Tokener) ParseToken(ctx context.Context) error {
	if t.Token == "" || t.UserSecretID == "" {
		return fmt.Errorf("token parse error, token or userSecretID is empty")
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(t.Token)
	if err != nil {
		return fmt.Errorf("token parse error, err:%s", err.Error())
	}

	appkey, _ := g.Cfg().Get(ctx, "custom.token.appkey")
	// 创建AES加密块
	block, err := aes.NewCipher([]byte(appkey.String()))
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
	decrypted, err := aesGCM.Open(nil, nonce, []byte(decodedBytes), nil)
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
	t.ExpireTime, _ = strconv.ParseInt(arrs[4], 10, 64)
	return nil
}
