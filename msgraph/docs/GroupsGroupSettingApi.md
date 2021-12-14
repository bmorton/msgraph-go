# \GroupsGroupSettingApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsCreateSettings**](GroupsGroupSettingApi.md#GroupsCreateSettings) | **Post** /groups/{group-id}/settings | Create new navigation property to settings for groups
[**GroupsDeleteSettings**](GroupsGroupSettingApi.md#GroupsDeleteSettings) | **Delete** /groups/{group-id}/settings/{groupSetting-id} | Delete navigation property settings for groups
[**GroupsGetSettings**](GroupsGroupSettingApi.md#GroupsGetSettings) | **Get** /groups/{group-id}/settings/{groupSetting-id} | Get settings from groups
[**GroupsListSettings**](GroupsGroupSettingApi.md#GroupsListSettings) | **Get** /groups/{group-id}/settings | Get settings from groups
[**GroupsUpdateSettings**](GroupsGroupSettingApi.md#GroupsUpdateSettings) | **Patch** /groups/{group-id}/settings/{groupSetting-id} | Update the navigation property settings in groups



## GroupsCreateSettings

> MicrosoftGraphGroupSetting GroupsCreateSettings(ctx, groupId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()

Create new navigation property to settings for groups



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
    groupId := "groupId_example" // string | key: id of group
    microsoftGraphGroupSetting := *openapiclient.NewMicrosoftGraphGroupSetting() // MicrosoftGraphGroupSetting | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupSettingApi.GroupsCreateSettings(context.Background(), groupId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupSettingApi.GroupsCreateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsCreateSettings`: MicrosoftGraphGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupSettingApi.GroupsCreateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsCreateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md) | New navigation property | 

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


## GroupsDeleteSettings

> GroupsDeleteSettings(ctx, groupId, groupSettingId).IfMatch(ifMatch).Execute()

Delete navigation property settings for groups



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
    groupId := "groupId_example" // string | key: id of group
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupSettingApi.GroupsDeleteSettings(context.Background(), groupId, groupSettingId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupSettingApi.GroupsDeleteSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteSettingsRequest struct via the builder pattern


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


## GroupsGetSettings

> MicrosoftGraphGroupSetting GroupsGetSettings(ctx, groupId, groupSettingId).Select_(select_).Expand(expand).Execute()

Get settings from groups



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
    groupId := "groupId_example" // string | key: id of group
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupSettingApi.GroupsGetSettings(context.Background(), groupId, groupSettingId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupSettingApi.GroupsGetSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetSettings`: MicrosoftGraphGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupSettingApi.GroupsGetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetSettingsRequest struct via the builder pattern


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


## GroupsListSettings

> CollectionOfGroupSetting GroupsListSettings(ctx, groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get settings from groups



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
    groupId := "groupId_example" // string | key: id of group
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
    resp, r, err := api_client.GroupsGroupSettingApi.GroupsListSettings(context.Background(), groupId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupSettingApi.GroupsListSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsListSettings`: CollectionOfGroupSetting
    fmt.Fprintf(os.Stdout, "Response from `GroupsGroupSettingApi.GroupsListSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsListSettingsRequest struct via the builder pattern


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


## GroupsUpdateSettings

> GroupsUpdateSettings(ctx, groupId, groupSettingId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()

Update the navigation property settings in groups



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
    groupId := "groupId_example" // string | key: id of group
    groupSettingId := "groupSettingId_example" // string | key: id of groupSetting
    microsoftGraphGroupSetting := *openapiclient.NewMicrosoftGraphGroupSetting() // MicrosoftGraphGroupSetting | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsGroupSettingApi.GroupsUpdateSettings(context.Background(), groupId, groupSettingId).MicrosoftGraphGroupSetting(microsoftGraphGroupSetting).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsGroupSettingApi.GroupsUpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 
**groupSettingId** | **string** | key: id of groupSetting | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphGroupSetting** | [**MicrosoftGraphGroupSetting**](MicrosoftGraphGroupSetting.md) | New navigation property values | 

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

