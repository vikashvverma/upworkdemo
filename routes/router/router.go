package router

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/vikashvverma/upworkdemo/routes/callback"
	"github.com/vikashvverma/upworkdemo/routes/home"
	"github.com/vikashvverma/upworkdemo/routes/login"
	"github.com/vikashvverma/upworkdemo/routes/logout"
	"github.com/vikashvverma/upworkdemo/routes/middlewares"
	"github.com/vikashvverma/upworkdemo/routes/user"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	return r
}
func notFound(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusForbidden)
}
