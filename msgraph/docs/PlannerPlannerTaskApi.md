# \PlannerPlannerTaskApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerCreateTasks**](PlannerPlannerTaskApi.md#PlannerCreateTasks) | **Post** /planner/tasks | Create new navigation property to tasks for planner
[**PlannerDeleteTasks**](PlannerPlannerTaskApi.md#PlannerDeleteTasks) | **Delete** /planner/tasks/{plannerTask-id} | Delete navigation property tasks for planner
[**PlannerGetTasks**](PlannerPlannerTaskApi.md#PlannerGetTasks) | **Get** /planner/tasks/{plannerTask-id} | Get tasks from planner
[**PlannerListTasks**](PlannerPlannerTaskApi.md#PlannerListTasks) | **Get** /planner/tasks | Get tasks from planner
[**PlannerTasksDeleteAssignedToTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksDeleteAssignedToTaskBoardFormat) | **Delete** /planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for planner
[**PlannerTasksDeleteBucketTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksDeleteBucketTaskBoardFormat) | **Delete** /planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for planner
[**PlannerTasksDeleteDetails**](PlannerPlannerTaskApi.md#PlannerTasksDeleteDetails) | **Delete** /planner/tasks/{plannerTask-id}/details | Delete navigation property details for planner
[**PlannerTasksDeleteProgressTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksDeleteProgressTaskBoardFormat) | **Delete** /planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for planner
[**PlannerTasksGetAssignedToTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetAssignedToTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from planner
[**PlannerTasksGetBucketTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetBucketTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from planner
[**PlannerTasksGetDetails**](PlannerPlannerTaskApi.md#PlannerTasksGetDetails) | **Get** /planner/tasks/{plannerTask-id}/details | Get details from planner
[**PlannerTasksGetProgressTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksGetProgressTaskBoardFormat) | **Get** /planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from planner
[**PlannerTasksUpdateAssignedToTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateAssignedToTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in planner
[**PlannerTasksUpdateBucketTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateBucketTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in planner
[**PlannerTasksUpdateDetails**](PlannerPlannerTaskApi.md#PlannerTasksUpdateDetails) | **Patch** /planner/tasks/{plannerTask-id}/details | Update the navigation property details in planner
[**PlannerTasksUpdateProgressTaskBoardFormat**](PlannerPlannerTaskApi.md#PlannerTasksUpdateProgressTaskBoardFormat) | **Patch** /planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in planner
[**PlannerUpdateTasks**](PlannerPlannerTaskApi.md#PlannerUpdateTasks) | **Patch** /planner/tasks/{plannerTask-id} | Update the navigation property tasks in planner



## PlannerCreateTasks

> MicrosoftGraphPlannerTask PlannerCreateTasks(ctx).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for planner



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
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerCreateTasks(context.Background()).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerCreateTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerCreateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | New navigation property | 

### Return type

[**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerDeleteTasks

> PlannerDeleteTasks(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerDeleteTasks(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerDeleteTasksRequest struct via the builder pattern


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


## PlannerGetTasks

> MicrosoftGraphPlannerTask PlannerGetTasks(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerGetTasks(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerGetTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerListTasks

> CollectionOfPlannerTask PlannerListTasks(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from planner



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
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerListTasks(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerListTasksRequest struct via the builder pattern


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

[**CollectionOfPlannerTask**](CollectionOfPlannerTask.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksDeleteAssignedToTaskBoardFormat

> PlannerTasksDeleteAssignedToTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksDeleteAssignedToTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## PlannerTasksDeleteBucketTaskBoardFormat

> PlannerTasksDeleteBucketTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksDeleteBucketTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## PlannerTasksDeleteDetails

> PlannerTasksDeleteDetails(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksDeleteDetails(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksDeleteDetailsRequest struct via the builder pattern


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


## PlannerTasksDeleteProgressTaskBoardFormat

> PlannerTasksDeleteProgressTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksDeleteProgressTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## PlannerTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat PlannerTasksGetAssignedToTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksGetAssignedToTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat PlannerTasksGetBucketTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksGetBucketTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](MicrosoftGraphPlannerBucketTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetDetails

> MicrosoftGraphPlannerTaskDetails PlannerTasksGetDetails(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksGetDetails(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerTaskDetails**](MicrosoftGraphPlannerTaskDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat PlannerTasksGetProgressTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksGetProgressTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerTaskApi.PlannerTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](MicrosoftGraphPlannerProgressTaskBoardTaskFormat.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerTasksUpdateAssignedToTaskBoardFormat

> PlannerTasksUpdateAssignedToTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksUpdateAssignedToTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerAssignedToTaskBoardTaskFormat** | [**MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat**](MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat.md) | New navigation property values | 

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


## PlannerTasksUpdateBucketTaskBoardFormat

> PlannerTasksUpdateBucketTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksUpdateBucketTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerBucketTaskBoardTaskFormat** | [**MicrosoftGraphPlannerBucketTaskBoardTaskFormat**](MicrosoftGraphPlannerBucketTaskBoardTaskFormat.md) | New navigation property values | 

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


## PlannerTasksUpdateDetails

> PlannerTasksUpdateDetails(ctx, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksUpdateDetails(context.Background(), plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksUpdateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerTaskDetails** | [**MicrosoftGraphPlannerTaskDetails**](MicrosoftGraphPlannerTaskDetails.md) | New navigation property values | 

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


## PlannerTasksUpdateProgressTaskBoardFormat

> PlannerTasksUpdateProgressTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerTasksUpdateProgressTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerProgressTaskBoardTaskFormat** | [**MicrosoftGraphPlannerProgressTaskBoardTaskFormat**](MicrosoftGraphPlannerProgressTaskBoardTaskFormat.md) | New navigation property values | 

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


## PlannerUpdateTasks

> PlannerUpdateTasks(ctx, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in planner



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerTaskApi.PlannerUpdateTasks(context.Background(), plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerTaskApi.PlannerUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerTask** | [**MicrosoftGraphPlannerTask**](MicrosoftGraphPlannerTask.md) | New navigation property values | 

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

