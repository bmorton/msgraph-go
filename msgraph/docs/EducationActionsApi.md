# \EducationActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EducationClassesEducationClassAssignmentsEducationAssignmentPublish**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentPublish) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/microsoft.graph.publish | Invoke action publish
[**EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/microsoft.graph.setUpResourcesFolder | Invoke action setUpResourcesFolder
[**EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/submissions/{educationSubmission-id}/microsoft.graph.return | Invoke action return
[**EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/submissions/{educationSubmission-id}/microsoft.graph.setUpResourcesFolder | Invoke action setUpResourcesFolder
[**EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/submissions/{educationSubmission-id}/microsoft.graph.submit | Invoke action submit
[**EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit**](EducationActionsApi.md#EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit) | **Post** /education/classes/{educationClass-id}/assignments/{educationAssignment-id}/submissions/{educationSubmission-id}/microsoft.graph.unsubmit | Invoke action unsubmit



## EducationClassesEducationClassAssignmentsEducationAssignmentPublish

> AnyOfmicrosoftGraphEducationAssignment EducationClassesEducationClassAssignmentsEducationAssignmentPublish(ctx, educationClassId, educationAssignmentId).Execute()

Invoke action publish

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentPublish(context.Background(), educationClassId, educationAssignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentPublish`: AnyOfmicrosoftGraphEducationAssignment
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnyOfmicrosoftGraphEducationAssignment**](anyOf&lt;microsoft.graph.educationAssignment&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder

> AnyOfmicrosoftGraphEducationAssignment EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder(ctx, educationClassId, educationAssignmentId).Execute()

Invoke action setUpResourcesFolder

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder(context.Background(), educationClassId, educationAssignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder`: AnyOfmicrosoftGraphEducationAssignment
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentSetUpResourcesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnyOfmicrosoftGraphEducationAssignment**](anyOf&lt;microsoft.graph.educationAssignment&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn

> AnyOfmicrosoftGraphEducationSubmission EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn(ctx, educationClassId, educationAssignmentId, educationSubmissionId).Execute()

Invoke action return

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment
    educationSubmissionId := "educationSubmissionId_example" // string | key: id of educationSubmission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn(context.Background(), educationClassId, educationAssignmentId, educationSubmissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn`: AnyOfmicrosoftGraphEducationSubmission
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 
**educationSubmissionId** | **string** | key: id of educationSubmission | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AnyOfmicrosoftGraphEducationSubmission**](anyOf&lt;microsoft.graph.educationSubmission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder

> AnyOfmicrosoftGraphEducationSubmission EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder(ctx, educationClassId, educationAssignmentId, educationSubmissionId).Execute()

Invoke action setUpResourcesFolder

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment
    educationSubmissionId := "educationSubmissionId_example" // string | key: id of educationSubmission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder(context.Background(), educationClassId, educationAssignmentId, educationSubmissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder`: AnyOfmicrosoftGraphEducationSubmission
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 
**educationSubmissionId** | **string** | key: id of educationSubmission | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSetUpResourcesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AnyOfmicrosoftGraphEducationSubmission**](anyOf&lt;microsoft.graph.educationSubmission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit

> AnyOfmicrosoftGraphEducationSubmission EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit(ctx, educationClassId, educationAssignmentId, educationSubmissionId).Execute()

Invoke action submit

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment
    educationSubmissionId := "educationSubmissionId_example" // string | key: id of educationSubmission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit(context.Background(), educationClassId, educationAssignmentId, educationSubmissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit`: AnyOfmicrosoftGraphEducationSubmission
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 
**educationSubmissionId** | **string** | key: id of educationSubmission | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AnyOfmicrosoftGraphEducationSubmission**](anyOf&lt;microsoft.graph.educationSubmission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit

> AnyOfmicrosoftGraphEducationSubmission EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit(ctx, educationClassId, educationAssignmentId, educationSubmissionId).Execute()

Invoke action unsubmit

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
    educationClassId := "educationClassId_example" // string | key: id of educationClass
    educationAssignmentId := "educationAssignmentId_example" // string | key: id of educationAssignment
    educationSubmissionId := "educationSubmissionId_example" // string | key: id of educationSubmission

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit(context.Background(), educationClassId, educationAssignmentId, educationSubmissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit`: AnyOfmicrosoftGraphEducationSubmission
    fmt.Fprintf(os.Stdout, "Response from `EducationActionsApi.EducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**educationClassId** | **string** | key: id of educationClass | 
**educationAssignmentId** | **string** | key: id of educationAssignment | 
**educationSubmissionId** | **string** | key: id of educationSubmission | 

### Other Parameters

Other parameters are passed through a pointer to a apiEducationClassesEducationClassAssignmentsEducationAssignmentSubmissionsEducationSubmissionUnsubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AnyOfmicrosoftGraphEducationSubmission**](anyOf&lt;microsoft.graph.educationSubmission&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

