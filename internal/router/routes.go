package router 

import(
	"net/http"
	"controle/internal/handler"
)

func Rotas()http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandlerHome)
	mux.HandleFunc("/equipamentos/api/v1/list", handler.GetEquipamentos)
	mux.HandleFunc("/equipamentos/api/v1/produto", handler.GetByProdutoEquipamentos)
	mux.HandleFunc("/equipamentos/api/v1/create", handler.PostEquipamentos)
	return mux
}
