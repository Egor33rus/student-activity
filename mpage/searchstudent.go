package mpage

import (
	"fmt"
	"net/http"
	MODMLOG "student-activity/mlog"
)

var Exdbmysqlg interface{}

func SearchStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "GET request not allowed")
		return
	}

	if MODMLOG.CheckLoginPOST(w, r) == 0 {
		fmt.Fprintf(w, "0###/")
		return
	}

	r.ParseMultipartForm(64 << 20)
	searchstr := r.FormValue("searchstr")

	sOut := "<table class='table'>台<th>ФИО</th><th>ВУЗ</th></thead><tbody>"
	sOut += "<tr><td>Иванов Иван</td><td>МГУ</td></tr>"
	sOut += "<tr><td>Петров Петр</td><td>СПбГУ</td></tr>"
	sOut += "</tbody></table>"

	fmt.Fprintf(w, "%v", sOut)
}
