# \UsersPlannerUserApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeletePlanner**](UsersPlannerUserApi.md#UsersDeletePlanner) | **Delete** /users/{user-id}/planner | Delete navigation property planner for users
[**UsersGetPlanner**](UsersPlannerUserApi.md#UsersGetPlanner) | **Get** /users/{user-id}/planner | Get planner from users
[**UsersPlannerCreatePlans**](UsersPlannerUserApi.md#UsersPlannerCreatePlans) | **Post** /users/{user-id}/planner/plans | Create new navigation property to plans for users
[**UsersPlannerCreateTasks**](UsersPlannerUserApi.md#UsersPlannerCreateTasks) | **Post** /users/{user-id}/planner/tasks | Create new navigation property to tasks for users
[**UsersPlannerDeletePlans**](UsersPlannerUserApi.md#UsersPlannerDeletePlans) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id} | Delete navigation property plans for users
[**UsersPlannerDeleteTasks**](UsersPlannerUserApi.md#UsersPlannerDeleteTasks) | **Delete** /users/{user-id}/planner/tasks/{plannerTask-id} | Delete navigation property tasks for users
[**UsersPlannerGetPlans**](UsersPlannerUserApi.md#UsersPlannerGetPlans) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id} | Get plans from users
[**UsersPlannerGetTasks**](UsersPlannerUserApi.md#UsersPlannerGetTasks) | **Get** /users/{user-id}/planner/tasks/{plannerTask-id} | Get tasks from users
[**UsersPlannerListPlans**](UsersPlannerUserApi.md#UsersPlannerListPlans) | **Get** /users/{user-id}/planner/plans | Get plans from users
[**UsersPlannerListTasks**](UsersPlannerUserApi.md#UsersPlannerListTasks) | **Get** /users/{user-id}/planner/tasks | Get tasks from users
[**UsersPlannerPlansBucketsCreateTasks**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsCreateTasks) | **Post** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks | Create new navigation property to tasks for users
[**UsersPlannerPlansBucketsDeleteTasks**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsDeleteTasks) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Delete navigation property tasks for users
[**UsersPlannerPlansBucketsGetTasks**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsGetTasks) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Get tasks from users
[**UsersPlannerPlansBucketsListTasks**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsListTasks) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks | Get tasks from users
[**UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for users
[**UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for users
[**UsersPlannerPlansBucketsTasksDeleteDetails**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksDeleteDetails) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Delete navigation property details for users
[**UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for users
[**UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from users
[**UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from users
[**UsersPlannerPlansBucketsTasksGetDetails**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksGetDetails) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Get details from users
[**UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from users
[**UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in users
[**UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in users
[**UsersPlannerPlansBucketsTasksUpdateDetails**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksUpdateDetails) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/details | Update the navigation property details in users
[**UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in users
[**UsersPlannerPlansBucketsUpdateTasks**](UsersPlannerUserApi.md#UsersPlannerPlansBucketsUpdateTasks) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id}/tasks/{plannerTask-id} | Update the navigation property tasks in users
[**UsersPlannerPlansCreateBuckets**](UsersPlannerUserApi.md#UsersPlannerPlansCreateBuckets) | **Post** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets | Create new navigation property to buckets for users
[**UsersPlannerPlansCreateTasks**](UsersPlannerUserApi.md#UsersPlannerPlansCreateTasks) | **Post** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks | Create new navigation property to tasks for users
[**UsersPlannerPlansDeleteBuckets**](UsersPlannerUserApi.md#UsersPlannerPlansDeleteBuckets) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Delete navigation property buckets for users
[**UsersPlannerPlansDeleteDetails**](UsersPlannerUserApi.md#UsersPlannerPlansDeleteDetails) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/details | Delete navigation property details for users
[**UsersPlannerPlansDeleteTasks**](UsersPlannerUserApi.md#UsersPlannerPlansDeleteTasks) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Delete navigation property tasks for users
[**UsersPlannerPlansGetBuckets**](UsersPlannerUserApi.md#UsersPlannerPlansGetBuckets) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Get buckets from users
[**UsersPlannerPlansGetDetails**](UsersPlannerUserApi.md#UsersPlannerPlansGetDetails) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/details | Get details from users
[**UsersPlannerPlansGetTasks**](UsersPlannerUserApi.md#UsersPlannerPlansGetTasks) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Get tasks from users
[**UsersPlannerPlansListBuckets**](UsersPlannerUserApi.md#UsersPlannerPlansListBuckets) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets | Get buckets from users
[**UsersPlannerPlansListTasks**](UsersPlannerUserApi.md#UsersPlannerPlansListTasks) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks | Get tasks from users
[**UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for users
[**UsersPlannerPlansTasksDeleteBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksDeleteBucketTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for users
[**UsersPlannerPlansTasksDeleteDetails**](UsersPlannerUserApi.md#UsersPlannerPlansTasksDeleteDetails) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Delete navigation property details for users
[**UsersPlannerPlansTasksDeleteProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksDeleteProgressTaskBoardFormat) | **Delete** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for users
[**UsersPlannerPlansTasksGetAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksGetAssignedToTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from users
[**UsersPlannerPlansTasksGetBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksGetBucketTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from users
[**UsersPlannerPlansTasksGetDetails**](UsersPlannerUserApi.md#UsersPlannerPlansTasksGetDetails) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Get details from users
[**UsersPlannerPlansTasksGetProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksGetProgressTaskBoardFormat) | **Get** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from users
[**UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in users
[**UsersPlannerPlansTasksUpdateBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksUpdateBucketTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in users
[**UsersPlannerPlansTasksUpdateDetails**](UsersPlannerUserApi.md#UsersPlannerPlansTasksUpdateDetails) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/details | Update the navigation property details in users
[**UsersPlannerPlansTasksUpdateProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerPlansTasksUpdateProgressTaskBoardFormat) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in users
[**UsersPlannerPlansUpdateBuckets**](UsersPlannerUserApi.md#UsersPlannerPlansUpdateBuckets) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/buckets/{plannerBucket-id} | Update the navigation property buckets in users
[**UsersPlannerPlansUpdateDetails**](UsersPlannerUserApi.md#UsersPlannerPlansUpdateDetails) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/details | Update the navigation property details in users
[**UsersPlannerPlansUpdateTasks**](UsersPlannerUserApi.md#UsersPlannerPlansUpdateTasks) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id}/tasks/{plannerTask-id} | Update the navigation property tasks in users
[**UsersPlannerTasksDeleteAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksDeleteAssignedToTaskBoardFormat) | **Delete** /users/{user-id}/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Delete navigation property assignedToTaskBoardFormat for users
[**UsersPlannerTasksDeleteBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksDeleteBucketTaskBoardFormat) | **Delete** /users/{user-id}/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Delete navigation property bucketTaskBoardFormat for users
[**UsersPlannerTasksDeleteDetails**](UsersPlannerUserApi.md#UsersPlannerTasksDeleteDetails) | **Delete** /users/{user-id}/planner/tasks/{plannerTask-id}/details | Delete navigation property details for users
[**UsersPlannerTasksDeleteProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksDeleteProgressTaskBoardFormat) | **Delete** /users/{user-id}/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Delete navigation property progressTaskBoardFormat for users
[**UsersPlannerTasksGetAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksGetAssignedToTaskBoardFormat) | **Get** /users/{user-id}/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Get assignedToTaskBoardFormat from users
[**UsersPlannerTasksGetBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksGetBucketTaskBoardFormat) | **Get** /users/{user-id}/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Get bucketTaskBoardFormat from users
[**UsersPlannerTasksGetDetails**](UsersPlannerUserApi.md#UsersPlannerTasksGetDetails) | **Get** /users/{user-id}/planner/tasks/{plannerTask-id}/details | Get details from users
[**UsersPlannerTasksGetProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksGetProgressTaskBoardFormat) | **Get** /users/{user-id}/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Get progressTaskBoardFormat from users
[**UsersPlannerTasksUpdateAssignedToTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksUpdateAssignedToTaskBoardFormat) | **Patch** /users/{user-id}/planner/tasks/{plannerTask-id}/assignedToTaskBoardFormat | Update the navigation property assignedToTaskBoardFormat in users
[**UsersPlannerTasksUpdateBucketTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksUpdateBucketTaskBoardFormat) | **Patch** /users/{user-id}/planner/tasks/{plannerTask-id}/bucketTaskBoardFormat | Update the navigation property bucketTaskBoardFormat in users
[**UsersPlannerTasksUpdateDetails**](UsersPlannerUserApi.md#UsersPlannerTasksUpdateDetails) | **Patch** /users/{user-id}/planner/tasks/{plannerTask-id}/details | Update the navigation property details in users
[**UsersPlannerTasksUpdateProgressTaskBoardFormat**](UsersPlannerUserApi.md#UsersPlannerTasksUpdateProgressTaskBoardFormat) | **Patch** /users/{user-id}/planner/tasks/{plannerTask-id}/progressTaskBoardFormat | Update the navigation property progressTaskBoardFormat in users
[**UsersPlannerUpdatePlans**](UsersPlannerUserApi.md#UsersPlannerUpdatePlans) | **Patch** /users/{user-id}/planner/plans/{plannerPlan-id} | Update the navigation property plans in users
[**UsersPlannerUpdateTasks**](UsersPlannerUserApi.md#UsersPlannerUpdateTasks) | **Patch** /users/{user-id}/planner/tasks/{plannerTask-id} | Update the navigation property tasks in users
[**UsersUpdatePlanner**](UsersPlannerUserApi.md#UsersUpdatePlanner) | **Patch** /users/{user-id}/planner | Update the navigation property planner in users



## UsersDeletePlanner

> UsersDeletePlanner(ctx, userId).IfMatch(ifMatch).Execute()

Delete navigation property planner for users



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
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersDeletePlanner(context.Background(), userId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersDeletePlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersDeletePlannerRequest struct via the builder pattern


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


## UsersGetPlanner

> MicrosoftGraphPlannerUser UsersGetPlanner(ctx, userId).Select_(select_).Expand(expand).Execute()

Get planner from users



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
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersGetPlanner(context.Background(), userId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersGetPlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetPlanner`: MicrosoftGraphPlannerUser
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersGetPlanner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetPlannerRequest struct via the builder pattern


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


## UsersPlannerCreatePlans

> MicrosoftGraphPlannerPlan UsersPlannerCreatePlans(ctx, userId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()

Create new navigation property to plans for users



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
    microsoftGraphPlannerPlan := *openapiclient.NewMicrosoftGraphPlannerPlan() // MicrosoftGraphPlannerPlan | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerCreatePlans(context.Background(), userId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerCreatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerCreatePlans`: MicrosoftGraphPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerCreatePlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerCreatePlansRequest struct via the builder pattern


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


## UsersPlannerCreateTasks

> MicrosoftGraphPlannerTask UsersPlannerCreateTasks(ctx, userId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for users



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
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerCreateTasks(context.Background(), userId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerCreateTasksRequest struct via the builder pattern


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


## UsersPlannerDeletePlans

> UsersPlannerDeletePlans(ctx, userId, plannerPlanId).IfMatch(ifMatch).Execute()

Delete navigation property plans for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerDeletePlans(context.Background(), userId, plannerPlanId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerDeletePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerDeletePlansRequest struct via the builder pattern


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


## UsersPlannerDeleteTasks

> UsersPlannerDeleteTasks(ctx, userId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerDeleteTasks(context.Background(), userId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerDeleteTasksRequest struct via the builder pattern


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


## UsersPlannerGetPlans

> MicrosoftGraphPlannerPlan UsersPlannerGetPlans(ctx, userId, plannerPlanId).Select_(select_).Expand(expand).Execute()

Get plans from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerGetPlans(context.Background(), userId, plannerPlanId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerGetPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerGetPlans`: MicrosoftGraphPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerGetPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerGetPlansRequest struct via the builder pattern


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


## UsersPlannerGetTasks

> MicrosoftGraphPlannerTask UsersPlannerGetTasks(ctx, userId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerGetTasks(context.Background(), userId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerGetTasksRequest struct via the builder pattern


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


## UsersPlannerListPlans

> CollectionOfPlannerPlan UsersPlannerListPlans(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get plans from users



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
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerListPlans(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerListPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerListPlans`: CollectionOfPlannerPlan
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerListPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerListPlansRequest struct via the builder pattern


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


## UsersPlannerListTasks

> CollectionOfPlannerTask UsersPlannerListTasks(ctx, userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerListTasks(context.Background(), userId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerListTasksRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsCreateTasks

> MicrosoftGraphPlannerTask UsersPlannerPlansBucketsCreateTasks(ctx, userId, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsCreateTasks(context.Background(), userId, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsCreateTasksRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsDeleteTasks

> UsersPlannerPlansBucketsDeleteTasks(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsDeleteTasks(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsDeleteTasksRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsGetTasks

> MicrosoftGraphPlannerTask UsersPlannerPlansBucketsGetTasks(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsGetTasks(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsGetTasksRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsListTasks

> CollectionOfPlannerTask UsersPlannerPlansBucketsListTasks(ctx, userId, plannerPlanId, plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsListTasks(context.Background(), userId, plannerPlanId, plannerBucketId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsListTasksRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat

> UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat

> UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksDeleteDetails

> UsersPlannerPlansBucketsTasksDeleteDetails(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteDetails(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksDeleteDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat

> UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksGetDetails

> MicrosoftGraphPlannerTaskDetails UsersPlannerPlansBucketsTasksGetDetails(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetDetails(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksGetDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat

> UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat

> UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksUpdateDetails

> UsersPlannerPlansBucketsTasksUpdateDetails(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateDetails(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksUpdateDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat

> UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansBucketsUpdateTasks

> UsersPlannerPlansBucketsUpdateTasks(ctx, userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansBucketsUpdateTasks(context.Background(), userId, plannerPlanId, plannerBucketId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansBucketsUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansBucketsUpdateTasksRequest struct via the builder pattern


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


## UsersPlannerPlansCreateBuckets

> MicrosoftGraphPlannerBucket UsersPlannerPlansCreateBuckets(ctx, userId, plannerPlanId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Create new navigation property to buckets for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansCreateBuckets(context.Background(), userId, plannerPlanId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansCreateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansCreateBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansCreateBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansCreateBucketsRequest struct via the builder pattern


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


## UsersPlannerPlansCreateTasks

> MicrosoftGraphPlannerTask UsersPlannerPlansCreateTasks(ctx, userId, plannerPlanId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Create new navigation property to tasks for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansCreateTasks(context.Background(), userId, plannerPlanId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansCreateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansCreateTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansCreateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansCreateTasksRequest struct via the builder pattern


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


## UsersPlannerPlansDeleteBuckets

> UsersPlannerPlansDeleteBuckets(ctx, userId, plannerPlanId, plannerBucketId).IfMatch(ifMatch).Execute()

Delete navigation property buckets for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansDeleteBuckets(context.Background(), userId, plannerPlanId, plannerBucketId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansDeleteBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansDeleteBucketsRequest struct via the builder pattern


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


## UsersPlannerPlansDeleteDetails

> UsersPlannerPlansDeleteDetails(ctx, userId, plannerPlanId).IfMatch(ifMatch).Execute()

Delete navigation property details for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansDeleteDetails(context.Background(), userId, plannerPlanId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansDeleteDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansDeleteTasks

> UsersPlannerPlansDeleteTasks(ctx, userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property tasks for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansDeleteTasks(context.Background(), userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansDeleteTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansDeleteTasksRequest struct via the builder pattern


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


## UsersPlannerPlansGetBuckets

> MicrosoftGraphPlannerBucket UsersPlannerPlansGetBuckets(ctx, userId, plannerPlanId, plannerBucketId).Select_(select_).Expand(expand).Execute()

Get buckets from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansGetBuckets(context.Background(), userId, plannerPlanId, plannerBucketId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansGetBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansGetBuckets`: MicrosoftGraphPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansGetBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansGetBucketsRequest struct via the builder pattern


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


## UsersPlannerPlansGetDetails

> MicrosoftGraphPlannerPlanDetails UsersPlannerPlansGetDetails(ctx, userId, plannerPlanId).Select_(select_).Expand(expand).Execute()

Get details from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansGetDetails(context.Background(), userId, plannerPlanId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansGetDetails`: MicrosoftGraphPlannerPlanDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansGetDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansGetTasks

> MicrosoftGraphPlannerTask UsersPlannerPlansGetTasks(ctx, userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansGetTasks(context.Background(), userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansGetTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansGetTasks`: MicrosoftGraphPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansGetTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansGetTasksRequest struct via the builder pattern


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


## UsersPlannerPlansListBuckets

> CollectionOfPlannerBucket UsersPlannerPlansListBuckets(ctx, userId, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get buckets from users



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
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansListBuckets(context.Background(), userId, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansListBuckets`: CollectionOfPlannerBucket
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansListBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansListBucketsRequest struct via the builder pattern


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


## UsersPlannerPlansListTasks

> CollectionOfPlannerTask UsersPlannerPlansListTasks(ctx, userId, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()

Get tasks from users



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
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansListTasks(context.Background(), userId, plannerPlanId).Top(top).Skip(skip).Search(search).Filter(filter).Count(count).Orderby(orderby).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansListTasks`: CollectionOfPlannerTask
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansListTasksRequest struct via the builder pattern


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


## UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat

> UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksDeleteBucketTaskBoardFormat

> UsersPlannerPlansTasksDeleteBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksDeleteBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksDeleteDetails

> UsersPlannerPlansTasksDeleteDetails(ctx, userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksDeleteDetails(context.Background(), userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksDeleteDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansTasksDeleteProgressTaskBoardFormat

> UsersPlannerPlansTasksDeleteProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksDeleteProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat UsersPlannerPlansTasksGetAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksGetAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat UsersPlannerPlansTasksGetBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksGetBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksGetDetails

> MicrosoftGraphPlannerTaskDetails UsersPlannerPlansTasksGetDetails(ctx, userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksGetDetails(context.Background(), userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksGetDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat UsersPlannerPlansTasksGetProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksGetProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerPlansTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerPlansTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat

> UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksUpdateBucketTaskBoardFormat

> UsersPlannerPlansTasksUpdateBucketTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksUpdateBucketTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansTasksUpdateDetails

> UsersPlannerPlansTasksUpdateDetails(ctx, userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksUpdateDetails(context.Background(), userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksUpdateDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansTasksUpdateProgressTaskBoardFormat

> UsersPlannerPlansTasksUpdateProgressTaskBoardFormat(ctx, userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansTasksUpdateProgressTaskBoardFormat(context.Background(), userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerPlansUpdateBuckets

> UsersPlannerPlansUpdateBuckets(ctx, userId, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()

Update the navigation property buckets in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerBucketId := "plannerBucketId_example" // string | key: id of plannerBucket
    microsoftGraphPlannerBucket := *openapiclient.NewMicrosoftGraphPlannerBucket() // MicrosoftGraphPlannerBucket | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansUpdateBuckets(context.Background(), userId, plannerPlanId, plannerBucketId).MicrosoftGraphPlannerBucket(microsoftGraphPlannerBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansUpdateBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerBucketId** | **string** | key: id of plannerBucket | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansUpdateBucketsRequest struct via the builder pattern


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


## UsersPlannerPlansUpdateDetails

> UsersPlannerPlansUpdateDetails(ctx, userId, plannerPlanId).MicrosoftGraphPlannerPlanDetails(microsoftGraphPlannerPlanDetails).Execute()

Update the navigation property details in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerPlanDetails := *openapiclient.NewMicrosoftGraphPlannerPlanDetails() // MicrosoftGraphPlannerPlanDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansUpdateDetails(context.Background(), userId, plannerPlanId).MicrosoftGraphPlannerPlanDetails(microsoftGraphPlannerPlanDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansUpdateDetailsRequest struct via the builder pattern


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


## UsersPlannerPlansUpdateTasks

> UsersPlannerPlansUpdateTasks(ctx, userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerPlansUpdateTasks(context.Background(), userId, plannerPlanId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerPlansUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerPlansUpdateTasksRequest struct via the builder pattern


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


## UsersPlannerTasksDeleteAssignedToTaskBoardFormat

> UsersPlannerTasksDeleteAssignedToTaskBoardFormat(ctx, userId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property assignedToTaskBoardFormat for users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksDeleteAssignedToTaskBoardFormat(context.Background(), userId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksDeleteAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksDeleteAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksDeleteBucketTaskBoardFormat

> UsersPlannerTasksDeleteBucketTaskBoardFormat(ctx, userId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property bucketTaskBoardFormat for users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksDeleteBucketTaskBoardFormat(context.Background(), userId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksDeleteBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksDeleteBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksDeleteDetails

> UsersPlannerTasksDeleteDetails(ctx, userId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property details for users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksDeleteDetails(context.Background(), userId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksDeleteDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksDeleteDetailsRequest struct via the builder pattern


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


## UsersPlannerTasksDeleteProgressTaskBoardFormat

> UsersPlannerTasksDeleteProgressTaskBoardFormat(ctx, userId, plannerTaskId).IfMatch(ifMatch).Execute()

Delete navigation property progressTaskBoardFormat for users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    ifMatch := "ifMatch_example" // string | ETag (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksDeleteProgressTaskBoardFormat(context.Background(), userId, plannerTaskId).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksDeleteProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksDeleteProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksGetAssignedToTaskBoardFormat

> MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat UsersPlannerTasksGetAssignedToTaskBoardFormat(ctx, userId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get assignedToTaskBoardFormat from users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksGetAssignedToTaskBoardFormat(context.Background(), userId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksGetAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerTasksGetAssignedToTaskBoardFormat`: MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerTasksGetAssignedToTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksGetAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksGetBucketTaskBoardFormat

> MicrosoftGraphPlannerBucketTaskBoardTaskFormat UsersPlannerTasksGetBucketTaskBoardFormat(ctx, userId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get bucketTaskBoardFormat from users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksGetBucketTaskBoardFormat(context.Background(), userId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksGetBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerTasksGetBucketTaskBoardFormat`: MicrosoftGraphPlannerBucketTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerTasksGetBucketTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksGetBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksGetDetails

> MicrosoftGraphPlannerTaskDetails UsersPlannerTasksGetDetails(ctx, userId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get details from users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksGetDetails(context.Background(), userId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksGetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerTasksGetDetails`: MicrosoftGraphPlannerTaskDetails
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerTasksGetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksGetDetailsRequest struct via the builder pattern


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


## UsersPlannerTasksGetProgressTaskBoardFormat

> MicrosoftGraphPlannerProgressTaskBoardTaskFormat UsersPlannerTasksGetProgressTaskBoardFormat(ctx, userId, plannerTaskId).Select_(select_).Expand(expand).Execute()

Get progressTaskBoardFormat from users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    select_ := []string{"Select_example"} // []string | Select properties to be returned (optional)
    expand := []string{"Expand_example"} // []string | Expand related entities (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksGetProgressTaskBoardFormat(context.Background(), userId, plannerTaskId).Select_(select_).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksGetProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersPlannerTasksGetProgressTaskBoardFormat`: MicrosoftGraphPlannerProgressTaskBoardTaskFormat
    fmt.Fprintf(os.Stdout, "Response from `UsersPlannerUserApi.UsersPlannerTasksGetProgressTaskBoardFormat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksGetProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksUpdateAssignedToTaskBoardFormat

> UsersPlannerTasksUpdateAssignedToTaskBoardFormat(ctx, userId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()

Update the navigation property assignedToTaskBoardFormat in users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerAssignedToTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerAssignedToTaskBoardTaskFormat() // MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksUpdateAssignedToTaskBoardFormat(context.Background(), userId, plannerTaskId).MicrosoftGraphPlannerAssignedToTaskBoardTaskFormat(microsoftGraphPlannerAssignedToTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksUpdateAssignedToTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksUpdateAssignedToTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksUpdateBucketTaskBoardFormat

> UsersPlannerTasksUpdateBucketTaskBoardFormat(ctx, userId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()

Update the navigation property bucketTaskBoardFormat in users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerBucketTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerBucketTaskBoardTaskFormat() // MicrosoftGraphPlannerBucketTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksUpdateBucketTaskBoardFormat(context.Background(), userId, plannerTaskId).MicrosoftGraphPlannerBucketTaskBoardTaskFormat(microsoftGraphPlannerBucketTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksUpdateBucketTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksUpdateBucketTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerTasksUpdateDetails

> UsersPlannerTasksUpdateDetails(ctx, userId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()

Update the navigation property details in users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTaskDetails := *openapiclient.NewMicrosoftGraphPlannerTaskDetails() // MicrosoftGraphPlannerTaskDetails | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksUpdateDetails(context.Background(), userId, plannerTaskId).MicrosoftGraphPlannerTaskDetails(microsoftGraphPlannerTaskDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksUpdateDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksUpdateDetailsRequest struct via the builder pattern


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


## UsersPlannerTasksUpdateProgressTaskBoardFormat

> UsersPlannerTasksUpdateProgressTaskBoardFormat(ctx, userId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()

Update the navigation property progressTaskBoardFormat in users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerProgressTaskBoardTaskFormat := *openapiclient.NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() // MicrosoftGraphPlannerProgressTaskBoardTaskFormat | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerTasksUpdateProgressTaskBoardFormat(context.Background(), userId, plannerTaskId).MicrosoftGraphPlannerProgressTaskBoardTaskFormat(microsoftGraphPlannerProgressTaskBoardTaskFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerTasksUpdateProgressTaskBoardFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerTasksUpdateProgressTaskBoardFormatRequest struct via the builder pattern


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


## UsersPlannerUpdatePlans

> UsersPlannerUpdatePlans(ctx, userId, plannerPlanId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()

Update the navigation property plans in users



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
    plannerPlanId := "plannerPlanId_example" // string | key: id of plannerPlan
    microsoftGraphPlannerPlan := *openapiclient.NewMicrosoftGraphPlannerPlan() // MicrosoftGraphPlannerPlan | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerUpdatePlans(context.Background(), userId, plannerPlanId).MicrosoftGraphPlannerPlan(microsoftGraphPlannerPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerUpdatePlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerPlanId** | **string** | key: id of plannerPlan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerUpdatePlansRequest struct via the builder pattern


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


## UsersPlannerUpdateTasks

> UsersPlannerUpdateTasks(ctx, userId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()

Update the navigation property tasks in users



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
    plannerTaskId := "plannerTaskId_example" // string | key: id of plannerTask
    microsoftGraphPlannerTask := *openapiclient.NewMicrosoftGraphPlannerTask() // MicrosoftGraphPlannerTask | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersPlannerUpdateTasks(context.Background(), userId, plannerTaskId).MicrosoftGraphPlannerTask(microsoftGraphPlannerTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersPlannerUpdateTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 
**plannerTaskId** | **string** | key: id of plannerTask | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersPlannerUpdateTasksRequest struct via the builder pattern


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


## UsersUpdatePlanner

> UsersUpdatePlanner(ctx, userId).MicrosoftGraphPlannerUser(microsoftGraphPlannerUser).Execute()

Update the navigation property planner in users



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
    microsoftGraphPlannerUser := *openapiclient.NewMicrosoftGraphPlannerUser() // MicrosoftGraphPlannerUser | New navigation property values

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersPlannerUserApi.UsersUpdatePlanner(context.Background(), userId).MicrosoftGraphPlannerUser(microsoftGraphPlannerUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersPlannerUserApi.UsersUpdatePlanner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | key: id of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdatePlannerRequest struct via the builder pattern


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

