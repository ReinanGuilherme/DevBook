package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

// preparando o struct
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

// recusando campos em branco
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é um campo obrigatório e não pode estar em branco.")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é um campo obrigatório e não pode estar em branco.")
	}

	if usuario.Email == "" {
		return errors.New("O email é um campo obrigatório e não pode estar em branco.")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O senha é um campo obrigatório e não pode estar em branco.")
	}

	return nil
}

// removendo espaços no texto
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
