# \TeamsTeamsAsyncOperationApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamsCreateOperations**](TeamsTeamsAsyncOperationApi.md#TeamsCreateOperations) | **Post** /teams/{team-id}/operations | Create new navigation property to operations for teams
[**TeamsDeleteOperations**](TeamsTeamsAsyncOperationApi.md#TeamsDeleteOperations) | **Delete** /teams/{team-id}/operations/{teamsAsyncOperation-id} | Delete navigation property operations for teams
[**TeamsGetOperations**](TeamsTeamsAsyncOperationApi.md#TeamsGetOperations) | **Get** /teams/{team-id}/operations/{teamsAsyncOperation-id} | Get operations from teams
[**TeamsListOperations**](TeamsTeamsAsyncOperationApi.md#TeamsListOperations) | **Get** /teams/{team-id}/operations | Get operations from teams
[**TeamsUpdateOperations**](TeamsTeamsAsyncOperationApi.md#TeamsUpdateOperations) | **Patch** /teams/{team-id}/operations/{teamsAsyncOperation-id} | Update the navigation property operations in teams



## TeamsCreateOperations

> MicrosoftGraphTeamsAsyncOperation TeamsCreateOperations(ctx, teamId).MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation).Execute()

Create new navigation property to operations for teams



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
    microsoftGraphTeamsAsyncOperation := *openapiclient.NewMicrosoftGraphTeamsAsyncOperation() // MicrosoftGraphTeamsAsyncOperation | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAsyncOperationApi.TeamsCreateOperations(context.Background(), teamId).MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAsyncOperationApi.TeamsCreateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreateOperations`: MicrosoftGraphTeamsAsyncOperation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAsyncOperationApi.TeamsCreateOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphTeamsAsyncOperation** | [**MicrosoftGraphTeamsAsyncOperation**](MicrosoftGraphTeamsAsyncOperation.md) | New navigation property | 

### Return type

[**MicrosoftGraphTeamsAsyncOperation**](MicrosoftGraphTeamsAsyncOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDeleteOperations

> TeamsDeleteOperations(ctx, teamId, teamsAsyncOperationId).IfMatch(ifMatch).Execute()

Delete navigation property operations for teams



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
    teamsAsyncOperationId := "teamsAsyncOperationId_example" // string | key: id of teamsAsyncOperation
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAsyncOperationApi.TeamsDeleteOperations(context.Background(), teamId, teamsAsyncOperationId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAsyncOperationApi.TeamsDeleteOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAsyncOperationId** | **string** | key: id of teamsAsyncOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDeleteOperationsRequest struct via the builder pattern


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


## TeamsGetOperations

> MicrosoftGraphTeamsAsyncOperation TeamsGetOperations(ctx, teamId, teamsAsyncOperationId).Select_(select_).Expand(expand).Execute()

Get operations from teams



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
    teamsAsyncOperationId := "teamsAsyncOperationId_example" // string | key: id of teamsAsyncOperation
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAsyncOperationApi.TeamsGetOperations(context.Background(), teamId, teamsAsyncOperationId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAsyncOperationApi.TeamsGetOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsGetOperations`: MicrosoftGraphTeamsAsyncOperation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAsyncOperationApi.TeamsGetOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAsyncOperationId** | **string** | key: id of teamsAsyncOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsGetOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamsAsyncOperation**](MicrosoftGraphTeamsAsyncOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsListOperations

> CollectionOfTeamsAsyncOperation TeamsListOperations(ctx, teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get operations from teams



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
    resp, r, err := api_client.TeamsTeamsAsyncOperationApi.TeamsListOperations(context.Background(), teamId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAsyncOperationApi.TeamsListOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsListOperations`: CollectionOfTeamsAsyncOperation
    fmt.Fprintf(os.Stdout, "Response from `TeamsTeamsAsyncOperationApi.TeamsListOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsListOperationsRequest struct via the builder pattern


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

[**CollectionOfTeamsAsyncOperation**](CollectionOfTeamsAsyncOperation.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdateOperations

> TeamsUpdateOperations(ctx, teamId, teamsAsyncOperationId).MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation).Execute()

Update the navigation property operations in teams



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
    teamsAsyncOperationId := "teamsAsyncOperationId_example" // string | key: id of teamsAsyncOperation
    microsoftGraphTeamsAsyncOperation := *openapiclient.NewMicrosoftGraphTeamsAsyncOperation() // MicrosoftGraphTeamsAsyncOperation | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamsTeamsAsyncOperationApi.TeamsUpdateOperations(context.Background(), teamId, teamsAsyncOperationId).MicrosoftGraphTeamsAsyncOperation(microsoftGraphTeamsAsyncOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsTeamsAsyncOperationApi.TeamsUpdateOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | key: id of team | 
**teamsAsyncOperationId** | **string** | key: id of teamsAsyncOperation | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **microsoftGraphTeamsAsyncOperation** | [**MicrosoftGraphTeamsAsyncOperation**](MicrosoftGraphTeamsAsyncOperation.md) | New navigation property values | 

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

