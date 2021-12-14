# MicrosoftGraphAccessPackageCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CatalogType** | Pointer to [**AnyOfmicrosoftGraphAccessPackageCatalogType**](anyOf&lt;microsoft.graph.accessPackageCatalogType&gt;.md) | Whether the catalog is created by a user or entitlement management. The possible values are: userManaged, serviceDefault, serviceManaged, unknownFutureValue. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Description** | Pointer to **NullableString** | The description of the access package catalog. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the access package catalog. | [optional] 
**IsExternallyVisible** | Pointer to **NullableBool** | Whether the access packages in this catalog can be requested by users outside of the tenant. | [optional] 
**ModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAccessPackageCatalogState**](anyOf&lt;microsoft.graph.accessPackageCatalogState&gt;.md) | Has the value published if the access packages are available for management. The possible values are: unpublished, published, unknownFutureValue. | [optional] 
**AccessPackages** | Pointer to [**[]MicrosoftGraphAccessPackage**](MicrosoftGraphAccessPackage.md) | The access packages in this catalog. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphAccessPackageCatalog

`func NewMicrosoftGraphAccessPackageCatalog() *MicrosoftGraphAccessPackageCatalog`

NewMicrosoftGraphAccessPackageCatalog instantiates a new MicrosoftGraphAccessPackageCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessPackageCatalogWithDefaults

`func NewMicrosoftGraphAccessPackageCatalogWithDefaults() *MicrosoftGraphAccessPackageCatalog`

NewMicrosoftGraphAccessPackageCatalogWithDefaults instantiates a new MicrosoftGraphAccessPackageCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAccessPackageCatalog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAccessPackageCatalog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAccessPackageCatalog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCatalogType

`func (o *MicrosoftGraphAccessPackageCatalog) GetCatalogType() AnyOfmicrosoftGraphAccessPackageCatalogType`

GetCatalogType returns the CatalogType field if non-nil, zero value otherwise.

### GetCatalogTypeOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetCatalogTypeOk() (*AnyOfmicrosoftGraphAccessPackageCatalogType, bool)`

GetCatalogTypeOk returns a tuple with the CatalogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogType

`func (o *MicrosoftGraphAccessPackageCatalog) SetCatalogType(v AnyOfmicrosoftGraphAccessPackageCatalogType)`

SetCatalogType sets CatalogType field to given value.

### HasCatalogType

`func (o *MicrosoftGraphAccessPackageCatalog) HasCatalogType() bool`

HasCatalogType returns a boolean if a field has been set.

### SetCatalogTypeNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetCatalogTypeNil(b bool)`

 SetCatalogTypeNil sets the value for CatalogType to be an explicit nil

### UnsetCatalogType
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetCatalogType()`

UnsetCatalogType ensures that no value is present for CatalogType, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphAccessPackageCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphAccessPackageCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphAccessPackageCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphAccessPackageCatalog) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphAccessPackageCatalog) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphAccessPackageCatalog) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsExternallyVisible

`func (o *MicrosoftGraphAccessPackageCatalog) GetIsExternallyVisible() bool`

GetIsExternallyVisible returns the IsExternallyVisible field if non-nil, zero value otherwise.

### GetIsExternallyVisibleOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetIsExternallyVisibleOk() (*bool, bool)`

GetIsExternallyVisibleOk returns a tuple with the IsExternallyVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyVisible

`func (o *MicrosoftGraphAccessPackageCatalog) SetIsExternallyVisible(v bool)`

SetIsExternallyVisible sets IsExternallyVisible field to given value.

### HasIsExternallyVisible

`func (o *MicrosoftGraphAccessPackageCatalog) HasIsExternallyVisible() bool`

HasIsExternallyVisible returns a boolean if a field has been set.

### SetIsExternallyVisibleNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetIsExternallyVisibleNil(b bool)`

 SetIsExternallyVisibleNil sets the value for IsExternallyVisible to be an explicit nil

### UnsetIsExternallyVisible
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetIsExternallyVisible()`

UnsetIsExternallyVisible ensures that no value is present for IsExternallyVisible, not even an explicit nil
### GetModifiedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) GetModifiedDateTime() time.Time`

GetModifiedDateTime returns the ModifiedDateTime field if non-nil, zero value otherwise.

### GetModifiedDateTimeOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetModifiedDateTimeOk() (*time.Time, bool)`

GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) SetModifiedDateTime(v time.Time)`

SetModifiedDateTime sets ModifiedDateTime field to given value.

### HasModifiedDateTime

`func (o *MicrosoftGraphAccessPackageCatalog) HasModifiedDateTime() bool`

HasModifiedDateTime returns a boolean if a field has been set.

### SetModifiedDateTimeNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetModifiedDateTimeNil(b bool)`

 SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil

### UnsetModifiedDateTime
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetModifiedDateTime()`

UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
### GetState

`func (o *MicrosoftGraphAccessPackageCatalog) GetState() AnyOfmicrosoftGraphAccessPackageCatalogState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetStateOk() (*AnyOfmicrosoftGraphAccessPackageCatalogState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphAccessPackageCatalog) SetState(v AnyOfmicrosoftGraphAccessPackageCatalogState)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphAccessPackageCatalog) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphAccessPackageCatalog) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphAccessPackageCatalog) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetAccessPackages

`func (o *MicrosoftGraphAccessPackageCatalog) GetAccessPackages() []MicrosoftGraphAccessPackage`

GetAccessPackages returns the AccessPackages field if non-nil, zero value otherwise.

### GetAccessPackagesOk

`func (o *MicrosoftGraphAccessPackageCatalog) GetAccessPackagesOk() (*[]MicrosoftGraphAccessPackage, bool)`

GetAccessPackagesOk returns a tuple with the AccessPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackages

`func (o *MicrosoftGraphAccessPackageCatalog) SetAccessPackages(v []MicrosoftGraphAccessPackage)`

SetAccessPackages sets AccessPackages field to given value.

### HasAccessPackages

`func (o *MicrosoftGraphAccessPackageCatalog) HasAccessPackages() bool`

HasAccessPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


