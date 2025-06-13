package router

import (
	"controle/internal/handler"
	"net/http"

	"github.com/rs/cors"
)

func Rotas() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/equipamentos/api/v1/list", handler.GetEquipamentos)
	mux.HandleFunc("/equipamentos/api/v1/produto", handler.GetByProdutoEquipamentos)
	mux.HandleFunc("/equipamentos/api/v1/produto/delete", handler.DeleteByIdEquipamentos)
	mux.HandleFunc("/equipamentos/api/v1/create", handler.PostEquipamentos)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})

	return c.Handler(mux)
}
