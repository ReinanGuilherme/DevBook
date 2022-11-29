package middlewares

import (
	"api/src/auth"
	msgresponse "api/src/msgResponse"
	"net/http"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidarToken(r); erro != nil {
			msgresponse.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}

}
