package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/edaaltuntas/gomatching/pkg/sqlite"
	"github.com/edaaltuntas/gomatching/pkg/sqlite/models"
)

// SignUpHandler in charge of registration steps, after a person
// triggers this handler, stores persons information at database,
// it sets user_id cookie
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := models.Person{
		Name:     r.FormValue("name"),
		Category: r.FormValue("category"),
		City:     r.FormValue("city"),
	}
	if err := sqlite.Conn.Create(&p).Error; err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    strconv.FormatUint(uint64(p.ID), 10),
		Path:     "/",
		HttpOnly: true,
	})
}

// HomeHandler is very simple file serving handler. It shows
// the signup page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pkg/web/templates/register.html")
}

// ResultHandler is very simple file serving handler. It shows
// the result page.
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pkg/web/templates/result.html")
}

// MatchingHandler uses SuperMatchingEngine and brings people
// closer :)
func MatchingHandler(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	var p models.Person
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	if res := sqlite.Conn.Where("ID = ?", cookie.Value).Find(&p); res.Error != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	if res := sqlite.Conn.Where("ID != ?", cookie.Value).Find(&persons); res.Error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		for i := range persons {
			if p.City == persons[i].City && p.Category != persons[i].Category {
				persons[i].MatchingScore = 100
			} else if p.City == persons[i].City && p.Category == persons[i].Category {
				persons[i].MatchingScore = 50
			} else if p.City != persons[i].City && p.Category != persons[i].Category {
				persons[i].MatchingScore = 50
			} else {
				persons[i].MatchingScore = 0
			}
		}
		ret, err := json.Marshal(&persons)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(ret)
		}
	}
}
