# \AuditLogsSignInApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateSignIns**](AuditLogsSignInApi.md#AuditLogsCreateSignIns) | **Post** /auditLogs/signIns | Create new navigation property to signIns for auditLogs
[**AuditLogsDeleteSignIns**](AuditLogsSignInApi.md#AuditLogsDeleteSignIns) | **Delete** /auditLogs/signIns/{signIn-id} | Delete navigation property signIns for auditLogs
[**AuditLogsGetSignIns**](AuditLogsSignInApi.md#AuditLogsGetSignIns) | **Get** /auditLogs/signIns/{signIn-id} | Get signIns from auditLogs
[**AuditLogsListSignIns**](AuditLogsSignInApi.md#AuditLogsListSignIns) | **Get** /auditLogs/signIns | Get signIns from auditLogs
[**AuditLogsUpdateSignIns**](AuditLogsSignInApi.md#AuditLogsUpdateSignIns) | **Patch** /auditLogs/signIns/{signIn-id} | Update the navigation property signIns in auditLogs



## AuditLogsCreateSignIns

> MicrosoftGraphSignIn AuditLogsCreateSignIns(ctx).MicrosoftGraphSignIn(microsoftGraphSignIn).Execute()

Create new navigation property to signIns for auditLogs



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
    microsoftGraphSignIn := *openapiclient.NewMicrosoftGraphSignIn() // MicrosoftGraphSignIn | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsSignInApi.AuditLogsCreateSignIns(context.Background()).MicrosoftGraphSignIn(microsoftGraphSignIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsSignInApi.AuditLogsCreateSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsCreateSignIns`: MicrosoftGraphSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsSignInApi.AuditLogsCreateSignIns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsCreateSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphSignIn** | [**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md) | New navigation property | 

### Return type

[**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsDeleteSignIns

> AuditLogsDeleteSignIns(ctx, signInId).IfMatch(ifMatch).Execute()

Delete navigation property signIns for auditLogs



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
    signInId := "signInId_example" // string | key: id of signIn
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsSignInApi.AuditLogsDeleteSignIns(context.Background(), signInId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsSignInApi.AuditLogsDeleteSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signInId** | **string** | key: id of signIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsDeleteSignInsRequest struct via the builder pattern


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


## AuditLogsGetSignIns

> MicrosoftGraphSignIn AuditLogsGetSignIns(ctx, signInId).Select_(select_).Expand(expand).Execute()

Get signIns from auditLogs



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
    signInId := "signInId_example" // string | key: id of signIn
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsSignInApi.AuditLogsGetSignIns(context.Background(), signInId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsSignInApi.AuditLogsGetSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsGetSignIns`: MicrosoftGraphSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsSignInApi.AuditLogsGetSignIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signInId** | **string** | key: id of signIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsGetSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListSignIns

> CollectionOfSignIn AuditLogsListSignIns(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get signIns from auditLogs



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
    resp, r, err := api_client.AuditLogsSignInApi.AuditLogsListSignIns(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsSignInApi.AuditLogsListSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsListSignIns`: CollectionOfSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsSignInApi.AuditLogsListSignIns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsListSignInsRequest struct via the builder pattern


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

[**CollectionOfSignIn**](CollectionOfSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateSignIns

> AuditLogsUpdateSignIns(ctx, signInId).MicrosoftGraphSignIn(microsoftGraphSignIn).Execute()

Update the navigation property signIns in auditLogs



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
    signInId := "signInId_example" // string | key: id of signIn
    microsoftGraphSignIn := *openapiclient.NewMicrosoftGraphSignIn() // MicrosoftGraphSignIn | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsSignInApi.AuditLogsUpdateSignIns(context.Background(), signInId).MicrosoftGraphSignIn(microsoftGraphSignIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsSignInApi.AuditLogsUpdateSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signInId** | **string** | key: id of signIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsUpdateSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphSignIn** | [**MicrosoftGraphSignIn**](MicrosoftGraphSignIn.md) | New navigation property values | 

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

