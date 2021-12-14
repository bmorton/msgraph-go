# MicrosoftGraphDeviceManagementReports

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ExportJobs** | Pointer to [**[]MicrosoftGraphDeviceManagementExportJob**](MicrosoftGraphDeviceManagementExportJob.md) | Entity representing a job to export a report | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementReports

`func NewMicrosoftGraphDeviceManagementReports() *MicrosoftGraphDeviceManagementReports`

NewMicrosoftGraphDeviceManagementReports instantiates a new MicrosoftGraphDeviceManagementReports object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementReportsWithDefaults

`func NewMicrosoftGraphDeviceManagementReportsWithDefaults() *MicrosoftGraphDeviceManagementReports`

NewMicrosoftGraphDeviceManagementReportsWithDefaults instantiates a new MicrosoftGraphDeviceManagementReports object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagementReports) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementReports) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementReports) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagementReports) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExportJobs

`func (o *MicrosoftGraphDeviceManagementReports) GetExportJobs() []MicrosoftGraphDeviceManagementExportJob`

GetExportJobs returns the ExportJobs field if non-nil, zero value otherwise.

### GetExportJobsOk

`func (o *MicrosoftGraphDeviceManagementReports) GetExportJobsOk() (*[]MicrosoftGraphDeviceManagementExportJob, bool)`

GetExportJobsOk returns a tuple with the ExportJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportJobs

`func (o *MicrosoftGraphDeviceManagementReports) SetExportJobs(v []MicrosoftGraphDeviceManagementExportJob)`

SetExportJobs sets ExportJobs field to given value.

### HasExportJobs

`func (o *MicrosoftGraphDeviceManagementReports) HasExportJobs() bool`

HasExportJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


