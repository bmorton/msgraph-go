# MicrosoftGraphContactFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The folder&#39;s display name. | [optional] 
**ParentFolderId** | Pointer to **NullableString** | The ID of the folder&#39;s parent folder. | [optional] 
**ChildFolders** | Pointer to [**[]MicrosoftGraphContactFolder**](MicrosoftGraphContactFolder.md) | The collection of child folders in the folder. Navigation property. Read-only. Nullable. | [optional] 
**Contacts** | Pointer to [**[]MicrosoftGraphContact**](MicrosoftGraphContact.md) | The contacts in the folder. Navigation property. Read-only. Nullable. | [optional] 
**MultiValueExtendedProperties** | Pointer to [**[]MicrosoftGraphMultiValueLegacyExtendedProperty**](MicrosoftGraphMultiValueLegacyExtendedProperty.md) | The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable. | [optional] 
**SingleValueExtendedProperties** | Pointer to [**[]MicrosoftGraphSingleValueLegacyExtendedProperty**](MicrosoftGraphSingleValueLegacyExtendedProperty.md) | The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphContactFolder

`func NewMicrosoftGraphContactFolder() *MicrosoftGraphContactFolder`

NewMicrosoftGraphContactFolder instantiates a new MicrosoftGraphContactFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphContactFolderWithDefaults

`func NewMicrosoftGraphContactFolderWithDefaults() *MicrosoftGraphContactFolder`

NewMicrosoftGraphContactFolderWithDefaults instantiates a new MicrosoftGraphContactFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphContactFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContactFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphContactFolder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphContactFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphContactFolder) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphContactFolder) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphContactFolder) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphContactFolder) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphContactFolder) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphContactFolder) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetParentFolderId

`func (o *MicrosoftGraphContactFolder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *MicrosoftGraphContactFolder) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *MicrosoftGraphContactFolder) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *MicrosoftGraphContactFolder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *MicrosoftGraphContactFolder) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *MicrosoftGraphContactFolder) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetChildFolders

`func (o *MicrosoftGraphContactFolder) GetChildFolders() []MicrosoftGraphContactFolder`

GetChildFolders returns the ChildFolders field if non-nil, zero value otherwise.

### GetChildFoldersOk

`func (o *MicrosoftGraphContactFolder) GetChildFoldersOk() (*[]MicrosoftGraphContactFolder, bool)`

GetChildFoldersOk returns a tuple with the ChildFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildFolders

`func (o *MicrosoftGraphContactFolder) SetChildFolders(v []MicrosoftGraphContactFolder)`

SetChildFolders sets ChildFolders field to given value.

### HasChildFolders

`func (o *MicrosoftGraphContactFolder) HasChildFolders() bool`

HasChildFolders returns a boolean if a field has been set.

### GetContacts

`func (o *MicrosoftGraphContactFolder) GetContacts() []MicrosoftGraphContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *MicrosoftGraphContactFolder) GetContactsOk() (*[]MicrosoftGraphContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *MicrosoftGraphContactFolder) SetContacts(v []MicrosoftGraphContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *MicrosoftGraphContactFolder) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty`

GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field if non-nil, zero value otherwise.

### GetMultiValueExtendedPropertiesOk

`func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool)`

GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty)`

SetMultiValueExtendedProperties sets MultiValueExtendedProperties field to given value.

### HasMultiValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) HasMultiValueExtendedProperties() bool`

HasMultiValueExtendedProperties returns a boolean if a field has been set.

### GetSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty`

GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field if non-nil, zero value otherwise.

### GetSingleValueExtendedPropertiesOk

`func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool)`

GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty)`

SetSingleValueExtendedProperties sets SingleValueExtendedProperties field to given value.

### HasSingleValueExtendedProperties

`func (o *MicrosoftGraphContactFolder) HasSingleValueExtendedProperties() bool`

HasSingleValueExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


