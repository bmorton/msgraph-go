# OnenoteSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **NullableBool** | Indicates whether this is the user&#39;s default section. Read-only. | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphSectionLinks**](anyOf&lt;microsoft.graph.sectionLinks&gt;.md) | Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it&#39;s installed. The oneNoteWebURL link opens the section in OneNote on the web. | [optional] 
**PagesUrl** | Pointer to **NullableString** | The pages endpoint where you can get details for all the pages in the section. Read-only. | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md) | The collection of pages in the section.  Read-only. Nullable. | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) | The notebook that contains the section.  Read-only. | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) | The section group that contains the section.  Read-only. | [optional] 

## Methods

### NewOnenoteSection

`func NewOnenoteSection() *OnenoteSection`

NewOnenoteSection instantiates a new OnenoteSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnenoteSectionWithDefaults

`func NewOnenoteSectionWithDefaults() *OnenoteSection`

NewOnenoteSectionWithDefaults instantiates a new OnenoteSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *OnenoteSection) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OnenoteSection) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OnenoteSection) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OnenoteSection) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *OnenoteSection) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *OnenoteSection) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetLinks

`func (o *OnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OnenoteSection) GetLinksOk() (*AnyOfmicrosoftGraphSectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OnenoteSection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *OnenoteSection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *OnenoteSection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPagesUrl

`func (o *OnenoteSection) GetPagesUrl() string`

GetPagesUrl returns the PagesUrl field if non-nil, zero value otherwise.

### GetPagesUrlOk

`func (o *OnenoteSection) GetPagesUrlOk() (*string, bool)`

GetPagesUrlOk returns a tuple with the PagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesUrl

`func (o *OnenoteSection) SetPagesUrl(v string)`

SetPagesUrl sets PagesUrl field to given value.

### HasPagesUrl

`func (o *OnenoteSection) HasPagesUrl() bool`

HasPagesUrl returns a boolean if a field has been set.

### SetPagesUrlNil

`func (o *OnenoteSection) SetPagesUrlNil(b bool)`

 SetPagesUrlNil sets the value for PagesUrl to be an explicit nil

### UnsetPagesUrl
`func (o *OnenoteSection) UnsetPagesUrl()`

UnsetPagesUrl ensures that no value is present for PagesUrl, not even an explicit nil
### GetPages

`func (o *OnenoteSection) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OnenoteSection) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OnenoteSection) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OnenoteSection) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetParentNotebook

`func (o *OnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *OnenoteSection) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *OnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *OnenoteSection) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *OnenoteSection) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *OnenoteSection) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSectionGroup

`func (o *OnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *OnenoteSection) GetParentSectionGroupOk() (*AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSectionGroup

`func (o *OnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup sets ParentSectionGroup field to given value.

### HasParentSectionGroup

`func (o *OnenoteSection) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroupNil

`func (o *OnenoteSection) SetParentSectionGroupNil(b bool)`

 SetParentSectionGroupNil sets the value for ParentSectionGroup to be an explicit nil

### UnsetParentSectionGroup
`func (o *OnenoteSection) UnsetParentSectionGroup()`

UnsetParentSectionGroup ensures that no value is present for ParentSectionGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


