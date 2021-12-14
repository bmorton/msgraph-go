# MicrosoftGraphSharingDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedBy** | Pointer to [**AnyOfmicrosoftGraphInsightIdentity**](anyOf&lt;microsoft.graph.insightIdentity&gt;.md) | The user who shared the document. | [optional] 
**SharedDateTime** | Pointer to **NullableTime** | The date and time the file was last shared. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**SharingReference** | Pointer to [**AnyOfmicrosoftGraphResourceReference**](anyOf&lt;microsoft.graph.resourceReference&gt;.md) |  | [optional] 
**SharingSubject** | Pointer to **NullableString** | The subject with which the document was shared. | [optional] 
**SharingType** | Pointer to **NullableString** | Determines the way the document was shared, can be by a &#39;Link&#39;, &#39;Attachment&#39;, &#39;Group&#39;, &#39;Site&#39;. | [optional] 

## Methods

### NewMicrosoftGraphSharingDetail

`func NewMicrosoftGraphSharingDetail() *MicrosoftGraphSharingDetail`

NewMicrosoftGraphSharingDetail instantiates a new MicrosoftGraphSharingDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharingDetailWithDefaults

`func NewMicrosoftGraphSharingDetailWithDefaults() *MicrosoftGraphSharingDetail`

NewMicrosoftGraphSharingDetailWithDefaults instantiates a new MicrosoftGraphSharingDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedBy

`func (o *MicrosoftGraphSharingDetail) GetSharedBy() AnyOfmicrosoftGraphInsightIdentity`

GetSharedBy returns the SharedBy field if non-nil, zero value otherwise.

### GetSharedByOk

`func (o *MicrosoftGraphSharingDetail) GetSharedByOk() (*AnyOfmicrosoftGraphInsightIdentity, bool)`

GetSharedByOk returns a tuple with the SharedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBy

`func (o *MicrosoftGraphSharingDetail) SetSharedBy(v AnyOfmicrosoftGraphInsightIdentity)`

SetSharedBy sets SharedBy field to given value.

### HasSharedBy

`func (o *MicrosoftGraphSharingDetail) HasSharedBy() bool`

HasSharedBy returns a boolean if a field has been set.

### SetSharedByNil

`func (o *MicrosoftGraphSharingDetail) SetSharedByNil(b bool)`

 SetSharedByNil sets the value for SharedBy to be an explicit nil

### UnsetSharedBy
`func (o *MicrosoftGraphSharingDetail) UnsetSharedBy()`

UnsetSharedBy ensures that no value is present for SharedBy, not even an explicit nil
### GetSharedDateTime

`func (o *MicrosoftGraphSharingDetail) GetSharedDateTime() time.Time`

GetSharedDateTime returns the SharedDateTime field if non-nil, zero value otherwise.

### GetSharedDateTimeOk

`func (o *MicrosoftGraphSharingDetail) GetSharedDateTimeOk() (*time.Time, bool)`

GetSharedDateTimeOk returns a tuple with the SharedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDateTime

`func (o *MicrosoftGraphSharingDetail) SetSharedDateTime(v time.Time)`

SetSharedDateTime sets SharedDateTime field to given value.

### HasSharedDateTime

`func (o *MicrosoftGraphSharingDetail) HasSharedDateTime() bool`

HasSharedDateTime returns a boolean if a field has been set.

### SetSharedDateTimeNil

`func (o *MicrosoftGraphSharingDetail) SetSharedDateTimeNil(b bool)`

 SetSharedDateTimeNil sets the value for SharedDateTime to be an explicit nil

### UnsetSharedDateTime
`func (o *MicrosoftGraphSharingDetail) UnsetSharedDateTime()`

UnsetSharedDateTime ensures that no value is present for SharedDateTime, not even an explicit nil
### GetSharingReference

`func (o *MicrosoftGraphSharingDetail) GetSharingReference() AnyOfmicrosoftGraphResourceReference`

GetSharingReference returns the SharingReference field if non-nil, zero value otherwise.

### GetSharingReferenceOk

`func (o *MicrosoftGraphSharingDetail) GetSharingReferenceOk() (*AnyOfmicrosoftGraphResourceReference, bool)`

GetSharingReferenceOk returns a tuple with the SharingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingReference

`func (o *MicrosoftGraphSharingDetail) SetSharingReference(v AnyOfmicrosoftGraphResourceReference)`

SetSharingReference sets SharingReference field to given value.

### HasSharingReference

`func (o *MicrosoftGraphSharingDetail) HasSharingReference() bool`

HasSharingReference returns a boolean if a field has been set.

### SetSharingReferenceNil

`func (o *MicrosoftGraphSharingDetail) SetSharingReferenceNil(b bool)`

 SetSharingReferenceNil sets the value for SharingReference to be an explicit nil

### UnsetSharingReference
`func (o *MicrosoftGraphSharingDetail) UnsetSharingReference()`

UnsetSharingReference ensures that no value is present for SharingReference, not even an explicit nil
### GetSharingSubject

`func (o *MicrosoftGraphSharingDetail) GetSharingSubject() string`

GetSharingSubject returns the SharingSubject field if non-nil, zero value otherwise.

### GetSharingSubjectOk

`func (o *MicrosoftGraphSharingDetail) GetSharingSubjectOk() (*string, bool)`

GetSharingSubjectOk returns a tuple with the SharingSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingSubject

`func (o *MicrosoftGraphSharingDetail) SetSharingSubject(v string)`

SetSharingSubject sets SharingSubject field to given value.

### HasSharingSubject

`func (o *MicrosoftGraphSharingDetail) HasSharingSubject() bool`

HasSharingSubject returns a boolean if a field has been set.

### SetSharingSubjectNil

`func (o *MicrosoftGraphSharingDetail) SetSharingSubjectNil(b bool)`

 SetSharingSubjectNil sets the value for SharingSubject to be an explicit nil

### UnsetSharingSubject
`func (o *MicrosoftGraphSharingDetail) UnsetSharingSubject()`

UnsetSharingSubject ensures that no value is present for SharingSubject, not even an explicit nil
### GetSharingType

`func (o *MicrosoftGraphSharingDetail) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *MicrosoftGraphSharingDetail) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *MicrosoftGraphSharingDetail) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.

### HasSharingType

`func (o *MicrosoftGraphSharingDetail) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.

### SetSharingTypeNil

`func (o *MicrosoftGraphSharingDetail) SetSharingTypeNil(b bool)`

 SetSharingTypeNil sets the value for SharingType to be an explicit nil

### UnsetSharingType
`func (o *MicrosoftGraphSharingDetail) UnsetSharingType()`

UnsetSharingType ensures that no value is present for SharingType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


