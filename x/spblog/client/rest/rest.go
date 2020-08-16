package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers spblog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/spblog/user", listUserHandler(cliCtx, "spblog")).Methods("GET")
	r.HandleFunc("/spblog/user", createUserHandler(cliCtx)).Methods("POST")
}
