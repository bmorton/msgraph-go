# MicrosoftGraphProvisioningSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The identity&#39;s display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won&#39;t show up as having changed when using delta. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier for the identity. | [optional] 
**Details** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Details of the system. | [optional] 

## Methods

### NewMicrosoftGraphProvisioningSystem

`func NewMicrosoftGraphProvisioningSystem() *MicrosoftGraphProvisioningSystem`

NewMicrosoftGraphProvisioningSystem instantiates a new MicrosoftGraphProvisioningSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisioningSystemWithDefaults

`func NewMicrosoftGraphProvisioningSystemWithDefaults() *MicrosoftGraphProvisioningSystem`

NewMicrosoftGraphProvisioningSystemWithDefaults instantiates a new MicrosoftGraphProvisioningSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *MicrosoftGraphProvisioningSystem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphProvisioningSystem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphProvisioningSystem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphProvisioningSystem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphProvisioningSystem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphProvisioningSystem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetId

`func (o *MicrosoftGraphProvisioningSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphProvisioningSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphProvisioningSystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphProvisioningSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphProvisioningSystem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphProvisioningSystem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphProvisioningSystem) GetDetails() AnyOfobject`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphProvisioningSystem) GetDetailsOk() (*AnyOfobject, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphProvisioningSystem) SetDetails(v AnyOfobject)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphProvisioningSystem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *MicrosoftGraphProvisioningSystem) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *MicrosoftGraphProvisioningSystem) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


