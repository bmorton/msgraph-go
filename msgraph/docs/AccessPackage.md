# AccessPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | The description of the access package. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the access package. | [optional] 
**IsHidden** | Pointer to **NullableBool** | Whether the access package is hidden from the requestor. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Catalog** | Pointer to [**AnyOfmicrosoftGraphAccessPackageCatalog**](anyOf&lt;microsoft.graph.accessPackageCatalog&gt;.md) | Read-only. Nullable. | [optional] 

## Methods

### NewAccessPackage

`func NewAccessPackage() *AccessPackage`

NewAccessPackage instantiates a new AccessPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPackageWithDefaults

`func NewAccessPackageWithDefaults() *AccessPackage`

NewAccessPackageWithDefaults instantiates a new AccessPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *AccessPackage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AccessPackage) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AccessPackage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AccessPackage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *AccessPackage) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *AccessPackage) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *AccessPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessPackage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessPackage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *AccessPackage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccessPackage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccessPackage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccessPackage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AccessPackage) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AccessPackage) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsHidden

`func (o *AccessPackage) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AccessPackage) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AccessPackage) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AccessPackage) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### SetIsHiddenNil

`func (o *AccessPackage) SetIsHiddenNil(b bool)`

 SetIsHiddenNil sets the value for IsHidden to be an explicit nil

### UnsetIsHidden
`func (o *AccessPackage) UnsetIsHidden()`

UnsetIsHidden ensures that no value is present for IsHidden, not even an explicit nil
### GetModifiedDateTime

`func (o *AccessPackage) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *AccessPackage) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *AccessPackage) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *AccessPackage) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *AccessPackage) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *AccessPackage) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetCatalog

`func (o *AccessPackage) GetCatalog() AnyOfmicrosoftGraphAccessPackageCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *AccessPackage) GetCatalogOk() (*AnyOfmicrosoftGraphAccessPackageCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *AccessPackage) SetCatalog(v AnyOfmicrosoftGraphAccessPackageCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *AccessPackage) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *AccessPackage) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *AccessPackage) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


