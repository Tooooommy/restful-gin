package main

import (
	"fmt"
	"github.com/muesli/regommend"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InitTest(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(t, 123, 123, "they should be equal")
}

func main() {
	books := regommend.Table("books")
	booksChrisRead := make(map[interface{}]float64)
	booksChrisRead["1984"] = 5.0
	booksChrisRead["Robinson Crusoe"] = 4.0
	booksChrisRead["Moby-Dick"] = 3.0
	books.Add("Chris", booksChrisRead)

	booksJayRead := make(map[interface{}]float64)
	booksJayRead["1984"] = 5.0
	booksJayRead["Robinson Crusoe"] = 4.0
	booksJayRead["Gulliver's Travels"] = 4.5
	books.Add("Jay", booksJayRead)

	booksRayRead := make(map[interface{}]float64)
	booksRayRead["1984"] = 1.0
	booksRayRead["Crusoe"] = 4.0
	booksRayRead["Travels"] = 3.0
	books.Add("Ray", booksRayRead)

	recs, _ := books.Recommend("Chris")
	for _, rec := range recs {
		fmt.Println("Recommending", rec.Key, "with score", rec.Distance)
	}
}
