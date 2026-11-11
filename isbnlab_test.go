package isbnlab_test

// Each client method against a local test server: the request it sends and the
// typed result it returns. Fixtures in testdata/ are the example responses from
// the API documentation, in plain JSON. No test calls the real API.

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	isbnlab "github.com/isbnlab/isbnlab-go"
)

type recorded struct {
	path, query, auth, accept string
}

// newClient returns a client whose requests go to a test server that answers
// every request with the given status and testdata file (or raw body).
func newClient(t *testing.T, status int, fixture, rawBody string) (*isbnlab.APIClient, context.Context, *recorded) {
	t.Helper()
	body := []byte(rawBody)
	if fixture != "" {
		var err error
		if body, err = os.ReadFile("testdata/" + fixture); err != nil {
			t.Fatal(err)
		}
	}
	rec := &recorded{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec.path, rec.query = r.URL.EscapedPath(), r.URL.RawQuery
		rec.auth, rec.accept = r.Header.Get("Authorization"), r.Header.Get("Accept")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(body)
	}))
	t.Cleanup(server.Close)

	cfg := isbnlab.NewConfiguration()
	cfg.Servers = isbnlab.ServerConfigurations{{URL: server.URL}}
	ctx := context.WithValue(context.Background(), isbnlab.ContextAccessToken, "test-token")

	return isbnlab.NewAPIClient(cfg), ctx, rec
}

func expectRequest(t *testing.T, rec *recorded, path, query string) {
	t.Helper()
	if rec.path != path {
		t.Errorf("path = %q, want %q", rec.path, path)
	}
	if rec.query != query {
		t.Errorf("query = %q, want %q", rec.query, query)
	}
	if rec.auth != "Bearer test-token" {
		t.Errorf("Authorization = %q, want the bearer token", rec.auth)
	}
	if rec.accept != "application/json" {
		t.Errorf("Accept = %q, want application/json", rec.accept)
	}
}

func TestGetBookByIsbn(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "book.json", "")

	book, _, err := client.BooksAPI.GetBookByIsbn(ctx, "9780765302304").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/book/isbn/9780765302304", "")

	if got := book.GetTitle(); got != "The Gathering Storm" {
		t.Errorf("title = %q", got)
	}
	if got := book.GetId(); got != 15596406 {
		t.Errorf("id = %d", got)
	}
	if got := book.GetIsbn13(); got != "9780765302304" {
		t.Errorf("isbn13 = %q", got)
	}
	if len(book.Authors) == 0 || book.Authors[0].GetName() != "Robert Jordan" {
		t.Errorf("authors = %+v", book.Authors)
	}
	if date, ok := book.GetPublishDateOk(); !ok || date.Year() != 2009 {
		t.Errorf("publishDate = %v, %v", date, ok)
	}
	series, ok := book.GetSeriesOk()
	if !ok || series.GetNumber() != 12 || series.GetName() == "" {
		t.Errorf("series = %+v", series)
	}
	if len(book.GetOtherEditions()) == 0 || book.GetOtherEditions()[0].GetIsbn13() == "" {
		t.Error("no other editions")
	}
	if details, ok := book.GetIsbnModelOk(); !ok || details.GetIsbn10() != "0765302306" {
		t.Errorf("isbnModel = %+v", details)
	}
}

func TestABookWithMostFieldsEmptyStillDecodes(t *testing.T) {
	client, ctx, _ := newClient(t, 200, "book-sparse.json", "")

	book, _, err := client.BooksAPI.GetBookByIsbn(ctx, "9780000000002").Execute()
	if err != nil {
		t.Fatal(err)
	}
	if book.GetTitle() != "A Sparse Book" {
		t.Errorf("title = %q", book.GetTitle())
	}
	// A null field: GetXOk returns a nil pointer, GetX the zero value.
	if ratings, _ := book.GetRatingsOk(); ratings != nil {
		t.Error("ratings should be null")
	}
	if date, _ := book.GetPublishDateOk(); date != nil || !book.GetPublishDate().IsZero() {
		t.Error("publishDate should be null")
	}
	if book.GetContributors() != nil {
		t.Error("contributors should be null")
	}
	if len(book.GetAuthors()) != 0 {
		t.Error("authors should be empty")
	}
	series, ok := book.GetSeriesOk()
	if !ok || series == nil || series.GetName() != "Legacy Series Name" {
		t.Errorf("series = %+v", series)
	}
	if number, _ := series.GetNumberOk(); number != nil {
		t.Error("series number should be null")
	}
}

func TestSearchBooksSendsPagingParameters(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "book-search.json", "")

	books, _, err := client.BooksAPI.SearchBooks(ctx, "gathering storm").Limit(5).Page(2).Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/search/book/gathering%20storm", "limit=5&page=2")
	if len(books) == 0 || books[0].GetTitle() == "" {
		t.Errorf("books = %+v", books)
	}
}

func TestBookCoverAndEditions(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "book-cover.json", "")
	cover, _, err := client.BooksAPI.GetBookCoverByIsbn(ctx, "9780765302304").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/book/cover/isbn/9780765302304", "")
	if len(cover.GetCovers()) == 0 {
		t.Error("no covers")
	}

	client, ctx, rec = newClient(t, 200, "book-editions.json", "")
	editions, _, err := client.BooksAPI.GetEditionsByIsbn(ctx, "9780765302304").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/editions/isbn/9780765302304", "")
	if len(editions.GetOtherEditions()) == 0 {
		t.Error("no editions")
	}
}

func TestGetAuthorDecodesPhotos(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "author.json", "")

	author, _, err := client.AuthorsAPI.GetAuthor(ctx, "1").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/author/1", "")
	if author.GetName() != "Robert Jordan" {
		t.Errorf("name = %q", author.GetName())
	}
	if birth, ok := author.GetBirthDateOk(); !ok || birth.Year() != 1948 {
		t.Errorf("birthDate = %v", birth)
	}
	photos := author.GetPhotos()
	if len(photos) == 0 || photos[0].GetUrl() == "" {
		t.Errorf("photos = %+v", photos)
	}
}

func TestPublisherSubjectAndGenre(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "publishers-by-prefix.json", "")
	publishers, _, err := client.PublishersAPI.GetPublishersByPrefix(ctx, "978-0-7653").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/publisher/prefix/978-0-7653", "")
	if len(publishers) == 0 || publishers[0].GetName() == "" {
		t.Errorf("publishers = %+v", publishers)
	}

	client, ctx, _ = newClient(t, 200, "subject.json", "")
	if subject, _, err := client.SubjectsAPI.GetSubject(ctx, "1").Execute(); err != nil || subject.GetName() == "" {
		t.Errorf("subject = %+v, %v", subject, err)
	}

	client, ctx, rec = newClient(t, 200, "genre.json", "")
	if genre, _, err := client.GenresAPI.FindGenre(ctx, "FIC009000").Execute(); err != nil || genre.GetName() == "" {
		t.Errorf("genre = %+v, %v", genre, err)
	}
	expectRequest(t, rec, "/api/genre/search/FIC009000", "")
}

func TestParseIsbn(t *testing.T) {
	client, ctx, rec := newClient(t, 200, "isbn-details.json", "")

	details, _, err := client.IsbnAPI.ParseIsbn(ctx, "0765302306").Execute()
	if err != nil {
		t.Fatal(err)
	}
	expectRequest(t, rec, "/api/isbn/isbn-lookup/0765302306", "")
	if details.GetIsbn13() != "9780765302304" || details.GetRegistrationGroupName() != "English language" {
		t.Errorf("details = %+v", details)
	}
}

func TestErrorsCarryStatusAndDetail(t *testing.T) {
	for _, tc := range []struct {
		status int
		detail string
	}{
		{404, "No book found for ISBN 9780000000000."},
		{403, "This endpoint needs a Standard or Pro plan."},
		{429, "Daily limit reached."},
	} {
		client, ctx, _ := newClient(t, tc.status, "", `{"title":"Error","detail":"`+tc.detail+`","status":0}`)

		_, response, err := client.BooksAPI.GetBookByIsbn(ctx, "9780000000000").Execute()

		var apiErr *isbnlab.GenericOpenAPIError
		if !errors.As(err, &apiErr) {
			t.Fatalf("%d: err = %v, want *GenericOpenAPIError", tc.status, err)
		}
		if response == nil || response.StatusCode != tc.status {
			t.Errorf("%d: response = %v", tc.status, response)
		}
		if got := isbnlab.ErrorDetail(err); got != tc.detail {
			t.Errorf("%d: detail = %q, want %q", tc.status, got, tc.detail)
		}
	}
}
