# PrintJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**MicrosoftGraphPrintJobConfiguration**](MicrosoftGraphPrintJobConfiguration.md) |  | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | Read-only. Nullable. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The DateTimeOffset when the job was created. Read-only. | [optional] 
**IsFetchable** | Pointer to **bool** | If true, document can be fetched by printer. | [optional] 
**RedirectedFrom** | Pointer to **NullableString** | Contains the source job URL, if the job has been redirected from another printer. | [optional] 
**RedirectedTo** | Pointer to **NullableString** | Contains the destination job URL, if the job has been redirected to another printer. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrintJobStatus**](MicrosoftGraphPrintJobStatus.md) |  | [optional] 
**Documents** | Pointer to [**[]MicrosoftGraphPrintDocument**](MicrosoftGraphPrintDocument.md) | Read-only. | [optional] 
**Tasks** | Pointer to [**[]MicrosoftGraphPrintTask**](MicrosoftGraphPrintTask.md) | A list of printTasks that were triggered by this print job. | [optional] 

## Methods

### NewPrintJob

`func NewPrintJob() *PrintJob`

NewPrintJob instantiates a new PrintJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintJobWithDefaults

`func NewPrintJobWithDefaults() *PrintJob`

NewPrintJobWithDefaults instantiates a new PrintJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PrintJob) GetConfiguration() MicrosoftGraphPrintJobConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PrintJob) GetConfigurationOk() (*MicrosoftGraphPrintJobConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PrintJob) SetConfiguration(v MicrosoftGraphPrintJobConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PrintJob) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrintJob) GetCreatedBy() AnyOfmicrosoftGraphUserIdentity`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrintJob) GetCreatedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrintJob) SetCreatedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrintJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PrintJob) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PrintJob) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *PrintJob) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PrintJob) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PrintJob) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PrintJob) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetIsFetchable

`func (o *PrintJob) GetIsFetchable() bool`

GetIsFetchable returns the IsFetchable field if non-nil, zero value otherwise.

### GetIsFetchableOk

`func (o *PrintJob) GetIsFetchableOk() (*bool, bool)`

GetIsFetchableOk returns a tuple with the IsFetchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFetchable

`func (o *PrintJob) SetIsFetchable(v bool)`

SetIsFetchable sets IsFetchable field to given value.

### HasIsFetchable

`func (o *PrintJob) HasIsFetchable() bool`

HasIsFetchable returns a boolean if a field has been set.

### GetRedirectedFrom

`func (o *PrintJob) GetRedirectedFrom() string`

GetRedirectedFrom returns the RedirectedFrom field if non-nil, zero value otherwise.

### GetRedirectedFromOk

`func (o *PrintJob) GetRedirectedFromOk() (*string, bool)`

GetRedirectedFromOk returns a tuple with the RedirectedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectedFrom

`func (o *PrintJob) SetRedirectedFrom(v string)`

SetRedirectedFrom sets RedirectedFrom field to given value.

### HasRedirectedFrom

`func (o *PrintJob) HasRedirectedFrom() bool`

HasRedirectedFrom returns a boolean if a field has been set.

### SetRedirectedFromNil

`func (o *PrintJob) SetRedirectedFromNil(b bool)`

 SetRedirectedFromNil sets the value for RedirectedFrom to be an explicit nil

### UnsetRedirectedFrom
`func (o *PrintJob) UnsetRedirectedFrom()`

UnsetRedirectedFrom ensures that no value is present for RedirectedFrom, not even an explicit nil
### GetRedirectedTo

`func (o *PrintJob) GetRedirectedTo() string`

GetRedirectedTo returns the RedirectedTo field if non-nil, zero value otherwise.

### GetRedirectedToOk

`func (o *PrintJob) GetRedirectedToOk() (*string, bool)`

GetRedirectedToOk returns a tuple with the RedirectedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectedTo

`func (o *PrintJob) SetRedirectedTo(v string)`

SetRedirectedTo sets RedirectedTo field to given value.

### HasRedirectedTo

`func (o *PrintJob) HasRedirectedTo() bool`

HasRedirectedTo returns a boolean if a field has been set.

### SetRedirectedToNil

`func (o *PrintJob) SetRedirectedToNil(b bool)`

 SetRedirectedToNil sets the value for RedirectedTo to be an explicit nil

### UnsetRedirectedTo
`func (o *PrintJob) UnsetRedirectedTo()`

UnsetRedirectedTo ensures that no value is present for RedirectedTo, not even an explicit nil
### GetStatus

`func (o *PrintJob) GetStatus() MicrosoftGraphPrintJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrintJob) GetStatusOk() (*MicrosoftGraphPrintJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrintJob) SetStatus(v MicrosoftGraphPrintJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrintJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDocuments

`func (o *PrintJob) GetDocuments() []MicrosoftGraphPrintDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *PrintJob) GetDocumentsOk() (*[]MicrosoftGraphPrintDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *PrintJob) SetDocuments(v []MicrosoftGraphPrintDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *PrintJob) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetTasks

`func (o *PrintJob) GetTasks() []MicrosoftGraphPrintTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *PrintJob) GetTasksOk() (*[]MicrosoftGraphPrintTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *PrintJob) SetTasks(v []MicrosoftGraphPrintTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *PrintJob) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


