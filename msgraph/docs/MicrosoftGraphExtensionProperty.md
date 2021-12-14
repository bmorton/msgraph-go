# MicrosoftGraphExtensionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**AppDisplayName** | Pointer to **NullableString** | Display name of the application object on which this extension property is defined. Read-only. | [optional] 
**DataType** | Pointer to **string** | Specifies the data type of the value the extension property can hold. Following values are supported. Not nullable. Binary - 256 bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit value.LargeInteger - 64-bit value.String - 256 characters maximum | [optional] 
**IsSyncedFromOnPremises** | Pointer to **NullableBool** | Indicates if this extension property was sycned from onpremises directory using Azure AD Connect. Read-only. | [optional] 
**Name** | Pointer to **string** | Name of the extension property. Not nullable. | [optional] 
**TargetObjects** | Pointer to **[]string** | Following values are supported. Not nullable. UserGroupOrganizationDeviceApplication | [optional] 

## Methods

### NewMicrosoftGraphExtensionProperty

`func NewMicrosoftGraphExtensionProperty() *MicrosoftGraphExtensionProperty`

NewMicrosoftGraphExtensionProperty instantiates a new MicrosoftGraphExtensionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExtensionPropertyWithDefaults

`func NewMicrosoftGraphExtensionPropertyWithDefaults() *MicrosoftGraphExtensionProperty`

NewMicrosoftGraphExtensionPropertyWithDefaults instantiates a new MicrosoftGraphExtensionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphExtensionProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphExtensionProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphExtensionProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphExtensionProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphExtensionProperty) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphExtensionProperty) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphExtensionProperty) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphExtensionProperty) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphExtensionProperty) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphExtensionProperty) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAppDisplayName

`func (o *MicrosoftGraphExtensionProperty) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphExtensionProperty) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphExtensionProperty) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphExtensionProperty) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphExtensionProperty) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphExtensionProperty) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetDataType

`func (o *MicrosoftGraphExtensionProperty) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *MicrosoftGraphExtensionProperty) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *MicrosoftGraphExtensionProperty) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *MicrosoftGraphExtensionProperty) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetIsSyncedFromOnPremises

`func (o *MicrosoftGraphExtensionProperty) GetIsSyncedFromOnPremises() bool`

GetIsSyncedFromOnPremises returns the IsSyncedFromOnPremises field if non-nil, zero value otherwise.

### GetIsSyncedFromOnPremisesOk

`func (o *MicrosoftGraphExtensionProperty) GetIsSyncedFromOnPremisesOk() (*bool, bool)`

GetIsSyncedFromOnPremisesOk returns a tuple with the IsSyncedFromOnPremises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncedFromOnPremises

`func (o *MicrosoftGraphExtensionProperty) SetIsSyncedFromOnPremises(v bool)`

SetIsSyncedFromOnPremises sets IsSyncedFromOnPremises field to given value.

### HasIsSyncedFromOnPremises

`func (o *MicrosoftGraphExtensionProperty) HasIsSyncedFromOnPremises() bool`

HasIsSyncedFromOnPremises returns a boolean if a field has been set.

### SetIsSyncedFromOnPremisesNil

`func (o *MicrosoftGraphExtensionProperty) SetIsSyncedFromOnPremisesNil(b bool)`

 SetIsSyncedFromOnPremisesNil sets the value for IsSyncedFromOnPremises to be an explicit nil

### UnsetIsSyncedFromOnPremises
`func (o *MicrosoftGraphExtensionProperty) UnsetIsSyncedFromOnPremises()`

UnsetIsSyncedFromOnPremises ensures that no value is present for IsSyncedFromOnPremises, not even an explicit nil
### GetName

`func (o *MicrosoftGraphExtensionProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphExtensionProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphExtensionProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphExtensionProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetObjects

`func (o *MicrosoftGraphExtensionProperty) GetTargetObjects() []string`

GetTargetObjects returns the TargetObjects field if non-nil, zero value otherwise.

### GetTargetObjectsOk

`func (o *MicrosoftGraphExtensionProperty) GetTargetObjectsOk() (*[]string, bool)`

GetTargetObjectsOk returns a tuple with the TargetObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetObjects

`func (o *MicrosoftGraphExtensionProperty) SetTargetObjects(v []string)`

SetTargetObjects sets TargetObjects field to given value.

### HasTargetObjects

`func (o *MicrosoftGraphExtensionProperty) HasTargetObjects() bool`

HasTargetObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


