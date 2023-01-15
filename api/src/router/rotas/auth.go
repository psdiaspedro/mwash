package rotas

import (
	"api/src/controllers"
)

/*
Rota principal para acesso à plataforma
- Não precisa estar logado
- Todos os tipos de usuário estão liberados
*/
var rotaAuth = Route{
	URI:        "/usuario/auth",
	Metodo:     "GET",
	Funcao:     controllers.ValidarToken,
	RequerAuth: true,
}
