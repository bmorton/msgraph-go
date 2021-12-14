# \TeamsTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsTeamCreateTeam**](TeamsTeamApi.md#TeamsTeamCreateTeam) | **Post** /teams | Add new entity to teams
[**TeamsTeamDeleteTeam**](TeamsTeamApi.md#TeamsTeamDeleteTeam) | **Delete** /teams/{team-id} | Delete entity from teams
[**TeamsTeamGetTeam**](TeamsTeamApi.md#TeamsTeamGetTeam) | **Get** /teams/{team-id} | Get entity from teams by key
[**TeamsTeamListTeam**](TeamsTeamApi.md#TeamsTeamListTeam) | **Get** /teams | Get entities from teams
[**TeamsTeamUpdateTeam**](TeamsTeamApi.md#TeamsTeamUpdateTeam) | **Patch** /teams/{team-id} | Update entity in teams



## TeamsTeamCreateTeam

> MicrosoftGraphTeam TeamsTeamCreateTeam(ctx).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Add new entity to teams

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
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New entity

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamApi.TeamsTeamCreateTeam(context.Background()).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamApi.TeamsTeamCreateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamCreateTeam`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamApi.TeamsTeamCreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md) | New entity | 

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


## TeamsTeamDeleteTeam

> TeamsTeamDeleteTeam(ctx, teamId).IfMatch(ifMatch).Execute()

Delete entity from teams

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
    resp, r, err := api_client.TeamsTeamApi.TeamsTeamDeleteTeam(context.Background(), teamId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamApi.TeamsTeamDeleteTeam``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsTeamDeleteTeamRequest struct via the builder pattern


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


## TeamsTeamGetTeam

> MicrosoftGraphTeam TeamsTeamGetTeam(ctx, teamId).Select_(select_).Expand(expand).Execute()

Get entity from teams by key

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
    resp, r, err := api_client.TeamsTeamApi.TeamsTeamGetTeam(context.Background(), teamId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamApi.TeamsTeamGetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamGetTeam`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamApi.TeamsTeamGetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamGetTeamRequest struct via the builder pattern


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


## TeamsTeamListTeam

> CollectionOfTeam TeamsTeamListTeam(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get entities from teams

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
    resp, r, err := api_client.TeamsTeamApi.TeamsTeamListTeam(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamApi.TeamsTeamListTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsTeamListTeam`: CollectionOfTeam
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamApi.TeamsTeamListTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamsTeamListTeamRequest struct via the builder pattern


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


## TeamsTeamUpdateTeam

> TeamsTeamUpdateTeam(ctx, teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Update entity in teams

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
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamApi.TeamsTeamUpdateTeam(context.Background(), teamId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamApi.TeamsTeamUpdateTeam``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTeamsTeamUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeam** | [**MicrosoftGraphTeam**](MicrosoftGraphTeam.md) | New property values | 

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

