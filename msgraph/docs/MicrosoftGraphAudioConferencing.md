# MicrosoftGraphAudioConferencing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceId** | Pointer to **NullableString** | The conference id of the online meeting. | [optional] 
**DialinUrl** | Pointer to **NullableString** | A URL to the externally-accessible web page that contains dial-in information. | [optional] 
**TollFreeNumber** | Pointer to **NullableString** |  | [optional] 
**TollFreeNumbers** | Pointer to **[]string** | List of toll-free numbers that are displayed in the meeting invite. | [optional] 
**TollNumber** | Pointer to **NullableString** |  | [optional] 
**TollNumbers** | Pointer to **[]string** | List of toll numbers that are displayed in the meeting invite. | [optional] 

## Methods

### NewMicrosoftGraphAudioConferencing

`func NewMicrosoftGraphAudioConferencing() *MicrosoftGraphAudioConferencing`

NewMicrosoftGraphAudioConferencing instantiates a new MicrosoftGraphAudioConferencing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAudioConferencingWithDefaults

`func NewMicrosoftGraphAudioConferencingWithDefaults() *MicrosoftGraphAudioConferencing`

NewMicrosoftGraphAudioConferencingWithDefaults instantiates a new MicrosoftGraphAudioConferencing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConferenceId

`func (o *MicrosoftGraphAudioConferencing) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *MicrosoftGraphAudioConferencing) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *MicrosoftGraphAudioConferencing) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *MicrosoftGraphAudioConferencing) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### SetConferenceIdNil

`func (o *MicrosoftGraphAudioConferencing) SetConferenceIdNil(b bool)`

 SetConferenceIdNil sets the value for ConferenceId to be an explicit nil

### UnsetConferenceId
`func (o *MicrosoftGraphAudioConferencing) UnsetConferenceId()`

UnsetConferenceId ensures that no value is present for ConferenceId, not even an explicit nil
### GetDialinUrl

`func (o *MicrosoftGraphAudioConferencing) GetDialinUrl() string`

GetDialinUrl returns the DialinUrl field if non-nil, zero value otherwise.

### GetDialinUrlOk

`func (o *MicrosoftGraphAudioConferencing) GetDialinUrlOk() (*string, bool)`

GetDialinUrlOk returns a tuple with the DialinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialinUrl

`func (o *MicrosoftGraphAudioConferencing) SetDialinUrl(v string)`

SetDialinUrl sets DialinUrl field to given value.

### HasDialinUrl

`func (o *MicrosoftGraphAudioConferencing) HasDialinUrl() bool`

HasDialinUrl returns a boolean if a field has been set.

### SetDialinUrlNil

`func (o *MicrosoftGraphAudioConferencing) SetDialinUrlNil(b bool)`

 SetDialinUrlNil sets the value for DialinUrl to be an explicit nil

### UnsetDialinUrl
`func (o *MicrosoftGraphAudioConferencing) UnsetDialinUrl()`

UnsetDialinUrl ensures that no value is present for DialinUrl, not even an explicit nil
### GetTollFreeNumber

`func (o *MicrosoftGraphAudioConferencing) GetTollFreeNumber() string`

GetTollFreeNumber returns the TollFreeNumber field if non-nil, zero value otherwise.

### GetTollFreeNumberOk

`func (o *MicrosoftGraphAudioConferencing) GetTollFreeNumberOk() (*string, bool)`

GetTollFreeNumberOk returns a tuple with the TollFreeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollFreeNumber

`func (o *MicrosoftGraphAudioConferencing) SetTollFreeNumber(v string)`

SetTollFreeNumber sets TollFreeNumber field to given value.

### HasTollFreeNumber

`func (o *MicrosoftGraphAudioConferencing) HasTollFreeNumber() bool`

HasTollFreeNumber returns a boolean if a field has been set.

### SetTollFreeNumberNil

`func (o *MicrosoftGraphAudioConferencing) SetTollFreeNumberNil(b bool)`

 SetTollFreeNumberNil sets the value for TollFreeNumber to be an explicit nil

### UnsetTollFreeNumber
`func (o *MicrosoftGraphAudioConferencing) UnsetTollFreeNumber()`

UnsetTollFreeNumber ensures that no value is present for TollFreeNumber, not even an explicit nil
### GetTollFreeNumbers

`func (o *MicrosoftGraphAudioConferencing) GetTollFreeNumbers() []*string`

GetTollFreeNumbers returns the TollFreeNumbers field if non-nil, zero value otherwise.

### GetTollFreeNumbersOk

`func (o *MicrosoftGraphAudioConferencing) GetTollFreeNumbersOk() (*[]*string, bool)`

GetTollFreeNumbersOk returns a tuple with the TollFreeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollFreeNumbers

`func (o *MicrosoftGraphAudioConferencing) SetTollFreeNumbers(v []*string)`

SetTollFreeNumbers sets TollFreeNumbers field to given value.

### HasTollFreeNumbers

`func (o *MicrosoftGraphAudioConferencing) HasTollFreeNumbers() bool`

HasTollFreeNumbers returns a boolean if a field has been set.

### GetTollNumber

`func (o *MicrosoftGraphAudioConferencing) GetTollNumber() string`

GetTollNumber returns the TollNumber field if non-nil, zero value otherwise.

### GetTollNumberOk

`func (o *MicrosoftGraphAudioConferencing) GetTollNumberOk() (*string, bool)`

GetTollNumberOk returns a tuple with the TollNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollNumber

`func (o *MicrosoftGraphAudioConferencing) SetTollNumber(v string)`

SetTollNumber sets TollNumber field to given value.

### HasTollNumber

`func (o *MicrosoftGraphAudioConferencing) HasTollNumber() bool`

HasTollNumber returns a boolean if a field has been set.

### SetTollNumberNil

`func (o *MicrosoftGraphAudioConferencing) SetTollNumberNil(b bool)`

 SetTollNumberNil sets the value for TollNumber to be an explicit nil

### UnsetTollNumber
`func (o *MicrosoftGraphAudioConferencing) UnsetTollNumber()`

UnsetTollNumber ensures that no value is present for TollNumber, not even an explicit nil
### GetTollNumbers

`func (o *MicrosoftGraphAudioConferencing) GetTollNumbers() []*string`

GetTollNumbers returns the TollNumbers field if non-nil, zero value otherwise.

### GetTollNumbersOk

`func (o *MicrosoftGraphAudioConferencing) GetTollNumbersOk() (*[]*string, bool)`

GetTollNumbersOk returns a tuple with the TollNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollNumbers

`func (o *MicrosoftGraphAudioConferencing) SetTollNumbers(v []*string)`

SetTollNumbers sets TollNumbers field to given value.

### HasTollNumbers

`func (o *MicrosoftGraphAudioConferencing) HasTollNumbers() bool`

HasTollNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


