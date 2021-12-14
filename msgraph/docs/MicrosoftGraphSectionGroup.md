# MicrosoftGraphSectionGroup

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
**SectionGroupsUrl** | Pointer to **NullableString** | The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only. | [optional] 
**SectionsUrl** | Pointer to **NullableString** | The URL for the sections navigation property, which returns all the sections in the section group. Read-only. | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) | The notebook that contains the section group. Read-only. | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) | The section group that contains the section group. Read-only. | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md) | The section groups in the section. Read-only. Nullable. | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md) | The sections in the section group. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphSectionGroup

`func NewMicrosoftGraphSectionGroup() *MicrosoftGraphSectionGroup`

NewMicrosoftGraphSectionGroup instantiates a new MicrosoftGraphSectionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSectionGroupWithDefaults

`func NewMicrosoftGraphSectionGroupWithDefaults() *MicrosoftGraphSectionGroup`

NewMicrosoftGraphSectionGroupWithDefaults instantiates a new MicrosoftGraphSectionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphSectionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphSectionGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphSectionGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphSectionGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSelf

`func (o *MicrosoftGraphSectionGroup) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphSectionGroup) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MicrosoftGraphSectionGroup) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MicrosoftGraphSectionGroup) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *MicrosoftGraphSectionGroup) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *MicrosoftGraphSectionGroup) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphSectionGroup) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphSectionGroup) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphSectionGroup) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphSectionGroup) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphSectionGroup) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphSectionGroup) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphSectionGroup) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphSectionGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphSectionGroup) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphSectionGroup) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphSectionGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphSectionGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphSectionGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphSectionGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphSectionGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphSectionGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphSectionGroup) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphSectionGroup) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphSectionGroup) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphSectionGroup) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphSectionGroup) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphSectionGroup) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsUrlOk() (*string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl sets SectionGroupsUrl field to given value.

### HasSectionGroupsUrl

`func (o *MicrosoftGraphSectionGroup) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrlNil

`func (o *MicrosoftGraphSectionGroup) SetSectionGroupsUrlNil(b bool)`

 SetSectionGroupsUrlNil sets the value for SectionGroupsUrl to be an explicit nil

### UnsetSectionGroupsUrl
`func (o *MicrosoftGraphSectionGroup) UnsetSectionGroupsUrl()`

UnsetSectionGroupsUrl ensures that no value is present for SectionGroupsUrl, not even an explicit nil
### GetSectionsUrl

`func (o *MicrosoftGraphSectionGroup) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *MicrosoftGraphSectionGroup) GetSectionsUrlOk() (*string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionsUrl

`func (o *MicrosoftGraphSectionGroup) SetSectionsUrl(v string)`

SetSectionsUrl sets SectionsUrl field to given value.

### HasSectionsUrl

`func (o *MicrosoftGraphSectionGroup) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrlNil

`func (o *MicrosoftGraphSectionGroup) SetSectionsUrlNil(b bool)`

 SetSectionsUrlNil sets the value for SectionsUrl to be an explicit nil

### UnsetSectionsUrl
`func (o *MicrosoftGraphSectionGroup) UnsetSectionsUrl()`

UnsetSectionsUrl ensures that no value is present for SectionsUrl, not even an explicit nil
### GetParentNotebook

`func (o *MicrosoftGraphSectionGroup) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphSectionGroup) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *MicrosoftGraphSectionGroup) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *MicrosoftGraphSectionGroup) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *MicrosoftGraphSectionGroup) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *MicrosoftGraphSectionGroup) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *MicrosoftGraphSectionGroup) GetParentSectionGroupOk() (*AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup sets ParentSectionGroup field to given value.

### HasParentSectionGroup

`func (o *MicrosoftGraphSectionGroup) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroupNil

`func (o *MicrosoftGraphSectionGroup) SetParentSectionGroupNil(b bool)`

 SetParentSectionGroupNil sets the value for ParentSectionGroup to be an explicit nil

### UnsetParentSectionGroup
`func (o *MicrosoftGraphSectionGroup) UnsetParentSectionGroup()`

UnsetParentSectionGroup ensures that no value is present for ParentSectionGroup, not even an explicit nil
### GetSectionGroups

`func (o *MicrosoftGraphSectionGroup) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *MicrosoftGraphSectionGroup) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroups

`func (o *MicrosoftGraphSectionGroup) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups sets SectionGroups field to given value.

### HasSectionGroups

`func (o *MicrosoftGraphSectionGroup) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### GetSections

`func (o *MicrosoftGraphSectionGroup) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *MicrosoftGraphSectionGroup) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *MicrosoftGraphSectionGroup) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *MicrosoftGraphSectionGroup) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


