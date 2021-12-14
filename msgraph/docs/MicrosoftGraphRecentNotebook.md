# MicrosoftGraphRecentNotebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The name of the notebook. | [optional] 
**LastAccessedTime** | Pointer to **NullableTime** | The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphRecentNotebookLinks**](anyOf&lt;microsoft.graph.recentNotebookLinks&gt;.md) | Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it&#39;s installed. The oneNoteWebURL link opens the notebook in OneNote on the web. | [optional] 
**SourceService** | Pointer to [**AnyOfmicrosoftGraphOnenoteSourceService**](anyOf&lt;microsoft.graph.onenoteSourceService&gt;.md) | The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive. | [optional] 

## Methods

### NewMicrosoftGraphRecentNotebook

`func NewMicrosoftGraphRecentNotebook() *MicrosoftGraphRecentNotebook`

NewMicrosoftGraphRecentNotebook instantiates a new MicrosoftGraphRecentNotebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecentNotebookWithDefaults

`func NewMicrosoftGraphRecentNotebookWithDefaults() *MicrosoftGraphRecentNotebook`

NewMicrosoftGraphRecentNotebookWithDefaults instantiates a new MicrosoftGraphRecentNotebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphRecentNotebook) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRecentNotebook) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRecentNotebook) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRecentNotebook) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphRecentNotebook) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphRecentNotebook) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTime() time.Time`

GetLastAccessedTime returns the LastAccessedTime field if non-nil, zero value otherwise.

### GetLastAccessedTimeOk

`func (o *MicrosoftGraphRecentNotebook) GetLastAccessedTimeOk() (*time.Time, bool)`

GetLastAccessedTimeOk returns a tuple with the LastAccessedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTime(v time.Time)`

SetLastAccessedTime sets LastAccessedTime field to given value.

### HasLastAccessedTime

`func (o *MicrosoftGraphRecentNotebook) HasLastAccessedTime() bool`

HasLastAccessedTime returns a boolean if a field has been set.

### SetLastAccessedTimeNil

`func (o *MicrosoftGraphRecentNotebook) SetLastAccessedTimeNil(b bool)`

 SetLastAccessedTimeNil sets the value for LastAccessedTime to be an explicit nil

### UnsetLastAccessedTime
`func (o *MicrosoftGraphRecentNotebook) UnsetLastAccessedTime()`

UnsetLastAccessedTime ensures that no value is present for LastAccessedTime, not even an explicit nil
### GetLinks

`func (o *MicrosoftGraphRecentNotebook) GetLinks() AnyOfmicrosoftGraphRecentNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphRecentNotebook) GetLinksOk() (*AnyOfmicrosoftGraphRecentNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MicrosoftGraphRecentNotebook) SetLinks(v AnyOfmicrosoftGraphRecentNotebookLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MicrosoftGraphRecentNotebook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MicrosoftGraphRecentNotebook) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MicrosoftGraphRecentNotebook) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSourceService

`func (o *MicrosoftGraphRecentNotebook) GetSourceService() AnyOfmicrosoftGraphOnenoteSourceService`

GetSourceService returns the SourceService field if non-nil, zero value otherwise.

### GetSourceServiceOk

`func (o *MicrosoftGraphRecentNotebook) GetSourceServiceOk() (*AnyOfmicrosoftGraphOnenoteSourceService, bool)`

GetSourceServiceOk returns a tuple with the SourceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceService

`func (o *MicrosoftGraphRecentNotebook) SetSourceService(v AnyOfmicrosoftGraphOnenoteSourceService)`

SetSourceService sets SourceService field to given value.

### HasSourceService

`func (o *MicrosoftGraphRecentNotebook) HasSourceService() bool`

HasSourceService returns a boolean if a field has been set.

### SetSourceServiceNil

`func (o *MicrosoftGraphRecentNotebook) SetSourceServiceNil(b bool)`

 SetSourceServiceNil sets the value for SourceService to be an explicit nil

### UnsetSourceService
`func (o *MicrosoftGraphRecentNotebook) UnsetSourceService()`

UnsetSourceService ensures that no value is present for SourceService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


