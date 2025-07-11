package repository

import (
	"controle/internal/model"
	"database/sql"
)

func ListaEquipamentos(db *sql.DB) ([]model.Equipamentos, error) {
	rows, err := db.Query(`SELECT id, produto, equipamento, modelo, numero_de_serie,
		serial_dsp,localizacao, status, descricao FROM cadastro_equipamentos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipamentos []model.Equipamentos

	for rows.Next() {
		var e model.Equipamentos
		err := rows.Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
			&e.SerialDSP, &e.Localizacao, &e.Status, &e.Descricao)
		if err != nil {
			return nil, err
		}
		equipamentos = append(equipamentos, e)
	}
	return equipamentos, nil
}

func EditarEquipamentos(db *sql.DB, e model.Equipamentos) error {
	query := ` UPDATE cadastro_equipamentos SET produto = $1, equipamento = $2, 
	modelo = $3, numero_de_serie = $4, serial_dsp = $5, localizacao = $6, 
	status = $7, descricao = $8 WHERE id = $9 `
	_, err := db.Exec(query, e.Produto, e.Equipamento, e.Modelo, e.NumeroSerie,
		e.SerialDSP, e.Localizacao, e.Status, e.Descricao, e.ID)
	return err
}

func BuscaEquipamentoPorID(db *sql.DB, id int) (model.Equipamentos, error) {
	var e model.Equipamentos
	err := db.QueryRow(`
		SELECT id, produto, equipamento, modelo, numero_de_serie,
		serial_dsp, localizacao, status, descricao FROM cadastro_equipamentos
		WHERE id = $1
		`, id).Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
		&e.SerialDSP, &e.Localizacao, &e.Status, &e.Descricao)
	return e, err
}

func BuscarEquipamentoPorProduto(db *sql.DB, produto string) (*model.Equipamentos, error) {
	row := db.QueryRow(`
	SELECT id, produto, equipamento, modelo, numero_de_serie,
	serial_dsp, localizacao, status, descricao FROM cadastro_equipamentos
	WHERE produto = $1
	`, produto)
	var e model.Equipamentos

	err := row.Scan(&e.ID, &e.Produto, &e.Equipamento, &e.Modelo, &e.NumeroSerie,
		&e.SerialDSP, &e.Localizacao, &e.Status, &e.Descricao)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func CriarEquipamento(db *sql.DB, e model.Equipamentos) error {
	_, err := db.Exec(`INSERT INTO cadastro_equipamentos (produto, equipamento,
	modelo, numero_de_serie, serial_dsp, localizacao, status, descricao) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8)`,
		e.Produto, e.Equipamento, e.Modelo, e.NumeroSerie, e.SerialDSP, e.Localizacao,
		e.Status, e.Descricao)
	return err
}

func DeletarEquipamento(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM cadastro_equipamentos WHERE id = $1`, id)
	return err
}
