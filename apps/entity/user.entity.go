package entity

import (
	"strings"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Role_Admin Role = "admin"
	Role_User  Role = "user"
)

type UserEntity struct {
	Id        uint      `db:"id"`
	PublicID  uuid.UUID `db:"public_id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u UserEntity) Validate() (err error) {
	if err = u.ValidateEmail(); err != nil {
		return
	}
	if err = u.ValidatePassword(); err != nil {
		return
	}
	return
}
func (u UserEntity) ValidateEmail() (err error) {
	if u.Email == "" {
		return pkg.ErrEmailRequired
	}

	email := strings.Split(u.Email, "@")
	if len(email) != 2 {
		return pkg.ErrEmailInvalid
	}
	return
}

func (u UserEntity) ValidatePassword() (err error) {
	if u.Password == "" {
		return pkg.ErrPasswordRequired
	}

	if len(u.Password) < 6 {
		return pkg.ErrPasswordInvalidLength
	}
	return
}

func NewFromRegisterRequest(req request.RegisterRequestPayload) UserEntity {
	return UserEntity{
		PublicID:  uuid.New(),
		Email:     req.Email,
		Password:  req.Password,
		Role:      Role_User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFromLoginRequest(req request.LoginRequestPayload) UserEntity {
	return UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (u *UserEntity) EncryptPassword(salt int) (err error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(encryptedPass)

	return nil
}

func (u UserEntity) VerifyPasswordFromEncrypted(plainpass string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainpass))
}
func (u UserEntity) VerifyPasswordFromPlain(encryptedpass string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encryptedpass), []byte(u.Password))
}

func (u *UserEntity) IsExists() bool {
	return u.Id != 0
}

func (u UserEntity) GenerateToken(secret string) (token string, err error) {
	return pkg.GenerateToken(u.PublicID.String(), string(u.Role), secret)

}
