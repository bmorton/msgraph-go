# \PlannerPlannerBucketApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlannerBucketsCreateTasks**](PlannerPlannerBucketApi.md#PlannerBucketsCreateTasks) | **Post** /planner/buckets/{plannerBucket-id}/tasks | Create new navigation property to tasks for planner
[**PlannerBucketsDeleteTasks**](PlannerPlannerBucketApi.md#PlannerBucketsDeleteTasks) | **Delete** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Delete navigation property tasks for planner
[**PlannerBucketsGetTasks**](PlannerPlannerBucketApi.md#PlannerBucketsGetTasks) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Get tasks from planner
[**PlannerBucketsListTasks**](PlannerPlannerBucketApi.md#PlannerBucketsListTasks) | **Get** /planner/buckets/{plannerBucket-id}/tasks | Get tasks from planner
[**PlannerBucketsTasksDeleteAssignedToTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksDeleteAssignedToTaskBoardFormat) | **Delete** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for planner
[**PlannerBucketsTasksDeleteBucketTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksDeleteBucketTaskBoardFormat) | **Delete** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for planner
[**PlannerBucketsTasksDeleteDetails**](PlannerPlannerBucketApi.md#PlannerBucketsTasksDeleteDetails) | **Delete** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Delete navigation property details for planner
[**PlannerBucketsTasksDeleteProgressTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksDeleteProgressTaskBoardFormat) | **Delete** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for planner
[**PlannerBucketsTasksGetAssignedToTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksGetAssignedToTaskBoardFormat) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from planner
[**PlannerBucketsTasksGetBucketTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksGetBucketTaskBoardFormat) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from planner
[**PlannerBucketsTasksGetDetails**](PlannerPlannerBucketApi.md#PlannerBucketsTasksGetDetails) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Get details from planner
[**PlannerBucketsTasksGetProgressTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksGetProgressTaskBoardFormat) | **Get** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from planner
[**PlannerBucketsTasksUpdateAssignedToTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksUpdateAssignedToTaskBoardFormat) | **Patch** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in planner
[**PlannerBucketsTasksUpdateBucketTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksUpdateBucketTaskBoardFormat) | **Patch** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in planner
[**PlannerBucketsTasksUpdateDetails**](PlannerPlannerBucketApi.md#PlannerBucketsTasksUpdateDetails) | **Patch** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Update the navigation property details in planner
[**PlannerBucketsTasksUpdateProgressTaskBoardFormat**](PlannerPlannerBucketApi.md#PlannerBucketsTasksUpdateProgressTaskBoardFormat) | **Patch** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in planner
[**PlannerBucketsUpdateTasks**](PlannerPlannerBucketApi.md#PlannerBucketsUpdateTasks) | **Patch** /planner/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Update the navigation property tasks in planner
[**PlannerCreateBuckets**](PlannerPlannerBucketApi.md#PlannerCreateBuckets) | **Post** /planner/buckets | Create new navigation property to buckets for planner
[**PlannerDeleteBuckets**](PlannerPlannerBucketApi.md#PlannerDeleteBuckets) | **Delete** /planner/buckets/{plannerBucket-id} | Delete navigation property buckets for planner
[**PlannerGetBuckets**](PlannerPlannerBucketApi.md#PlannerGetBuckets) | **Get** /planner/buckets/{plannerBucket-id} | Get buckets from planner
[**PlannerListBuckets**](PlannerPlannerBucketApi.md#PlannerListBuckets) | **Get** /planner/buckets | Get buckets from planner
[**PlannerUpdateBuckets**](PlannerPlannerBucketApi.md#PlannerUpdateBuckets) | **Patch** /planner/buckets/{plannerBucket-id} | Update the navigation property buckets in planner



## PlannerBucketsCreateTasks

> MicrosoftGraphPlannerTask PlannerBucketsCreateTasks(ctx, plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsCreateTasks(context.Background(), plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsCreateTasksRequest struct via the builder pattern


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


## PlannerBucketsDeleteTasks

> PlannerBucketsDeleteTasks(ctx, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsDeleteTasks(context.Background(), plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsDeleteTasksRequest struct via the builder pattern


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


## PlannerBucketsGetTasks

> MicrosoftGraphPlannerTask PlannerBucketsGetTasks(ctx, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsGetTasks(context.Background(), plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsGetTasksRequest struct via the builder pattern


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


## PlannerBucketsListTasks

> CollectionOfPlannerTask PlannerBucketsListTasks(ctx, plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
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
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsListTasks(context.Background(), plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsListTasksRequest struct via the builder pattern


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


## PlannerBucketsTasksDeleteAssignedToTaskBoardFormat

> PlannerBucketsTasksDeleteAssignedToTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksDeleteAssignedToTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksDeleteBucketTaskBoardFormat

> PlannerBucketsTasksDeleteBucketTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksDeleteBucketTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksDeleteDetails

> PlannerBucketsTasksDeleteDetails(ctx, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksDeleteDetails(context.Background(), plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksDeleteDetailsRequest struct via the builder pattern


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


## PlannerBucketsTasksDeleteProgressTaskBoardFormat

> PlannerBucketsTasksDeleteProgressTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksDeleteProgressTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat PlannerBucketsTasksGetAssignedToTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksGetAssignedToTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat PlannerBucketsTasksGetBucketTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksGetBucketTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksGetDetails

> MicrosoftGraphPlannerTaskDetails PlannerBucketsTasksGetDetails(ctx, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksGetDetails(context.Background(), plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksGetDetailsRequest struct via the builder pattern


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


## PlannerBucketsTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat PlannerBucketsTasksGetProgressTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksGetProgressTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerBucketsTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerBucketsTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksUpdateAssignedToTaskBoardFormat

> PlannerBucketsTasksUpdateAssignedToTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksUpdateAssignedToTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksUpdateBucketTaskBoardFormat

> PlannerBucketsTasksUpdateBucketTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksUpdateBucketTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsTasksUpdateDetails

> PlannerBucketsTasksUpdateDetails(ctx, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksUpdateDetails(context.Background(), plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksUpdateDetailsRequest struct via the builder pattern


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


## PlannerBucketsTasksUpdateProgressTaskBoardFormat

> PlannerBucketsTasksUpdateProgressTaskBoardFormat(ctx, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsTasksUpdateProgressTaskBoardFormat(context.Background(), plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## PlannerBucketsUpdateTasks

> PlannerBucketsUpdateTasks(ctx, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerBucketsUpdateTasks(context.Background(), plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerBucketsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerBucketsUpdateTasksRequest struct via the builder pattern


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


## PlannerCreateBuckets

> MicrosoftGraphPlannerBucket PlannerCreateBuckets(ctx).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Create new navigation property to buckets for planner



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
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerCreateBuckets(context.Background()).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerCreateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerCreateBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerCreateBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerCreateBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlannerBucket** | [**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | New navigation property | 

### Return type

[**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerDeleteBuckets

> PlannerDeleteBuckets(ctx, plannerBucketId).IfMatch(ifMatch).Execute()

Delete navigation property buckets for planner



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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerDeleteBuckets(context.Background(), plannerBucketId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerDeleteBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerDeleteBucketsRequest struct via the builder pattern


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


## PlannerGetBuckets

> MicrosoftGraphPlannerBucket PlannerGetBuckets(ctx, plannerBucketId).Select_(select_).Expand(expand).Execute()

Get buckets from planner



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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerGetBuckets(context.Background(), plannerBucketId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerGetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerGetBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerGetBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerGetBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerListBuckets

> CollectionOfPlannerBucket PlannerListBuckets(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get buckets from planner



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
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerListBuckets(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlannerListBuckets`: CollectionOfPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `PlannerPlannerBucketApi.PlannerListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlannerListBucketsRequest struct via the builder pattern


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

[**CollectionOfPlannerBucket**](CollectionOfPlannerBucket.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlannerUpdateBuckets

> PlannerUpdateBuckets(ctx, plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Update the navigation property buckets in planner



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
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlannerPlannerBucketApi.PlannerUpdateBuckets(context.Background(), plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlannerPlannerBucketApi.PlannerUpdateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlannerUpdateBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerBucket** | [**MicrosoftGraphPlannerBucket**](MicrosoftGraphPlannerBucket.md) | New navigation property values | 

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

