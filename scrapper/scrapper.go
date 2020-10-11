package scrapper

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vpoletaev11/avitoParser/errhand"
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
	minsToLoopStr := os.Getenv("MIN_TO_SCRAPPING_ALL_LINKS")
	minsToLoop, err := strconv.Atoi(minsToLoopStr)
	if err != nil {
		panic(err)
	}

	for {
		links := getLinksAndPriceFromDB(db)

		changedPriceLinks := linksWithChangedPrice(db, links)

		sendMails(db, changedPriceLinks)

		time.Sleep(time.Duration(minsToLoop) * time.Minute)
	}
}

func getLinksAndPriceFromDB(db *sql.DB) []linkPrice {
	rows, err := db.Query(getLinksAndPrice)
	if err != nil {
		errhand.InternalErrorLog(err)
		return []linkPrice{}
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
			errhand.InternalErrorLog(err)
			continue
		}

		links = append(links, lp)
	}

	return links
}

func linksWithChangedPrice(db *sql.DB, links []linkPrice) []linkPrice {
	secToGetOnePageStr := os.Getenv("SEC_TO_GET_ONE_PAGE")
	secToGetOnePage, err := strconv.Atoi(secToGetOnePageStr)
	if err != nil {
		panic(err)
	}

	changedPriceLinks := []linkPrice{}
	for _, lp := range links {
		price, err := GetPrice(lp.link)
		if err != nil {
			errhand.InternalErrorLog(err)
			continue
		}

		if price != lp.price {
			lp.price = price
			_, err := db.Exec(updatePrice, lp.price, lp.link)
			if err != nil {
				errhand.InternalErrorLog(err)
				continue
			}
			changedPriceLinks = append(changedPriceLinks, lp)
		}

		time.Sleep(time.Duration(secToGetOnePage) * time.Second)
	}

	return changedPriceLinks
}

func sendMails(db *sql.DB, changedPriceLinks []linkPrice) {
	for _, lp := range changedPriceLinks {
		rows, err := db.Query(getEmailsRelWithLink, lp.link)
		if err != nil {
			errhand.InternalErrorLog(err)
			return
		}
		defer rows.Close()

		email := ""
		emails := []string{}
		for rows.Next() {
			err := rows.Scan(
				&email,
			)
			if err != nil {
				errhand.InternalErrorLog(err)
				continue
			}

			emails = append(emails, email)
		}

		sendMail(emails, lp.link, lp.price)
	}
}

func sendMail(receiver []string, link string, price int) {
	// Sender data.
	from := os.Getenv("SENDER_MAIL")
	password := os.Getenv("MAIL_PASSWORD")

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	fromL := fmt.Sprintf("From: <%s>\r\n", from)
	subject := "Subject: Avito parser\r\n"
	body := "New price of " + link + " amount " + strconv.Itoa(price) + " rub."

	message := []byte(fromL + subject + "\r\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, receiver, message)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	bodyString := string(body)

	priceStr := getStringBetweenTwoStrings(bodyString, `"dynx_price":`, ",")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func getStringBetweenTwoStrings(str string, startS string, endS string) (result string) {
	s := strings.Index(str, startS)
	if s == -1 {
		return ""
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return ""
	}
	result = newS[:e]
	return result
}
