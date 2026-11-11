# \SubjectsAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindSubject**](SubjectsAPI.md#FindSubject) | **Get** /api/subject/search/{search} | Find a subject by name.
[**GetSubject**](SubjectsAPI.md#GetSubject) | **Get** /api/subject/{id} | Look up a subject by id.



## FindSubject

> Subject FindSubject(ctx, search).Execute()

Find a subject by name.

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
	resp, r, err := apiClient.SubjectsAPI.FindSubject(context.Background(), search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.FindSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindSubject`: Subject
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.FindSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**search** | **string** | Name to search for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subject**](Subject.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubject

> Subject GetSubject(ctx, id).Execute()

Look up a subject by id.

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
	id := "id_example" // string | Subject identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubjectsAPI.GetSubject(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubjectsAPI.GetSubject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubject`: Subject
	fmt.Fprintf(os.Stdout, "Response from `SubjectsAPI.GetSubject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subject identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subject**](Subject.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

