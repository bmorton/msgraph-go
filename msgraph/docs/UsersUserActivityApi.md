# \UsersUserActivityApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersActivitiesCreateHistoryItems**](UsersUserActivityApi.md#UsersActivitiesCreateHistoryItems) | **Post** /users/{user-id}/activities/{userActivity-id}/historyItems | Create new navigation property to historyItems for users
[**UsersActivitiesDeleteHistoryItems**](UsersUserActivityApi.md#UsersActivitiesDeleteHistoryItems) | **Delete** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Delete navigation property historyItems for users
[**UsersActivitiesGetHistoryItems**](UsersUserActivityApi.md#UsersActivitiesGetHistoryItems) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Get historyItems from users
[**UsersActivitiesHistoryItemsDeleteRefActivity**](UsersUserActivityApi.md#UsersActivitiesHistoryItemsDeleteRefActivity) | **Delete** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Delete ref of navigation property activity for users
[**UsersActivitiesHistoryItemsGetActivity**](UsersUserActivityApi.md#UsersActivitiesHistoryItemsGetActivity) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity | Get activity from users
[**UsersActivitiesHistoryItemsGetRefActivity**](UsersUserActivityApi.md#UsersActivitiesHistoryItemsGetRefActivity) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Get ref of activity from users
[**UsersActivitiesHistoryItemsUpdateRefActivity**](UsersUserActivityApi.md#UsersActivitiesHistoryItemsUpdateRefActivity) | **Put** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id}/activity/$ref | Update the ref of navigation property activity in users
[**UsersActivitiesListHistoryItems**](UsersUserActivityApi.md#UsersActivitiesListHistoryItems) | **Get** /users/{user-id}/activities/{userActivity-id}/historyItems | Get historyItems from users
[**UsersActivitiesUpdateHistoryItems**](UsersUserActivityApi.md#UsersActivitiesUpdateHistoryItems) | **Patch** /users/{user-id}/activities/{userActivity-id}/historyItems/{activityHistoryItem-id} | Update the navigation property historyItems in users
[**UsersCreateActivities**](UsersUserActivityApi.md#UsersCreateActivities) | **Post** /users/{user-id}/activities | Create new navigation property to activities for users
[**UsersDeleteActivities**](UsersUserActivityApi.md#UsersDeleteActivities) | **Delete** /users/{user-id}/activities/{userActivity-id} | Delete navigation property activities for users
[**UsersGetActivities**](UsersUserActivityApi.md#UsersGetActivities) | **Get** /users/{user-id}/activities/{userActivity-id} | Get activities from users
[**UsersListActivities**](UsersUserActivityApi.md#UsersListActivities) | **Get** /users/{user-id}/activities | Get activities from users
[**UsersUpdateActivities**](UsersUserActivityApi.md#UsersUpdateActivities) | **Patch** /users/{user-id}/activities/{userActivity-id} | Update the navigation property activities in users



## UsersActivitiesCreateHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesCreateHistoryItems(ctx, userId, userActivityId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()

Create new navigation property to historyItems for users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    microsoftGraphActivityHistoryItem := *openapiclient.NewMicrosoftGraphActivityHistoryItem() // MicrosoftGraphActivityHistoryItem | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesCreateHistoryItems(context.Background(), userId, userActivityId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesCreateHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersActivitiesCreateHistoryItems`: MicrosoftGraphActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersActivitiesCreateHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesCreateHistoryItemsRequest struct via the builder pattern


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


## UsersActivitiesDeleteHistoryItems

> UsersActivitiesDeleteHistoryItems(ctx, userId, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()

Delete navigation property historyItems for users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesDeleteHistoryItems(context.Background(), userId, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesDeleteHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesDeleteHistoryItemsRequest struct via the builder pattern


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


## UsersActivitiesGetHistoryItems

> MicrosoftGraphActivityHistoryItem UsersActivitiesGetHistoryItems(ctx, userId, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()

Get historyItems from users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesGetHistoryItems(context.Background(), userId, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesGetHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersActivitiesGetHistoryItems`: MicrosoftGraphActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersActivitiesGetHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesGetHistoryItemsRequest struct via the builder pattern


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


## UsersActivitiesHistoryItemsDeleteRefActivity

> UsersActivitiesHistoryItemsDeleteRefActivity(ctx, userId, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()

Delete ref of navigation property activity for users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesHistoryItemsDeleteRefActivity(context.Background(), userId, userActivityId, activityHistoryItemId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesHistoryItemsDeleteRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesHistoryItemsDeleteRefActivityRequest struct via the builder pattern


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


## UsersActivitiesHistoryItemsGetActivity

> MicrosoftGraphUserActivity UsersActivitiesHistoryItemsGetActivity(ctx, userId, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()

Get activity from users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesHistoryItemsGetActivity(context.Background(), userId, userActivityId, activityHistoryItemId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesHistoryItemsGetActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersActivitiesHistoryItemsGetActivity`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersActivitiesHistoryItemsGetActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesHistoryItemsGetActivityRequest struct via the builder pattern


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


## UsersActivitiesHistoryItemsGetRefActivity

> string UsersActivitiesHistoryItemsGetRefActivity(ctx, userId, userActivityId, activityHistoryItemId).Execute()

Get ref of activity from users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesHistoryItemsGetRefActivity(context.Background(), userId, userActivityId, activityHistoryItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesHistoryItemsGetRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersActivitiesHistoryItemsGetRefActivity`: string
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersActivitiesHistoryItemsGetRefActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesHistoryItemsGetRefActivityRequest struct via the builder pattern


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


## UsersActivitiesHistoryItemsUpdateRefActivity

> UsersActivitiesHistoryItemsUpdateRefActivity(ctx, userId, userActivityId, activityHistoryItemId).RequestBody(requestBody).Execute()

Update the ref of navigation property activity in users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | New navigation property ref values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesHistoryItemsUpdateRefActivity(context.Background(), userId, userActivityId, activityHistoryItemId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesHistoryItemsUpdateRefActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesHistoryItemsUpdateRefActivityRequest struct via the builder pattern


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


## UsersActivitiesListHistoryItems

> CollectionOfActivityHistoryItem UsersActivitiesListHistoryItems(ctx, userId, userActivityId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get historyItems from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesListHistoryItems(context.Background(), userId, userActivityId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesListHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersActivitiesListHistoryItems`: CollectionOfActivityHistoryItem
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersActivitiesListHistoryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesListHistoryItemsRequest struct via the builder pattern


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


## UsersActivitiesUpdateHistoryItems

> UsersActivitiesUpdateHistoryItems(ctx, userId, userActivityId, activityHistoryItemId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()

Update the navigation property historyItems in users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    activityHistoryItemId := "activityHistoryItemId_example" // string | key: id of activityHistoryItem
    microsoftGraphActivityHistoryItem := *openapiclient.NewMicrosoftGraphActivityHistoryItem() // MicrosoftGraphActivityHistoryItem | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersActivitiesUpdateHistoryItems(context.Background(), userId, userActivityId, activityHistoryItemId).MicrosoftGraphActivityHistoryItem(microsoftGraphActivityHistoryItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersActivitiesUpdateHistoryItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 
**activityHistoryItemId** | **string** | key: id of activityHistoryItem | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersActivitiesUpdateHistoryItemsRequest struct via the builder pattern


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


## UsersCreateActivities

> MicrosoftGraphUserActivity UsersCreateActivities(ctx, userId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()

Create new navigation property to activities for users



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
    userId := "userId_example" // string | key: id of user
    microsoftGraphUserActivity := *openapiclient.NewMicrosoftGraphUserActivity() // MicrosoftGraphUserActivity | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersCreateActivities(context.Background(), userId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersCreateActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateActivities`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersCreateActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateActivitiesRequest struct via the builder pattern


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


## UsersDeleteActivities

> UsersDeleteActivities(ctx, userId, userActivityId).IfMatch(ifMatch).Execute()

Delete navigation property activities for users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersDeleteActivities(context.Background(), userId, userActivityId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersDeleteActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteActivitiesRequest struct via the builder pattern


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


## UsersGetActivities

> MicrosoftGraphUserActivity UsersGetActivities(ctx, userId, userActivityId).Select_(select_).Expand(expand).Execute()

Get activities from users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersGetActivities(context.Background(), userId, userActivityId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersGetActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetActivities`: MicrosoftGraphUserActivity
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersGetActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetActivitiesRequest struct via the builder pattern


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


## UsersListActivities

> CollectionOfUserActivity UsersListActivities(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get activities from users



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
    userId := "userId_example" // string | key: id of user
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
    resp, r, err := api_client.UsersUserActivityApi.UsersListActivities(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersListActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListActivities`: CollectionOfUserActivity
    fmt.Fprintf(os.Stdout, "Response from `UsersUserActivityApi.UsersListActivities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListActivitiesRequest struct via the builder pattern


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


## UsersUpdateActivities

> UsersUpdateActivities(ctx, userId, userActivityId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()

Update the navigation property activities in users



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
    userId := "userId_example" // string | key: id of user
    userActivityId := "userActivityId_example" // string | key: id of userActivity
    microsoftGraphUserActivity := *openapiclient.NewMicrosoftGraphUserActivity() // MicrosoftGraphUserActivity | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersUserActivityApi.UsersUpdateActivities(context.Background(), userId, userActivityId).MicrosoftGraphUserActivity(microsoftGraphUserActivity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserActivityApi.UsersUpdateActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**userActivityId** | **string** | key: id of userActivity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateActivitiesRequest struct via the builder pattern


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

