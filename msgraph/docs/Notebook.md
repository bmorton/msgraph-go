# Notebook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | Pointer to **NullableBool** | Indicates whether this is the user&#39;s default notebook. Read-only. | [optional] 
**IsShared** | Pointer to **NullableBool** | Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only. | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphNotebookLinks**](anyOf&lt;microsoft.graph.notebookLinks&gt;.md) | Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it&#39;s installed. The oneNoteWebURL link opens the notebook in OneNote on the web. | [optional] 
**SectionGroupsUrl** | Pointer to **NullableString** | The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only. | [optional] 
**SectionsUrl** | Pointer to **NullableString** | The URL for the sections navigation property, which returns all the sections in the notebook. Read-only. | [optional] 
**UserRole** | Pointer to [**AnyOfmicrosoftGraphOnenoteUserRole**](anyOf&lt;microsoft.graph.onenoteUserRole&gt;.md) | Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only. | [optional] 
**SectionGroups** | Pointer to [**[]MicrosoftGraphSectionGroup**](MicrosoftGraphSectionGroup.md) | The section groups in the notebook. Read-only. Nullable. | [optional] 
**Sections** | Pointer to [**[]MicrosoftGraphOnenoteSection**](MicrosoftGraphOnenoteSection.md) | The sections in the notebook. Read-only. Nullable. | [optional] 

## Methods

### NewNotebook

`func NewNotebook() *Notebook`

NewNotebook instantiates a new Notebook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotebookWithDefaults

`func NewNotebookWithDefaults() *Notebook`

NewNotebookWithDefaults instantiates a new Notebook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDefault

`func (o *Notebook) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Notebook) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Notebook) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Notebook) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *Notebook) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *Notebook) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsShared

`func (o *Notebook) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *Notebook) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *Notebook) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *Notebook) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *Notebook) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *Notebook) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetLinks

`func (o *Notebook) GetLinks() AnyOfmicrosoftGraphNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Notebook) GetLinksOk() (*AnyOfmicrosoftGraphNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Notebook) SetLinks(v AnyOfmicrosoftGraphNotebookLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Notebook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *Notebook) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *Notebook) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSectionGroupsUrl

`func (o *Notebook) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *Notebook) GetSectionGroupsUrlOk() (*string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroupsUrl

`func (o *Notebook) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl sets SectionGroupsUrl field to given value.

### HasSectionGroupsUrl

`func (o *Notebook) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrlNil

`func (o *Notebook) SetSectionGroupsUrlNil(b bool)`

 SetSectionGroupsUrlNil sets the value for SectionGroupsUrl to be an explicit nil

### UnsetSectionGroupsUrl
`func (o *Notebook) UnsetSectionGroupsUrl()`

UnsetSectionGroupsUrl ensures that no value is present for SectionGroupsUrl, not even an explicit nil
### GetSectionsUrl

`func (o *Notebook) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *Notebook) GetSectionsUrlOk() (*string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionsUrl

`func (o *Notebook) SetSectionsUrl(v string)`

SetSectionsUrl sets SectionsUrl field to given value.

### HasSectionsUrl

`func (o *Notebook) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrlNil

`func (o *Notebook) SetSectionsUrlNil(b bool)`

 SetSectionsUrlNil sets the value for SectionsUrl to be an explicit nil

### UnsetSectionsUrl
`func (o *Notebook) UnsetSectionsUrl()`

UnsetSectionsUrl ensures that no value is present for SectionsUrl, not even an explicit nil
### GetUserRole

`func (o *Notebook) GetUserRole() AnyOfmicrosoftGraphOnenoteUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *Notebook) GetUserRoleOk() (*AnyOfmicrosoftGraphOnenoteUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *Notebook) SetUserRole(v AnyOfmicrosoftGraphOnenoteUserRole)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *Notebook) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *Notebook) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *Notebook) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetSectionGroups

`func (o *Notebook) GetSectionGroups() []MicrosoftGraphSectionGroup`

GetSectionGroups returns the SectionGroups field if non-nil, zero value otherwise.

### GetSectionGroupsOk

`func (o *Notebook) GetSectionGroupsOk() (*[]MicrosoftGraphSectionGroup, bool)`

GetSectionGroupsOk returns a tuple with the SectionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroups

`func (o *Notebook) SetSectionGroups(v []MicrosoftGraphSectionGroup)`

SetSectionGroups sets SectionGroups field to given value.

### HasSectionGroups

`func (o *Notebook) HasSectionGroups() bool`

HasSectionGroups returns a boolean if a field has been set.

### GetSections

`func (o *Notebook) GetSections() []MicrosoftGraphOnenoteSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Notebook) GetSectionsOk() (*[]MicrosoftGraphOnenoteSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Notebook) SetSections(v []MicrosoftGraphOnenoteSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Notebook) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


