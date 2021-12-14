# \TeamworkTeamworkApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamworkTeamworkGetTeamwork**](TeamworkTeamworkApi.md#TeamworkTeamworkGetTeamwork) | **Get** /teamwork | Get teamwork
[**TeamworkTeamworkUpdateTeamwork**](TeamworkTeamworkApi.md#TeamworkTeamworkUpdateTeamwork) | **Patch** /teamwork | Update teamwork



## TeamworkTeamworkGetTeamwork

> MicrosoftGraphTeamwork TeamworkTeamworkGetTeamwork(ctx).Select_(select_).Expand(expand).Execute()

Get teamwork

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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkTeamworkApi.TeamworkTeamworkGetTeamwork(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkTeamworkApi.TeamworkTeamworkGetTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamworkTeamworkGetTeamwork`: MicrosoftGraphTeamwork
    fmt.Fprintf(os.Stdout, "Response from `TeamworkTeamworkApi.TeamworkTeamworkGetTeamwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkTeamworkGetTeamworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphTeamwork**](MicrosoftGraphTeamwork.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamworkTeamworkUpdateTeamwork

> TeamworkTeamworkUpdateTeamwork(ctx).MicrosoftGraphTeamwork(microsoftGraphTeamwork).Execute()

Update teamwork

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
    microsoftGraphTeamwork := *openapiclient.NewMicrosoftGraphTeamwork() // MicrosoftGraphTeamwork | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TeamworkTeamworkApi.TeamworkTeamworkUpdateTeamwork(context.Background()).MicrosoftGraphTeamwork(microsoftGraphTeamwork).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamworkTeamworkApi.TeamworkTeamworkUpdateTeamwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTeamworkTeamworkUpdateTeamworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphTeamwork** | [**MicrosoftGraphTeamwork**](MicrosoftGraphTeamwork.md) | New property values | 

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

