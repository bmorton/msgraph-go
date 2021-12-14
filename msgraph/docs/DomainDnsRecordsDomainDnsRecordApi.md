# \DomainDnsRecordsDomainDnsRecordApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord) | **Post** /domainDnsRecords | Add new entity to domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord) | **Delete** /domainDnsRecords/{domainDnsRecord-id} | Delete entity from domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord) | **Get** /domainDnsRecords/{domainDnsRecord-id} | Get entity from domainDnsRecords by key
[**DomainDnsRecordsDomainDnsRecordListDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordListDomainDnsRecord) | **Get** /domainDnsRecords | Get entities from domainDnsRecords
[**DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord**](DomainDnsRecordsDomainDnsRecordApi.md#DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord) | **Patch** /domainDnsRecords/{domainDnsRecord-id} | Update entity in domainDnsRecords



## DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord

> MicrosoftGraphDomainDnsRecord DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord(ctx).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Add new entity to domainDnsRecords

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
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord(context.Background()).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordCreateDomainDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDnsRecordsDomainDnsRecordCreateDomainDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New entity | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord

> DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord(ctx, domainDnsRecordId).IfMatch(ifMatch).Execute()

Delete entity from domainDnsRecords

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
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord(context.Background(), domainDnsRecordId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDnsRecordsDomainDnsRecordDeleteDomainDnsRecordRequest struct via the builder pattern


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


## DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord

> MicrosoftGraphDomainDnsRecord DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord(ctx, domainDnsRecordId).Select_(select_).Expand(expand).Execute()

Get entity from domainDnsRecords by key

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
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord(context.Background(), domainDnsRecordId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord`: MicrosoftGraphDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordGetDomainDnsRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDnsRecordsDomainDnsRecordGetDomainDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordListDomainDnsRecord

> CollectionOfDomainDnsRecord DomainDnsRecordsDomainDnsRecordListDomainDnsRecord(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from domainDnsRecords

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
    resp, r, err := api_client.DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordListDomainDnsRecord(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordListDomainDnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainDnsRecordsDomainDnsRecordListDomainDnsRecord`: CollectionOfDomainDnsRecord
    fmt.Fprintf(os.Stdout, "Response from `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordListDomainDnsRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDomainDnsRecordsDomainDnsRecordListDomainDnsRecordRequest struct via the builder pattern


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

[**CollectionOfDomainDnsRecord**](CollectionOfDomainDnsRecord.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord

> DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord(ctx, domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()

Update entity in domainDnsRecords

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
    domainDnsRecordId := "domainDnsRecordId_example" // string | key: id of domainDnsRecord
    microsoftGraphDomainDnsRecord := *openapiclient.NewMicrosoftGraphDomainDnsRecord() // MicrosoftGraphDomainDnsRecord | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord(context.Background(), domainDnsRecordId).MicrosoftGraphDomainDnsRecord(microsoftGraphDomainDnsRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainDnsRecordsDomainDnsRecordApi.DomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainDnsRecordId** | **string** | key: id of domainDnsRecord | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainDnsRecordsDomainDnsRecordUpdateDomainDnsRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDomainDnsRecord** | [**MicrosoftGraphDomainDnsRecord**](MicrosoftGraphDomainDnsRecord.md) | New property values | 

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

