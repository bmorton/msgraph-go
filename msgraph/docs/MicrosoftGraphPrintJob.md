# MicrosoftGraphPrintJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
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

### NewMicrosoftGraphPrintJob

`func NewMicrosoftGraphPrintJob() *MicrosoftGraphPrintJob`

NewMicrosoftGraphPrintJob instantiates a new MicrosoftGraphPrintJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintJobWithDefaults

`func NewMicrosoftGraphPrintJobWithDefaults() *MicrosoftGraphPrintJob`

NewMicrosoftGraphPrintJobWithDefaults instantiates a new MicrosoftGraphPrintJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfiguration

`func (o *MicrosoftGraphPrintJob) GetConfiguration() MicrosoftGraphPrintJobConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MicrosoftGraphPrintJob) GetConfigurationOk() (*MicrosoftGraphPrintJobConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MicrosoftGraphPrintJob) SetConfiguration(v MicrosoftGraphPrintJobConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MicrosoftGraphPrintJob) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MicrosoftGraphPrintJob) GetCreatedBy() AnyOfmicrosoftGraphUserIdentity`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphPrintJob) GetCreatedByOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphPrintJob) SetCreatedBy(v AnyOfmicrosoftGraphUserIdentity)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphPrintJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphPrintJob) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphPrintJob) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphPrintJob) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPrintJob) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPrintJob) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphPrintJob) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetIsFetchable

`func (o *MicrosoftGraphPrintJob) GetIsFetchable() bool`

GetIsFetchable returns the IsFetchable field if non-nil, zero value otherwise.

### GetIsFetchableOk

`func (o *MicrosoftGraphPrintJob) GetIsFetchableOk() (*bool, bool)`

GetIsFetchableOk returns a tuple with the IsFetchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFetchable

`func (o *MicrosoftGraphPrintJob) SetIsFetchable(v bool)`

SetIsFetchable sets IsFetchable field to given value.

### HasIsFetchable

`func (o *MicrosoftGraphPrintJob) HasIsFetchable() bool`

HasIsFetchable returns a boolean if a field has been set.

### GetRedirectedFrom

`func (o *MicrosoftGraphPrintJob) GetRedirectedFrom() string`

GetRedirectedFrom returns the RedirectedFrom field if non-nil, zero value otherwise.

### GetRedirectedFromOk

`func (o *MicrosoftGraphPrintJob) GetRedirectedFromOk() (*string, bool)`

GetRedirectedFromOk returns a tuple with the RedirectedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectedFrom

`func (o *MicrosoftGraphPrintJob) SetRedirectedFrom(v string)`

SetRedirectedFrom sets RedirectedFrom field to given value.

### HasRedirectedFrom

`func (o *MicrosoftGraphPrintJob) HasRedirectedFrom() bool`

HasRedirectedFrom returns a boolean if a field has been set.

### SetRedirectedFromNil

`func (o *MicrosoftGraphPrintJob) SetRedirectedFromNil(b bool)`

 SetRedirectedFromNil sets the value for RedirectedFrom to be an explicit nil

### UnsetRedirectedFrom
`func (o *MicrosoftGraphPrintJob) UnsetRedirectedFrom()`

UnsetRedirectedFrom ensures that no value is present for RedirectedFrom, not even an explicit nil
### GetRedirectedTo

`func (o *MicrosoftGraphPrintJob) GetRedirectedTo() string`

GetRedirectedTo returns the RedirectedTo field if non-nil, zero value otherwise.

### GetRedirectedToOk

`func (o *MicrosoftGraphPrintJob) GetRedirectedToOk() (*string, bool)`

GetRedirectedToOk returns a tuple with the RedirectedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectedTo

`func (o *MicrosoftGraphPrintJob) SetRedirectedTo(v string)`

SetRedirectedTo sets RedirectedTo field to given value.

### HasRedirectedTo

`func (o *MicrosoftGraphPrintJob) HasRedirectedTo() bool`

HasRedirectedTo returns a boolean if a field has been set.

### SetRedirectedToNil

`func (o *MicrosoftGraphPrintJob) SetRedirectedToNil(b bool)`

 SetRedirectedToNil sets the value for RedirectedTo to be an explicit nil

### UnsetRedirectedTo
`func (o *MicrosoftGraphPrintJob) UnsetRedirectedTo()`

UnsetRedirectedTo ensures that no value is present for RedirectedTo, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphPrintJob) GetStatus() MicrosoftGraphPrintJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphPrintJob) GetStatusOk() (*MicrosoftGraphPrintJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphPrintJob) SetStatus(v MicrosoftGraphPrintJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphPrintJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDocuments

`func (o *MicrosoftGraphPrintJob) GetDocuments() []MicrosoftGraphPrintDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *MicrosoftGraphPrintJob) GetDocumentsOk() (*[]MicrosoftGraphPrintDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *MicrosoftGraphPrintJob) SetDocuments(v []MicrosoftGraphPrintDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *MicrosoftGraphPrintJob) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetTasks

`func (o *MicrosoftGraphPrintJob) GetTasks() []MicrosoftGraphPrintTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MicrosoftGraphPrintJob) GetTasksOk() (*[]MicrosoftGraphPrintTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MicrosoftGraphPrintJob) SetTasks(v []MicrosoftGraphPrintTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MicrosoftGraphPrintJob) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


