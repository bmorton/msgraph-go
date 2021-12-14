# SectionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionGroupsUrl** | Pointer to **NullableString** | The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only. | [optional] 
**SectionsUrl** | Pointer to **NullableString** | The URL for the sections navigation property, which returns all the sections in the section group. Read-only. | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) | The notebook that contains the section group. Read-only. | [optional] 
**ParentSectionGroup** | Pointer to [**AnyOfmicrosoftGraphSectionGroup**](anyOf&lt;microsoft.graph.sectionGroup&gt;.md) | The section group that contains the section group. Read-only. | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md) | The section groups in the section. Read-only. Nullable. | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md) | The sections in the section group. Read-only. Nullable. | [optional] 

## Methods

### NewSectionGroup

`func NewSectionGroup() *SectionGroup`

NewSectionGroup instantiates a new SectionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionGroupWithDefaults

`func NewSectionGroupWithDefaults() *SectionGroup`

NewSectionGroupWithDefaults instantiates a new SectionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionGroupsUrl

`func (o *SectionGroup) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *SectionGroup) GetSectionGroupsUrlOk() (*string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroupsUrl

`func (o *SectionGroup) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl sets SectionGroupsUrl field to given value.

### HasSectionGroupsUrl

`func (o *SectionGroup) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrlNil

`func (o *SectionGroup) SetSectionGroupsUrlNil(b bool)`

 SetSectionGroupsUrlNil sets the value for SectionGroupsUrl to be an explicit nil

### UnsetSectionGroupsUrl
`func (o *SectionGroup) UnsetSectionGroupsUrl()`

UnsetSectionGroupsUrl ensures that no value is present for SectionGroupsUrl, not even an explicit nil
### GetSectionsUrl

`func (o *SectionGroup) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *SectionGroup) GetSectionsUrlOk() (*string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionsUrl

`func (o *SectionGroup) SetSectionsUrl(v string)`

SetSectionsUrl sets SectionsUrl field to given value.

### HasSectionsUrl

`func (o *SectionGroup) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrlNil

`func (o *SectionGroup) SetSectionsUrlNil(b bool)`

 SetSectionsUrlNil sets the value for SectionsUrl to be an explicit nil

### UnsetSectionsUrl
`func (o *SectionGroup) UnsetSectionsUrl()`

UnsetSectionsUrl ensures that no value is present for SectionsUrl, not even an explicit nil
### GetParentNotebook

`func (o *SectionGroup) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *SectionGroup) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *SectionGroup) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *SectionGroup) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *SectionGroup) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *SectionGroup) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSectionGroup

`func (o *SectionGroup) GetParentSectionGroup() AnyOfmicrosoftGraphSectionGroup`

GetParentSectionGroup returns the ParentSectionGroup field if non-nil, zero value otherwise.

### GetParentSectionGroupOk

`func (o *SectionGroup) GetParentSectionGroupOk() (*AnyOfmicrosoftGraphSectionGroup, bool)`

GetParentSectionGroupOk returns a tuple with the ParentSectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSectionGroup

`func (o *SectionGroup) SetParentSectionGroup(v AnyOfmicrosoftGraphSectionGroup)`

SetParentSectionGroup sets ParentSectionGroup field to given value.

### HasParentSectionGroup

`func (o *SectionGroup) HasParentSectionGroup() bool`

HasParentSectionGroup returns a boolean if a field has been set.

### SetParentSectionGroupNil

`func (o *SectionGroup) SetParentSectionGroupNil(b bool)`

 SetParentSectionGroupNil sets the value for ParentSectionGroup to be an explicit nil

### UnsetParentSectionGroup
`func (o *SectionGroup) UnsetParentSectionGroup()`

UnsetParentSectionGroup ensures that no value is present for ParentSectionGroup, not even an explicit nil
### GetSectionGroups

`func (o *SectionGroup) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *SectionGroup) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroups

`func (o *SectionGroup) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups sets SectionGroups field to given value.

### HasSectionGroups

`func (o *SectionGroup) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### GetSections

`func (o *SectionGroup) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SectionGroup) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SectionGroup) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SectionGroup) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


