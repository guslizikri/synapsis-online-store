package entity

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPass(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		userEntity := UserEntity{
			Email:    "my@email.com",
			Password: "pass",
		}
		err := userEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", userEntity)
	})
}
