# \DirectoryObjectsDirectoryObjectApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DirectoryObjectsDirectoryObjectCreateDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectCreateDirectoryObject) | **Post** /directoryObjects | Add new entity to directoryObjects
[**DirectoryObjectsDirectoryObjectDeleteDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectDeleteDirectoryObject) | **Delete** /directoryObjects/{directoryObject-id} | Delete entity from directoryObjects
[**DirectoryObjectsDirectoryObjectGetDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectGetDirectoryObject) | **Get** /directoryObjects/{directoryObject-id} | Get entity from directoryObjects by key
[**DirectoryObjectsDirectoryObjectListDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectListDirectoryObject) | **Get** /directoryObjects | Get entities from directoryObjects
[**DirectoryObjectsDirectoryObjectUpdateDirectoryObject**](DirectoryObjectsDirectoryObjectApi.md#DirectoryObjectsDirectoryObjectUpdateDirectoryObject) | **Patch** /directoryObjects/{directoryObject-id} | Update entity in directoryObjects



## DirectoryObjectsDirectoryObjectCreateDirectoryObject

> MicrosoftGraphDirectoryObject DirectoryObjectsDirectoryObjectCreateDirectoryObject(ctx).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Add new entity to directoryObjects

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
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectCreateDirectoryObject(context.Background()).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectCreateDirectoryObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectCreateDirectoryObject`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectCreateDirectoryObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectCreateDirectoryObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New entity | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsDirectoryObjectDeleteDirectoryObject

> DirectoryObjectsDirectoryObjectDeleteDirectoryObject(ctx, directoryObjectId).IfMatch(ifMatch).Execute()

Delete entity from directoryObjects

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectDeleteDirectoryObject(context.Background(), directoryObjectId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectDeleteDirectoryObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectDeleteDirectoryObjectRequest struct via the builder pattern


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


## DirectoryObjectsDirectoryObjectGetDirectoryObject

> MicrosoftGraphDirectoryObject DirectoryObjectsDirectoryObjectGetDirectoryObject(ctx, directoryObjectId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()

Get entity from directoryObjects by key

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectGetDirectoryObject(context.Background(), directoryObjectId).ConsistencyLevel(consistencyLevel).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectGetDirectoryObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectGetDirectoryObject`: MicrosoftGraphDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectGetDirectoryObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectGetDirectoryObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsDirectoryObjectListDirectoryObject

> CollectionOfDirectoryObject DirectoryObjectsDirectoryObjectListDirectoryObject(ctx).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from directoryObjects

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
    consistencyLevel := "eventual" // string | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ (optional)
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
    resp, r, err := api_client.DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectListDirectoryObject(context.Background()).ConsistencyLevel(consistencyLevel).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectListDirectoryObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DirectoryObjectsDirectoryObjectListDirectoryObject`: CollectionOfDirectoryObject
    fmt.Fprintf(os.Stdout, "Response from `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectListDirectoryObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectListDirectoryObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consistencyLevel** | **string** | Indicates the requested consistency level. Documentation URL: https://developer.microsoft.com/en-us/office/blogs/microsoft-graph-advanced-queries-for-directory-objects-are-now-generally-available/ | 
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**CollectionOfDirectoryObject**](CollectionOfDirectoryObject.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DirectoryObjectsDirectoryObjectUpdateDirectoryObject

> DirectoryObjectsDirectoryObjectUpdateDirectoryObject(ctx, directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()

Update entity in directoryObjects

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
    directoryObjectId := "directoryObjectId_example" // string | key: id of directoryObject
    microsoftGraphDirectoryObject := *openapiclient.NewMicrosoftGraphDirectoryObject() // MicrosoftGraphDirectoryObject | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectUpdateDirectoryObject(context.Background(), directoryObjectId).MicrosoftGraphDirectoryObject(microsoftGraphDirectoryObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DirectoryObjectsDirectoryObjectApi.DirectoryObjectsDirectoryObjectUpdateDirectoryObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**directoryObjectId** | **string** | key: id of directoryObject | 

### Other Parameters

Other parameters are passed through a pointer to a apiDirectoryObjectsDirectoryObjectUpdateDirectoryObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphDirectoryObject** | [**MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | New property values | 

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

