# ManagedMobileApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileAppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The identifier for an app with it&#39;s operating system type. | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 

## Methods

### NewManagedMobileApp

`func NewManagedMobileApp() *ManagedMobileApp`

NewManagedMobileApp instantiates a new ManagedMobileApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedMobileAppWithDefaults

`func NewManagedMobileAppWithDefaults() *ManagedMobileApp`

NewManagedMobileAppWithDefaults instantiates a new ManagedMobileApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileAppIdentifier

`func (o *ManagedMobileApp) GetMobileAppIdentifier() AnyOfobject`

GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.

### GetMobileAppIdentifierOk

`func (o *ManagedMobileApp) GetMobileAppIdentifierOk() (*AnyOfobject, bool)`

GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppIdentifier

`func (o *ManagedMobileApp) SetMobileAppIdentifier(v AnyOfobject)`

SetMobileAppIdentifier sets MobileAppIdentifier field to given value.

### HasMobileAppIdentifier

`func (o *ManagedMobileApp) HasMobileAppIdentifier() bool`

HasMobileAppIdentifier returns a boolean if a field has been set.

### SetMobileAppIdentifierNil

`func (o *ManagedMobileApp) SetMobileAppIdentifierNil(b bool)`

 SetMobileAppIdentifierNil sets the value for MobileAppIdentifier to be an explicit nil

### UnsetMobileAppIdentifier
`func (o *ManagedMobileApp) UnsetMobileAppIdentifier()`

UnsetMobileAppIdentifier ensures that no value is present for MobileAppIdentifier, not even an explicit nil
### GetVersion

`func (o *ManagedMobileApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedMobileApp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedMobileApp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagedMobileApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ManagedMobileApp) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ManagedMobileApp) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


