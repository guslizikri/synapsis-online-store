package pkg

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		tokenString, err := GenerateToken(publicId, "user", "niecret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)
		log.Println(tokenString)
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		publicId := uuid.NewString()
		role := "user"
		tokenString, err := GenerateToken(publicId, role, "niecret")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtID, jwtRole, err := ValidateToken(tokenString, "niecret")
		require.Nil(t, err)
		require.NotEmpty(t, jwtID)
		require.NotEmpty(t, jwtRole)

		require.Equal(t, publicId, jwtID)
		require.Equal(t, role, jwtRole)
	})

}
