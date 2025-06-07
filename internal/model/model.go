package model

type Equipamentos struct {
	ID int `json:"id"`
	Produto string `json:"produto"` 
	Equipamento  string `json:"equipamento"`
	Modelo string `json:"modelo"`
	NumeroSerie string `json:"numero_serie"`
	SerialDSP string `json:"serial_dsp"`
	Descricao string `json:"descricao"`
}
