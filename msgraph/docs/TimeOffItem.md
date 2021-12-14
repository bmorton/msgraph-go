# TimeOffItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOffReasonId** | Pointer to **NullableString** | ID of the timeOffReason for this timeOffItem. Required. | [optional] 

## Methods

### NewTimeOffItem

`func NewTimeOffItem() *TimeOffItem`

NewTimeOffItem instantiates a new TimeOffItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeOffItemWithDefaults

`func NewTimeOffItemWithDefaults() *TimeOffItem`

NewTimeOffItemWithDefaults instantiates a new TimeOffItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOffReasonId

`func (o *TimeOffItem) GetTimeOffReasonId() string`

GetTimeOffReasonId returns the TimeOffReasonId field if non-nil, zero value otherwise.

### GetTimeOffReasonIdOk

`func (o *TimeOffItem) GetTimeOffReasonIdOk() (*string, bool)`

GetTimeOffReasonIdOk returns a tuple with the TimeOffReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffReasonId

`func (o *TimeOffItem) SetTimeOffReasonId(v string)`

SetTimeOffReasonId sets TimeOffReasonId field to given value.

### HasTimeOffReasonId

`func (o *TimeOffItem) HasTimeOffReasonId() bool`

HasTimeOffReasonId returns a boolean if a field has been set.

### SetTimeOffReasonIdNil

`func (o *TimeOffItem) SetTimeOffReasonIdNil(b bool)`

 SetTimeOffReasonIdNil sets the value for TimeOffReasonId to be an explicit nil

### UnsetTimeOffReasonId
`func (o *TimeOffItem) UnsetTimeOffReasonId()`

UnsetTimeOffReasonId ensures that no value is present for TimeOffReasonId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


