package services

import (
	"context"
	"database/sql"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"time"
)

type RepoUserIF interface {
	GetUserByEmail(ctx context.Context, email string) (model entity.UserEntity, err error)
	CreateUser(ctx context.Context, model entity.UserEntity) (err error)
	BlacklistToken(token string, expiration time.Duration) (err error)
}

type ServiceUser struct {
	repo RepoUserIF
}

func NewServiceUser(repo RepoUserIF) *ServiceUser {
	return &ServiceUser{
		repo: repo,
	}
}

func (s *ServiceUser) Register(ctx context.Context, req request.RegisterRequestPayload) (err error) {
	userEntity := entity.NewFromRegisterRequest(req)

	if err = userEntity.Validate(); err != nil {
		return
	}

	// validasi apakah email sudah terdaftar/belum
	model, err := s.repo.GetUserByEmail(ctx, userEntity.Email)
	if err != nil {
		// jika terdapat error not found maka lanjut melakukan create user
		// jika terjadi error selain error not found maka akan me return error tersebut
		if err != pkg.ErrNotFound {
			return err
		}
	}

	if model.IsExists() {
		return pkg.ErrEmailExist
	}

	if err = userEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	return s.repo.CreateUser(ctx, userEntity)
}

func (s *ServiceUser) Login(ctx context.Context, req request.LoginRequestPayload) (token string, err error) {
	userEntity := entity.NewFromLoginRequest(req)

	if err = userEntity.Validate(); err != nil {
		return
	}

	// model ini berisi data dari database
	model, err := s.repo.GetUserByEmail(ctx, userEntity.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			err = pkg.ErrNotFound
			return
		}
		return
	}

	err = userEntity.VerifyPasswordFromPlain(model.Password)
	if err != nil {
		err = pkg.ErrPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JwtSecret)

	return
}

func (s *ServiceUser) Logout(token string, expiration time.Duration) error {
	// Blacklist token dengan menyimpannya ke Redis
	return s.repo.BlacklistToken(token, expiration)
}
