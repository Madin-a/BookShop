package middlware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"HumosBooks/cmd/app/tokens"
)

func Authorized() func(next httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			bearerToken := request.Header.Get("Authorization")
			Token := bearerToken[len("Bearer "):]
			claims := tokens.ParseToken(Token)
			if claims.Login == "" {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(writer, request, params)
		}
	}
}
