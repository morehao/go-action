package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
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

// loadPublicKeyFromPEM 从 PEM 格式加载公钥
func loadPublicKeyFromPEM(pemData string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

// 解密 License
func verifyLicense(publicKey *rsa.PublicKey, encryptedLicense string) (*License, error) {
	// // base64 解码
	// encryptedData, err := base64.StdEncoding.DecodeString(encryptedLicense)
	// if err != nil {
	// 	return nil, err
	// }

	// // 使用私钥解密
	// licenseData, err := rsa.VerifyPKCS1v15(publicKey)
	// if err != nil {
	// 	return nil, err
	// }

	// // 反序列化
	// var license License
	// err = json.Unmarshal(licenseData, &license)
	// if err != nil {
	// 	return nil, err
	// }

	// return &license, nil
	return nil, nil
}

func getServerFingerprint() string {
	hostname, _ := os.Hostname()
	return fmt.Sprintf("SERVER-%s", hostname)
}
