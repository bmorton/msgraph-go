# \InformationProtectionThreatAssessmentRequestApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InformationProtectionCreateThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionCreateThreatAssessmentRequests) | **Post** /informationProtection/threatAssessmentRequests | Create new navigation property to threatAssessmentRequests for informationProtection
[**InformationProtectionDeleteThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionDeleteThreatAssessmentRequests) | **Delete** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id} | Delete navigation property threatAssessmentRequests for informationProtection
[**InformationProtectionGetThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionGetThreatAssessmentRequests) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id} | Get threatAssessmentRequests from informationProtection
[**InformationProtectionListThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionListThreatAssessmentRequests) | **Get** /informationProtection/threatAssessmentRequests | Get threatAssessmentRequests from informationProtection
[**InformationProtectionThreatAssessmentRequestsCreateResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsCreateResults) | **Post** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results | Create new navigation property to results for informationProtection
[**InformationProtectionThreatAssessmentRequestsDeleteResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsDeleteResults) | **Delete** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results/{threatAssessmentResult-id} | Delete navigation property results for informationProtection
[**InformationProtectionThreatAssessmentRequestsGetResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsGetResults) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results/{threatAssessmentResult-id} | Get results from informationProtection
[**InformationProtectionThreatAssessmentRequestsListResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsListResults) | **Get** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results | Get results from informationProtection
[**InformationProtectionThreatAssessmentRequestsUpdateResults**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionThreatAssessmentRequestsUpdateResults) | **Patch** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id}/results/{threatAssessmentResult-id} | Update the navigation property results in informationProtection
[**InformationProtectionUpdateThreatAssessmentRequests**](InformationProtectionThreatAssessmentRequestApi.md#InformationProtectionUpdateThreatAssessmentRequests) | **Patch** /informationProtection/threatAssessmentRequests/{threatAssessmentRequest-id} | Update the navigation property threatAssessmentRequests in informationProtection



## InformationProtectionCreateThreatAssessmentRequests

> MicrosoftGraphThreatAssessmentRequest InformationProtectionCreateThreatAssessmentRequests(ctx).MicrosoftGraphThreatAssessmentRequest(microsoftGraphThreatAssessmentRequest).Execute()

Create new navigation property to threatAssessmentRequests for informationProtection

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
    microsoftGraphThreatAssessmentRequest := *openapiclient.NewMicrosoftGraphThreatAssessmentRequest() // MicrosoftGraphThreatAssessmentRequest | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionCreateThreatAssessmentRequests(context.Background()).MicrosoftGraphThreatAssessmentRequest(microsoftGraphThreatAssessmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionCreateThreatAssessmentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionCreateThreatAssessmentRequests`: MicrosoftGraphThreatAssessmentRequest
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionCreateThreatAssessmentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionCreateThreatAssessmentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphThreatAssessmentRequest** | [**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md) | New navigation property | 

### Return type

[**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionDeleteThreatAssessmentRequests

> InformationProtectionDeleteThreatAssessmentRequests(ctx, threatAssessmentRequestId).IfMatch(ifMatch).Execute()

Delete navigation property threatAssessmentRequests for informationProtection

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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionDeleteThreatAssessmentRequests(context.Background(), threatAssessmentRequestId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionDeleteThreatAssessmentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionDeleteThreatAssessmentRequestsRequest struct via the builder pattern


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


## InformationProtectionGetThreatAssessmentRequests

> MicrosoftGraphThreatAssessmentRequest InformationProtectionGetThreatAssessmentRequests(ctx, threatAssessmentRequestId).Select_(select_).Expand(expand).Execute()

Get threatAssessmentRequests from informationProtection

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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionGetThreatAssessmentRequests(context.Background(), threatAssessmentRequestId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionGetThreatAssessmentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionGetThreatAssessmentRequests`: MicrosoftGraphThreatAssessmentRequest
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionGetThreatAssessmentRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionGetThreatAssessmentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionListThreatAssessmentRequests

> CollectionOfThreatAssessmentRequest InformationProtectionListThreatAssessmentRequests(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get threatAssessmentRequests from informationProtection

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
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionListThreatAssessmentRequests(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionListThreatAssessmentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionListThreatAssessmentRequests`: CollectionOfThreatAssessmentRequest
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionListThreatAssessmentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionListThreatAssessmentRequestsRequest struct via the builder pattern


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

[**CollectionOfThreatAssessmentRequest**](CollectionOfThreatAssessmentRequest.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsCreateResults

> MicrosoftGraphThreatAssessmentResult InformationProtectionThreatAssessmentRequestsCreateResults(ctx, threatAssessmentRequestId).MicrosoftGraphThreatAssessmentResult(microsoftGraphThreatAssessmentResult).Execute()

Create new navigation property to results for informationProtection



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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    microsoftGraphThreatAssessmentResult := *openapiclient.NewMicrosoftGraphThreatAssessmentResult() // MicrosoftGraphThreatAssessmentResult | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsCreateResults(context.Background(), threatAssessmentRequestId).MicrosoftGraphThreatAssessmentResult(microsoftGraphThreatAssessmentResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsCreateResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionThreatAssessmentRequestsCreateResults`: MicrosoftGraphThreatAssessmentResult
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsCreateResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionThreatAssessmentRequestsCreateResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphThreatAssessmentResult** | [**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md) | New navigation property | 

### Return type

[**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsDeleteResults

> InformationProtectionThreatAssessmentRequestsDeleteResults(ctx, threatAssessmentRequestId, threatAssessmentResultId).IfMatch(ifMatch).Execute()

Delete navigation property results for informationProtection



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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    threatAssessmentResultId := "threatAssessmentResultId_example" // string | key: id of threatAssessmentResult
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsDeleteResults(context.Background(), threatAssessmentRequestId, threatAssessmentResultId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsDeleteResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 
**threatAssessmentResultId** | **string** | key: id of threatAssessmentResult | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionThreatAssessmentRequestsDeleteResultsRequest struct via the builder pattern


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


## InformationProtectionThreatAssessmentRequestsGetResults

> MicrosoftGraphThreatAssessmentResult InformationProtectionThreatAssessmentRequestsGetResults(ctx, threatAssessmentRequestId, threatAssessmentResultId).Select_(select_).Expand(expand).Execute()

Get results from informationProtection



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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    threatAssessmentResultId := "threatAssessmentResultId_example" // string | key: id of threatAssessmentResult
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsGetResults(context.Background(), threatAssessmentRequestId, threatAssessmentResultId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsGetResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionThreatAssessmentRequestsGetResults`: MicrosoftGraphThreatAssessmentResult
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsGetResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 
**threatAssessmentResultId** | **string** | key: id of threatAssessmentResult | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionThreatAssessmentRequestsGetResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsListResults

> CollectionOfThreatAssessmentResult InformationProtectionThreatAssessmentRequestsListResults(ctx, threatAssessmentRequestId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get results from informationProtection



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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
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
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsListResults(context.Background(), threatAssessmentRequestId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsListResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InformationProtectionThreatAssessmentRequestsListResults`: CollectionOfThreatAssessmentResult
    fmt.Fprintf(os.Stdout, "Response from `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsListResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionThreatAssessmentRequestsListResultsRequest struct via the builder pattern


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

[**CollectionOfThreatAssessmentResult**](CollectionOfThreatAssessmentResult.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InformationProtectionThreatAssessmentRequestsUpdateResults

> InformationProtectionThreatAssessmentRequestsUpdateResults(ctx, threatAssessmentRequestId, threatAssessmentResultId).MicrosoftGraphThreatAssessmentResult(microsoftGraphThreatAssessmentResult).Execute()

Update the navigation property results in informationProtection



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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    threatAssessmentResultId := "threatAssessmentResultId_example" // string | key: id of threatAssessmentResult
    microsoftGraphThreatAssessmentResult := *openapiclient.NewMicrosoftGraphThreatAssessmentResult() // MicrosoftGraphThreatAssessmentResult | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsUpdateResults(context.Background(), threatAssessmentRequestId, threatAssessmentResultId).MicrosoftGraphThreatAssessmentResult(microsoftGraphThreatAssessmentResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionThreatAssessmentRequestsUpdateResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 
**threatAssessmentResultId** | **string** | key: id of threatAssessmentResult | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionThreatAssessmentRequestsUpdateResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphThreatAssessmentResult** | [**MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md) | New navigation property values | 

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


## InformationProtectionUpdateThreatAssessmentRequests

> InformationProtectionUpdateThreatAssessmentRequests(ctx, threatAssessmentRequestId).MicrosoftGraphThreatAssessmentRequest(microsoftGraphThreatAssessmentRequest).Execute()

Update the navigation property threatAssessmentRequests in informationProtection

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
    threatAssessmentRequestId := "threatAssessmentRequestId_example" // string | key: id of threatAssessmentRequest
    microsoftGraphThreatAssessmentRequest := *openapiclient.NewMicrosoftGraphThreatAssessmentRequest() // MicrosoftGraphThreatAssessmentRequest | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InformationProtectionThreatAssessmentRequestApi.InformationProtectionUpdateThreatAssessmentRequests(context.Background(), threatAssessmentRequestId).MicrosoftGraphThreatAssessmentRequest(microsoftGraphThreatAssessmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationProtectionThreatAssessmentRequestApi.InformationProtectionUpdateThreatAssessmentRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**threatAssessmentRequestId** | **string** | key: id of threatAssessmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiInformationProtectionUpdateThreatAssessmentRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphThreatAssessmentRequest** | [**MicrosoftGraphThreatAssessmentRequest**](MicrosoftGraphThreatAssessmentRequest.md) | New navigation property values | 

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

