package scrapper

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
)

const (
	getLinksAndPrice     = "SELECT * FROM link;"
	updatePrice          = "UPDATE link SET price=? WHERE link=?;"
	getEmailsRelWithLink = "SELECT email FROM email WHERE link=?"
)

type linkPrice struct {
	link  string
	price int
}

func ComparePrices(db *sql.DB) {
	links, err := getLinksAndPriceFromDB(db)
	if err != nil {
		log.Println(err)
		return
	}

	changedPriceLinks, err := linksWithChangedPrice(db, links)
	if err != nil {
		log.Println(err)
		return
	}

	err = sendMails(db, changedPriceLinks)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(links)
	fmt.Println(changedPriceLinks)
}

func getLinksAndPriceFromDB(db *sql.DB) ([]linkPrice, error) {
	rows, err := db.Query(getLinksAndPrice)
	if err != nil {
		return []linkPrice{}, err
	}
	defer rows.Close()

	lp := linkPrice{}
	links := []linkPrice{}
	for rows.Next() {
		err := rows.Scan(
			&lp.link,
			&lp.price,
		)
		if err != nil {
			return []linkPrice{}, err
		}

		links = append(links, lp)
	}

	return links, nil
}

func linksWithChangedPrice(db *sql.DB, links []linkPrice) ([]linkPrice, error) {
	changedPriceLinks := []linkPrice{}
	for _, lp := range links {
		price, err := GetPrice(lp.link)
		if err != nil {
			log.Println(err)
			continue
		}

		if price != lp.price {
			lp.price = price
			_, err := db.Exec(updatePrice, lp.price, lp.link)
			if err != nil {
				return []linkPrice{}, err
			}
			changedPriceLinks = append(changedPriceLinks, lp)
		}
	}

	return changedPriceLinks, nil
}

func sendMails(db *sql.DB, changedPriceLinks []linkPrice) error {
	for _, lp := range changedPriceLinks {
		rows, err := db.Query(getEmailsRelWithLink, lp.link)
		if err != nil {
			return err
		}
		defer rows.Close()

		email := ""
		emails := []string{}
		for rows.Next() {
			err := rows.Scan(
				&email,
			)
			if err != nil {
				return err
			}

			emails = append(emails, email)
		}

		sendMail(emails, lp.link, lp.price)
	}

	return nil
}

func sendMail(receiver []string, link string, cost int) {
	// Sender data.
	from := "SENDER_MAIL"
	password := "MAIL_PASSWORD"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	fromL := fmt.Sprintf("From: <%s>\r\n", from)
	subject := "Avito parser\r\n"
	body := "New price of " + link + " amount " + strconv.Itoa(cost) + " rub."

	message := []byte(fromL + subject + "\r\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, receiver, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func GetPrice(link string) (int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	bodyString := string(body)

	priceStr, _ := getStringBetweenTwoStrings(bodyString, `"dynx_price":`, ",")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func getStringBetweenTwoStrings(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}
