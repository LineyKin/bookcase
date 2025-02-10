package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a AuthData) CalcPwdHash(pwd string) string {
	saltedPwd := fmt.Sprintf("%s%s%s", os.Getenv("SALT_1"), pwd, os.Getenv("SALT_2"))
	result := sha256.Sum256([]byte(saltedPwd))

	return hex.EncodeToString(result[:])
}

func (a *AuthData) HashPwd() {
	a.Password = a.CalcPwdHash(a.Password)
}
