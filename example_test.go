package isbnlab_test

// The README's examples, compiled by go test (they are not run: they would
// call the real API).

import (
	"context"
	"fmt"
	"log"
	"os"

	isbnlab "github.com/isbnlab/isbnlab-go"
)

func Example() {
	client := isbnlab.NewAPIClient(isbnlab.NewConfiguration())
	ctx := context.WithValue(context.Background(), isbnlab.ContextAccessToken, os.Getenv("ISBNLAB_API_TOKEN"))

	book, _, err := client.BooksAPI.GetBookByIsbn(ctx, "9780765302304").Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(book.GetTitle())
	fmt.Println(book.GetAuthors()[0].GetName())
	fmt.Println(book.GetPublishDate().Year())

	for _, edition := range book.GetOtherEditions() {
		fmt.Println(edition.GetIsbn13(), edition.GetTitle())
	}

	if ratings, _ := book.GetRatingsOk(); ratings != nil {
		if summary, _ := ratings.GetSummaryOk(); summary != nil {
			fmt.Println(summary.GetAverage(), summary.GetCount())
		}
	}

	author, _, err := client.AuthorsAPI.GetAuthor(ctx, "123").Execute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(author.GetBio())
}

func Example_errors() {
	client := isbnlab.NewAPIClient(isbnlab.NewConfiguration())
	ctx := context.WithValue(context.Background(), isbnlab.ContextAccessToken, os.Getenv("ISBNLAB_API_TOKEN"))

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
}
