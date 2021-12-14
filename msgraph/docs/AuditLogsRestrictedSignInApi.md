# \AuditLogsRestrictedSignInApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsCreateRestrictedSignIns) | **Post** /auditLogs/restrictedSignIns | Create new navigation property to restrictedSignIns for auditLogs
[**AuditLogsDeleteRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsDeleteRestrictedSignIns) | **Delete** /auditLogs/restrictedSignIns/{restrictedSignIn-id} | Delete navigation property restrictedSignIns for auditLogs
[**AuditLogsGetRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsGetRestrictedSignIns) | **Get** /auditLogs/restrictedSignIns/{restrictedSignIn-id} | Get restrictedSignIns from auditLogs
[**AuditLogsListRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsListRestrictedSignIns) | **Get** /auditLogs/restrictedSignIns | Get restrictedSignIns from auditLogs
[**AuditLogsUpdateRestrictedSignIns**](AuditLogsRestrictedSignInApi.md#AuditLogsUpdateRestrictedSignIns) | **Patch** /auditLogs/restrictedSignIns/{restrictedSignIn-id} | Update the navigation property restrictedSignIns in auditLogs



## AuditLogsCreateRestrictedSignIns

> MicrosoftGraphRestrictedSignIn AuditLogsCreateRestrictedSignIns(ctx).MicrosoftGraphRestrictedSignIn(microsoftGraphRestrictedSignIn).Execute()

Create new navigation property to restrictedSignIns for auditLogs

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
    microsoftGraphRestrictedSignIn := *openapiclient.NewMicrosoftGraphRestrictedSignIn() // MicrosoftGraphRestrictedSignIn | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsRestrictedSignInApi.AuditLogsCreateRestrictedSignIns(context.Background()).MicrosoftGraphRestrictedSignIn(microsoftGraphRestrictedSignIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsRestrictedSignInApi.AuditLogsCreateRestrictedSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsCreateRestrictedSignIns`: MicrosoftGraphRestrictedSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsRestrictedSignInApi.AuditLogsCreateRestrictedSignIns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsCreateRestrictedSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphRestrictedSignIn** | [**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md) | New navigation property | 

### Return type

[**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsDeleteRestrictedSignIns

> AuditLogsDeleteRestrictedSignIns(ctx, restrictedSignInId).IfMatch(ifMatch).Execute()

Delete navigation property restrictedSignIns for auditLogs

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
    restrictedSignInId := "restrictedSignInId_example" // string | key: id of restrictedSignIn
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsRestrictedSignInApi.AuditLogsDeleteRestrictedSignIns(context.Background(), restrictedSignInId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsRestrictedSignInApi.AuditLogsDeleteRestrictedSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restrictedSignInId** | **string** | key: id of restrictedSignIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsDeleteRestrictedSignInsRequest struct via the builder pattern


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


## AuditLogsGetRestrictedSignIns

> MicrosoftGraphRestrictedSignIn AuditLogsGetRestrictedSignIns(ctx, restrictedSignInId).Select_(select_).Expand(expand).Execute()

Get restrictedSignIns from auditLogs

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
    restrictedSignInId := "restrictedSignInId_example" // string | key: id of restrictedSignIn
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsRestrictedSignInApi.AuditLogsGetRestrictedSignIns(context.Background(), restrictedSignInId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsRestrictedSignInApi.AuditLogsGetRestrictedSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsGetRestrictedSignIns`: MicrosoftGraphRestrictedSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsRestrictedSignInApi.AuditLogsGetRestrictedSignIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restrictedSignInId** | **string** | key: id of restrictedSignIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsGetRestrictedSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListRestrictedSignIns

> CollectionOfRestrictedSignIn AuditLogsListRestrictedSignIns(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get restrictedSignIns from auditLogs

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
    resp, r, err := api_client.AuditLogsRestrictedSignInApi.AuditLogsListRestrictedSignIns(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsRestrictedSignInApi.AuditLogsListRestrictedSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsListRestrictedSignIns`: CollectionOfRestrictedSignIn
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsRestrictedSignInApi.AuditLogsListRestrictedSignIns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsListRestrictedSignInsRequest struct via the builder pattern


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

[**CollectionOfRestrictedSignIn**](CollectionOfRestrictedSignIn.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateRestrictedSignIns

> AuditLogsUpdateRestrictedSignIns(ctx, restrictedSignInId).MicrosoftGraphRestrictedSignIn(microsoftGraphRestrictedSignIn).Execute()

Update the navigation property restrictedSignIns in auditLogs

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
    restrictedSignInId := "restrictedSignInId_example" // string | key: id of restrictedSignIn
    microsoftGraphRestrictedSignIn := *openapiclient.NewMicrosoftGraphRestrictedSignIn() // MicrosoftGraphRestrictedSignIn | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsRestrictedSignInApi.AuditLogsUpdateRestrictedSignIns(context.Background(), restrictedSignInId).MicrosoftGraphRestrictedSignIn(microsoftGraphRestrictedSignIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsRestrictedSignInApi.AuditLogsUpdateRestrictedSignIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restrictedSignInId** | **string** | key: id of restrictedSignIn | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsUpdateRestrictedSignInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphRestrictedSignIn** | [**MicrosoftGraphRestrictedSignIn**](MicrosoftGraphRestrictedSignIn.md) | New navigation property values | 

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

