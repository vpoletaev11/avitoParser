package subscribe

import (
	"net/http"
	"strings"

	"github.com/vpoletaev11/avitoParser/errhand"
	"github.com/vpoletaev11/avitoParser/scrapper"
)

const (
	insertEmail            = "INSERT INTO email(email, link) VALUES(?, ?);"
	insertLink             = "INSERT INTO link(link, price) VALUES(?, ?);"
	updateLinkForEmail     = "UPDATE email SET link=? WHERE email=?;"
	getEmailsRelatedToLink = "SELECT COUNT(email) FROM email WHERE link=?;"
	deleteLink             = "DELETE FROM link WHERE link=?;"
)

// Handler receives subscriptions to price updating and puts them into database after validation.
// Subscriptions should come via POST method and contains [email] and [link] fields.
// Handler can store only one link for one email addres.
func Handler(dep scrapper.Dep) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			errhand.InternalError(err, w)
			return
		}

		email := r.PostForm.Get("email")
		link := r.PostForm.Get("link")

		// link validation
		price, err := scrapper.GetPrice(dep, link)
		if err != nil {
			errhand.InternalError(err, w)
			return
		}

		_, err = dep.DB.Exec(insertLink, link, price)
		if err != nil {
			// if link already exists in database- do nothing
			if strings.Contains(err.Error(), "Error 1062") {

			} else {
				errhand.InternalError(err, w)
				return
			}
		}

		_, err = dep.DB.Exec(insertEmail, email, link)
		if err != nil {
			// if email already exists in database- rewrite related to it link
			if strings.Contains(err.Error(), "Error 1062") {
				_, err := dep.DB.Exec(updateLinkForEmail, link, email)
				if err != nil {
					errhand.InternalError(err, w)
					return
				}

				countEmails := 0
				err = dep.DB.QueryRow(getEmailsRelatedToLink, link).Scan(&countEmails)
				if err != nil {
					errhand.InternalError(err, w)
					return
				}
				if countEmails <= 1 {
					_, err := dep.DB.Exec(deleteLink, link)
					if err != nil {
						errhand.InternalError(err, w)
						return
					}
				}
				return
			}
			errhand.InternalError(err, w)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
