package auth

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LicenseAuthRequest struct {
	License   string `json:"license"`   // base64 编码的加密 license
	Signature string `json:"signature"` // 请求签名
	Timestamp int64  `json:"timestamp"` // 请求时间戳
	Nonce     string `json:"nonce"`     // 随机数防重放
}

type LicenseAuthResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Signature string `json:"signature"` // 响应签名
	Timestamp int64  `json:"timestamp"`
}

type Client struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	cache      map[string]*LicenseAuthResponse // 简单内存缓存
}

// 创建认证中心实例
func NewClient(privateKey *rsa.PrivateKey) *Client {
	return &Client{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
		cache:      make(map[string]*LicenseAuthResponse),
	}
}

// 验证 License
func (client *Client) validateLicense(encryptedLicense string) (*License, error) {
	license, err := verifyLicense(nil, encryptedLicense)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt license: %v", err)
	}

	// 检查是否过期
	if time.Now().After(license.ExpiresAt) {
		return nil, fmt.Errorf("license expired")
	}

	// 检查服务器指纹
	currentFingerprint := getServerFingerprint()
	if license.ServerFingerprint != currentFingerprint {
		return nil, fmt.Errorf("server fingerprint mismatch")
	}

	return license, nil
}

// 认证 API 处理函数
func (client *Client) authenticate(c *gin.Context) {
	var req LicenseAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// 检查时间戳（防止重放攻击）
	if time.Now().Unix()-req.Timestamp > 300 { // 5分钟超时
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request expired"})
		return
	}

	// 构建签名数据
	// reqSignature := fmt.Sprintf("%s|%d|%s", req.License, req.Timestamp, req.Nonce)

	// 验证请求签名（这里简化，实际应该用客户端公钥验证）
	// err := verifySignature(client.publicKey, []byte(reqSignature), req.Signature)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
	// 	return
	// }

	// 检查缓存
	cacheKey := fmt.Sprintf("%s-%s", req.License, req.Nonce)
	if cachedResp, exists := client.cache[cacheKey]; exists {
		c.JSON(http.StatusOK, cachedResp)
		return
	}

	// 验证 License
	license, err := client.validateLicense(req.License)

	response := &LicenseAuthResponse{
		Timestamp: time.Now().Unix(),
	}

	if err != nil {
		response.Success = false
		response.Message = err.Error()
	} else {
		response.Success = true
		response.Message = fmt.Sprintf("Authentication successful for %s", license.CustomerName)
	}

	// 签名响应
	_ = fmt.Sprintf("%t|%s|%d", response.Success, response.Message, response.Timestamp)
	// resSignature, err := signData(client.privateKey, []byte(responseData))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign response"})
	// 	return
	// }
	// response.Signature = resSignature

	// // 缓存结果
	// client.cache[cacheKey] = response

	// c.JSON(http.StatusOK, response)
}
