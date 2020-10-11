package subscribe

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/vpoletaev11/avitoParser/errhand"
	"github.com/vpoletaev11/avitoParser/scrapper"
)

const (
	insertEmail        = "INSERT INTO email(email, link) VALUES(?, ?);"
	insertLink         = "INSERT INTO link(link, price) VALUES(?, ?);"
	updateLinkForEmail = "UPDATE email SET link=? WHERE email=?;"
)

func Handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			errhand.InternalError(err, w)
			return
		}

		email := r.PostForm.Get("email")
		link := r.PostForm.Get("link")

		price, err := scrapper.GetPrice(link)
		if err != nil {
			errhand.InternalError(err, w)
			return
		}

		_, err = db.Exec(insertLink, link, price)
		if err != nil {
			// if link already exists in database- do nothing
			if strings.Contains(err.Error(), "Error 1062") {

			} else {
				errhand.InternalError(err, w)
				return
			}
		}

		_, err = db.Exec(insertEmail, email, link)
		if err != nil {
			// if email already exists in database- rewrite related to it link
			if strings.Contains(err.Error(), "Error 1062") {
				_, err := db.Exec(updateLinkForEmail, link, email)
				if err != nil {
					errhand.InternalError(err, w)
					return
				}
				return
			}
			errhand.InternalError(err, w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
