package authorization

import (
	"github.com/casbin/casbin"
	"net/http"
	"x-archives/http-csbin-example/model"
)

func Authorizer(e *casbin.Enforcer,users model.Users) func(next http.Handler) http.Handler  {
	return func(next http.Handler) http.Handler {
		
	}
}