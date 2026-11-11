# \GenresAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindGenre**](GenresAPI.md#FindGenre) | **Get** /api/genre/search/{search} | Find a BISAC genre by name or BISAC code.
[**GetGenre**](GenresAPI.md#GetGenre) | **Get** /api/genre/{id} | Look up a BISAC genre by id.



## FindGenre

> Genre FindGenre(ctx, search).Execute()

Find a BISAC genre by name or BISAC code.

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
	search := "search_example" // string | Name to search for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresAPI.FindGenre(context.Background(), search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.FindGenre``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindGenre`: Genre
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.FindGenre`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**search** | **string** | Name to search for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindGenreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Genre**](Genre.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenre

> Genre GetGenre(ctx, id).Execute()

Look up a BISAC genre by id.

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
	id := "id_example" // string | Genre identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenresAPI.GetGenre(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.GetGenre``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenre`: Genre
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.GetGenre`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Genre identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Genre**](Genre.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

