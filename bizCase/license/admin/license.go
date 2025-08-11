package admin

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type License struct {
	LicenseID         string    `json:"license_id"`
	CustomerName      string    `json:"customer_name"`
	ServerFingerprint string    `json:"server_fingerprint"`
	ProductVersion    string    `json:"product_version"`
	GeneratedAt       time.Time `json:"generated_at"`
	ExpiresAt         time.Time `json:"expires_at"`
}

// 生成 License
func generateLicense(privateKey *rsa.PrivateKey, customerName, serverFingerprint string, validDays int) (string, error) {
	license := License{
		GeneratedAt:       time.Now(),
		ServerFingerprint: serverFingerprint,
		ExpiresAt:         time.Now().AddDate(0, 0, validDays),
		CustomerName:      customerName,
		LicenseID:         fmt.Sprintf("LIC-%d", time.Now().Unix()),
	}

	// 序列化 license
	licenseData, err := json.Marshal(license)
	if err != nil {
		return "", err
	}

	// 使用私钥加密
	encryptedLicense, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, licenseData)
	if err != nil {
		return "", err
	}

	// 返回 base64 编码的加密数据
	return base64.StdEncoding.EncodeToString(encryptedLicense), nil
}
