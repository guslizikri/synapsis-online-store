package services

import (
	"context"
	"fmt"
	"log"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var svcUser *ServiceUser

func init() {
	filename := "../../cmd/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := pkg.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		return
	}
	repo := repository.NewRepoUser(db)
	svcUser = NewServiceUser(repo)
}

func TestRegisterSuccess(t *testing.T) {
	req := request.RegisterRequestPayload{
		// ini menggunakan uuid agar emailnya selalu beda
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "abcd1234",
	}
	err := svcUser.Register(context.Background(), req)
	require.Nil(t, err)
}
func TestRegisterFail(t *testing.T) {
	t.Run("Error email already used", func(t *testing.T) {
		// prepare for duplicate email
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		req := request.RegisterRequestPayload{
			// ini menggunakan uuid agar emailnya selalu beda
			Email:    email,
			Password: "abcd1234",
		}
		err := svcUser.Register(context.Background(), req)
		require.Nil(t, err)
		// end preparation

		//  sebelumnya regist berhasil dulu,
		// kemudian regist lagi menggunakan email yg sama
		err = svcUser.Register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, "email already exists", err.Error())
	})
}

func TestLogin(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// prepare for login
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		pass := "abcd1234"
		registReq := request.RegisterRequestPayload{
			Email:    email,
			Password: pass,
		}
		err := svcUser.Register(context.Background(), registReq)
		require.Nil(t, err)
		// end preparation
		loginReq := request.LoginRequestPayload{
			Email:    email,
			Password: pass,
		}

		token, err := svcUser.Login(context.Background(), loginReq)
		require.Nil(t, err)
		require.NotEmpty(t, token)
		log.Println(token)
	})
}
