package query

import (
	"context"
	"yatter-backend-go/app/domain/object/auth"
	"yatter-backend-go/app/usecase/query"

	"github.com/jmoiron/sqlx"
)

var _ query.AuthQueryService = (*AuthQueryServiceImpl)(nil)

type AuthQueryServiceImpl struct {
	db *sqlx.DB
}

func NewAuthQueryService(db *sqlx.DB) *AuthQueryServiceImpl {
	return &AuthQueryServiceImpl{db: db}
}

type FindCredentialByUsernameDTO struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

func (s *AuthQueryServiceImpl) FindCredentialByUsername(
	ctx context.Context,
	username string,
) (*auth.Credential, error) {
	var dbCredential FindCredentialByUsernameDTO
	err := s.db.GetContext(ctx, &dbCredential,
		`SELECT username, password_hash FROM user WHERE username = ?`,
		username,
	)

	if err != nil {
		return nil, err
	}

	credential, err := auth.ReconstructCredential(dbCredential.Username, dbCredential.PasswordHash)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
