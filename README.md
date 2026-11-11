# ISBNLab Go client

The official Go client for the [ISBNLab book metadata API](https://www.isbnlab.com/api-information). Look up a book by ISBN and get its title, authors, publishers, series, BISAC genres, formats, page count, description and every other edition, from a catalog of more than 100 million editions. It also validates and converts ISBN-10, ISBN-13, EAN and Amazon ASIN.

## Install

```bash
go get github.com/isbnlab/isbnlab-go
```

Go 1.21 or later. No dependencies outside the standard library.

## Get an API token

[Create a free ISBNLab account](https://www.isbnlab.com/register) and copy your API token from your profile. The Free plan includes 1,000 calls a day; see [pricing](https://www.isbnlab.com/pricing) for more.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	isbnlab "github.com/isbnlab/isbnlab-go"
)

func main() {
	client := isbnlab.NewAPIClient(isbnlab.NewConfiguration())
	ctx := context.WithValue(context.Background(), isbnlab.ContextAccessToken, os.Getenv("ISBNLAB_API_TOKEN"))

	book, _, err := client.BooksAPI.GetBookByIsbn(ctx, "9780765302304").Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book.GetTitle())                // The Gathering Storm
	fmt.Println(book.GetAuthors()[0].GetName()) // Robert Jordan
	fmt.Println(book.GetPublishDate().Year())

	for _, edition := range book.GetOtherEditions() {
		fmt.Println(edition.GetIsbn13(), edition.GetTitle())
	}
}
```

The token travels in the request's `context.Context` (`isbnlab.ContextAccessToken`) and is sent as `Authorization: Bearer <token>`. Create the context once and reuse it; add a deadline with `context.WithTimeout` as usual.

### Empty fields

Many fields are empty for some books, so every field is nullable:

- `book.GetTitle()` returns the value, or the zero value (`""`, `0`, a zero `time.Time`) when the field is null.
- `book.GetTitleOk()` returns a pointer, which is `nil` when the field is null.

```go
if ratings, _ := book.GetRatingsOk(); ratings != nil {
	if summary, _ := ratings.GetSummaryOk(); summary != nil {
		fmt.Println(summary.GetAverage(), summary.GetCount())
	}
}
```

### Linked records

Linked records carry ids you can look up directly. Author, publisher, subject and genre lookups need a Standard or Pro plan:

```go
author, _, err := client.AuthorsAPI.GetAuthor(ctx, "123").Execute()
if err != nil {
	log.Fatal(err)
}
fmt.Println(author.GetBio())
```

## Errors

Any response other than 2xx returns an error of type `*isbnlab.GenericOpenAPIError`. The HTTP status is on the `*http.Response` returned next to it, and `isbnlab.ErrorDetail(err)` gives the API's explanation:

```go
_, response, err := client.BooksAPI.GetBookByIsbn(ctx, "9780000000000").Execute()
if err != nil {
	switch {
	case response != nil && response.StatusCode == 404:
		fmt.Println("Not in the catalog")
	case response != nil && response.StatusCode == 429:
		fmt.Println("Daily limit reached:", isbnlab.ErrorDetail(err))
	case response != nil && response.StatusCode == 403:
		fmt.Println("Your plan does not include this endpoint:", isbnlab.ErrorDetail(err))
	default:
		log.Fatal(err)
	}
}
```

`response` is `nil` when the request never got an answer (for example a network error).

## What you can call

| Service | Method | Plan |
|---|---|---|
| `BooksAPI` | `GetBookByIsbn`, `GetBookByAsin`, `GetBookById` | Free |
| `IsbnAPI` | `ParseIsbn`, `Isbn10ToIsbn13`, `Isbn13ToIsbn10`, `IsbnToEan`, `EanToIsbn`, `IsbnToAsin`, `AsinToIsbn` | Free |
| `BooksAPI` | `GetBookCoverByIsbn`, `GetBookCoverByAsin` | Standard |
| `AuthorsAPI` | `GetAuthor`, `FindAuthor` | Standard |
| `PublishersAPI` | `GetPublisher`, `FindPublisher`, `GetPublishersByPrefix` | Standard |
| `SubjectsAPI` | `GetSubject`, `FindSubject` | Standard |
| `GenresAPI` | `GetGenre`, `FindGenre` (name or BISAC code) | Standard |
| `BooksAPI` | `SearchBooks(ctx, query).Limit(10).Page(1).MaxResults(100)` | Pro |
| `BooksAPI` | `GetEditionsByIsbn`, `GetEditionsByAsin` | Pro |

Each method returns a request builder; call `.Execute()` to send it. It returns the result, the `*http.Response` (with the rate-limit headers, such as `X-RateLimit-Remaining`) and an error.

Reference for every method and model: [docs/](docs/). Field-by-field documentation: [API documentation](https://www.isbnlab.com/api-information).

## Development

The `api_*.go`, `model_*.go`, `client.go`, `configuration.go`, `response.go` and `utils.go` files and `docs/` are generated from the API's [OpenAPI document](https://www.isbnlab.com/api/docs.jsonopenapi) with [OpenAPI Generator](https://openapi-generator.tech/), so changes to them are made in the generator and arrive with the next release; please open an issue rather than a pull request for those files. `README.md`, `errors.go`, `go.mod` and the tests are maintained by hand; `example_test.go` holds this README's examples, so `go test` checks they compile.

```bash
go vet ./...
go test ./...
```

## License

MIT. Maintained by [ISBNLab](https://www.isbnlab.com/), the book metadata API.
