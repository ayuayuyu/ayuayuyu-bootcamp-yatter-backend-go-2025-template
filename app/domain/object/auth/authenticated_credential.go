package auth

import "golang.org/x/crypto/bcrypt"

// AuthenticatedCredential は認証済みのユーザー情報を表す構造体
// 与えられたpasswordとpasswordHashが一致した認証情報
type AuthenticatedCredential struct {
	username     string
	passwordHash string
}

func NewAuthenticatedCredential(credential *Credential, password string) (*AuthenticatedCredential, error) {
	err := bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash()), []byte(password))
	if err != nil {
		return nil, err
	}

	return &AuthenticatedCredential{
		username:     credential.Username(),
		passwordHash: credential.PasswordHash(),
	}, nil
}

func (c *AuthenticatedCredential) Username() string {
	return c.username
}

func (c *AuthenticatedCredential) PasswordHash() string {
	return c.passwordHash
}
