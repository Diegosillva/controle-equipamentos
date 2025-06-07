package handler

import (
	"controle/internal/model"
	"controle/internal/service"
	"database/sql"
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
		http.Error(w, "Erro ao conectar ao Banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, produto, equipamento, modelo, numero_de_serie,
		serial_dsp, descricao FROM cadastro_equipamentos`)
	if err != nil {
		http.Error(w, "Erro ao consultar no Banco", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipamentos []model.Equipamentos

	for rows.Next() {
		var e model.Equipamentos
		err := rows.Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
			&e.SerialDSP, &e.Descricao)
		if err != nil {
			http.Error(w, "Erro ao ler os dados", http.StatusInternalServerError)
			return
		}
		equipamentos = append(equipamentos, e)
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
		http.Error(w, "Erro ao conectar ao Banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
	SELECT id, produto, equipamento, modelo, numero_de_serie,
	serial_dsp, descricao FROM cadastro_equipamentos
	WHERE produto = $1
	`
	rows, err := db.Query(query, produto)
	if err != nil {
		http.Error(w, "Erro ao consultar no Banco", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipamentos []model.Equipamentos

	for rows.Next() {
		var e model.Equipamentos
		err := rows.Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
			&e.SerialDSP, &e.Descricao)
		if err != nil {
			http.Error(w, "Erro ao ler os dados", http.StatusInternalServerError)
			return
		}
		equipamentos = append(equipamentos, e)
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

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao Banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	var id int

	query := `
		INSERT INTO cadastro_equipamentos (produto,equipamento,modelo,
		numero_de_serie,serial_dsp,descricao)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	err = db.QueryRow(query, e.Produto, e.Equipamento, e.Modelo, e.NumeroSerie,
		e.SerialDSP, e.Descricao).Scan(&id)
	if err != nil {
		http.Error(w, "Erro ao inserir no Banco", http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"id":           id,
		"produto":      e.Produto,
		"equipamento":  e.Equipamento,
		"modelo":       e.Modelo,
		"numero_serie": e.NumeroSerie,
		"serial_dsp":   e.SerialDSP,
		"descricao":    e.Descricao,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func DeleteByIdEquipamentos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Parametro 'id' e obrigatorio", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	db, err := service.OpenDB()
	if err != nil {
		http.Error(w, "Erro ao conectar ao Banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
	DELETE FROM cadastro_equipamentos
	WHERE id = $1
	RETURNING id, produto, equipamento, modelo, 
	numero_de_serie, serial_dsp, descricao
	`
	var e model.Equipamentos

	row := db.QueryRow(query, id)
	err = row.Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
		&e.SerialDSP, &e.Descricao)

	if err == sql.ErrNoRows {
		http.Error(w, "Equipamento não encontrado.", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Erro ao excluir no banco", http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonFormat, err := json.MarshalIndent(e, "", " ")
	if err != nil {
		http.Error(w, "Error ao formatar JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonFormat)
}
