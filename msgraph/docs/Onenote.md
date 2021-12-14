# Onenote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notebooks** | Pointer to [**[]MicrosoftGraphNotebook**](MicrosoftGraphNotebook.md) | The collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphOnenoteOperation**](MicrosoftGraphOnenoteOperation.md) | The status of OneNote operations. Getting an operations collection is not supported, but you can get the status of long-running operations if the Operation-Location header is returned in the response. Read-only. Nullable. | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md) | The pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 
**Resources** | Pointer to [**[]MicrosoftGraphOnenoteResource**](MicrosoftGraphOnenoteResource.md) | The image and other file resources in OneNote pages. Getting a resources collection is not supported, but you can get the binary content of a specific resource. Read-only. Nullable. | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md) | The section groups in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md) | The sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable. | [optional] 

## Methods

### NewOnenote

`func NewOnenote() *Onenote`

NewOnenote instantiates a new Onenote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnenoteWithDefaults

`func NewOnenoteWithDefaults() *Onenote`

NewOnenoteWithDefaults instantiates a new Onenote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotebooks

`func (o *Onenote) GetNotebooks() []MicrosoftGraphNotebook`

GetNotebooks returns the Notebooks field if non-nil, zero value otherwise.

### GetNotebooksOk

`func (o *Onenote) GetNotebooksOk() (*[]MicrosoftGraphNotebook, bool)`

GetNotebooksOk returns a tuple with the Notebooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotebooks

`func (o *Onenote) SetNotebooks(v []MicrosoftGraphNotebook)`

SetNotebooks sets Notebooks field to given value.

### HasNotebooks

`func (o *Onenote) HasNotebooks() bool`

HasNotebooks returns a boolean if a field has been set.

### GetOperations

`func (o *Onenote) GetOperations() []MicrosoftGraphOnenoteOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Onenote) GetOperationsOk() (*[]MicrosoftGraphOnenoteOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Onenote) SetOperations(v []MicrosoftGraphOnenoteOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *Onenote) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetPages

`func (o *Onenote) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *Onenote) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *Onenote) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *Onenote) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetResources

`func (o *Onenote) GetResources() []MicrosoftGraphOnenoteResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Onenote) GetResourcesOk() (*[]MicrosoftGraphOnenoteResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Onenote) SetResources(v []MicrosoftGraphOnenoteResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Onenote) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSectionGroups

`func (o *Onenote) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *Onenote) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroups

`func (o *Onenote) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups sets SectionGroups field to given value.

### HasSectionGroups

`func (o *Onenote) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### GetSections

`func (o *Onenote) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Onenote) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Onenote) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Onenote) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


