# MicrosoftGraphDriveRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **NullableString** | The alias of the domain object, for cases where an email address is unavailable (e.g. security groups). | [optional] 
**Email** | Pointer to **NullableString** | The email address for the recipient, if the recipient has an associated email address. | [optional] 
**ObjectId** | Pointer to **NullableString** | The unique identifier for the recipient in the directory. | [optional] 

## Methods

### NewMicrosoftGraphDriveRecipient

`func NewMicrosoftGraphDriveRecipient() *MicrosoftGraphDriveRecipient`

NewMicrosoftGraphDriveRecipient instantiates a new MicrosoftGraphDriveRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDriveRecipientWithDefaults

`func NewMicrosoftGraphDriveRecipientWithDefaults() *MicrosoftGraphDriveRecipient`

NewMicrosoftGraphDriveRecipientWithDefaults instantiates a new MicrosoftGraphDriveRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *MicrosoftGraphDriveRecipient) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *MicrosoftGraphDriveRecipient) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *MicrosoftGraphDriveRecipient) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *MicrosoftGraphDriveRecipient) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *MicrosoftGraphDriveRecipient) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *MicrosoftGraphDriveRecipient) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetEmail

`func (o *MicrosoftGraphDriveRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MicrosoftGraphDriveRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MicrosoftGraphDriveRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MicrosoftGraphDriveRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MicrosoftGraphDriveRecipient) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MicrosoftGraphDriveRecipient) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetObjectId

`func (o *MicrosoftGraphDriveRecipient) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MicrosoftGraphDriveRecipient) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MicrosoftGraphDriveRecipient) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MicrosoftGraphDriveRecipient) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *MicrosoftGraphDriveRecipient) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *MicrosoftGraphDriveRecipient) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


