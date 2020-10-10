package scrapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/avitoParser/scrapper"
)

func TestScrapperSuccess(t *testing.T) {
	price, err := scrapper.GetPrice("https://www.avito.ru/moskva/kollektsionirovanie/moneta_iz_makdonaldsa_2009709110")
	assert.Nil(t, err)
	assert.Equal(t, 1000000, price)
}
