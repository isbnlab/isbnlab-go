# \BooksAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBookByAsin**](BooksAPI.md#GetBookByAsin) | **Get** /api/book/asin/{asin} | Look up a book by Amazon ASIN.
[**GetBookById**](BooksAPI.md#GetBookById) | **Get** /api/book/{id} | Look up a book by ISBNLab book id.
[**GetBookByIsbn**](BooksAPI.md#GetBookByIsbn) | **Get** /api/book/isbn/{isbn} | Look up a book by ISBN-10 or ISBN-13.
[**GetBookCoverByAsin**](BooksAPI.md#GetBookCoverByAsin) | **Get** /api/book/cover/asin/{asin} | Get the cover image URLs for a book by ASIN.
[**GetBookCoverByIsbn**](BooksAPI.md#GetBookCoverByIsbn) | **Get** /api/book/cover/isbn/{isbn} | Get the cover image URLs for a book by ISBN.
[**GetEditionsByAsin**](BooksAPI.md#GetEditionsByAsin) | **Get** /api/editions/asin/{asin} | Get every other edition of a book by ASIN.
[**GetEditionsByIsbn**](BooksAPI.md#GetEditionsByIsbn) | **Get** /api/editions/isbn/{isbn} | Get every other edition of a book by ISBN.
[**SearchBooks**](BooksAPI.md#SearchBooks) | **Get** /api/search/book/{query} | Search books by title, subtitle, author or ISBN.



## GetBookByAsin

> Book GetBookByAsin(ctx, asin).Execute()

Look up a book by Amazon ASIN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	asin := "asin_example" // string | ASIN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookByAsin(context.Background(), asin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookByAsin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookByAsin`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookByAsin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | ASIN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookByAsinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Book**](Book.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookById

> Book GetBookById(ctx, id).Execute()

Look up a book by ISBNLab book id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	id := "id_example" // string | Book identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookById`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Book identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Book**](Book.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookByIsbn

> Book GetBookByIsbn(ctx, isbn).Execute()

Look up a book by ISBN-10 or ISBN-13.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	isbn := "isbn_example" // string | ISBN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookByIsbn(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookByIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookByIsbn`: Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookByIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookByIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Book**](Book.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookCoverByAsin

> BookCover GetBookCoverByAsin(ctx, asin).Execute()

Get the cover image URLs for a book by ASIN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	asin := "asin_example" // string | ASIN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookCoverByAsin(context.Background(), asin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookCoverByAsin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookCoverByAsin`: BookCover
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookCoverByAsin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | ASIN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookCoverByAsinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookCover**](BookCover.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookCoverByIsbn

> BookCover GetBookCoverByIsbn(ctx, isbn).Execute()

Get the cover image URLs for a book by ISBN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	isbn := "isbn_example" // string | ISBN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetBookCoverByIsbn(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetBookCoverByIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookCoverByIsbn`: BookCover
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetBookCoverByIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookCoverByIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookCover**](BookCover.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditionsByAsin

> BookEditions GetEditionsByAsin(ctx, asin).Execute()

Get every other edition of a book by ASIN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	asin := "asin_example" // string | ASIN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetEditionsByAsin(context.Background(), asin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetEditionsByAsin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditionsByAsin`: BookEditions
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetEditionsByAsin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | ASIN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditionsByAsinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookEditions**](BookEditions.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditionsByIsbn

> BookEditions GetEditionsByIsbn(ctx, isbn).Execute()

Get every other edition of a book by ISBN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	isbn := "isbn_example" // string | ISBN identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.GetEditionsByIsbn(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.GetEditionsByIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditionsByIsbn`: BookEditions
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.GetEditionsByIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditionsByIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookEditions**](BookEditions.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBooks

> []Book SearchBooks(ctx, query).Limit(limit).Page(page).MaxResults(maxResults).Execute()

Search books by title, subtitle, author or ISBN.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/isbnlab/isbnlab-go"
)

func main() {
	query := "query_example" // string | Search text
	limit := int64(56) // int64 | Books per page, 1-100. Default 10. (optional) (default to 10)
	page := int64(56) // int64 | Page number, starting at 1. Default 1. (optional) (default to 1)
	maxResults := int64(56) // int64 | Maximum number of results across all pages. Default 100. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BooksAPI.SearchBooks(context.Background(), query).Limit(limit).Page(page).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BooksAPI.SearchBooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBooks`: []Book
	fmt.Fprintf(os.Stdout, "Response from `BooksAPI.SearchBooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string** | Search text | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchBooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Books per page, 1-100. Default 10. | [default to 10]
 **page** | **int64** | Page number, starting at 1. Default 1. | [default to 1]
 **maxResults** | **int64** | Maximum number of results across all pages. Default 100. | [default to 100]

### Return type

[**[]Book**](Book.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

