package structembedding

import "encoding/json"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// Author represents information about the book's author
type Author struct {
	// TODO: Define the Author struct fields
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Book represents information about a book
type Book struct {
	// TODO: Define the Book struct fields, embedding the Author struct
	Title  string `json:"title"`
	Author `json:"author"`
	Pages  int    `json:"pages"`
	ISBN   string `json:"ISBN"`
}

// Article represents information about a article
type Article struct {
	// TODO: Define the Article struct fields, embedding the Author struct
	Title   string `json:"title"`
	Author  `json:"author"`
	Journal string `json:"journal"`
	Year    int    `json:"year"`
}

// ParseBook parses the given JSON data into a Book struct
func ParseBook(jsonData []byte) (Book, error) {
	var book Book
	err := json.Unmarshal(jsonData, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

// ParseArticle parses the given JSON data into a Article struct
func ParseArticle(jsonData []byte) (Article, error) {
	var article Article
	err := json.Unmarshal(jsonData, &article)
	if err != nil {
		return article, err
	}
	return article, nil
}
