package subscribe_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vpoletaev11/avitoParser/scrapper"
	"github.com/vpoletaev11/avitoParser/subscribe"
)

func TestHandlerSuccess(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", bodyString)
}

func TestHandlerBadLink(t *testing.T) {
	dep, _, _ := newTestDepAndServer()
	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "wrong link")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

func TestHandlerLinkExists(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", bodyString)
}

func TestHandlerInsertLinkErrorDB(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

func TestHandlerEmailExists(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs(ts.URL, "example.mail.com").WillReturnResult(sqlmock.NewResult(1, 1))
	row := []string{"count"}
	sqlMock.ExpectQuery("SELECT COUNT").WithArgs(ts.URL).WillReturnRows(sqlmock.NewRows(row).AddRow(1))
	sqlMock.ExpectExec("DELETE FROM link").WithArgs(ts.URL).WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", bodyString)
}

func TestHandlerEmailExistsUpdateErrorDB(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs(ts.URL, "example.mail.com").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

func TestHandlerEmailExistsCountErrorDB(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs(ts.URL, "example.mail.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectQuery("SELECT COUNT").WithArgs(ts.URL).WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

func TestHandlerEmailExistsDeleteLinkErrorDB(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs(ts.URL, "example.mail.com").WillReturnResult(sqlmock.NewResult(1, 1))
	row := []string{"count"}
	sqlMock.ExpectQuery("SELECT COUNT").WithArgs(ts.URL).WillReturnRows(sqlmock.NewRows(row).AddRow(1))
	sqlMock.ExpectExec("DELETE FROM link").WithArgs(ts.URL).WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

func TestHandlerInsertEmailErrorDB(t *testing.T) {
	dep, sqlMock, ts := newTestDepAndServer()
	sqlMock.ExpectExec("INSERT INTO link").WithArgs(ts.URL, 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", ts.URL).WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", ts.URL)
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(dep)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "INTERNAL ERROR. Please try later\n", bodyString)
}

// NewTestDepAndServer returns: dependencies, Sqlmock interface to add mocks and httptest server, that emulating avito ad page
func newTestDepAndServer() (scrapper.Dep, sqlmock.Sqlmock, *httptest.Server) {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, avitoAdHTML)
	}))

	dep := scrapper.Dep{
		DB:     db,
		Client: ts.Client(),
	}

	return dep, sqlMock, ts
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
