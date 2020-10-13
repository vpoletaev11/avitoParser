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
	"github.com/vpoletaev11/avitoParser/test.go"
)

func TestHandlerSuccess(t *testing.T) {
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, _, _ := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
	dep, sqlMock, ts := test.NewDepAndServer()
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
