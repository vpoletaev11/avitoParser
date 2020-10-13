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
	getEmailsRelWithLink = "SELECT email FROM email WHERE link=?;"
)

type linkPrice struct {
	link  string
	price int
}

// Dep stores dependencies
type Dep struct {
	DB     *sql.DB
	Client *http.Client

	minsToLoop      int
	secToGetOnePage int
	senderEmail     string
	senderEmailPass string
}

// NewDep initializes Dep
func NewDep() Dep {
	// Set DB
	mySQLAddr := os.Getenv("MYSQL_ADDR")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", mySQLAddr+"/"+dbName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to MySql database")

	// Set http client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}

	// Set timings
	minsToLoopStr := os.Getenv("MIN_TO_SCRAPPING_ALL_LINKS")
	minsToLoop, err := strconv.Atoi(minsToLoopStr)
	if err != nil {
		panic(err)
	}
	secToGetOnePageStr := os.Getenv("SEC_TO_GET_ONE_PAGE")
	secToGetOnePage, err := strconv.Atoi(secToGetOnePageStr)
	if err != nil {
		panic(err)
	}

	// Set sender email configs
	senderEmail := os.Getenv("SENDER_MAIL")
	senderEmailPass := os.Getenv("MAIL_PASSWORD")

	fmt.Println(minsToLoop, secToGetOnePage, senderEmail, senderEmailPass)
	return Dep{
		DB:     db,
		Client: client,

		minsToLoop:      minsToLoop,
		secToGetOnePage: secToGetOnePage,
		senderEmail:     senderEmail,
		senderEmailPass: senderEmailPass,
	}
}

// ComparePrices checks if links in database have actual prices using parsing avito.ru.
// If links prices aren't actual:
// 1) rewrite link prices in db;
// 2) send emails to addresses related with changed price links.
func ComparePrices(dep Dep) {
	for {
		links := getLinksAndPriceFromDB(dep.DB)

		changedPriceLinks := linksWithChangedPrice(dep, links)

		sendMails(dep, changedPriceLinks)

		time.Sleep(time.Duration(dep.minsToLoop) * time.Minute)
	}
}

// getLinksAndPriceFromDB gets all links and prices from database.
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

// linksWithChangedPrice checks if inputted links have actual price by comparing it prices with prices getted by parsing avito.ru.
// If inputted links haven't actual price linksWithChangedPrice writes for them actual price in db.
func linksWithChangedPrice(dep Dep, links []linkPrice) []linkPrice {
	changedPriceLinks := []linkPrice{}
	for _, lp := range links {
		price, err := dep.GetPrice(lp.link)
		if err != nil {
			errhand.InternalErrorLog(err)
			continue
		}

		if price != lp.price {
			lp.price = price
			_, err := dep.DB.Exec(updatePrice, lp.price, lp.link)
			if err != nil {
				errhand.InternalErrorLog(err)
				continue
			}
			changedPriceLinks = append(changedPriceLinks, lp)
		}

		time.Sleep(time.Duration(dep.secToGetOnePage) * time.Second)
	}

	return changedPriceLinks
}

// sendMails gets from db email addresses related with inserted links and sends to them notifications about changing price.
func sendMails(dep Dep, changedPriceLinks []linkPrice) {
	for _, lp := range changedPriceLinks {
		rows, err := dep.DB.Query(getEmailsRelWithLink, lp.link)
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

		dep.sendMail(emails, lp.link, lp.price)
	}
}

// sendMail sends notifications about changing price to receivers.
// This func uses system variables to get sender email-addres and password (you can configure it in docker-compose.yml).
func (dep Dep) sendMail(receiver []string, link string, price int) {
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	fromL := fmt.Sprintf("From: <%s>\r\n", dep.senderEmail)
	subject := "Subject: Avito parser\r\n"
	body := "New price of " + link + " amount " + strconv.Itoa(price) + " rub."

	message := []byte(fromL + subject + "\r\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", dep.senderEmail, dep.senderEmailPass, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, dep.senderEmail, receiver, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// GetPrice parses avito.ru ad page to get price.
func (dep Dep) GetPrice(link string) (int, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return 0, err
	}

	resp, err := dep.Client.Do(req)
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
