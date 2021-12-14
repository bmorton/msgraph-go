# \UsersTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersCreateJoinedTeams**](UsersTeamApi.md#UsersCreateJoinedTeams) | **Post** /users/{user-id}/joinedTeams | Create new navigation property to joinedTeams for users
[**UsersDeleteJoinedTeams**](UsersTeamApi.md#UsersDeleteJoinedTeams) | **Delete** /users/{user-id}/joinedTeams/{team-id} | Delete navigation property joinedTeams for users
[**UsersGetJoinedTeams**](UsersTeamApi.md#UsersGetJoinedTeams) | **Get** /users/{user-id}/joinedTeams/{team-id} | Get joinedTeams from users
[**UsersListJoinedTeams**](UsersTeamApi.md#UsersListJoinedTeams) | **Get** /users/{user-id}/joinedTeams | Get joinedTeams from users
[**UsersUpdateJoinedTeams**](UsersTeamApi.md#UsersUpdateJoinedTeams) | **Patch** /users/{user-id}/joinedTeams/{team-id} | Update the navigation property joinedTeams in users



## UsersCreateJoinedTeams

> MicrosoftGraphTeam UsersCreateJoinedTeams(ctx, userId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Create new navigation property to joinedTeams for users



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
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTeamApi.UsersCreateJoinedTeams(context.Background(), userId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTeamApi.UsersCreateJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersCreateJoinedTeams`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `UsersTeamApi.UsersCreateJoinedTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersCreateJoinedTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersDeleteJoinedTeams

> UsersDeleteJoinedTeams(ctx, userId, teamId).IfMatch(ifMatch).Execute()

Delete navigation property joinedTeams for users



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
    teamId := "teamId_example" // string | key: id of team
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTeamApi.UsersDeleteJoinedTeams(context.Background(), userId, teamId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTeamApi.UsersDeleteJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeleteJoinedTeamsRequest struct via the builder pattern


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


## UsersGetJoinedTeams

> MicrosoftGraphTeam UsersGetJoinedTeams(ctx, userId, teamId).Select_(select_).Expand(expand).Execute()

Get joinedTeams from users



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
    teamId := "teamId_example" // string | key: id of team
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTeamApi.UsersGetJoinedTeams(context.Background(), userId, teamId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTeamApi.UsersGetJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetJoinedTeams`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `UsersTeamApi.UsersGetJoinedTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetJoinedTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeam**](MicrosoftGraphTeam.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersListJoinedTeams

> CollectionOfTeam UsersListJoinedTeams(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get joinedTeams from users



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
    resp, r, err := api_client.UsersTeamApi.UsersListJoinedTeams(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTeamApi.UsersListJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersListJoinedTeams`: CollectionOfTeam
    fmt.Fprintf(os.Stdout, "Response from `UsersTeamApi.UsersListJoinedTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersListJoinedTeamsRequest struct via the builder pattern


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

[**CollectionOfTeam**](CollectionOfTeam.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateJoinedTeams

> UsersUpdateJoinedTeams(ctx, userId, teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Update the navigation property joinedTeams in users



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
    teamId := "teamId_example" // string | key: id of team
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersTeamApi.UsersUpdateJoinedTeams(context.Background(), userId, teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersTeamApi.UsersUpdateJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateJoinedTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md) | New navigation property values | 

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

