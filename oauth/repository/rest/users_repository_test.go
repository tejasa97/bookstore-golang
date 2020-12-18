package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting mock server...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:        "localhost:8001/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody:    `{"email":"email@gmail.com", "password":"password"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "password")
	fmt.Println(user, err)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserNoError(t *testing.T) {

}
