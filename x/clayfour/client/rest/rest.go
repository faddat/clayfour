package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers clayfour-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/clayfour/post", createPostHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/clayfour/post", listPostHandler(cliCtx, "clayfour")).Methods("GET")
		r.HandleFunc("/clayfour/post/{key}", getPostHandler(cliCtx, "clayfour")).Methods("GET")
		r.HandleFunc("/clayfour/post", setPostHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/clayfour/post", deletePostHandler(cliCtx)).Methods("DELETE")

		
}
