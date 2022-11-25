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
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	if erro = usuario.Preparar("cadastro"); erro != nil {
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
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := database.Conectar()
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositories := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositories.Buscar(nomeOuNick)
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	msgresponse.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	msgresponse.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		msgresponse.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	msgresponse.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	msgresponse.JSON(w, http.StatusNoContent, nil)
}
