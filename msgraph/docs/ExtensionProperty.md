# ExtensionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDisplayName** | Pointer to **NullableString** | Display name of the application object on which this extension property is defined. Read-only. | [optional] 
**DataType** | Pointer to **string** | Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum | [optional] 
**IsSyncedFromOnPremises** | Pointer to **NullableBool** | Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only. | [optional] 
**Name** | Pointer to **string** | Name of the extension property. Not nullable. | [optional] 
**TargetObjects** | Pointer to **[]string** | Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication | [optional] 

## Methods

### NewExtensionProperty

`func NewExtensionProperty() *ExtensionProperty`

NewExtensionProperty instantiates a new ExtensionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionPropertyWithDefaults

`func NewExtensionPropertyWithDefaults() *ExtensionProperty`

NewExtensionPropertyWithDefaults instantiates a new ExtensionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDisplayName

`func (o *ExtensionProperty) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *ExtensionProperty) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *ExtensionProperty) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *ExtensionProperty) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *ExtensionProperty) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *ExtensionProperty) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetDataType

`func (o *ExtensionProperty) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ExtensionProperty) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ExtensionProperty) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ExtensionProperty) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsSyncedFromOnPremises

`func (o *ExtensionProperty) GetIsSyncedFromOnPremises() bool`

GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field if non-nil, zero value otherwise.

### GetIsSyncedFromOnPremisesOk

`func (o *ExtensionProperty) GetIsSyncedFromOnPremisesOk() (*bool, bool)`

GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncedFromOnPremises

`func (o *ExtensionProperty) SetIsSyncedFromOnPremises(v bool)`

SetIsSyncedFromOnPremises sets IsSyncedFromOnPremises field to given value.

### HasIsSyncedFromOnPremises

`func (o *ExtensionProperty) HasIsSyncedFromOnPremises() bool`

HasIsSyncedFromOnPremises returns a boolean if a field has been set.

### SetIsSyncedFromOnPremisesNil

`func (o *ExtensionProperty) SetIsSyncedFromOnPremisesNil(b bool)`

 SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil

### UnsetIsSyncedFromOnPremises
`func (o *ExtensionProperty) UnsetIsSyncedFromOnPremises()`

UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
### GetName

`func (o *ExtensionProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtensionProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtensionProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtensionProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetObjects

`func (o *ExtensionProperty) GetTargetObjects() []string`

GetTargetObjects returns the TargetObjects field if non-nil, zero value otherwise.

### GetTargetObjectsOk

`func (o *ExtensionProperty) GetTargetObjectsOk() (*[]string, bool)`

GetTargetObjectsOk returns a tuple with the TargetObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjects

`func (o *ExtensionProperty) SetTargetObjects(v []string)`

SetTargetObjects sets TargetObjects field to given value.

### HasTargetObjects

`func (o *ExtensionProperty) HasTargetObjects() bool`

HasTargetObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


