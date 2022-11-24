package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// configurar, coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	//pegando o slice das rotasUsuarios
	rotas := rotasUsuarios

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r

}
