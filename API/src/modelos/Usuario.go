package modelos

import "time"

// Usuario representa os dados de um usuário que está utilizando a rede social
type Usuario struct {
	ID       uint   `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Senha    string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}