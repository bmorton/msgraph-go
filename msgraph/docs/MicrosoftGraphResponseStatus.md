# MicrosoftGraphResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**AnyOfmicrosoftGraphResponseType**](anyOf&lt;microsoft.graph.responseType&gt;.md) | The response type. Possible values are: none, organizer, tentativelyAccepted, accepted, declined, notResponded.To differentiate between none and notResponded:  none – from organizer&#39;s perspective. This value is used when the status of an attendee/participant is reported to the organizer of a meeting.  notResponded – from attendde&#39;s perspective. Indicates the attendee has not responded to the meeting request.  Clients can treat notResponded &#x3D;&#x3D; none.  As an example, if attendee Alex hasn&#39;t responded to a meeting request, getting Alex&#39; response status for that event in Alex&#39; calendar returns notResponded. Getting Alex&#39; response from the calendar of any other attendee or the organizer&#39;s returns none. Getting the organizer&#39;s response for the event in anybody&#39;s calendar also returns none. | [optional] 
**Time** | Pointer to **NullableTime** | The date and time that the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 

## Methods

### NewMicrosoftGraphResponseStatus

`func NewMicrosoftGraphResponseStatus() *MicrosoftGraphResponseStatus`

NewMicrosoftGraphResponseStatus instantiates a new MicrosoftGraphResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphResponseStatusWithDefaults

`func NewMicrosoftGraphResponseStatusWithDefaults() *MicrosoftGraphResponseStatus`

NewMicrosoftGraphResponseStatusWithDefaults instantiates a new MicrosoftGraphResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *MicrosoftGraphResponseStatus) GetResponse() AnyOfmicrosoftGraphResponseType`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MicrosoftGraphResponseStatus) GetResponseOk() (*AnyOfmicrosoftGraphResponseType, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MicrosoftGraphResponseStatus) SetResponse(v AnyOfmicrosoftGraphResponseType)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MicrosoftGraphResponseStatus) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *MicrosoftGraphResponseStatus) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *MicrosoftGraphResponseStatus) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil
### GetTime

`func (o *MicrosoftGraphResponseStatus) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MicrosoftGraphResponseStatus) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MicrosoftGraphResponseStatus) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *MicrosoftGraphResponseStatus) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *MicrosoftGraphResponseStatus) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *MicrosoftGraphResponseStatus) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


