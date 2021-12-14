# \SecuritySecureScoreApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecurityCreateSecureScores**](SecuritySecureScoreApi.md#SecurityCreateSecureScores) | **Post** /security/secureScores | Create new navigation property to secureScores for security
[**SecurityDeleteSecureScores**](SecuritySecureScoreApi.md#SecurityDeleteSecureScores) | **Delete** /security/secureScores/{secureScore-id} | Delete navigation property secureScores for security
[**SecurityGetSecureScores**](SecuritySecureScoreApi.md#SecurityGetSecureScores) | **Get** /security/secureScores/{secureScore-id} | Get secureScores from security
[**SecurityListSecureScores**](SecuritySecureScoreApi.md#SecurityListSecureScores) | **Get** /security/secureScores | Get secureScores from security
[**SecurityUpdateSecureScores**](SecuritySecureScoreApi.md#SecurityUpdateSecureScores) | **Patch** /security/secureScores/{secureScore-id} | Update the navigation property secureScores in security



## SecurityCreateSecureScores

> MicrosoftGraphSecureScore SecurityCreateSecureScores(ctx).MicrosoftGraphSecureScore(microsoftGraphSecureScore).Execute()

Create new navigation property to secureScores for security

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    microsoftGraphSecureScore := *openapiclient.NewMicrosoftGraphSecureScore() // MicrosoftGraphSecureScore | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreApi.SecurityCreateSecureScores(context.Background()).MicrosoftGraphSecureScore(microsoftGraphSecureScore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreApi.SecurityCreateSecureScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityCreateSecureScores`: MicrosoftGraphSecureScore
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreApi.SecurityCreateSecureScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityCreateSecureScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSecureScore** | [**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md) | New navigation property | 

### Return type

[**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityDeleteSecureScores

> SecurityDeleteSecureScores(ctx, secureScoreId).IfMatch(ifMatch).Execute()

Delete navigation property secureScores for security

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secureScoreId := "secureScoreId_example" // string | key: id of secureScore
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreApi.SecurityDeleteSecureScores(context.Background(), secureScoreId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreApi.SecurityDeleteSecureScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreId** | **string** | key: id of secureScore | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityDeleteSecureScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | ETag | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityGetSecureScores

> MicrosoftGraphSecureScore SecurityGetSecureScores(ctx, secureScoreId).Select_(select_).Expand(expand).Execute()

Get secureScores from security

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secureScoreId := "secureScoreId_example" // string | key: id of secureScore
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreApi.SecurityGetSecureScores(context.Background(), secureScoreId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreApi.SecurityGetSecureScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityGetSecureScores`: MicrosoftGraphSecureScore
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreApi.SecurityGetSecureScores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreId** | **string** | key: id of secureScore | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityGetSecureScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityListSecureScores

> CollectionOfSecureScore SecurityListSecureScores(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get secureScores from security

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    top := int32(50) // int32 | Show only the first n items (optional)
    skip := int32(56) // int32 | Skip the first n items (optional)
    search := "search_example" // string | Search items by search phrases (optional)
    filter := "filter_example" // string | Filter items by property values (optional)
    count := true // bool | Include count of items (optional)
    orderby := []string{"Orderby_example"} // []string | Order items by property values (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreApi.SecurityListSecureScores(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreApi.SecurityListSecureScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecurityListSecureScores`: CollectionOfSecureScore
    fmt.Fprintf(os.Stdout, "Response from `SecuritySecureScoreApi.SecurityListSecureScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecurityListSecureScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfSecureScore**](CollectionOfSecureScore.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecurityUpdateSecureScores

> SecurityUpdateSecureScores(ctx, secureScoreId).MicrosoftGraphSecureScore(microsoftGraphSecureScore).Execute()

Update the navigation property secureScores in security

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secureScoreId := "secureScoreId_example" // string | key: id of secureScore
    microsoftGraphSecureScore := *openapiclient.NewMicrosoftGraphSecureScore() // MicrosoftGraphSecureScore | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SecuritySecureScoreApi.SecurityUpdateSecureScores(context.Background(), secureScoreId).MicrosoftGraphSecureScore(microsoftGraphSecureScore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecuritySecureScoreApi.SecurityUpdateSecureScores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secureScoreId** | **string** | key: id of secureScore | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecurityUpdateSecureScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSecureScore** | [**MicrosoftGraphSecureScore**](MicrosoftGraphSecureScore.md) | New navigation property values | 

### Return type

 (empty response body)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

