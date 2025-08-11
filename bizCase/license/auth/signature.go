package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func SignData(data any, secret string) (string, error) {
	return signData(data, secret)
}

func signData(data any, secret string) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(jsonBytes)
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func VerifySign(data any, signature string, secret string) (bool, error) {
	return verifySign(data, signature, secret)
}

func verifySign(data any, signature string, secret string) (bool, error) {
	expected, err := signData(data, secret)
	if err != nil {
		return false, err
	}
	if expected != signature {
		return false, fmt.Errorf("signature is invalid, expected: %s, got: %s", expected, signature)
	}
	return true, nil
}

// // 验证签名
// func verifySignature(publicKey *rsa.PublicKey, data []byte, signature string) error {
// 	sig, err := base64.StdEncoding.DecodeString(signature)
// 	if err != nil {
// 		return err
// 	}
// 	hashed := sha256.Sum256(data)
// 	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], sig)
// }
