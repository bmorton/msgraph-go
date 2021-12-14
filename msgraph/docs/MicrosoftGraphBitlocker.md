# MicrosoftGraphBitlocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**RecoveryKeys** | Pointer to [**[]MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md) | The recovery keys associated with the bitlocker entity. | [optional] 

## Methods

### NewMicrosoftGraphBitlocker

`func NewMicrosoftGraphBitlocker() *MicrosoftGraphBitlocker`

NewMicrosoftGraphBitlocker instantiates a new MicrosoftGraphBitlocker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBitlockerWithDefaults

`func NewMicrosoftGraphBitlockerWithDefaults() *MicrosoftGraphBitlocker`

NewMicrosoftGraphBitlockerWithDefaults instantiates a new MicrosoftGraphBitlocker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphBitlocker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphBitlocker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphBitlocker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphBitlocker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecoveryKeys

`func (o *MicrosoftGraphBitlocker) GetRecoveryKeys() []MicrosoftGraphBitlockerRecoveryKey`

GetRecoveryKeys returns the RecoveryKeys field if non-nil, zero value otherwise.

### GetRecoveryKeysOk

`func (o *MicrosoftGraphBitlocker) GetRecoveryKeysOk() (*[]MicrosoftGraphBitlockerRecoveryKey, bool)`

GetRecoveryKeysOk returns a tuple with the RecoveryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeys

`func (o *MicrosoftGraphBitlocker) SetRecoveryKeys(v []MicrosoftGraphBitlockerRecoveryKey)`

SetRecoveryKeys sets RecoveryKeys field to given value.

### HasRecoveryKeys

`func (o *MicrosoftGraphBitlocker) HasRecoveryKeys() bool`

HasRecoveryKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


