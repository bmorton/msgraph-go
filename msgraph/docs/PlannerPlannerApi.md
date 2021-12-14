# \PlannerPlannerApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerPlannerGetPlanner**](PlannerPlannerApi.md#PlannerPlannerGetPlanner) | **Get** /planner | Get planner
[**PlannerPlannerUpdatePlanner**](PlannerPlannerApi.md#PlannerPlannerUpdatePlanner) | **Patch** /planner | Update planner



## PlannerPlannerGetPlanner

> MicrosoftGraphPlanner PlannerPlannerGetPlanner(ctx).Select_(select_).Expand(expand).Execute()

Get planner

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
    resp, r, err := api_client.PlannerPlannerApi.PlannerPlannerGetPlanner(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerApi.PlannerPlannerGetPlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerPlannerGetPlanner`: MicrosoftGraphPlanner
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerApi.PlannerPlannerGetPlanner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerPlannerGetPlannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlanner**](MicrosoftGraphPlanner.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerPlannerUpdatePlanner

> PlannerPlannerUpdatePlanner(ctx).MicrosoftGraphPlanner(microsoftGraphPlanner).Execute()

Update planner

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
    microsoftGraphPlanner := *openapiclient.NewMicrosoftGraphPlanner() // MicrosoftGraphPlanner | New property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerApi.PlannerPlannerUpdatePlanner(context.Background()).MicrosoftGraphPlanner(microsoftGraphPlanner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerApi.PlannerPlannerUpdatePlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerPlannerUpdatePlannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlanner** | [**MicrosoftGraphPlanner**](MicrosoftGraphPlanner.md) | New property values | 

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

