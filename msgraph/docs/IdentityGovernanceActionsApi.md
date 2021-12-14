# \IdentityGovernanceActionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.acceptRecommendations | Invoke action acceptRecommendations
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.applyDecisions | Invoke action applyDecisions
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.batchRecordDecisions | Invoke action batchRecordDecisions
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.resetDecisions | Invoke action resetDecisions
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.sendReminder | Invoke action sendReminder
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/microsoft.graph.stop | Invoke action stop
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop**](IdentityGovernanceActionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop) | **Post** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/microsoft.graph.stop | Invoke action stop
[**IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements**](IdentityGovernanceActionsApi.md#IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements) | **Post** /identityGovernance/entitlementManagement/accessPackages/{accessPackage-id}/microsoft.graph.getApplicablePolicyRequirements | Invoke action getApplicablePolicyRequirements
[**IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements**](IdentityGovernanceActionsApi.md#IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements) | **Post** /identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest-id}/accessPackage/microsoft.graph.getApplicablePolicyRequirements | Invoke action getApplicablePolicyRequirements
[**IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel**](IdentityGovernanceActionsApi.md#IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel) | **Post** /identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest-id}/microsoft.graph.cancel | Invoke action cancel
[**IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements**](IdentityGovernanceActionsApi.md#IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements) | **Post** /identityGovernance/entitlementManagement/assignments/{accessPackageAssignment-id}/accessPackage/microsoft.graph.getApplicablePolicyRequirements | Invoke action getApplicablePolicyRequirements
[**IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements**](IdentityGovernanceActionsApi.md#IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements) | **Post** /identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog-id}/accessPackages/{accessPackage-id}/microsoft.graph.getApplicablePolicyRequirements | Invoke action getApplicablePolicyRequirements



## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()

Invoke action acceptRecommendations

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceAcceptRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()

Invoke action applyDecisions

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceApplyDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).InlineObject358(inlineObject358).Execute()

Invoke action batchRecordDecisions

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance
    inlineObject358 := *openapiclient.NewInlineObject358() // InlineObject358 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).InlineObject358(inlineObject358).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceBatchRecordDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject358** | [**InlineObject358**](InlineObject358.md) |  | 

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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()

Invoke action resetDecisions

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceResetDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()

Invoke action sendReminder

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceSendReminderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()

Invoke action stop

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition
    accessReviewInstanceId := "accessReviewInstanceId_example" // string | key: id of accessReviewInstance

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop

> IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop(ctx, accessReviewScheduleDefinitionId).Execute()

Invoke action stop

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
    accessReviewScheduleDefinitionId := "accessReviewScheduleDefinitionId_example" // string | key: id of accessReviewScheduleDefinition

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop(context.Background(), accessReviewScheduleDefinitionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements

> []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements(ctx, accessPackageId).Execute()

Invoke action getApplicablePolicyRequirements

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
    accessPackageId := "accessPackageId_example" // string | key: id of accessPackage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements(context.Background(), accessPackageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements`: []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageId** | **string** | key: id of accessPackage | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAccessPackagesAccessPackageGetApplicablePolicyRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements**](anyOf&lt;microsoft.graph.accessPackageAssignmentRequestRequirements&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements

> []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements(ctx, accessPackageAssignmentRequestId).Execute()

Invoke action getApplicablePolicyRequirements

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
    accessPackageAssignmentRequestId := "accessPackageAssignmentRequestId_example" // string | key: id of accessPackageAssignmentRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements(context.Background(), accessPackageAssignmentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements`: []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageAssignmentRequestId** | **string** | key: id of accessPackageAssignmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestAccessPackageGetApplicablePolicyRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements**](anyOf&lt;microsoft.graph.accessPackageAssignmentRequestRequirements&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel

> IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel(ctx, accessPackageAssignmentRequestId).Execute()

Invoke action cancel

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
    accessPackageAssignmentRequestId := "accessPackageAssignmentRequestId_example" // string | key: id of accessPackageAssignmentRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel(context.Background(), accessPackageAssignmentRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageAssignmentRequestId** | **string** | key: id of accessPackageAssignmentRequest | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements

> []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements(ctx, accessPackageAssignmentId).Execute()

Invoke action getApplicablePolicyRequirements

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
    accessPackageAssignmentId := "accessPackageAssignmentId_example" // string | key: id of accessPackageAssignment

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements(context.Background(), accessPackageAssignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements`: []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageAssignmentId** | **string** | key: id of accessPackageAssignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentAccessPackageGetApplicablePolicyRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements**](anyOf&lt;microsoft.graph.accessPackageAssignmentRequestRequirements&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements

> []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements(ctx, accessPackageCatalogId, accessPackageId).Execute()

Invoke action getApplicablePolicyRequirements

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
    accessPackageCatalogId := "accessPackageCatalogId_example" // string | key: id of accessPackageCatalog
    accessPackageId := "accessPackageId_example" // string | key: id of accessPackage

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements(context.Background(), accessPackageCatalogId, accessPackageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements`: []*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceActionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageCatalogId** | **string** | key: id of accessPackageCatalog | 
**accessPackageId** | **string** | key: id of accessPackage | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesAccessPackageGetApplicablePolicyRequirementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignmentRequestRequirements**](anyOf&lt;microsoft.graph.accessPackageAssignmentRequestRequirements&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

