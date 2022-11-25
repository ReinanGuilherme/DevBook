package controllers

import (
	"api/src/database"
	"api/src/models"
	msgresponse "api/src/msgResponse"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		msgresponse.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//convertendo de JSON para STRUCT
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//validando struct usuario
	if erro = usuario.Preparar(); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//abrindo conexão com o banco de dados
	db, erro := database.Conectar()
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	//passando conexão para o repositorio
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario.ID = usuarioID
	msgresponse.JSON(w, http.StatusCreated, usuario)

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário!"))
}
