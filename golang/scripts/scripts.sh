# para rodar um banco de dados atraves de um arquivo.
psql -U {usuario} -d {banco} -f {arquivo .sql} 

# POST
curl http://localhost:8080/equipamentos/api/v1/create \
-H "Content-Type: application/json" \
-d '{
"produto":"Manila",
"equipamento":"Prensa",
"modelo":"BE-30",
"numero_serie":"00000000000",
"serial_dsp":"DSP-1020",
"descricao":"Problema na chave fim de curso."
}'


# GET por produto
curl "http://localhost:8080/equipamentos/api/v1/produto?nome=Manila"

# DELETE por ID
curl -X DELETE "http://localhost:8080/equipamentos/api/v1/produto/delete?id=3"
