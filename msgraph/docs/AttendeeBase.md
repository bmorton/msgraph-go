# AttendeeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AnyOfmicrosoftGraphAttendeeType**](anyOf&lt;microsoft.graph.attendeeType&gt;.md) | The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type. | [optional] 

## Methods

### NewAttendeeBase

`func NewAttendeeBase() *AttendeeBase`

NewAttendeeBase instantiates a new AttendeeBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttendeeBaseWithDefaults

`func NewAttendeeBaseWithDefaults() *AttendeeBase`

NewAttendeeBaseWithDefaults instantiates a new AttendeeBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttendeeBase) GetType() AnyOfmicrosoftGraphAttendeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttendeeBase) GetTypeOk() (*AnyOfmicrosoftGraphAttendeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttendeeBase) SetType(v AnyOfmicrosoftGraphAttendeeType)`

SetType sets Type field to given value.

### HasType

`func (o *AttendeeBase) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AttendeeBase) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AttendeeBase) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


