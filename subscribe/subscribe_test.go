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
	"github.com/vpoletaev11/avitoParser/subscribe"
)

func TestHandlerSuccess(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}

func TestHandlerLinkExists(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}

func TestHandlerInsertLinkErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}

func TestHandlerEmailExists(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.example.com").WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs("https://www.example.com", "example.mail.com").WillReturnResult(sqlmock.NewResult(1, 1))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}

func TestHandlerEmailExistsErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.example.com").WillReturnError(fmt.Errorf("Error 1062 ....."))
	sqlMock.ExpectExec("UPDATE email").WithArgs("https://www.example.com", "example.mail.com").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}

func TestHandlerInsertEmailErrorDB(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	require.NoError(t, err)
	sqlMock.ExpectExec("INSERT INTO link").WithArgs("https://www.example.com").WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec("INSERT INTO email").WithArgs("example.mail.com", "https://www.example.com").WillReturnError(fmt.Errorf("Some error with DB connection"))

	data := url.Values{}
	data.Set("email", "example.mail.com")
	data.Add("link", "https://www.example.com")
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
}
