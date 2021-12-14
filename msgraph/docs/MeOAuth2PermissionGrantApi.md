# \MeOAuth2PermissionGrantApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateRefOauth2PermissionGrants**](MeOAuth2PermissionGrantApi.md#MeCreateRefOauth2PermissionGrants) | **Post** /me/oauth2PermissionGrants/$ref | Create new navigation property ref to oauth2PermissionGrants for me
[**MeListOauth2PermissionGrants**](MeOAuth2PermissionGrantApi.md#MeListOauth2PermissionGrants) | **Get** /me/oauth2PermissionGrants | Get oauth2PermissionGrants from me
[**MeListRefOauth2PermissionGrants**](MeOAuth2PermissionGrantApi.md#MeListRefOauth2PermissionGrants) | **Get** /me/oauth2PermissionGrants/$ref | Get ref of oauth2PermissionGrants from me



## MeCreateRefOauth2PermissionGrants

> map[string]interface{} MeCreateRefOauth2PermissionGrants(ctx).RequestBody(requestBody).Execute()

Create new navigation property ref to oauth2PermissionGrants for me

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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref value

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOAuth2PermissionGrantApi.MeCreateRefOauth2PermissionGrants(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOAuth2PermissionGrantApi.MeCreateRefOauth2PermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateRefOauth2PermissionGrants`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MeOAuth2PermissionGrantApi.MeCreateRefOauth2PermissionGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateRefOauth2PermissionGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref value | 

### Return type

**map[string]interface{}**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListOauth2PermissionGrants

> CollectionOfOAuth2PermissionGrant MeListOauth2PermissionGrants(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get oauth2PermissionGrants from me

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
    resp, r, err := api_client.MeOAuth2PermissionGrantApi.MeListOauth2PermissionGrants(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOAuth2PermissionGrantApi.MeListOauth2PermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListOauth2PermissionGrants`: CollectionOfOAuth2PermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `MeOAuth2PermissionGrantApi.MeListOauth2PermissionGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListOauth2PermissionGrantsRequest struct via the builder pattern


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

[**CollectionOfOAuth2PermissionGrant**](CollectionOfOAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListRefOauth2PermissionGrants

> CollectionOfLinksOfOAuth2PermissionGrant MeListRefOauth2PermissionGrants(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()

Get ref of oauth2PermissionGrants from me

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeOAuth2PermissionGrantApi.MeListRefOauth2PermissionGrants(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeOAuth2PermissionGrantApi.MeListRefOauth2PermissionGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListRefOauth2PermissionGrants`: CollectionOfLinksOfOAuth2PermissionGrant
    fmt.Fprintf(os.Stdout, "Response from `MeOAuth2PermissionGrantApi.MeListRefOauth2PermissionGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListRefOauth2PermissionGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Show only the first n items | 
 **skip** | **int32** | Skip the first n items | 
 **search** | **string** | Search items by search phrases | 
 **filter** | **string** | Filter items by property values | 
 **count** | **bool** | Include count of items | 
 **orderby** | **[]string** | Order items by property values | 

### Return type

[**CollectionOfLinksOfOAuth2PermissionGrant**](CollectionOfLinksOfOAuth2PermissionGrant.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

