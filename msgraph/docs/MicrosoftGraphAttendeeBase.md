# MicrosoftGraphAttendeeBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The recipient&#39;s email address. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphAttendeeType**](anyOf&lt;microsoft.graph.attendeeType&gt;.md) | The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type. | [optional] 

## Methods

### NewMicrosoftGraphAttendeeBase

`func NewMicrosoftGraphAttendeeBase() *MicrosoftGraphAttendeeBase`

NewMicrosoftGraphAttendeeBase instantiates a new MicrosoftGraphAttendeeBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAttendeeBaseWithDefaults

`func NewMicrosoftGraphAttendeeBaseWithDefaults() *MicrosoftGraphAttendeeBase`

NewMicrosoftGraphAttendeeBaseWithDefaults instantiates a new MicrosoftGraphAttendeeBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *MicrosoftGraphAttendeeBase) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphAttendeeBase) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphAttendeeBase) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphAttendeeBase) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *MicrosoftGraphAttendeeBase) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MicrosoftGraphAttendeeBase) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetType

`func (o *MicrosoftGraphAttendeeBase) GetType() AnyOfmicrosoftGraphAttendeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphAttendeeBase) GetTypeOk() (*AnyOfmicrosoftGraphAttendeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphAttendeeBase) SetType(v AnyOfmicrosoftGraphAttendeeType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphAttendeeBase) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphAttendeeBase) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphAttendeeBase) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


