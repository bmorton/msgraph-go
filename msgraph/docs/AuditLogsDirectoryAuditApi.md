# \AuditLogsDirectoryAuditApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsCreateDirectoryAudits) | **Post** /auditLogs/directoryAudits | Create new navigation property to directoryAudits for auditLogs
[**AuditLogsDeleteDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsDeleteDirectoryAudits) | **Delete** /auditLogs/directoryAudits/{directoryAudit-id} | Delete navigation property directoryAudits for auditLogs
[**AuditLogsGetDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsGetDirectoryAudits) | **Get** /auditLogs/directoryAudits/{directoryAudit-id} | Get directoryAudits from auditLogs
[**AuditLogsListDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsListDirectoryAudits) | **Get** /auditLogs/directoryAudits | Get directoryAudits from auditLogs
[**AuditLogsUpdateDirectoryAudits**](AuditLogsDirectoryAuditApi.md#AuditLogsUpdateDirectoryAudits) | **Patch** /auditLogs/directoryAudits/{directoryAudit-id} | Update the navigation property directoryAudits in auditLogs



## AuditLogsCreateDirectoryAudits

> MicrosoftGraphDirectoryAudit AuditLogsCreateDirectoryAudits(ctx).MicrosoftGraphDirectoryAudit(microsoftGraphDirectoryAudit).Execute()

Create new navigation property to directoryAudits for auditLogs



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
    microsoftGraphDirectoryAudit := *openapiclient.NewMicrosoftGraphDirectoryAudit() // MicrosoftGraphDirectoryAudit | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsDirectoryAuditApi.AuditLogsCreateDirectoryAudits(context.Background()).MicrosoftGraphDirectoryAudit(microsoftGraphDirectoryAudit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsDirectoryAuditApi.AuditLogsCreateDirectoryAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsCreateDirectoryAudits`: MicrosoftGraphDirectoryAudit
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsDirectoryAuditApi.AuditLogsCreateDirectoryAudits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsCreateDirectoryAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectoryAudit** | [**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md) | New navigation property | 

### Return type

[**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsDeleteDirectoryAudits

> AuditLogsDeleteDirectoryAudits(ctx, directoryAuditId).IfMatch(ifMatch).Execute()

Delete navigation property directoryAudits for auditLogs



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
    directoryAuditId := "directoryAuditId_example" // string | key: id of directoryAudit
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsDirectoryAuditApi.AuditLogsDeleteDirectoryAudits(context.Background(), directoryAuditId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsDirectoryAuditApi.AuditLogsDeleteDirectoryAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryAuditId** | **string** | key: id of directoryAudit | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsDeleteDirectoryAuditsRequest struct via the builder pattern


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


## AuditLogsGetDirectoryAudits

> MicrosoftGraphDirectoryAudit AuditLogsGetDirectoryAudits(ctx, directoryAuditId).Select_(select_).Expand(expand).Execute()

Get directoryAudits from auditLogs



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
    directoryAuditId := "directoryAuditId_example" // string | key: id of directoryAudit
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsDirectoryAuditApi.AuditLogsGetDirectoryAudits(context.Background(), directoryAuditId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsDirectoryAuditApi.AuditLogsGetDirectoryAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsGetDirectoryAudits`: MicrosoftGraphDirectoryAudit
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsDirectoryAuditApi.AuditLogsGetDirectoryAudits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryAuditId** | **string** | key: id of directoryAudit | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsGetDirectoryAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListDirectoryAudits

> CollectionOfDirectoryAudit AuditLogsListDirectoryAudits(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get directoryAudits from auditLogs



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
    resp, r, err := api_client.AuditLogsDirectoryAuditApi.AuditLogsListDirectoryAudits(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsDirectoryAuditApi.AuditLogsListDirectoryAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsListDirectoryAudits`: CollectionOfDirectoryAudit
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsDirectoryAuditApi.AuditLogsListDirectoryAudits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsListDirectoryAuditsRequest struct via the builder pattern


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

[**CollectionOfDirectoryAudit**](CollectionOfDirectoryAudit.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateDirectoryAudits

> AuditLogsUpdateDirectoryAudits(ctx, directoryAuditId).MicrosoftGraphDirectoryAudit(microsoftGraphDirectoryAudit).Execute()

Update the navigation property directoryAudits in auditLogs



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
    directoryAuditId := "directoryAuditId_example" // string | key: id of directoryAudit
    microsoftGraphDirectoryAudit := *openapiclient.NewMicrosoftGraphDirectoryAudit() // MicrosoftGraphDirectoryAudit | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsDirectoryAuditApi.AuditLogsUpdateDirectoryAudits(context.Background(), directoryAuditId).MicrosoftGraphDirectoryAudit(microsoftGraphDirectoryAudit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsDirectoryAuditApi.AuditLogsUpdateDirectoryAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryAuditId** | **string** | key: id of directoryAudit | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsUpdateDirectoryAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryAudit** | [**MicrosoftGraphDirectoryAudit**](MicrosoftGraphDirectoryAudit.md) | New navigation property values | 

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

