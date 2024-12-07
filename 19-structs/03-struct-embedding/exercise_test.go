package structembedding

import (
	"reflect"
	"testing"
)

func TestParseBook(t *testing.T) {
	jsonData := []byte(`
	{
		"title": "The Go Programming Language",
		"author": {
			"name": "Alan A. A. Donovan",
                        "email": "alan@example.com"
		},
		"pages": 380,
		"ISBN": "978-0134190440"
	}`)

	expected := Book{
		Title: "The Go Programming Language",
		Author: Author{
			Name:  "Alan A. A. Donovan",
                        Email: "alan@example.com",
		},
		Pages: 380,
		ISBN:  "978-0134190440",
	}

	result, err := ParseBook(jsonData)
	if err != nil {
		t.Fatalf("ParseBook returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseBook result doesn't match expected.\nGot: %+v\nWant: %+v", result, expected)
	}
}

func TestParseBookInvalidInput(t *testing.T) {
	invalidJSON := []byte(`{"title": "Incomplete Book"`)

	_, err := ParseBook(invalidJSON)
	if err == nil {
		t.Error("ParseBook should return an error for invalid JSON input")
	}
}

func TestParseArticle(t *testing.T) {
	jsonData := []byte(`
        {
        	"title": "Smashing The Kernel Stack For Fun And Profit",
        	"author": {
        		"name": "Sinan Eren",
        		"email": "noir@olympos.org"
        	},
        	"journal": "Phrack Magazine",
        	"year": 2002
        }`)

	expected := Article{
		Title: "Smashing The Kernel Stack For Fun And Profit",
		Author: Author{
			Name:  "Sinan Eren",
			Email: "noir@olympos.org",
		},
		Journal: "Phrack Magazine",
		Year:  2002,
	}

	result, err := ParseArticle(jsonData)
	if err != nil {
		t.Fatalf("ParseArticle returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseArticle result doesn't match expected.\nGot: %+v\nWant: %+v", result, expected)
	}
}

func TestParseArticleInvalidInput(t *testing.T) {
	invalidJSON := []byte(`{"title": "Incomplete Article"`)

	_, err := ParseArticle(invalidJSON)
	if err == nil {
		t.Error("ParseArticle should return an error for invalid JSON input")
	}
}
