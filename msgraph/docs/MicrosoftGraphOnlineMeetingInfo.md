# MicrosoftGraphOnlineMeetingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConferenceId** | Pointer to **NullableString** | The ID of the conference. | [optional] 
**JoinUrl** | Pointer to **NullableString** | The external link that launches the online meeting. This is a URL that clients will launch into a browser and will redirect the user to join the meeting. | [optional] 
**Phones** | Pointer to [**[]AnyOfmicrosoftGraphPhone**](AnyOfmicrosoftGraphPhone.md) | All of the phone numbers associated with this conference. | [optional] 
**QuickDial** | Pointer to **NullableString** | The pre-formatted quickdial for this call. | [optional] 
**TollFreeNumbers** | Pointer to **[]string** | The toll free numbers that can be used to join the conference. | [optional] 
**TollNumber** | Pointer to **NullableString** | The toll number that can be used to join the conference. | [optional] 

## Methods

### NewMicrosoftGraphOnlineMeetingInfo

`func NewMicrosoftGraphOnlineMeetingInfo() *MicrosoftGraphOnlineMeetingInfo`

NewMicrosoftGraphOnlineMeetingInfo instantiates a new MicrosoftGraphOnlineMeetingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnlineMeetingInfoWithDefaults

`func NewMicrosoftGraphOnlineMeetingInfoWithDefaults() *MicrosoftGraphOnlineMeetingInfo`

NewMicrosoftGraphOnlineMeetingInfoWithDefaults instantiates a new MicrosoftGraphOnlineMeetingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConferenceId

`func (o *MicrosoftGraphOnlineMeetingInfo) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *MicrosoftGraphOnlineMeetingInfo) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *MicrosoftGraphOnlineMeetingInfo) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### SetConferenceIdNil

`func (o *MicrosoftGraphOnlineMeetingInfo) SetConferenceIdNil(b bool)`

 SetConferenceIdNil sets the value for ConferenceId to be an explicit nil

### UnsetConferenceId
`func (o *MicrosoftGraphOnlineMeetingInfo) UnsetConferenceId()`

UnsetConferenceId ensures that no value is present for ConferenceId, not even an explicit nil
### GetJoinUrl

`func (o *MicrosoftGraphOnlineMeetingInfo) GetJoinUrl() string`

GetJoinUrl returns the JoinUrl field if non-nil, zero value otherwise.

### GetJoinUrlOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetJoinUrlOk() (*string, bool)`

GetJoinUrlOk returns a tuple with the JoinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinUrl

`func (o *MicrosoftGraphOnlineMeetingInfo) SetJoinUrl(v string)`

SetJoinUrl sets JoinUrl field to given value.

### HasJoinUrl

`func (o *MicrosoftGraphOnlineMeetingInfo) HasJoinUrl() bool`

HasJoinUrl returns a boolean if a field has been set.

### SetJoinUrlNil

`func (o *MicrosoftGraphOnlineMeetingInfo) SetJoinUrlNil(b bool)`

 SetJoinUrlNil sets the value for JoinUrl to be an explicit nil

### UnsetJoinUrl
`func (o *MicrosoftGraphOnlineMeetingInfo) UnsetJoinUrl()`

UnsetJoinUrl ensures that no value is present for JoinUrl, not even an explicit nil
### GetPhones

`func (o *MicrosoftGraphOnlineMeetingInfo) GetPhones() []*AnyOfmicrosoftGraphPhone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetPhonesOk() (*[]*AnyOfmicrosoftGraphPhone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *MicrosoftGraphOnlineMeetingInfo) SetPhones(v []*AnyOfmicrosoftGraphPhone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *MicrosoftGraphOnlineMeetingInfo) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetQuickDial

`func (o *MicrosoftGraphOnlineMeetingInfo) GetQuickDial() string`

GetQuickDial returns the QuickDial field if non-nil, zero value otherwise.

### GetQuickDialOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetQuickDialOk() (*string, bool)`

GetQuickDialOk returns a tuple with the QuickDial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDial

`func (o *MicrosoftGraphOnlineMeetingInfo) SetQuickDial(v string)`

SetQuickDial sets QuickDial field to given value.

### HasQuickDial

`func (o *MicrosoftGraphOnlineMeetingInfo) HasQuickDial() bool`

HasQuickDial returns a boolean if a field has been set.

### SetQuickDialNil

`func (o *MicrosoftGraphOnlineMeetingInfo) SetQuickDialNil(b bool)`

 SetQuickDialNil sets the value for QuickDial to be an explicit nil

### UnsetQuickDial
`func (o *MicrosoftGraphOnlineMeetingInfo) UnsetQuickDial()`

UnsetQuickDial ensures that no value is present for QuickDial, not even an explicit nil
### GetTollFreeNumbers

`func (o *MicrosoftGraphOnlineMeetingInfo) GetTollFreeNumbers() []*string`

GetTollFreeNumbers returns the TollFreeNumbers field if non-nil, zero value otherwise.

### GetTollFreeNumbersOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetTollFreeNumbersOk() (*[]*string, bool)`

GetTollFreeNumbersOk returns a tuple with the TollFreeNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollFreeNumbers

`func (o *MicrosoftGraphOnlineMeetingInfo) SetTollFreeNumbers(v []*string)`

SetTollFreeNumbers sets TollFreeNumbers field to given value.

### HasTollFreeNumbers

`func (o *MicrosoftGraphOnlineMeetingInfo) HasTollFreeNumbers() bool`

HasTollFreeNumbers returns a boolean if a field has been set.

### GetTollNumber

`func (o *MicrosoftGraphOnlineMeetingInfo) GetTollNumber() string`

GetTollNumber returns the TollNumber field if non-nil, zero value otherwise.

### GetTollNumberOk

`func (o *MicrosoftGraphOnlineMeetingInfo) GetTollNumberOk() (*string, bool)`

GetTollNumberOk returns a tuple with the TollNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTollNumber

`func (o *MicrosoftGraphOnlineMeetingInfo) SetTollNumber(v string)`

SetTollNumber sets TollNumber field to given value.

### HasTollNumber

`func (o *MicrosoftGraphOnlineMeetingInfo) HasTollNumber() bool`

HasTollNumber returns a boolean if a field has been set.

### SetTollNumberNil

`func (o *MicrosoftGraphOnlineMeetingInfo) SetTollNumberNil(b bool)`

 SetTollNumberNil sets the value for TollNumber to be an explicit nil

### UnsetTollNumber
`func (o *MicrosoftGraphOnlineMeetingInfo) UnsetTollNumber()`

UnsetTollNumber ensures that no value is present for TollNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


