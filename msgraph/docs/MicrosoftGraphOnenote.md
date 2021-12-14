# MicrosoftGraphOnenote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Notebooks** | Pointer to [**[]MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md) | The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md) | The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable. | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md) | The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 
**Resources** | Pointer to [**[]MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md) | The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable. | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md) | The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md) | The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphOnenote

`func NewMicrosoftGraphOnenote() *MicrosoftGraphOnenote`

NewMicrosoftGraphOnenote instantiates a new MicrosoftGraphOnenote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenoteWithDefaults

`func NewMicrosoftGraphOnenoteWithDefaults() *MicrosoftGraphOnenote`

NewMicrosoftGraphOnenoteWithDefaults instantiates a new MicrosoftGraphOnenote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOnenote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOnenote) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOnenote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotebooks

`func (o *MicrosoftGraphOnenote) GetNotebooks() []MicrosoftGraphNotebook`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *MicrosoftGraphOnenote) GetNotebooksOk() (*[]MicrosoftGraphNotebook, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *MicrosoftGraphOnenote) SetNotebooks(v []MicrosoftGraphNotebook)`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *MicrosoftGraphOnenote) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.

### GetOperations

`func (o *MicrosoftGraphOnenote) GetOperations() []MicrosoftGraphOnenoteOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *MicrosoftGraphOnenote) GetOperationsOk() (*[]MicrosoftGraphOnenoteOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *MicrosoftGraphOnenote) SetOperations(v []MicrosoftGraphOnenoteOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *MicrosoftGraphOnenote) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPages

`func (o *MicrosoftGraphOnenote) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *MicrosoftGraphOnenote) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *MicrosoftGraphOnenote) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *MicrosoftGraphOnenote) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResources

`func (o *MicrosoftGraphOnenote) GetResources() []MicrosoftGraphOnenoteResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *MicrosoftGraphOnenote) GetResourcesOk() (*[]MicrosoftGraphOnenoteResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *MicrosoftGraphOnenote) SetResources(v []MicrosoftGraphOnenoteResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *MicrosoftGraphOnenote) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSectionGroups

`func (o *MicrosoftGraphOnenote) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *MicrosoftGraphOnenote) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroups

`func (o *MicrosoftGraphOnenote) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups sets SectionGroups field to given value.

### HasSectionGroups

`func (o *MicrosoftGraphOnenote) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### GetSections

`func (o *MicrosoftGraphOnenote) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *MicrosoftGraphOnenote) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *MicrosoftGraphOnenote) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *MicrosoftGraphOnenote) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


