# \GroupsTeamApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsDeleteTeam**](GroupsTeamApi.md#GroupsDeleteTeam) | **Delete** /groups/{group-id}/team | Delete navigation property team for groups
[**GroupsGetTeam**](GroupsTeamApi.md#GroupsGetTeam) | **Get** /groups/{group-id}/team | Get team from groups
[**GroupsUpdateTeam**](GroupsTeamApi.md#GroupsUpdateTeam) | **Patch** /groups/{group-id}/team | Update the navigation property team in groups



## GroupsDeleteTeam

> GroupsDeleteTeam(ctx, groupId).IfMatch(ifMatch).Execute()

Delete navigation property team for groups

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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsTeamApi.GroupsDeleteTeam(context.Background(), groupId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsTeamApi.GroupsDeleteTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsDeleteTeamRequest struct via the builder pattern


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


## GroupsGetTeam

> MicrosoftGraphTeam GroupsGetTeam(ctx, groupId).Select_(select_).Expand(expand).Execute()

Get team from groups

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsTeamApi.GroupsGetTeam(context.Background(), groupId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsTeamApi.GroupsGetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsGetTeam`: MicrosoftGraphTeam
    fmt.Fprintf(os.Stdout, "Response from `GroupsTeamApi.GroupsGetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetTeamRequest struct via the builder pattern


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


## GroupsUpdateTeam

> GroupsUpdateTeam(ctx, groupId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()

Update the navigation property team in groups

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
    microsoftGraphTeam := *openapiclient.NewMicrosoftGraphTeam() // MicrosoftGraphTeam | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupsTeamApi.GroupsUpdateTeam(context.Background(), groupId).MicrosoftGraphTeam(microsoftGraphTeam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsTeamApi.GroupsUpdateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | key: id of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsUpdateTeamRequest struct via the builder pattern


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

