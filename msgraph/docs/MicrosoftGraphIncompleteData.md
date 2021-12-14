# MicrosoftGraphIncompleteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissingDataBeforeDateTime** | Pointer to **NullableTime** | The service does not have source data before the specified time. | [optional] 
**WasThrottled** | Pointer to **NullableBool** | Some data was not recorded due to excessive activity. | [optional] 

## Methods

### NewMicrosoftGraphIncompleteData

`func NewMicrosoftGraphIncompleteData() *MicrosoftGraphIncompleteData`

NewMicrosoftGraphIncompleteData instantiates a new MicrosoftGraphIncompleteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIncompleteDataWithDefaults

`func NewMicrosoftGraphIncompleteDataWithDefaults() *MicrosoftGraphIncompleteData`

NewMicrosoftGraphIncompleteDataWithDefaults instantiates a new MicrosoftGraphIncompleteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTime() time.Time`

GetMissingDataBeforeDateTime returns the MissingDataBeforeDateTime field if non-nil, zero value otherwise.

### GetMissingDataBeforeDateTimeOk

`func (o *MicrosoftGraphIncompleteData) GetMissingDataBeforeDateTimeOk() (*time.Time, bool)`

GetMissingDataBeforeDateTimeOk returns a tuple with the MissingDataBeforeDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTime(v time.Time)`

SetMissingDataBeforeDateTime sets MissingDataBeforeDateTime field to given value.

### HasMissingDataBeforeDateTime

`func (o *MicrosoftGraphIncompleteData) HasMissingDataBeforeDateTime() bool`

HasMissingDataBeforeDateTime returns a boolean if a field has been set.

### SetMissingDataBeforeDateTimeNil

`func (o *MicrosoftGraphIncompleteData) SetMissingDataBeforeDateTimeNil(b bool)`

 SetMissingDataBeforeDateTimeNil sets the value for MissingDataBeforeDateTime to be an explicit nil

### UnsetMissingDataBeforeDateTime
`func (o *MicrosoftGraphIncompleteData) UnsetMissingDataBeforeDateTime()`

UnsetMissingDataBeforeDateTime ensures that no value is present for MissingDataBeforeDateTime, not even an explicit nil
### GetWasThrottled

`func (o *MicrosoftGraphIncompleteData) GetWasThrottled() bool`

GetWasThrottled returns the WasThrottled field if non-nil, zero value otherwise.

### GetWasThrottledOk

`func (o *MicrosoftGraphIncompleteData) GetWasThrottledOk() (*bool, bool)`

GetWasThrottledOk returns a tuple with the WasThrottled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasThrottled

`func (o *MicrosoftGraphIncompleteData) SetWasThrottled(v bool)`

SetWasThrottled sets WasThrottled field to given value.

### HasWasThrottled

`func (o *MicrosoftGraphIncompleteData) HasWasThrottled() bool`

HasWasThrottled returns a boolean if a field has been set.

### SetWasThrottledNil

`func (o *MicrosoftGraphIncompleteData) SetWasThrottledNil(b bool)`

 SetWasThrottledNil sets the value for WasThrottled to be an explicit nil

### UnsetWasThrottled
`func (o *MicrosoftGraphIncompleteData) UnsetWasThrottled()`

UnsetWasThrottled ensures that no value is present for WasThrottled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


