# MicrosoftGraphOnenoteSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Self** | Pointer to **NullableString** | The endpoint where you can get details about the page. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, and application which created the item. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the notebook. | [optional] 
**LastModifiedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of the user, device, and application which created the item. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**IsDefault** | Pointer to **NullableBool** | Indicates whether this is the user&#39;s default section. Read-only. | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphSectionLinks**](anyOf&lt;microsoft.graph.sectionLinks&gt;.md) | Links for opening the section. The oneNoteClientURL link opens the section in the OneNote native client if it&#39;s installed. The oneNoteWebURL link opens the section in OneNote on the web. | [optional] 
**PagesUrl** | Pointer to **NullableString** | The pages endpoint where you can get details for all the pages in the section. Read-only. | [optional] 
**Pages** | Pointer to [**[]MicrosoftGraphOnenotePage**](MicrosoftGraphOnenotePage.md) | The collection of pages in the section.  Read-only. Nullable. | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) | The notebook that contains the section.  Read-only. | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) | The section group that contains the section.  Read-only. | [optional] 

## Methods

### NewMicrosoftGraphOnenoteSection

`func NewMicrosoftGraphOnenoteSection() *MicrosoftGraphOnenoteSection`

NewMicrosoftGraphOnenoteSection instantiates a new MicrosoftGraphOnenoteSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenoteSectionWithDefaults

`func NewMicrosoftGraphOnenoteSectionWithDefaults() *MicrosoftGraphOnenoteSection`

NewMicrosoftGraphOnenoteSectionWithDefaults instantiates a new MicrosoftGraphOnenoteSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOnenoteSection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenoteSection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOnenoteSection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOnenoteSection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSelf

`func (o *MicrosoftGraphOnenoteSection) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphOnenoteSection) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MicrosoftGraphOnenoteSection) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MicrosoftGraphOnenoteSection) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *MicrosoftGraphOnenoteSection) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *MicrosoftGraphOnenoteSection) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenoteSection) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenoteSection) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOnenoteSection) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOnenoteSection) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphOnenoteSection) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphOnenoteSection) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphOnenoteSection) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphOnenoteSection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphOnenoteSection) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphOnenoteSection) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphOnenoteSection) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOnenoteSection) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphOnenoteSection) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphOnenoteSection) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphOnenoteSection) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphOnenoteSection) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphOnenoteSection) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphOnenoteSection) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOnenoteSection) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOnenoteSection) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphOnenoteSection) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphOnenoteSection) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphOnenoteSection) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphOnenoteSection) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphOnenoteSection) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphOnenoteSection) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphOnenoteSection) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphOnenoteSection) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetLinks

`func (o *MicrosoftGraphOnenoteSection) GetLinks() AnyOfmicrosoftGraphSectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphOnenoteSection) GetLinksOk() (*AnyOfmicrosoftGraphSectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MicrosoftGraphOnenoteSection) SetLinks(v AnyOfmicrosoftGraphSectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MicrosoftGraphOnenoteSection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MicrosoftGraphOnenoteSection) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MicrosoftGraphOnenoteSection) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPagesUrl

`func (o *MicrosoftGraphOnenoteSection) GetPagesUrl() string`

GetPagesUrl returns the PagesUrl field if non-nil, zero value otherwise.

### GetPagesUrlOk

`func (o *MicrosoftGraphOnenoteSection) GetPagesUrlOk() (*string, bool)`

GetPagesUrlOk returns a tuple with the PagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesUrl

`func (o *MicrosoftGraphOnenoteSection) SetPagesUrl(v string)`

SetPagesUrl sets PagesUrl field to given value.

### HasPagesUrl

`func (o *MicrosoftGraphOnenoteSection) HasPagesUrl() bool`

HasPagesUrl returns a boolean if a field has been set.

### SetPagesUrlNil

`func (o *MicrosoftGraphOnenoteSection) SetPagesUrlNil(b bool)`

 SetPagesUrlNil sets the value for PagesUrl to be an explicit nil

### UnsetPagesUrl
`func (o *MicrosoftGraphOnenoteSection) UnsetPagesUrl()`

UnsetPagesUrl ensures that no value is present for PagesUrl, not even an explicit nil
### GetPages

`func (o *MicrosoftGraphOnenoteSection) GetPages() []MicrosoftGraphOnenotePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *MicrosoftGraphOnenoteSection) GetPagesOk() (*[]MicrosoftGraphOnenotePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *MicrosoftGraphOnenoteSection) SetPages(v []MicrosoftGraphOnenotePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *MicrosoftGraphOnenoteSection) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetParentNotebook

`func (o *MicrosoftGraphOnenoteSection) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphOnenoteSection) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *MicrosoftGraphOnenoteSection) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *MicrosoftGraphOnenoteSection) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *MicrosoftGraphOnenoteSection) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *MicrosoftGraphOnenoteSection) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *MicrosoftGraphOnenoteSection) GetParentSectionGroupOk() (*AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup sets ParentSectionGroup field to given value.

### HasParentSectionGroup

`func (o *MicrosoftGraphOnenoteSection) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroupNil

`func (o *MicrosoftGraphOnenoteSection) SetParentSectionGroupNil(b bool)`

 SetParentSectionGroupNil sets the value for ParentSectionGroup to be an explicit nil

### UnsetParentSectionGroup
`func (o *MicrosoftGraphOnenoteSection) UnsetParentSectionGroup()`

UnsetParentSectionGroup ensures that no value is present for ParentSectionGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


