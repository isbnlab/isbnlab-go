# \PublishersAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindPublisher**](PublishersAPI.md#FindPublisher) | **Get** /api/publisher/search/{search} | Find a publisher by name.
[**GetPublisher**](PublishersAPI.md#GetPublisher) | **Get** /api/publisher/{id} | Look up a publisher by id.
[**GetPublishersByPrefix**](PublishersAPI.md#GetPublishersByPrefix) | **Get** /api/publisher/prefix/{prefix} | List the publishers that use an ISBN registrant prefix.



## FindPublisher

> Publisher FindPublisher(ctx, search).Execute()

Find a publisher by name.

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
	resp, r, err := apiClient.PublishersAPI.FindPublisher(context.Background(), search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.FindPublisher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindPublisher`: Publisher
	fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.FindPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**search** | **string** | Name to search for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Publisher**](Publisher.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublisher

> Publisher GetPublisher(ctx, id).Execute()

Look up a publisher by id.

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
	id := "id_example" // string | Publisher identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublishersAPI.GetPublisher(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.GetPublisher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublisher`: Publisher
	fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.GetPublisher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Publisher identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublisherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Publisher**](Publisher.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublishersByPrefix

> []Publisher GetPublishersByPrefix(ctx, prefix).Page(page).Execute()

List the publishers that use an ISBN registrant prefix.

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
	prefix := "prefix_example" // string | ISBN registrant prefix, for example 978-0-7653.
	page := int64(56) // int64 | The collection page number (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublishersAPI.GetPublishersByPrefix(context.Background(), prefix).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublishersAPI.GetPublishersByPrefix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublishersByPrefix`: []Publisher
	fmt.Fprintf(os.Stdout, "Response from `PublishersAPI.GetPublishersByPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**prefix** | **string** | ISBN registrant prefix, for example 978-0-7653. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublishersByPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | The collection page number | [default to 1]

### Return type

[**[]Publisher**](Publisher.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

