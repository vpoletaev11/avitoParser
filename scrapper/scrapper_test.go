package scrapper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLinksAndPriceFromDBSuccess(t *testing.T) {
	dep, sqlMock, ts := NewTestDepAndServer()
	link := []string{
		"link",
		"price",
	}
	sqlMock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(link).AddRow(
		ts.URL,
		1000000,
	))

	links := getLinksAndPriceFromDB(dep.DB)
	assert.Equal(t, []linkPrice{{ts.URL, 1000000}}, links)
}

func TestGetLinksAndPriceFromDBErrorDBConnect(t *testing.T) {
	dep, sqlMock, _ := NewTestDepAndServer()
	sqlMock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("Some error with DB connection"))

	links := getLinksAndPriceFromDB(dep.DB)
	assert.Equal(t, []linkPrice{}, links)
}

func TestGetLinksAndPriceFromDBBadValuesDBAnswer(t *testing.T) {
	dep, sqlMock, _ := NewTestDepAndServer()
	link := []string{
		"link",
		"price",
	}
	sqlMock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(link).AddRow(
		1,
		"wrong data",
	))

	links := getLinksAndPriceFromDB(dep.DB)
	assert.Equal(t, []linkPrice{}, links)
}

func TestLinksWithChangedPriceNonePriceChangesSuccess(t *testing.T) {
	dep, _, ts := NewTestDepAndServer()

	links := linksWithChangedPrice(dep, []linkPrice{{ts.URL, 1000000}})
	assert.Equal(t, []linkPrice{}, links)
}

func TestLinksWithChangedPriceGetPriceError(t *testing.T) {
	dep, _, _ := NewTestDepAndServer()

	links := linksWithChangedPrice(dep, []linkPrice{{"Bad link", 1000000}})
	assert.Equal(t, []linkPrice{}, links)
}

func TestLinksWithChangedPricePriceChangesSuccess(t *testing.T) {
	dep, sqlMock, ts := NewTestDepAndServer()
	sqlMock.ExpectExec("UPDATE").WithArgs(1000000, ts.URL).WillReturnResult(sqlmock.NewResult(1, 1))

	links := linksWithChangedPrice(dep, []linkPrice{{ts.URL, 10}})
	assert.Equal(t, []linkPrice{{ts.URL, 1000000}}, links)
}

func TestLinksWithChangedPricePriceChangesDBErrorUpdate(t *testing.T) {
	dep, sqlMock, ts := NewTestDepAndServer()
	sqlMock.ExpectExec("UPDATE").WithArgs(1000000, ts.URL).WillReturnError(fmt.Errorf("Some error with DB connection"))

	links := linksWithChangedPrice(dep, []linkPrice{{ts.URL, 10}})
	assert.Equal(t, []linkPrice{}, links)
}

func TestSendMailsSuccess(t *testing.T) {
	dep, sqlMock, _ := NewTestDepAndServer()
	email := []string{
		"email",
	}
	sqlMock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(email).AddRow(
		"exmaple@mail.com",
	))
	sendMails(dep, []linkPrice{{"expectedLink", 1000000}})
}

func TestSendMailsDBError(t *testing.T) {
	dep, sqlMock, _ := NewTestDepAndServer()
	sqlMock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("Some error with DB connection"))
	sendMails(dep, []linkPrice{{"expectedLink", 1000000}})
}

func TestSendMailsWrondDBOutput(t *testing.T) {
	dep, sqlMock, _ := NewTestDepAndServer()
	badData := []string{
		"cat",
		"dog",
	}
	sqlMock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(badData).AddRow(
		1,
		2,
	))
	sendMails(dep, []linkPrice{{"expectedLink", 1000000}})
}

func TestGetPriceSuccess(t *testing.T) {
	dep, _, ts := NewTestDepAndServer()
	price, err := GetPrice(dep, ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, 1000000, price)
}

// NewTestDepAndServer returns: dependencies, Sqlmock interface to add mocks and httptest server, that emulating avito ad page
func NewTestDepAndServer() (Dep, sqlmock.Sqlmock, *httptest.Server) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, avitoAdHTML)
	}))

	dep := Dep{
		DB:     db,
		Client: ts.Client(),

		sender: testEmailConfig{},

		minsToLoop:      0,
		secToGetOnePage: 0,
	}

	return dep, sqlMock, ts
}

type testEmailConfig struct {
	smtpHost        string
	smtpPort        string
	senderEmail     string
	senderEmailPass string
}

// sendEmail sends notifications about changing price to receivers.
func (t testEmailConfig) sendEmail(receiver []string, link string, price int) {
	if receiver[0] != "exmaple@mail.com" {
		panic("sendMail expects to get exmaple@mail.com as receiver")
	}
	if link != "expectedLink" {
		panic("sendMail expects to get expectedLink as link")
	}
	if price != 1000000 {
		panic("sendMail expects to get 1000000 as price")
	}
}

const avitoAdHTML = `<!DOCTYPE html>
            
<html> <head> <script>
 try {
 window.firstHiddenTime = document.visibilityState === 'hidden' ? 0 : Infinity;
 document.addEventListener('visibilitychange', function (event) {
 window.firstHiddenTime = Math.min(window.firstHiddenTime, event.timeStamp);
 }, { once: true });
 if ('PerformanceLongTaskTiming' in window) {
 var globalStats = window.__statsLongTasks = { tasks: [] };
 globalStats.observer = new PerformanceObserver(function(list) {
 globalStats.tasks = globalStats.tasks.concat(list.getEntries());
 });
 globalStats.observer.observe({ entryTypes: ['longtask'] });
 }
 if (PerformanceObserver && (PerformanceObserver.supportedEntryTypes || []).some(function(e) {
 return e === 'element'
 })) {
 if (!window.oet) {
 window.oet = [];
 }
 new PerformanceObserver(function(l) {
 window.oet.push.apply(window.oet, l.getEntries());
 }).observe({ entryTypes: ['element'] });
 }
 } catch (e) {
 console.error(e);
 }
 </script>
    <script>
 window.dataLayer = [{"dynx_user":"a","dynx_region":"moskva","dynx_prodid":2009709110,"dynx_price":1000000,"dynx_category":"kollektsionirovanie","dynx_vertical":4,"dynx_pagetype":"item"},{"pageType":"Item","itemID":2009709110,"vertical":"GENERAL","categoryId":36,"categorySlug":"kollektsionirovanie","microCategoryId":320,"locationId":637640,"isShop":0,"isClientType1":0,"itemPrice":1000000,"withDelivery":1,"sostoyanie":"Б\/у","vid_tovara":"Монеты","type_of_trade":"Продаю своё"}];`
