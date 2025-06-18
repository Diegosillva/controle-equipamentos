package handler

import (
	"controle/internal/model"
	"controle/internal/repository"
	"controle/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Minha Api...")
}

func GetEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Error ao conectar ao banco", http.StatusInternalServerError)
	}
	defer db.Close()

	equipamentos, err := repository.ListaEquipamentos(db)
	if err != nil {
		http.Error(w, "Error ao listar equipamentos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonFormat, err := json.MarshalIndent(equipamentos, "", " ")
	if err != nil {
		http.Error(w, "Error ao formatar JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonFormat)
}

func UpdateEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	var e model.Equipamentos
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if e.ID == 0 {
		http.Error(w, "ID do equipamento obrigatorio", http.StatusBadRequest)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Error ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = repository.EditarEquipamentos(db, e)
	if err != nil {
		http.Error(w, "Erro ao editar o Equipamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetByProdutoEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	produto := r.URL.Query().Get("nome")
	if produto == "" {
		http.Error(w, "Parametro 'nome' e obrigatorio", http.StatusBadRequest)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Error ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	equipamentos, err := repository.BuscarEquipamentoPorProduto(db, produto)
	if err != nil {
		http.Error(w, "Equipamento não econtrado.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonFormat, err := json.MarshalIndent(equipamentos, "", " ")
	if err != nil {
		http.Error(w, "Error ao formatar JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonFormat)
}

func PostEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	var e model.Equipamentos
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Erro ao conectar no banco.", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	err = repository.CriarEquipamento(db, e)
	if err != nil {
		http.Error(w, "Erro ao criar o Equipamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func DeleteByIdEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	db, _ := service.OpenDB()
	defer db.Close()

	equipamento, err := repository.BuscaEquipamentoPorID(db, id)
	if err != nil {
		http.Error(w, "Equipamento não encontrado.", http.StatusNotFound)
		return
	}

	err = repository.DeletarEquipamento(db, id)
	if err != nil {
		http.Error(w, "Error ao deletar equipamento", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonFormat, err := json.MarshalIndent(equipamento, "", " ")
	if err != nil {
		http.Error(w, "Error ao formatar JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonFormat)
}
