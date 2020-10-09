package subscribe

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

const (
	insertEmail        = "INSERT INTO email(email, link) VALUES(?, ?);"
	insertLink         = "INSERT INTO link(link) VALUES(?);"
	updateLinkForEmail = "UPDATE email SET link=? WHERE email=?;"
)

func Handler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		email := r.PostForm.Get("email")
		link := r.PostForm.Get("link")

		_, err = db.Exec(insertLink, link)
		if err != nil {
			// if link already exists in database- do nothing
			if strings.Contains(err.Error(), "Error 1062") {

			} else {
				fmt.Fprintf(w, err.Error())
				return
			}
		}

		_, err = db.Exec(insertEmail, email, link)
		if err != nil {
			// if email already exists in database- rewrite related to it link
			if strings.Contains(err.Error(), "Error 1062") {
				_, err := db.Exec(updateLinkForEmail, link, email)
				if err != nil {
					fmt.Fprintf(w, err.Error())
					return
				}
				return
			}
			fmt.Fprintf(w, err.Error())
			return
		}
	}
}
