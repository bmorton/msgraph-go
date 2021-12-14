# \MeUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeActivitiesCreateHistoryItems**](MeUserActivityApi.md#MeActivitiesCreateHistoryItems) | **Post** /me/activities/{userActivity-id}/historyItems | Create new navigation property to historyItems for me
[**MeActivitiesDeleteHistoryItems**](MeUserActivityApi.md#MeActivitiesDeleteHistoryItems) | **Delete** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Delete navigation property historyItems for me
[**MeActivitiesGetHistoryItems**](MeUserActivityApi.md#MeActivitiesGetHistoryItems) | **Get** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Get historyItems from me
[**MeActivitiesHistoryItemsDeleteRefActivity**](MeUserActivityApi.md#MeActivitiesHistoryItemsDeleteRefActivity) | **Delete** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Delete ref of navigation property activity for me
[**MeActivitiesHistoryItemsGetActivity**](MeUserActivityApi.md#MeActivitiesHistoryItemsGetActivity) | **Get** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity | Get activity from me
[**MeActivitiesHistoryItemsGetRefActivity**](MeUserActivityApi.md#MeActivitiesHistoryItemsGetRefActivity) | **Get** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Get ref of activity from me
[**MeActivitiesHistoryItemsUpdateRefActivity**](MeUserActivityApi.md#MeActivitiesHistoryItemsUpdateRefActivity) | **Put** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Update the ref of navigation property activity in me
[**MeActivitiesListHistoryItems**](MeUserActivityApi.md#MeActivitiesListHistoryItems) | **Get** /me/activities/{userActivity-id}/historyItems | Get historyItems from me
[**MeActivitiesUpdateHistoryItems**](MeUserActivityApi.md#MeActivitiesUpdateHistoryItems) | **Patch** /me/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Update the navigation property historyItems in me
[**MeCreateActivities**](MeUserActivityApi.md#MeCreateActivities) | **Post** /me/activities | Create new navigation property to activities for me
[**MeDeleteActivities**](MeUserActivityApi.md#MeDeleteActivities) | **Delete** /me/activities/{userActivity-id} | Delete navigation property activities for me
[**MeGetActivities**](MeUserActivityApi.md#MeGetActivities) | **Get** /me/activities/{userActivity-id} | Get activities from me
[**MeListActivities**](MeUserActivityApi.md#MeListActivities) | **Get** /me/activities | Get activities from me
[**MeUpdateActivities**](MeUserActivityApi.md#MeUpdateActivities) | **Patch** /me/activities/{userActivity-id} | Update the navigation property activities in me



## MeActivitiesCreateHistoryItems

> MicrosoftGraphActivityHistoryItem MeActivitiesCreateHistoryItems(ctx, userActivityId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()

Create new navigation property to historyItems for me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    microsoftGraphActivityHistoryItem := *openapiclient.NewMicrosoftGraphActivityHistoryItem() // MicrosoftGraphActivityHistoryItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesCreateHistoryItems(context.Background(), userActivityId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesCreateHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeActivitiesCreateHistoryItems`: MicrosoftGraphActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeActivitiesCreateHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesCreateHistoryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md) | New navigation property | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeActivitiesDeleteHistoryItems

> MeActivitiesDeleteHistoryItems(ctx, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()

Delete navigation property historyItems for me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesDeleteHistoryItems(context.Background(), userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesDeleteHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesDeleteHistoryItemsRequest struct via the builder pattern


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


## MeActivitiesGetHistoryItems

> MicrosoftGraphActivityHistoryItem MeActivitiesGetHistoryItems(ctx, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()

Get historyItems from me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesGetHistoryItems(context.Background(), userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesGetHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeActivitiesGetHistoryItems`: MicrosoftGraphActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeActivitiesGetHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesGetHistoryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeActivitiesHistoryItemsDeleteRefActivity

> MeActivitiesHistoryItemsDeleteRefActivity(ctx, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property activity for me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesHistoryItemsDeleteRefActivity(context.Background(), userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesHistoryItemsDeleteRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesHistoryItemsDeleteRefActivityRequest struct via the builder pattern


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


## MeActivitiesHistoryItemsGetActivity

> MicrosoftGraphUserActivity MeActivitiesHistoryItemsGetActivity(ctx, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()

Get activity from me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesHistoryItemsGetActivity(context.Background(), userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesHistoryItemsGetActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeActivitiesHistoryItemsGetActivity`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeActivitiesHistoryItemsGetActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesHistoryItemsGetActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeActivitiesHistoryItemsGetRefActivity

> string MeActivitiesHistoryItemsGetRefActivity(ctx, userActivityId, activityHistoryItemId).Execute()

Get ref of activity from me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesHistoryItemsGetRefActivity(context.Background(), userActivityId, activityHistoryItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesHistoryItemsGetRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeActivitiesHistoryItemsGetRefActivity`: string
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeActivitiesHistoryItemsGetRefActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesHistoryItemsGetRefActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeActivitiesHistoryItemsUpdateRefActivity

> MeActivitiesHistoryItemsUpdateRefActivity(ctx, userActivityId, activityHistoryItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property activity in me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesHistoryItemsUpdateRefActivity(context.Background(), userActivityId, activityHistoryItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesHistoryItemsUpdateRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesHistoryItemsUpdateRefActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]map[string]interface{}** | New navigation property ref values | 

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


## MeActivitiesListHistoryItems

> CollectionOfActivityHistoryItem MeActivitiesListHistoryItems(ctx, userActivityId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get historyItems from me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
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
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesListHistoryItems(context.Background(), userActivityId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesListHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeActivitiesListHistoryItems`: CollectionOfActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeActivitiesListHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesListHistoryItemsRequest struct via the builder pattern


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

[**CollectionOfActivityHistoryItem**](CollectionOfActivityHistoryItem.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeActivitiesUpdateHistoryItems

> MeActivitiesUpdateHistoryItems(ctx, userActivityId, activityHistoryItemId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()

Update the navigation property historyItems in me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    microsoftGraphActivityHistoryItem := *openapiclient.NewMicrosoftGraphActivityHistoryItem() // MicrosoftGraphActivityHistoryItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeActivitiesUpdateHistoryItems(context.Background(), userActivityId, activityHistoryItemId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeActivitiesUpdateHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeActivitiesUpdateHistoryItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphActivityHistoryItem** | [**MicrosoftGraphActivityHistoryItem**](MicrosoftGraphActivityHistoryItem.md) | New navigation property values | 

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


## MeCreateActivities

> MicrosoftGraphUserActivity MeCreateActivities(ctx).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()

Create new navigation property to activities for me



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
    microsoftGraphUserActivity := *openapiclient.NewMicrosoftGraphUserActivity() // MicrosoftGraphUserActivity | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeCreateActivities(context.Background()).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeCreateActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateActivities`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeCreateActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md) | New navigation property | 

### Return type

[**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeDeleteActivities

> MeDeleteActivities(ctx, userActivityId).IfMatch(ifMatch).Execute()

Delete navigation property activities for me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeDeleteActivities(context.Background(), userActivityId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeDeleteActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteActivitiesRequest struct via the builder pattern


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


## MeGetActivities

> MicrosoftGraphUserActivity MeGetActivities(ctx, userActivityId).Select_(select_).Expand(expand).Execute()

Get activities from me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeGetActivities(context.Background(), userActivityId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeGetActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetActivities`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeGetActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeListActivities

> CollectionOfUserActivity MeListActivities(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get activities from me



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
    resp, r, err := api_client.MeUserActivityApi.MeListActivities(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeListActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListActivities`: CollectionOfUserActivity
    fmt.Fprintf(os.Stdout, "Response from `MeUserActivityApi.MeListActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListActivitiesRequest struct via the builder pattern


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

[**CollectionOfUserActivity**](CollectionOfUserActivity.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeUpdateActivities

> MeUpdateActivities(ctx, userActivityId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()

Update the navigation property activities in me



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
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    microsoftGraphUserActivity := *openapiclient.NewMicrosoftGraphUserActivity() // MicrosoftGraphUserActivity | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeUserActivityApi.MeUpdateActivities(context.Background(), userActivityId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeUserActivityApi.MeUpdateActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphUserActivity** | [**MicrosoftGraphUserActivity**](MicrosoftGraphUserActivity.md) | New navigation property values | 

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

