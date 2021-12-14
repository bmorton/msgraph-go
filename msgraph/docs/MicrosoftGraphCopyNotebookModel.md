# MicrosoftGraphCopyNotebookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedByIdentity** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**CreatedTime** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**IsShared** | Pointer to **NullableBool** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 
**LastModifiedByIdentity** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) |  | [optional] 
**LastModifiedTime** | Pointer to **NullableTime** |  | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphNotebookLinks**](anyOf&lt;microsoft.graph.notebookLinks&gt;.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SectionGroupsUrl** | Pointer to **NullableString** |  | [optional] 
**SectionsUrl** | Pointer to **NullableString** |  | [optional] 
**Self** | Pointer to **NullableString** |  | [optional] 
**UserRole** | Pointer to [**AnyOfmicrosoftGraphOnenoteUserRole**](anyOf&lt;microsoft.graph.onenoteUserRole&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphCopyNotebookModel

`func NewMicrosoftGraphCopyNotebookModel() *MicrosoftGraphCopyNotebookModel`

NewMicrosoftGraphCopyNotebookModel instantiates a new MicrosoftGraphCopyNotebookModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphCopyNotebookModelWithDefaults

`func NewMicrosoftGraphCopyNotebookModelWithDefaults() *MicrosoftGraphCopyNotebookModel`

NewMicrosoftGraphCopyNotebookModelWithDefaults instantiates a new MicrosoftGraphCopyNotebookModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphCopyNotebookModel) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByIdentity() AnyOfmicrosoftGraphIdentitySet`

GetCreatedByIdentity returns the CreatedByIdentity field if non-nil, zero value otherwise.

### GetCreatedByIdentityOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedByIdentityOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByIdentityOk returns a tuple with the CreatedByIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByIdentity(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedByIdentity sets CreatedByIdentity field to given value.

### HasCreatedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedByIdentity() bool`

HasCreatedByIdentity returns a boolean if a field has been set.

### SetCreatedByIdentityNil

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedByIdentityNil(b bool)`

 SetCreatedByIdentityNil sets the value for CreatedByIdentity to be an explicit nil

### UnsetCreatedByIdentity
`func (o *MicrosoftGraphCopyNotebookModel) UnsetCreatedByIdentity()`

UnsetCreatedByIdentity ensures that no value is present for CreatedByIdentity, not even an explicit nil
### GetCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *MicrosoftGraphCopyNotebookModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *MicrosoftGraphCopyNotebookModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *MicrosoftGraphCopyNotebookModel) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *MicrosoftGraphCopyNotebookModel) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetId

`func (o *MicrosoftGraphCopyNotebookModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphCopyNotebookModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphCopyNotebookModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphCopyNotebookModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphCopyNotebookModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphCopyNotebookModel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphCopyNotebookModel) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphCopyNotebookModel) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsShared

`func (o *MicrosoftGraphCopyNotebookModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *MicrosoftGraphCopyNotebookModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *MicrosoftGraphCopyNotebookModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *MicrosoftGraphCopyNotebookModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *MicrosoftGraphCopyNotebookModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *MicrosoftGraphCopyNotebookModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *MicrosoftGraphCopyNotebookModel) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByIdentity() AnyOfmicrosoftGraphIdentitySet`

GetLastModifiedByIdentity returns the LastModifiedByIdentity field if non-nil, zero value otherwise.

### GetLastModifiedByIdentityOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedByIdentityOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetLastModifiedByIdentityOk returns a tuple with the LastModifiedByIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByIdentity(v AnyOfmicrosoftGraphIdentitySet)`

SetLastModifiedByIdentity sets LastModifiedByIdentity field to given value.

### HasLastModifiedByIdentity

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedByIdentity() bool`

HasLastModifiedByIdentity returns a boolean if a field has been set.

### SetLastModifiedByIdentityNil

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedByIdentityNil(b bool)`

 SetLastModifiedByIdentityNil sets the value for LastModifiedByIdentity to be an explicit nil

### UnsetLastModifiedByIdentity
`func (o *MicrosoftGraphCopyNotebookModel) UnsetLastModifiedByIdentity()`

UnsetLastModifiedByIdentity ensures that no value is present for LastModifiedByIdentity, not even an explicit nil
### GetLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *MicrosoftGraphCopyNotebookModel) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### SetLastModifiedTimeNil

`func (o *MicrosoftGraphCopyNotebookModel) SetLastModifiedTimeNil(b bool)`

 SetLastModifiedTimeNil sets the value for LastModifiedTime to be an explicit nil

### UnsetLastModifiedTime
`func (o *MicrosoftGraphCopyNotebookModel) UnsetLastModifiedTime()`

UnsetLastModifiedTime ensures that no value is present for LastModifiedTime, not even an explicit nil
### GetLinks

`func (o *MicrosoftGraphCopyNotebookModel) GetLinks() AnyOfmicrosoftGraphNotebookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphCopyNotebookModel) GetLinksOk() (*AnyOfmicrosoftGraphNotebookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MicrosoftGraphCopyNotebookModel) SetLinks(v AnyOfmicrosoftGraphNotebookLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MicrosoftGraphCopyNotebookModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MicrosoftGraphCopyNotebookModel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MicrosoftGraphCopyNotebookModel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetName

`func (o *MicrosoftGraphCopyNotebookModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphCopyNotebookModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphCopyNotebookModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphCopyNotebookModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphCopyNotebookModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphCopyNotebookModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionGroupsUrl() string`

GetSectionGroupsUrl returns the SectionGroupsUrl field if non-nil, zero value otherwise.

### GetSectionGroupsUrlOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionGroupsUrlOk() (*string, bool)`

GetSectionGroupsUrlOk returns a tuple with the SectionGroupsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionGroupsUrl(v string)`

SetSectionGroupsUrl sets SectionGroupsUrl field to given value.

### HasSectionGroupsUrl

`func (o *MicrosoftGraphCopyNotebookModel) HasSectionGroupsUrl() bool`

HasSectionGroupsUrl returns a boolean if a field has been set.

### SetSectionGroupsUrlNil

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionGroupsUrlNil(b bool)`

 SetSectionGroupsUrlNil sets the value for SectionGroupsUrl to be an explicit nil

### UnsetSectionGroupsUrl
`func (o *MicrosoftGraphCopyNotebookModel) UnsetSectionGroupsUrl()`

UnsetSectionGroupsUrl ensures that no value is present for SectionGroupsUrl, not even an explicit nil
### GetSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionsUrl() string`

GetSectionsUrl returns the SectionsUrl field if non-nil, zero value otherwise.

### GetSectionsUrlOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSectionsUrlOk() (*string, bool)`

GetSectionsUrlOk returns a tuple with the SectionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionsUrl(v string)`

SetSectionsUrl sets SectionsUrl field to given value.

### HasSectionsUrl

`func (o *MicrosoftGraphCopyNotebookModel) HasSectionsUrl() bool`

HasSectionsUrl returns a boolean if a field has been set.

### SetSectionsUrlNil

`func (o *MicrosoftGraphCopyNotebookModel) SetSectionsUrlNil(b bool)`

 SetSectionsUrlNil sets the value for SectionsUrl to be an explicit nil

### UnsetSectionsUrl
`func (o *MicrosoftGraphCopyNotebookModel) UnsetSectionsUrl()`

UnsetSectionsUrl ensures that no value is present for SectionsUrl, not even an explicit nil
### GetSelf

`func (o *MicrosoftGraphCopyNotebookModel) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphCopyNotebookModel) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MicrosoftGraphCopyNotebookModel) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MicrosoftGraphCopyNotebookModel) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *MicrosoftGraphCopyNotebookModel) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *MicrosoftGraphCopyNotebookModel) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetUserRole

`func (o *MicrosoftGraphCopyNotebookModel) GetUserRole() AnyOfmicrosoftGraphOnenoteUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *MicrosoftGraphCopyNotebookModel) GetUserRoleOk() (*AnyOfmicrosoftGraphOnenoteUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *MicrosoftGraphCopyNotebookModel) SetUserRole(v AnyOfmicrosoftGraphOnenoteUserRole)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *MicrosoftGraphCopyNotebookModel) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *MicrosoftGraphCopyNotebookModel) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *MicrosoftGraphCopyNotebookModel) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


