# \IdentityGovernanceFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/{accessReviewInstance-id}/decisions/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser) | **Get** /identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition-id}/instances/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser) | **Get** /identityGovernance/accessReviews/definitions/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser) | **Get** /identityGovernance/appConsent/appConsentRequests/{appConsentRequest-id}/userConsentRequests/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser) | **Get** /identityGovernance/appConsent/appConsentRequests/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser) | **Get** /identityGovernance/entitlementManagement/accessPackageAssignmentApprovals/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser) | **Get** /identityGovernance/entitlementManagement/accessPackages/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser) | **Get** /identityGovernance/entitlementManagement/assignmentRequests/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser) | **Get** /identityGovernance/entitlementManagement/assignments/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser
[**IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser**](IdentityGovernanceFunctionsApi.md#IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser) | **Get** /identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog-id}/accessPackages/microsoft.graph.filterByCurrentUser(on&#x3D;{on}) | Invoke function filterByCurrentUser



## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItem IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser(ctx, accessReviewScheduleDefinitionId, accessReviewInstanceId, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser(context.Background(), accessReviewScheduleDefinitionId, accessReviewInstanceId, true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItem
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**accessReviewInstanceId** | **string** | key: id of accessReviewInstance | 
**true** | [**AnyOfmicrosoftGraphAccessReviewInstanceDecisionItemFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesAccessReviewInstanceDecisionsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]*AnyOfmicrosoftGraphAccessReviewInstanceDecisionItem**](anyOf&lt;microsoft.graph.accessReviewInstanceDecisionItem&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessReviewInstance IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser(ctx, accessReviewScheduleDefinitionId, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessReviewInstanceFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser(context.Background(), accessReviewScheduleDefinitionId, true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessReviewInstance
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessReviewScheduleDefinitionId** | **string** | key: id of accessReviewScheduleDefinition | 
**true** | [**AnyOfmicrosoftGraphAccessReviewInstanceFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsAccessReviewScheduleDefinitionInstancesFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphAccessReviewInstance**](anyOf&lt;microsoft.graph.accessReviewInstance&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessReviewScheduleDefinition IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessReviewScheduleDefinition
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphAccessReviewScheduleDefinitionFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAccessReviewsDefinitionsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessReviewScheduleDefinition**](anyOf&lt;microsoft.graph.accessReviewScheduleDefinition&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser

> []*AnyOfmicrosoftGraphUserConsentRequest IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser(ctx, appConsentRequestId, true).Execute()

Invoke function filterByCurrentUser

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
    appConsentRequestId := "appConsentRequestId_example" // string | key: id of appConsentRequest
    true := TODO // AnyOfmicrosoftGraphConsentRequestFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser(context.Background(), appConsentRequestId, true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser`: []*AnyOfmicrosoftGraphUserConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appConsentRequestId** | **string** | key: id of appConsentRequest | 
**true** | [**AnyOfmicrosoftGraphConsentRequestFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAppConsentAppConsentRequestsAppConsentRequestUserConsentRequestsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphUserConsentRequest**](anyOf&lt;microsoft.graph.userConsentRequest&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser

> []*AnyOfmicrosoftGraphAppConsentRequest IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphConsentRequestFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser`: []*AnyOfmicrosoftGraphAppConsentRequest
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphConsentRequestFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceAppConsentAppConsentRequestsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAppConsentRequest**](anyOf&lt;microsoft.graph.appConsentRequest&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser

> []*AnyOfmicrosoftGraphApproval IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphApprovalFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser`: []*AnyOfmicrosoftGraphApproval
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphApprovalFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphApproval**](anyOf&lt;microsoft.graph.approval&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessPackage IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessPackageFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessPackage
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphAccessPackageFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAccessPackagesFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackage**](anyOf&lt;microsoft.graph.accessPackage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessPackageAssignmentRequest IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessPackageAssignmentRequestFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessPackageAssignmentRequest
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphAccessPackageAssignmentRequestFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAssignmentRequestsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignmentRequest**](anyOf&lt;microsoft.graph.accessPackageAssignmentRequest&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessPackageAssignment IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser(ctx, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessPackageAssignmentFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser(context.Background(), true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessPackageAssignment
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**true** | [**AnyOfmicrosoftGraphAccessPackageAssignmentFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementAssignmentsFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]*AnyOfmicrosoftGraphAccessPackageAssignment**](anyOf&lt;microsoft.graph.accessPackageAssignment&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser

> []*AnyOfmicrosoftGraphAccessPackage IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser(ctx, accessPackageCatalogId, true).Execute()

Invoke function filterByCurrentUser

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
    true := TODO // AnyOfmicrosoftGraphAccessPackageFilterByCurrentUserOptions | Usage: on={on}

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser(context.Background(), accessPackageCatalogId, true).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser`: []*AnyOfmicrosoftGraphAccessPackage
    fmt.Fprintf(os.Stdout, "Response from `IdentityGovernanceFunctionsApi.IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessPackageCatalogId** | **string** | key: id of accessPackageCatalog | 
**true** | [**AnyOfmicrosoftGraphAccessPackageFilterByCurrentUserOptions**](.md) | Usage: on&#x3D;{on} | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogAccessPackagesFilterByCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]*AnyOfmicrosoftGraphAccessPackage**](anyOf&lt;microsoft.graph.accessPackage&gt;.md)

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

