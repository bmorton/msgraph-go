# \AdminFunctionsApi

All URIs are relative to *https://graph.microsoft.com/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport**](AdminFunctionsApi.md#AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport) | **Get** /admin/serviceAnnouncement/healthOverviews/{serviceHealth-id}/issues/{serviceHealthIssue-id}/microsoft.graph.incidentReport() | Invoke function incidentReport
[**AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport**](AdminFunctionsApi.md#AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport) | **Get** /admin/serviceAnnouncement/issues/{serviceHealthIssue-id}/microsoft.graph.incidentReport() | Invoke function incidentReport



## AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport

> string AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport(ctx, serviceHealthId, serviceHealthIssueId).Execute()

Invoke function incidentReport

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
    serviceHealthId := "serviceHealthId_example" // string | key: id of serviceHealth
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminFunctionsApi.AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport(context.Background(), serviceHealthId, serviceHealthIssueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFunctionsApi.AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport`: string
    fmt.Fprintf(os.Stdout, "Response from `AdminFunctionsApi.AdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthId** | **string** | key: id of serviceHealth | 
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementHealthOverviewsServiceHealthIssuesServiceHealthIssueIncidentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport

> string AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport(ctx, serviceHealthIssueId).Execute()

Invoke function incidentReport

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
    serviceHealthIssueId := "serviceHealthIssueId_example" // string | key: id of serviceHealthIssue

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminFunctionsApi.AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport(context.Background(), serviceHealthIssueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminFunctionsApi.AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport`: string
    fmt.Fprintf(os.Stdout, "Response from `AdminFunctionsApi.AdminServiceAnnouncementIssuesServiceHealthIssueIncidentReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceHealthIssueId** | **string** | key: id of serviceHealthIssue | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminServiceAnnouncementIssuesServiceHealthIssueIncidentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[azureaadv2](../README.md#azureaadv2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

