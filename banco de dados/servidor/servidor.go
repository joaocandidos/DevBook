package servidor

import (
	"banco-de-dados/banco"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
)

type usuario struct {
	ID    int    `json: id`
	Nome  string `json: nome`
	Email string `json: email`
}

// CriarUsuariovinsere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("erro ao ler o corpo da requisicao"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("erro ao converter o usuario para struct"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	//evitar ataque sql injection
	statement, erro := db.Prepare(" insert into usuario (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("erro ao criar o statement"))
		return

	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("erro ao executar o statement"))
		return
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("erro ao obter o id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso! ID: %d", idInserido)))

}

// BuscarUsuarios busca todos os usuarios do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuario")
	if erro != nil {
		w.Write([]byte("erro ao buscar todos os usuarios "))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("erro ao escanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)

	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("erro ao converter o usuario para json"))
		return
	}
}

// BuscarUsuario busca usuario por id
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
