# Search

Write a function that reads the below text and returns `true` or `false` if "alive" appears in the text *more than once*, without respect to letter case (i.e., the each of words `Hello`, `heLLo`, and `HELlo` match for the search strings `HELLO` and `hello`).

> "The upper realm has plenary powers. He lived a long time ago, but he is still alive. Apollonius of Tyana, writing as Hermes Trismegistos, said, That which is above is that which is below. By this he meant to tell us that our universe is a hologram, but he lacked the term."

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

HINT:

- Make sure to remove punctuation from the input (like `.`, `;`, or `?`) and trim the text to alphanumeric characters. For instance, you can use [`regexp.ReplaceAllString`](https://pkg.go.dev/regexp#Regexp.ReplaceAllString) with the regular expression `[^a-zA-Z0-9 ]+` for this purpose.
- You can use a `map[string]int` to remember the number of times a particular word is found while iterating through the text word by word.
