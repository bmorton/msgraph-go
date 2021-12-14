# MicrosoftGraphProvisionedIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the identity. | [optional] 
**Details** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Details of the identity. | [optional] 
**IdentityType** | Pointer to **NullableString** | Type of identity that has been provisioned, such as &#39;user&#39; or &#39;group&#39;. | [optional] 

## Methods

### NewMicrosoftGraphProvisionedIdentity

`func NewMicrosoftGraphProvisionedIdentity() *MicrosoftGraphProvisionedIdentity`

NewMicrosoftGraphProvisionedIdentity instantiates a new MicrosoftGraphProvisionedIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisionedIdentityWithDefaults

`func NewMicrosoftGraphProvisionedIdentityWithDefaults() *MicrosoftGraphProvisionedIdentity`

NewMicrosoftGraphProvisionedIdentityWithDefaults instantiates a new MicrosoftGraphProvisionedIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphProvisionedIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphProvisionedIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphProvisionedIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphProvisionedIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphProvisionedIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphProvisionedIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphProvisionedIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphProvisionedIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphProvisionedIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphProvisionedIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphProvisionedIdentity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphProvisionedIdentity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphProvisionedIdentity) GetDetails() AnyOfobject`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphProvisionedIdentity) GetDetailsOk() (*AnyOfobject, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphProvisionedIdentity) SetDetails(v AnyOfobject)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphProvisionedIdentity) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *MicrosoftGraphProvisionedIdentity) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *MicrosoftGraphProvisionedIdentity) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetIdentityType

`func (o *MicrosoftGraphProvisionedIdentity) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *MicrosoftGraphProvisionedIdentity) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *MicrosoftGraphProvisionedIdentity) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *MicrosoftGraphProvisionedIdentity) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *MicrosoftGraphProvisionedIdentity) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *MicrosoftGraphProvisionedIdentity) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


