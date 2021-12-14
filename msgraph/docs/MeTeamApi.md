# \MeTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeCreateJoinedTeams**](MeTeamApi.md#MeCreateJoinedTeams) | **Post** /me/joinedTeams | Create new navigation property to joinedTeams for me
[**MeDeleteJoinedTeams**](MeTeamApi.md#MeDeleteJoinedTeams) | **Delete** /me/joinedTeams/{team-id} | Delete navigation property joinedTeams for me
[**MeGetJoinedTeams**](MeTeamApi.md#MeGetJoinedTeams) | **Get** /me/joinedTeams/{team-id} | Get joinedTeams from me
[**MeListJoinedTeams**](MeTeamApi.md#MeListJoinedTeams) | **Get** /me/joinedTeams | Get joinedTeams from me
[**MeUpdateJoinedTeams**](MeTeamApi.md#MeUpdateJoinedTeams) | **Patch** /me/joinedTeams/{team-id} | Update the navigation property joinedTeams in me



## MeCreateJoinedTeams

> MicrosoftGraphTeam MeCreateJoinedTeams(ctx).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Create new navigation property to joinedTeams for me



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
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTeamApi.MeCreateJoinedTeams(context.Background()).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTeamApi.MeCreateJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeCreateJoinedTeams`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `MeTeamApi.MeCreateJoinedTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeCreateJoinedTeamsRequest struct via the builder pattern


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


## MeDeleteJoinedTeams

> MeDeleteJoinedTeams(ctx, teamId).IfMatch(ifMatch).Execute()

Delete navigation property joinedTeams for me



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
    teamId := "teamId_example" // string | key: id of team
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTeamApi.MeDeleteJoinedTeams(context.Background(), teamId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTeamApi.MeDeleteJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeDeleteJoinedTeamsRequest struct via the builder pattern


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


## MeGetJoinedTeams

> MicrosoftGraphTeam MeGetJoinedTeams(ctx, teamId).Select_(select_).Expand(expand).Execute()

Get joinedTeams from me



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
    teamId := "teamId_example" // string | key: id of team
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTeamApi.MeGetJoinedTeams(context.Background(), teamId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTeamApi.MeGetJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetJoinedTeams`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `MeTeamApi.MeGetJoinedTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeGetJoinedTeamsRequest struct via the builder pattern


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


## MeListJoinedTeams

> CollectionOfTeam MeListJoinedTeams(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get joinedTeams from me



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
    resp, r, err := api_client.MeTeamApi.MeListJoinedTeams(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTeamApi.MeListJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeListJoinedTeams`: CollectionOfTeam
    fmt.Fprintf(os.Stdout, "Response from `MeTeamApi.MeListJoinedTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeListJoinedTeamsRequest struct via the builder pattern


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


## MeUpdateJoinedTeams

> MeUpdateJoinedTeams(ctx, teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Update the navigation property joinedTeams in me



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
    teamId := "teamId_example" // string | key: id of team
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MeTeamApi.MeUpdateJoinedTeams(context.Background(), teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MeTeamApi.MeUpdateJoinedTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdateJoinedTeamsRequest struct via the builder pattern


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

