package scrapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/avitoParser/test.go"
)

func TestScrapperSuccess(t *testing.T) {
	dep, _, ts := test.NewDepAndServer()
	price, err := dep.GetPrice(ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, 1000000, price)
}
