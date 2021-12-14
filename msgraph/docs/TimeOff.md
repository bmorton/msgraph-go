# TimeOff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DraftTimeOff** | Pointer to [**AnyOfmicrosoftGraphTimeOffItem**](anyOf&lt;microsoft.graph.timeOffItem&gt;.md) | The draft version of this timeOff that is viewable by managers. Required. | [optional] 
**SharedTimeOff** | Pointer to [**AnyOfmicrosoftGraphTimeOffItem**](anyOf&lt;microsoft.graph.timeOffItem&gt;.md) | The shared version of this timeOff that is viewable by both employees and managers. Required. | [optional] 
**UserId** | Pointer to **NullableString** | ID of the user assigned to the timeOff. Required. | [optional] 

## Methods

### NewTimeOff

`func NewTimeOff() *TimeOff`

NewTimeOff instantiates a new TimeOff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffWithDefaults

`func NewTimeOffWithDefaults() *TimeOff`

NewTimeOffWithDefaults instantiates a new TimeOff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraftTimeOff

`func (o *TimeOff) GetDraftTimeOff() AnyOfmicrosoftGraphTimeOffItem`

GetDraftTimeOff returns the DraftTimeOff field if non-nil, zero value otherwise.

### GetDraftTimeOffOk

`func (o *TimeOff) GetDraftTimeOffOk() (*AnyOfmicrosoftGraphTimeOffItem, bool)`

GetDraftTimeOffOk returns a tuple with the DraftTimeOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftTimeOff

`func (o *TimeOff) SetDraftTimeOff(v AnyOfmicrosoftGraphTimeOffItem)`

SetDraftTimeOff sets DraftTimeOff field to given value.

### HasDraftTimeOff

`func (o *TimeOff) HasDraftTimeOff() bool`

HasDraftTimeOff returns a boolean if a field has been set.

### SetDraftTimeOffNil

`func (o *TimeOff) SetDraftTimeOffNil(b bool)`

 SetDraftTimeOffNil sets the value for DraftTimeOff to be an explicit nil

### UnsetDraftTimeOff
`func (o *TimeOff) UnsetDraftTimeOff()`

UnsetDraftTimeOff ensures that no value is present for DraftTimeOff, not even an explicit nil
### GetSharedTimeOff

`func (o *TimeOff) GetSharedTimeOff() AnyOfmicrosoftGraphTimeOffItem`

GetSharedTimeOff returns the SharedTimeOff field if non-nil, zero value otherwise.

### GetSharedTimeOffOk

`func (o *TimeOff) GetSharedTimeOffOk() (*AnyOfmicrosoftGraphTimeOffItem, bool)`

GetSharedTimeOffOk returns a tuple with the SharedTimeOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTimeOff

`func (o *TimeOff) SetSharedTimeOff(v AnyOfmicrosoftGraphTimeOffItem)`

SetSharedTimeOff sets SharedTimeOff field to given value.

### HasSharedTimeOff

`func (o *TimeOff) HasSharedTimeOff() bool`

HasSharedTimeOff returns a boolean if a field has been set.

### SetSharedTimeOffNil

`func (o *TimeOff) SetSharedTimeOffNil(b bool)`

 SetSharedTimeOffNil sets the value for SharedTimeOff to be an explicit nil

### UnsetSharedTimeOff
`func (o *TimeOff) UnsetSharedTimeOff()`

UnsetSharedTimeOff ensures that no value is present for SharedTimeOff, not even an explicit nil
### GetUserId

`func (o *TimeOff) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TimeOff) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TimeOff) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TimeOff) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *TimeOff) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *TimeOff) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


