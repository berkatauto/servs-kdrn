package login

import (
	"net/http"

	"github.com/whatsauth/watoken"
)

var Privatekey = "56e4eb16f428e82cea21e5bceed2b078c0955ce1b8509631369dab20e1a952180a9ea5fae87b3895fba98c2b138c336ccfba886b0823fd774415ccc9394ae159"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not authorized.", http.StatusMethodNotAllowed)
		return
	}

	userid := r.FormValue("userid")
	userpass := r.FormValue("userpass")

	if checkCredentials(userid, userpass) {
		tokenString, err := watoken.Encode(userid, Privatekey)
		if err != nil {
			http.Error(w, "Token Generating Fail", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	} else {
		http.Error(w, "Login Failed", http.StatusUnauthorized)
	}
}

func checkCredentials(userid, userpass string) bool {
	return userid == "userid" && userpass == "userpass"
}
