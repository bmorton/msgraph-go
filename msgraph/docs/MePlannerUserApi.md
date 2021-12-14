# \MePlannerUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeDeletePlanner**](MePlannerUserApi.md#MeDeletePlanner) | **Delete** /me/planner | Delete navigation property planner for me
[**MeGetPlanner**](MePlannerUserApi.md#MeGetPlanner) | **Get** /me/planner | Get planner from me
[**MePlannerCreatePlans**](MePlannerUserApi.md#MePlannerCreatePlans) | **Post** /me/planner/plans | Create new navigation property to plans for me
[**MePlannerCreateTasks**](MePlannerUserApi.md#MePlannerCreateTasks) | **Post** /me/planner/tasks | Create new navigation property to tasks for me
[**MePlannerDeletePlans**](MePlannerUserApi.md#MePlannerDeletePlans) | **Delete** /me/planner/plans/{plannerPlan-id} | Delete navigation property plans for me
[**MePlannerDeleteTasks**](MePlannerUserApi.md#MePlannerDeleteTasks) | **Delete** /me/planner/tasks/{plannerTask-id} | Delete navigation property tasks for me
[**MePlannerGetPlans**](MePlannerUserApi.md#MePlannerGetPlans) | **Get** /me/planner/plans/{plannerPlan-id} | Get plans from me
[**MePlannerGetTasks**](MePlannerUserApi.md#MePlannerGetTasks) | **Get** /me/planner/tasks/{plannerTask-id} | Get tasks from me
[**MePlannerListPlans**](MePlannerUserApi.md#MePlannerListPlans) | **Get** /me/planner/plans | Get plans from me
[**MePlannerListTasks**](MePlannerUserApi.md#MePlannerListTasks) | **Get** /me/planner/tasks | Get tasks from me
[**MePlannerPlansBucketsCreateTasks**](MePlannerUserApi.md#MePlannerPlansBucketsCreateTasks) | **Post** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks | Create new navigation property to tasks for me
[**MePlannerPlansBucketsDeleteTasks**](MePlannerUserApi.md#MePlannerPlansBucketsDeleteTasks) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Delete navigation property tasks for me
[**MePlannerPlansBucketsGetTasks**](MePlannerUserApi.md#MePlannerPlansBucketsGetTasks) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Get tasks from me
[**MePlannerPlansBucketsListTasks**](MePlannerUserApi.md#MePlannerPlansBucketsListTasks) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks | Get tasks from me
[**MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for me
[**MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for me
[**MePlannerPlansBucketsTasksDeleteDetails**](MePlannerUserApi.md#MePlannerPlansBucketsTasksDeleteDetails) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Delete navigation property details for me
[**MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for me
[**MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from me
[**MePlannerPlansBucketsTasksGetBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksGetBucketTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from me
[**MePlannerPlansBucketsTasksGetDetails**](MePlannerUserApi.md#MePlannerPlansBucketsTasksGetDetails) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Get details from me
[**MePlannerPlansBucketsTasksGetProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksGetProgressTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from me
[**MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in me
[**MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in me
[**MePlannerPlansBucketsTasksUpdateDetails**](MePlannerUserApi.md#MePlannerPlansBucketsTasksUpdateDetails) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Update the navigation property details in me
[**MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in me
[**MePlannerPlansBucketsUpdateTasks**](MePlannerUserApi.md#MePlannerPlansBucketsUpdateTasks) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Update the navigation property tasks in me
[**MePlannerPlansCreateBuckets**](MePlannerUserApi.md#MePlannerPlansCreateBuckets) | **Post** /me/planner/plans/{plannerPlan-id}/buckets | Create new navigation property to buckets for me
[**MePlannerPlansCreateTasks**](MePlannerUserApi.md#MePlannerPlansCreateTasks) | **Post** /me/planner/plans/{plannerPlan-id}/tasks | Create new navigation property to tasks for me
[**MePlannerPlansDeleteBuckets**](MePlannerUserApi.md#MePlannerPlansDeleteBuckets) | **Delete** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Delete navigation property buckets for me
[**MePlannerPlansDeleteDetails**](MePlannerUserApi.md#MePlannerPlansDeleteDetails) | **Delete** /me/planner/plans/{plannerPlan-id}/details | Delete navigation property details for me
[**MePlannerPlansDeleteTasks**](MePlannerUserApi.md#MePlannerPlansDeleteTasks) | **Delete** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Delete navigation property tasks for me
[**MePlannerPlansGetBuckets**](MePlannerUserApi.md#MePlannerPlansGetBuckets) | **Get** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Get buckets from me
[**MePlannerPlansGetDetails**](MePlannerUserApi.md#MePlannerPlansGetDetails) | **Get** /me/planner/plans/{plannerPlan-id}/details | Get details from me
[**MePlannerPlansGetTasks**](MePlannerUserApi.md#MePlannerPlansGetTasks) | **Get** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Get tasks from me
[**MePlannerPlansListBuckets**](MePlannerUserApi.md#MePlannerPlansListBuckets) | **Get** /me/planner/plans/{plannerPlan-id}/buckets | Get buckets from me
[**MePlannerPlansListTasks**](MePlannerUserApi.md#MePlannerPlansListTasks) | **Get** /me/planner/plans/{plannerPlan-id}/tasks | Get tasks from me
[**MePlannerPlansTasksDeleteAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksDeleteAssignedToTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for me
[**MePlannerPlansTasksDeleteBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksDeleteBucketTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for me
[**MePlannerPlansTasksDeleteDetails**](MePlannerUserApi.md#MePlannerPlansTasksDeleteDetails) | **Delete** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Delete navigation property details for me
[**MePlannerPlansTasksDeleteProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksDeleteProgressTaskBoardFormat) | **Delete** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for me
[**MePlannerPlansTasksGetAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksGetAssignedToTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from me
[**MePlannerPlansTasksGetBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksGetBucketTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from me
[**MePlannerPlansTasksGetDetails**](MePlannerUserApi.md#MePlannerPlansTasksGetDetails) | **Get** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Get details from me
[**MePlannerPlansTasksGetProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksGetProgressTaskBoardFormat) | **Get** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from me
[**MePlannerPlansTasksUpdateAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksUpdateAssignedToTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in me
[**MePlannerPlansTasksUpdateBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksUpdateBucketTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in me
[**MePlannerPlansTasksUpdateDetails**](MePlannerUserApi.md#MePlannerPlansTasksUpdateDetails) | **Patch** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Update the navigation property details in me
[**MePlannerPlansTasksUpdateProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerPlansTasksUpdateProgressTaskBoardFormat) | **Patch** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in me
[**MePlannerPlansUpdateBuckets**](MePlannerUserApi.md#MePlannerPlansUpdateBuckets) | **Patch** /me/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Update the navigation property buckets in me
[**MePlannerPlansUpdateDetails**](MePlannerUserApi.md#MePlannerPlansUpdateDetails) | **Patch** /me/planner/plans/{plannerPlan-id}/details | Update the navigation property details in me
[**MePlannerPlansUpdateTasks**](MePlannerUserApi.md#MePlannerPlansUpdateTasks) | **Patch** /me/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Update the navigation property tasks in me
[**MePlannerTasksDeleteAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksDeleteAssignedToTaskBoardFormat) | **Delete** /me/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for me
[**MePlannerTasksDeleteBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksDeleteBucketTaskBoardFormat) | **Delete** /me/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for me
[**MePlannerTasksDeleteDetails**](MePlannerUserApi.md#MePlannerTasksDeleteDetails) | **Delete** /me/planner/tasks/{plannerTask-id}/details | Delete navigation property details for me
[**MePlannerTasksDeleteProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksDeleteProgressTaskBoardFormat) | **Delete** /me/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for me
[**MePlannerTasksGetAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksGetAssignedToTaskBoardFormat) | **Get** /me/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from me
[**MePlannerTasksGetBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksGetBucketTaskBoardFormat) | **Get** /me/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from me
[**MePlannerTasksGetDetails**](MePlannerUserApi.md#MePlannerTasksGetDetails) | **Get** /me/planner/tasks/{plannerTask-id}/details | Get details from me
[**MePlannerTasksGetProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksGetProgressTaskBoardFormat) | **Get** /me/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from me
[**MePlannerTasksUpdateAssignedToTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksUpdateAssignedToTaskBoardFormat) | **Patch** /me/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in me
[**MePlannerTasksUpdateBucketTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksUpdateBucketTaskBoardFormat) | **Patch** /me/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in me
[**MePlannerTasksUpdateDetails**](MePlannerUserApi.md#MePlannerTasksUpdateDetails) | **Patch** /me/planner/tasks/{plannerTask-id}/details | Update the navigation property details in me
[**MePlannerTasksUpdateProgressTaskBoardFormat**](MePlannerUserApi.md#MePlannerTasksUpdateProgressTaskBoardFormat) | **Patch** /me/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in me
[**MePlannerUpdatePlans**](MePlannerUserApi.md#MePlannerUpdatePlans) | **Patch** /me/planner/plans/{plannerPlan-id} | Update the navigation property plans in me
[**MePlannerUpdateTasks**](MePlannerUserApi.md#MePlannerUpdateTasks) | **Patch** /me/planner/tasks/{plannerTask-id} | Update the navigation property tasks in me
[**MeUpdatePlanner**](MePlannerUserApi.md#MeUpdatePlanner) | **Patch** /me/planner | Update the navigation property planner in me



## MeDeletePlanner

> MeDeletePlanner(ctx).IfMatch(ifMatch).Execute()

Delete navigation property planner for me



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MeDeletePlanner(context.Background()).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MeDeletePlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeDeletePlannerRequest struct via the builder pattern


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


## MeGetPlanner

> MicrosoftGraphPlannerUser MeGetPlanner(ctx).Select_(select_).Expand(expand).Execute()

Get planner from me



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
    resp, r, err := api_client.MePlannerUserApi.MeGetPlanner(context.Background()).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MeGetPlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeGetPlanner`: MicrosoftGraphPlannerUser
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MeGetPlanner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeGetPlannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerUser**](MicrosoftGraphPlannerUser.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MePlannerCreatePlans

> MicrosoftGraphPlannerPlan MePlannerCreatePlans(ctx).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()

Create new navigation property to plans for me



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
    microsoftGraphPlannerPlan := *openapiclient.NewMicrosoftGraphPlannerPlan() // MicrosoftGraphPlannerPlan | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerCreatePlans(context.Background()).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerCreatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerCreatePlans`: MicrosoftGraphPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerCreatePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerCreatePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlannerPlan** | [**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | New navigation property | 

### Return type

[**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MePlannerCreateTasks

> MicrosoftGraphPlannerTask MePlannerCreateTasks(ctx).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerCreateTasks(context.Background()).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerCreateTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerCreateTasksRequest struct via the builder pattern


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


## MePlannerDeletePlans

> MePlannerDeletePlans(ctx, plannerPlanId).IfMatch(ifMatch).Execute()

Delete navigation property plans for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerDeletePlans(context.Background(), plannerPlanId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerDeletePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerDeletePlansRequest struct via the builder pattern


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


## MePlannerDeleteTasks

> MePlannerDeleteTasks(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerDeleteTasks(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerDeleteTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerDeleteTasksRequest struct via the builder pattern


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


## MePlannerGetPlans

> MicrosoftGraphPlannerPlan MePlannerGetPlans(ctx, plannerPlanId).Select_(select_).Expand(expand).Execute()

Get plans from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerGetPlans(context.Background(), plannerPlanId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerGetPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerGetPlans`: MicrosoftGraphPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerGetPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MePlannerGetTasks

> MicrosoftGraphPlannerTask MePlannerGetTasks(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerGetTasks(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerGetTasksRequest struct via the builder pattern


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


## MePlannerListPlans

> CollectionOfPlannerPlan MePlannerListPlans(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get plans from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerListPlans(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerListPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerListPlans`: CollectionOfPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerListPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerListPlansRequest struct via the builder pattern


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

[**CollectionOfPlannerPlan**](CollectionOfPlannerPlan.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MePlannerListTasks

> CollectionOfPlannerTask MePlannerListTasks(ctx).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerListTasks(context.Background()).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerListTasksRequest struct via the builder pattern


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


## MePlannerPlansBucketsCreateTasks

> MicrosoftGraphPlannerTask MePlannerPlansBucketsCreateTasks(ctx, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsCreateTasks(context.Background(), plannerPlanId, plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsCreateTasksRequest struct via the builder pattern


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


## MePlannerPlansBucketsDeleteTasks

> MePlannerPlansBucketsDeleteTasks(ctx, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsDeleteTasks(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsDeleteTasksRequest struct via the builder pattern


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


## MePlannerPlansBucketsGetTasks

> MicrosoftGraphPlannerTask MePlannerPlansBucketsGetTasks(ctx, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsGetTasks(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsGetTasksRequest struct via the builder pattern


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


## MePlannerPlansBucketsListTasks

> CollectionOfPlannerTask MePlannerPlansBucketsListTasks(ctx, plannerPlanId, plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
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
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsListTasks(context.Background(), plannerPlanId, plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsListTasksRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat

> MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat

> MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksDeleteDetails

> MePlannerPlansBucketsTasksDeleteDetails(ctx, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksDeleteDetails(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksDeleteDetailsRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat

> MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat MePlannerPlansBucketsTasksGetBucketTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksGetBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksGetDetails

> MicrosoftGraphPlannerTaskDetails MePlannerPlansBucketsTasksGetDetails(ctx, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksGetDetails(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksGetDetailsRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat MePlannerPlansBucketsTasksGetProgressTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksGetProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansBucketsTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansBucketsTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat

> MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat

> MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksUpdateDetails

> MePlannerPlansBucketsTasksUpdateDetails(ctx, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksUpdateDetails(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksUpdateDetailsRequest struct via the builder pattern


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


## MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat

> MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat(ctx, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansBucketsUpdateTasks

> MePlannerPlansBucketsUpdateTasks(ctx, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansBucketsUpdateTasks(context.Background(), plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansBucketsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansBucketsUpdateTasksRequest struct via the builder pattern


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


## MePlannerPlansCreateBuckets

> MicrosoftGraphPlannerBucket MePlannerPlansCreateBuckets(ctx, plannerPlanId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Create new navigation property to buckets for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansCreateBuckets(context.Background(), plannerPlanId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansCreateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansCreateBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansCreateBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansCreateBucketsRequest struct via the builder pattern


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


## MePlannerPlansCreateTasks

> MicrosoftGraphPlannerTask MePlannerPlansCreateTasks(ctx, plannerPlanId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansCreateTasks(context.Background(), plannerPlanId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansCreateTasksRequest struct via the builder pattern


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


## MePlannerPlansDeleteBuckets

> MePlannerPlansDeleteBuckets(ctx, plannerPlanId, plannerBucketId).IfMatch(ifMatch).Execute()

Delete navigation property buckets for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansDeleteBuckets(context.Background(), plannerPlanId, plannerBucketId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansDeleteBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansDeleteBucketsRequest struct via the builder pattern


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


## MePlannerPlansDeleteDetails

> MePlannerPlansDeleteDetails(ctx, plannerPlanId).IfMatch(ifMatch).Execute()

Delete navigation property details for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansDeleteDetails(context.Background(), plannerPlanId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansDeleteDetailsRequest struct via the builder pattern


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


## MePlannerPlansDeleteTasks

> MePlannerPlansDeleteTasks(ctx, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansDeleteTasks(context.Background(), plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansDeleteTasksRequest struct via the builder pattern


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


## MePlannerPlansGetBuckets

> MicrosoftGraphPlannerBucket MePlannerPlansGetBuckets(ctx, plannerPlanId, plannerBucketId).Select_(select_).Expand(expand).Execute()

Get buckets from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansGetBuckets(context.Background(), plannerPlanId, plannerBucketId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansGetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansGetBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansGetBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansGetBucketsRequest struct via the builder pattern


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


## MePlannerPlansGetDetails

> MicrosoftGraphPlannerPlanDetails MePlannerPlansGetDetails(ctx, plannerPlanId).Select_(select_).Expand(expand).Execute()

Get details from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansGetDetails(context.Background(), plannerPlanId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansGetDetails`: MicrosoftGraphPlannerPlanDetails
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **[]string** | Select properties to be returned | 
 **expand** | **[]string** | Expand related entities | 

### Return type

[**MicrosoftGraphPlannerPlanDetails**](MicrosoftGraphPlannerPlanDetails.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MePlannerPlansGetTasks

> MicrosoftGraphPlannerTask MePlannerPlansGetTasks(ctx, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansGetTasks(context.Background(), plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansGetTasksRequest struct via the builder pattern


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


## MePlannerPlansListBuckets

> CollectionOfPlannerBucket MePlannerPlansListBuckets(ctx, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get buckets from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
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
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansListBuckets(context.Background(), plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansListBuckets`: CollectionOfPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansListBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansListBucketsRequest struct via the builder pattern


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


## MePlannerPlansListTasks

> CollectionOfPlannerTask MePlannerPlansListTasks(ctx, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
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
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansListTasks(context.Background(), plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansListTasksRequest struct via the builder pattern


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


## MePlannerPlansTasksDeleteAssignedToTaskBoardFormat

> MePlannerPlansTasksDeleteAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksDeleteAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksDeleteBucketTaskBoardFormat

> MePlannerPlansTasksDeleteBucketTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksDeleteBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksDeleteDetails

> MePlannerPlansTasksDeleteDetails(ctx, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksDeleteDetails(context.Background(), plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksDeleteDetailsRequest struct via the builder pattern


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


## MePlannerPlansTasksDeleteProgressTaskBoardFormat

> MePlannerPlansTasksDeleteProgressTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksDeleteProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat MePlannerPlansTasksGetAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksGetAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat MePlannerPlansTasksGetBucketTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksGetBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksGetDetails

> MicrosoftGraphPlannerTaskDetails MePlannerPlansTasksGetDetails(ctx, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksGetDetails(context.Background(), plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksGetDetailsRequest struct via the builder pattern


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


## MePlannerPlansTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat MePlannerPlansTasksGetProgressTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksGetProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerPlansTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerPlansTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksUpdateAssignedToTaskBoardFormat

> MePlannerPlansTasksUpdateAssignedToTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksUpdateAssignedToTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksUpdateBucketTaskBoardFormat

> MePlannerPlansTasksUpdateBucketTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksUpdateBucketTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansTasksUpdateDetails

> MePlannerPlansTasksUpdateDetails(ctx, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksUpdateDetails(context.Background(), plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksUpdateDetailsRequest struct via the builder pattern


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


## MePlannerPlansTasksUpdateProgressTaskBoardFormat

> MePlannerPlansTasksUpdateProgressTaskBoardFormat(ctx, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansTasksUpdateProgressTaskBoardFormat(context.Background(), plannerPlanId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerPlansUpdateBuckets

> MePlannerPlansUpdateBuckets(ctx, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Update the navigation property buckets in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansUpdateBuckets(context.Background(), plannerPlanId, plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansUpdateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansUpdateBucketsRequest struct via the builder pattern


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


## MePlannerPlansUpdateDetails

> MePlannerPlansUpdateDetails(ctx, plannerPlanId).MicrosoftGraphPlannerPlanDetails(microsoftGraphPlannerPlanDetails).Execute()

Update the navigation property details in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerPlanDetails := *openapiclient.NewMicrosoftGraphPlannerPlanDetails() // MicrosoftGraphPlannerPlanDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansUpdateDetails(context.Background(), plannerPlanId).MicrosoftGraphPlannerPlanDetails(microsoftGraphPlannerPlanDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansUpdateDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerPlanDetails** | [**MicrosoftGraphPlannerPlanDetails**](MicrosoftGraphPlannerPlanDetails.md) | New navigation property values | 

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


## MePlannerPlansUpdateTasks

> MePlannerPlansUpdateTasks(ctx, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerPlansUpdateTasks(context.Background(), plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerPlansUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerPlansUpdateTasksRequest struct via the builder pattern


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


## MePlannerTasksDeleteAssignedToTaskBoardFormat

> MePlannerTasksDeleteAssignedToTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksDeleteAssignedToTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksDeleteBucketTaskBoardFormat

> MePlannerTasksDeleteBucketTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksDeleteBucketTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksDeleteBucketTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksDeleteDetails

> MePlannerTasksDeleteDetails(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksDeleteDetails(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksDeleteDetails``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksDeleteDetailsRequest struct via the builder pattern


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


## MePlannerTasksDeleteProgressTaskBoardFormat

> MePlannerTasksDeleteProgressTaskBoardFormat(ctx, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksDeleteProgressTaskBoardFormat(context.Background(), plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksDeleteProgressTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat MePlannerTasksGetAssignedToTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksGetAssignedToTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat MePlannerTasksGetBucketTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksGetBucketTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksGetDetails

> MicrosoftGraphPlannerTaskDetails MePlannerTasksGetDetails(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksGetDetails(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerTasksGetDetailsRequest struct via the builder pattern


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


## MePlannerTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat MePlannerTasksGetProgressTaskBoardFormat(ctx, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksGetProgressTaskBoardFormat(context.Background(), plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MePlannerTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `MePlannerUserApi.MePlannerTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksUpdateAssignedToTaskBoardFormat

> MePlannerTasksUpdateAssignedToTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksUpdateAssignedToTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksUpdateBucketTaskBoardFormat

> MePlannerTasksUpdateBucketTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksUpdateBucketTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksUpdateBucketTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerTasksUpdateDetails

> MePlannerTasksUpdateDetails(ctx, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksUpdateDetails(context.Background(), plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksUpdateDetails``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksUpdateDetailsRequest struct via the builder pattern


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


## MePlannerTasksUpdateProgressTaskBoardFormat

> MePlannerTasksUpdateProgressTaskBoardFormat(ctx, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerTasksUpdateProgressTaskBoardFormat(context.Background(), plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerTasksUpdateProgressTaskBoardFormat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## MePlannerUpdatePlans

> MePlannerUpdatePlans(ctx, plannerPlanId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()

Update the navigation property plans in me



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerPlan := *openapiclient.NewMicrosoftGraphPlannerPlan() // MicrosoftGraphPlannerPlan | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MePlannerUpdatePlans(context.Background(), plannerPlanId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerUpdatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiMePlannerUpdatePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **microsoftGraphPlannerPlan** | [**MicrosoftGraphPlannerPlan**](MicrosoftGraphPlannerPlan.md) | New navigation property values | 

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


## MePlannerUpdateTasks

> MePlannerUpdateTasks(ctx, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in me



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
    resp, r, err := api_client.MePlannerUserApi.MePlannerUpdateTasks(context.Background(), plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MePlannerUpdateTasks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiMePlannerUpdateTasksRequest struct via the builder pattern


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


## MeUpdatePlanner

> MeUpdatePlanner(ctx).MicrosoftGraphPlannerUser(microsoftGraphPlannerUser).Execute()

Update the navigation property planner in me



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
    microsoftGraphPlannerUser := *openapiclient.NewMicrosoftGraphPlannerUser() // MicrosoftGraphPlannerUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MePlannerUserApi.MeUpdatePlanner(context.Background()).MicrosoftGraphPlannerUser(microsoftGraphPlannerUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MePlannerUserApi.MeUpdatePlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeUpdatePlannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **microsoftGraphPlannerUser** | [**MicrosoftGraphPlannerUser**](MicrosoftGraphPlannerUser.md) | New navigation property values | 

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

