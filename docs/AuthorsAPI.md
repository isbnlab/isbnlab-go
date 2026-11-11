# \AuthorsAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAuthor**](AuthorsAPI.md#FindAuthor) | **Get** /api/author/search/{search} | Find an author by name.
[**GetAuthor**](AuthorsAPI.md#GetAuthor) | **Get** /api/author/{id} | Look up an author by id.



## FindAuthor

> Author FindAuthor(ctx, search).Execute()

Find an author by name.

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
	resp, r, err := apiClient.AuthorsAPI.FindAuthor(context.Background(), search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.FindAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAuthor`: Author
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.FindAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**search** | **string** | Name to search for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Author**](Author.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthor

> Author GetAuthor(ctx, id).Execute()

Look up an author by id.

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
	id := "id_example" // string | Author identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.GetAuthor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.GetAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthor`: Author
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.GetAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Author identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Author**](Author.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

