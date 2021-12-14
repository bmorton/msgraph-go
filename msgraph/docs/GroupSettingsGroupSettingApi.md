# \GroupSettingsGroupSettingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupSettingsGroupSettingCreateGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingCreateGroupSetting) | **Post** /groupSettings | Add new entity to groupSettings
[**GroupSettingsGroupSettingDeleteGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingDeleteGroupSetting) | **Delete** /groupSettings/{groupSetting-id} | Delete entity from groupSettings
[**GroupSettingsGroupSettingGetGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingGetGroupSetting) | **Get** /groupSettings/{groupSetting-id} | Get entity from groupSettings by key
[**GroupSettingsGroupSettingListGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingListGroupSetting) | **Get** /groupSettings | Get entities from groupSettings
[**GroupSettingsGroupSettingUpdateGroupSetting**](GroupSettingsGroupSettingApi.md#GroupSettingsGroupSettingUpdateGroupSetting) | **Patch** /groupSettings/{groupSetting-id} | Update entity in groupSettings



## GroupSettingsGroupSettingCreateGroupSetting

> MicrosoftGraphGroupSetting GroupSettingsGroupSettingCreateGroupSetting(ctx).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()

Add new entity to groupSettings

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
    microsoftGraphGroupSetting := *openapiclient.NewMicrosoftGraphGroupSetting() // MicrosoftGraphGroupSetting | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingsGroupSettingApi.GroupSettingsGroupSettingCreateGroupSetting(context.Background()).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingCreateGroupSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingsGroupSettingCreateGroupSetting`: MicrosoftGraphGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingCreateGroupSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingsGroupSettingCreateGroupSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md) | New entity | 

### Return type

[**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingsGroupSettingDeleteGroupSetting

> GroupSettingsGroupSettingDeleteGroupSetting(ctx, groupSettingId).IfMatch(ifMatch).Execute()

Delete entity from groupSettings

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
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingsGroupSettingApi.GroupSettingsGroupSettingDeleteGroupSetting(context.Background(), groupSettingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingDeleteGroupSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingsGroupSettingDeleteGroupSettingRequest struct via the builder pattern


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


## GroupSettingsGroupSettingGetGroupSetting

> MicrosoftGraphGroupSetting GroupSettingsGroupSettingGetGroupSetting(ctx, groupSettingId).Select_(select_).Expand(expand).Execute()

Get entity from groupSettings by key

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
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingsGroupSettingApi.GroupSettingsGroupSettingGetGroupSetting(context.Background(), groupSettingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingGetGroupSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingsGroupSettingGetGroupSetting`: MicrosoftGraphGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingGetGroupSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingsGroupSettingGetGroupSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingsGroupSettingListGroupSetting

> CollectionOfGroupSetting GroupSettingsGroupSettingListGroupSetting(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from groupSettings

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
    resp, r, err := api_client.GroupSettingsGroupSettingApi.GroupSettingsGroupSettingListGroupSetting(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingListGroupSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupSettingsGroupSettingListGroupSetting`: CollectionOfGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingListGroupSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingsGroupSettingListGroupSettingRequest struct via the builder pattern


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

[**CollectionOfGroupSetting**](CollectionOfGroupSetting.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupSettingsGroupSettingUpdateGroupSetting

> GroupSettingsGroupSettingUpdateGroupSetting(ctx, groupSettingId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()

Update entity in groupSettings

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
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    microsoftGraphGroupSetting := *openapiclient.NewMicrosoftGraphGroupSetting() // MicrosoftGraphGroupSetting | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupSettingsGroupSettingApi.GroupSettingsGroupSettingUpdateGroupSetting(context.Background(), groupSettingId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupSettingsGroupSettingApi.GroupSettingsGroupSettingUpdateGroupSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupSettingsGroupSettingUpdateGroupSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md) | New property values | 

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

