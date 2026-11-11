# \IsbnAPI

All URIs are relative to *https://www.isbnlab.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsinToIsbn**](IsbnAPI.md#AsinToIsbn) | **Get** /api/isbn/asin-to-isbn/{asin} | Convert an Amazon ASIN to ISBN.
[**EanToIsbn**](IsbnAPI.md#EanToIsbn) | **Get** /api/isbn/ean-to-isbn/{ean} | Convert an EAN to ISBN.
[**Isbn10ToIsbn13**](IsbnAPI.md#Isbn10ToIsbn13) | **Get** /api/isbn/isbn10-to-isbn13/{isbn10} | Convert an ISBN-10 to ISBN-13.
[**Isbn13ToIsbn10**](IsbnAPI.md#Isbn13ToIsbn10) | **Get** /api/isbn/isbn13-to-isbn10/{isbn13} | Convert an ISBN-13 to ISBN-10.
[**IsbnToAsin**](IsbnAPI.md#IsbnToAsin) | **Get** /api/isbn/isbn-to-asin/{isbn} | Convert an ISBN to Amazon ASIN.
[**IsbnToEan**](IsbnAPI.md#IsbnToEan) | **Get** /api/isbn/isbn-to-ean/{isbn} | Convert an ISBN to EAN.
[**ParseIsbn**](IsbnAPI.md#ParseIsbn) | **Get** /api/isbn/isbn-lookup/{isbn} | Validate an ISBN and split it into its parts.



## AsinToIsbn

> IsbnDetails AsinToIsbn(ctx, asin).Execute()

Convert an Amazon ASIN to ISBN.

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
	asin := "asin_example" // string | ASIN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.AsinToIsbn(context.Background(), asin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.AsinToIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AsinToIsbn`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.AsinToIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**asin** | **string** | ASIN | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsinToIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EanToIsbn

> IsbnDetails EanToIsbn(ctx, ean).Execute()

Convert an EAN to ISBN.

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
	ean := "ean_example" // string | EAN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.EanToIsbn(context.Background(), ean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.EanToIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EanToIsbn`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.EanToIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ean** | **string** | EAN | 

### Other Parameters

Other parameters are passed through a pointer to a apiEanToIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Isbn10ToIsbn13

> IsbnDetails Isbn10ToIsbn13(ctx, isbn10).Execute()

Convert an ISBN-10 to ISBN-13.

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
	isbn10 := "isbn10_example" // string | ISBN-10

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.Isbn10ToIsbn13(context.Background(), isbn10).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.Isbn10ToIsbn13``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Isbn10ToIsbn13`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.Isbn10ToIsbn13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn10** | **string** | ISBN-10 | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsbn10ToIsbn13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Isbn13ToIsbn10

> IsbnDetails Isbn13ToIsbn10(ctx, isbn13).Execute()

Convert an ISBN-13 to ISBN-10.

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
	isbn13 := "isbn13_example" // string | ISBN-13

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.Isbn13ToIsbn10(context.Background(), isbn13).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.Isbn13ToIsbn10``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Isbn13ToIsbn10`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.Isbn13ToIsbn10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn13** | **string** | ISBN-13 | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsbn13ToIsbn10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsbnToAsin

> IsbnDetails IsbnToAsin(ctx, isbn).Execute()

Convert an ISBN to Amazon ASIN.

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
	isbn := "isbn_example" // string | ISBN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.IsbnToAsin(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.IsbnToAsin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsbnToAsin`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.IsbnToAsin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsbnToAsinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsbnToEan

> IsbnDetails IsbnToEan(ctx, isbn).Execute()

Convert an ISBN to EAN.

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
	isbn := "isbn_example" // string | ISBN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.IsbnToEan(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.IsbnToEan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsbnToEan`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.IsbnToEan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsbnToEanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseIsbn

> IsbnDetails ParseIsbn(ctx, isbn).Execute()

Validate an ISBN and split it into its parts.

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
	isbn := "isbn_example" // string | ISBN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IsbnAPI.ParseIsbn(context.Background(), isbn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IsbnAPI.ParseIsbn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseIsbn`: IsbnDetails
	fmt.Fprintf(os.Stdout, "Response from `IsbnAPI.ParseIsbn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isbn** | **string** | ISBN | 

### Other Parameters

Other parameters are passed through a pointer to a apiParseIsbnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IsbnDetails**](IsbnDetails.md)

### Authorization

[UrlKey](../README.md#UrlKey), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

