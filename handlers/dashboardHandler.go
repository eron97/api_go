package handlers

/*


Esta pasta é usada para armazenar os handlers (controladores) da sua API.
Os handlers são responsáveis por receber as solicitações HTTP, processá-las
e retornar as respostas apropriadas. Cada rota da sua API geralmente terá um handler associado.

*/

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem-vindo ao dashboard do seu perfil!")
}
