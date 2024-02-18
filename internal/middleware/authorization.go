package middleware

import (
	"errors",
	"net/http"

	"github.com/itsAakanksha/goapi/api"
	"github.com/itsAakanksha/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	
)

var UnAuthorizedError  = errors.New("Invalid Username or token")

funct Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
	
		if username='' || token = '' {
             log.Error(UnAuthorizedError)
			 api.RequestEroorHandler(w,UnAuthorizedError)
			 return
		}

		var database *tools.DatabaseInterface
		database,err = tools.NewDatabase()
		if err!=nil{
			api.InternalErrorHandler(w)
            return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if(loginDetails == nil || (token != (*loginDetails).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestEroorHandler(w,UnAuthorizedError)
			return
		}
		next.ServeHTTP(w,r)
	}

}
