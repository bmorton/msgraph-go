# MicrosoftGraphManagedMobileApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**MobileAppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The identifier for an app with it&#39;s operating system type. | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 

## Methods

### NewMicrosoftGraphManagedMobileApp

`func NewMicrosoftGraphManagedMobileApp() *MicrosoftGraphManagedMobileApp`

NewMicrosoftGraphManagedMobileApp instantiates a new MicrosoftGraphManagedMobileApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphManagedMobileAppWithDefaults

`func NewMicrosoftGraphManagedMobileAppWithDefaults() *MicrosoftGraphManagedMobileApp`

NewMicrosoftGraphManagedMobileAppWithDefaults instantiates a new MicrosoftGraphManagedMobileApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphManagedMobileApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphManagedMobileApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphManagedMobileApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphManagedMobileApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMobileAppIdentifier

`func (o *MicrosoftGraphManagedMobileApp) GetMobileAppIdentifier() AnyOfobject`

GetMobileAppIdentifier returns the MobileAppIdentifier field if non-nil, zero value otherwise.

### GetMobileAppIdentifierOk

`func (o *MicrosoftGraphManagedMobileApp) GetMobileAppIdentifierOk() (*AnyOfobject, bool)`

GetMobileAppIdentifierOk returns a tuple with the MobileAppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileAppIdentifier

`func (o *MicrosoftGraphManagedMobileApp) SetMobileAppIdentifier(v AnyOfobject)`

SetMobileAppIdentifier sets MobileAppIdentifier field to given value.

### HasMobileAppIdentifier

`func (o *MicrosoftGraphManagedMobileApp) HasMobileAppIdentifier() bool`

HasMobileAppIdentifier returns a boolean if a field has been set.

### SetMobileAppIdentifierNil

`func (o *MicrosoftGraphManagedMobileApp) SetMobileAppIdentifierNil(b bool)`

 SetMobileAppIdentifierNil sets the value for MobileAppIdentifier to be an explicit nil

### UnsetMobileAppIdentifier
`func (o *MicrosoftGraphManagedMobileApp) UnsetMobileAppIdentifier()`

UnsetMobileAppIdentifier ensures that no value is present for MobileAppIdentifier, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphManagedMobileApp) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphManagedMobileApp) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphManagedMobileApp) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphManagedMobileApp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphManagedMobileApp) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphManagedMobileApp) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


