package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}
func BuscarTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscando todos usuario"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando usuario"))
}
