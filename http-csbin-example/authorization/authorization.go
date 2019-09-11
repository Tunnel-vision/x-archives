package authorization

import (
	"github.com/alexedwards/scs/session"
	"github.com/casbin/casbin"
	"net/http"
	"x-archives/http-csbin-example/model"
	"log"
	"errors"
)

func Authorizer(e *casbin.Enforcer, users model.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role, err := session.GetString(r, "role")
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if role == "" {
				role = "anonymous"
			}
			// if it is a member,check if the user still exists
			if role == "member" {
				uid, err := session.GetInt(r, "")
				if err != nil{
					writeError(http.StatusInternalServerError, "ERROR", w, err)
					return
				}
				exists := users.Exists(uid)
				if !exists{
					writeError(http.StatusForbidden,"FORBIDDEN",w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			//res := true
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil{
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res{
				next.ServeHTTP(w,r)
			}else {
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", "--------",err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}
