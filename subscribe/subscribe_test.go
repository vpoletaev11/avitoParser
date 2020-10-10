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
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vpoletaev11/avitoParser/subscribe"
)

func TestHandlerSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerBadLink(t *testing.T) {
	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "wrong link")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(nil)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "Bad link", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerLinkExists(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerInsertLinkErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "Some error with DB connection", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerEmailExists(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110").WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", "example.mail.com").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerEmailExistsErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110").WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", "example.mail.com").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "Some error with DB connection", bodyString)

	time.Sleep(1 * time.Second)
}

func TestHandlerInsertEmailErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110", 1000000).WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	r, err := http.NewRequest("POST", "http://localhost/", strings.NewReader(data.Encode()))
	require.NoError(t, err)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	w := httptest.NewRecorder()

	sut := subscribe.Handler(db)
	sut(w, r)

	bodyBytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	assert.Equal(t, "Some error with DB connection", bodyString)

	time.Sleep(1 * time.Second)
}
