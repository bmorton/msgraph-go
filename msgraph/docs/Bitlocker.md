# Bitlocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryKeys** | Pointer to [**[]MicrosoftGraphBitlockerRecoveryKey**](MicrosoftGraphBitlockerRecoveryKey.md) | The recovery keys associated with the bitlocker entity. | [optional] 

## Methods

### NewBitlocker

`func NewBitlocker() *Bitlocker`

NewBitlocker instantiates a new Bitlocker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBitlockerWithDefaults

`func NewBitlockerWithDefaults() *Bitlocker`

NewBitlockerWithDefaults instantiates a new Bitlocker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryKeys

`func (o *Bitlocker) GetRecoveryKeys() []MicrosoftGraphBitlockerRecoveryKey`

GetRecoveryKeys returns the RecoveryKeys field if non-nil, zero value otherwise.

### GetRecoveryKeysOk

`func (o *Bitlocker) GetRecoveryKeysOk() (*[]MicrosoftGraphBitlockerRecoveryKey, bool)`

GetRecoveryKeysOk returns a tuple with the RecoveryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeys

`func (o *Bitlocker) SetRecoveryKeys(v []MicrosoftGraphBitlockerRecoveryKey)`

SetRecoveryKeys sets RecoveryKeys field to given value.

### HasRecoveryKeys

`func (o *Bitlocker) HasRecoveryKeys() bool`

HasRecoveryKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


