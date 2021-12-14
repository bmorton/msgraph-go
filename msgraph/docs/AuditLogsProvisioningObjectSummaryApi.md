# \AuditLogsProvisioningObjectSummaryApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuditLogsCreateProvisioning**](AuditLogsProvisioningObjectSummaryApi.md#AuditLogsCreateProvisioning) | **Post** /auditLogs/provisioning | Create new navigation property to provisioning for auditLogs
[**AuditLogsDeleteProvisioning**](AuditLogsProvisioningObjectSummaryApi.md#AuditLogsDeleteProvisioning) | **Delete** /auditLogs/provisioning/{provisioningObjectSummary-id} | Delete navigation property provisioning for auditLogs
[**AuditLogsGetProvisioning**](AuditLogsProvisioningObjectSummaryApi.md#AuditLogsGetProvisioning) | **Get** /auditLogs/provisioning/{provisioningObjectSummary-id} | Get provisioning from auditLogs
[**AuditLogsListProvisioning**](AuditLogsProvisioningObjectSummaryApi.md#AuditLogsListProvisioning) | **Get** /auditLogs/provisioning | Get provisioning from auditLogs
[**AuditLogsUpdateProvisioning**](AuditLogsProvisioningObjectSummaryApi.md#AuditLogsUpdateProvisioning) | **Patch** /auditLogs/provisioning/{provisioningObjectSummary-id} | Update the navigation property provisioning in auditLogs



## AuditLogsCreateProvisioning

> MicrosoftGraphProvisioningObjectSummary AuditLogsCreateProvisioning(ctx).MicrosoftGraphProvisioningObjectSummary(microsoftGraphProvisioningObjectSummary).Execute()

Create new navigation property to provisioning for auditLogs

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
    microsoftGraphProvisioningObjectSummary := *openapiclient.NewMicrosoftGraphProvisioningObjectSummary() // MicrosoftGraphProvisioningObjectSummary | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsProvisioningObjectSummaryApi.AuditLogsCreateProvisioning(context.Background()).MicrosoftGraphProvisioningObjectSummary(microsoftGraphProvisioningObjectSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsProvisioningObjectSummaryApi.AuditLogsCreateProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsCreateProvisioning`: MicrosoftGraphProvisioningObjectSummary
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsProvisioningObjectSummaryApi.AuditLogsCreateProvisioning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsCreateProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphProvisioningObjectSummary** | [**MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md) | New navigation property | 

### Return type

[**MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsDeleteProvisioning

> AuditLogsDeleteProvisioning(ctx, provisioningObjectSummaryId).IfMatch(ifMatch).Execute()

Delete navigation property provisioning for auditLogs

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
    provisioningObjectSummaryId := "provisioningObjectSummaryId_example" // string | key: id of provisioningObjectSummary
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsProvisioningObjectSummaryApi.AuditLogsDeleteProvisioning(context.Background(), provisioningObjectSummaryId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsProvisioningObjectSummaryApi.AuditLogsDeleteProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningObjectSummaryId** | **string** | key: id of provisioningObjectSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsDeleteProvisioningRequest struct via the builder pattern


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


## AuditLogsGetProvisioning

> MicrosoftGraphProvisioningObjectSummary AuditLogsGetProvisioning(ctx, provisioningObjectSummaryId).Select_(select_).Expand(expand).Execute()

Get provisioning from auditLogs

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
    provisioningObjectSummaryId := "provisioningObjectSummaryId_example" // string | key: id of provisioningObjectSummary
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsProvisioningObjectSummaryApi.AuditLogsGetProvisioning(context.Background(), provisioningObjectSummaryId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsProvisioningObjectSummaryApi.AuditLogsGetProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsGetProvisioning`: MicrosoftGraphProvisioningObjectSummary
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsProvisioningObjectSummaryApi.AuditLogsGetProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningObjectSummaryId** | **string** | key: id of provisioningObjectSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsGetProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsListProvisioning

> CollectionOfProvisioningObjectSummary AuditLogsListProvisioning(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get provisioning from auditLogs

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
    resp, r, err := api_client.AuditLogsProvisioningObjectSummaryApi.AuditLogsListProvisioning(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsProvisioningObjectSummaryApi.AuditLogsListProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuditLogsListProvisioning`: CollectionOfProvisioningObjectSummary
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsProvisioningObjectSummaryApi.AuditLogsListProvisioning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsListProvisioningRequest struct via the builder pattern


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

[**CollectionOfProvisioningObjectSummary**](CollectionOfProvisioningObjectSummary.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuditLogsUpdateProvisioning

> AuditLogsUpdateProvisioning(ctx, provisioningObjectSummaryId).MicrosoftGraphProvisioningObjectSummary(microsoftGraphProvisioningObjectSummary).Execute()

Update the navigation property provisioning in auditLogs

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
    provisioningObjectSummaryId := "provisioningObjectSummaryId_example" // string | key: id of provisioningObjectSummary
    microsoftGraphProvisioningObjectSummary := *openapiclient.NewMicrosoftGraphProvisioningObjectSummary() // MicrosoftGraphProvisioningObjectSummary | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogsProvisioningObjectSummaryApi.AuditLogsUpdateProvisioning(context.Background(), provisioningObjectSummaryId).MicrosoftGraphProvisioningObjectSummary(microsoftGraphProvisioningObjectSummary).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsProvisioningObjectSummaryApi.AuditLogsUpdateProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningObjectSummaryId** | **string** | key: id of provisioningObjectSummary | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuditLogsUpdateProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphProvisioningObjectSummary** | [**MicrosoftGraphProvisioningObjectSummary**](MicrosoftGraphProvisioningObjectSummary.md) | New navigation property values | 

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

