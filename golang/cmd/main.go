package main 

import (
	"net/http"
	"log"
	"fmt"

	"controle/internal/service"
	"controle/internal/router"
)

func main(){

	db,err := service.OpenDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no Banco de Dados: %v", err)
	}
	defer db.Close()

	fmt.Println("Banco de dados connectado com sucesso")

	rotas := router.Rotas()

	log.Println("Servido iniciado na porta :8080")
	err = http.ListenAndServe(":8080",rotas)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v",err)
	}
}
