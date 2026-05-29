package mlog

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

var Exdbmysqlg *sql.DB
var encryptionKey = "13OtdSecret"
var LoggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))
var sGDisplayName = ""

type Page struct {
	Date        string
	Username    string
	Displayname string
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	conditionsMap := map[string]interface{}{}

	if r.FormValue("Login") != "" && r.FormValue("Username") != "" {
		username := r.FormValue("Username")
		password := r.FormValue("Password")

		if username == "mal" && password == "parol" {
			sGDisplayName = "Малышев"
			conditionsMap["Username"] = username
			conditionsMap["LoginError"] = false
			session, _ := LoggedUserSession.New(r, "my-user-session")
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/index", http.StatusFound)
			return
		} else {
			conditionsMap["LoginError"] = true
		}
	}

	tmpl := template.Must(template.ParseFiles("public/html/login.html"))
	tmpl.Execute(w, conditionsMap)
}

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := LoggedUserSession.Get(r, "my-user-session")
	if session.Values["username"] == nil || session.Values["username"] == "" {
		http.Redirect(w, r, "/logout", http.StatusFound)
		return
	}

	year, month, day := time.Now().Date()
	curdate := fmt.Sprintf("%02d.%02d.%d", day, month, year)
	username := fmt.Sprintf("%v", session.Values["username"])

	p := &Page{
		Date:        curdate,
		Username:    username,
		Displayname: sGDisplayName,
	}

	t := template.Must(template.ParseFiles("public/html/index.html"))
	t.Execute(w, p)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := LoggedUserSession.Get(r, "my-user-session")
	session.Values["username"] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func CheckLoginPOST(w http.ResponseWriter, r *http.Request) int {
	session, _ := LoggedUserSession.Get(r, "my-user-session")
	if session.Values["username"] == nil || session.Values["username"] == "" {
		return 0
	}
	return 1
}
