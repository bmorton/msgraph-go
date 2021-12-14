# MicrosoftGraphServiceHealthIssuePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | The published time of the post. | [optional] 
**Description** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The content of the service issue post. | [optional] 
**PostType** | Pointer to [**AnyOfmicrosoftGraphPostType**](anyOf&lt;microsoft.graph.postType&gt;.md) | The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphServiceHealthIssuePost

`func NewMicrosoftGraphServiceHealthIssuePost() *MicrosoftGraphServiceHealthIssuePost`

NewMicrosoftGraphServiceHealthIssuePost instantiates a new MicrosoftGraphServiceHealthIssuePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceHealthIssuePostWithDefaults

`func NewMicrosoftGraphServiceHealthIssuePostWithDefaults() *MicrosoftGraphServiceHealthIssuePost`

NewMicrosoftGraphServiceHealthIssuePostWithDefaults instantiates a new MicrosoftGraphServiceHealthIssuePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *MicrosoftGraphServiceHealthIssuePost) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphServiceHealthIssuePost) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphServiceHealthIssuePost) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphServiceHealthIssuePost) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphServiceHealthIssuePost) GetDescription() AnyOfmicrosoftGraphItemBody`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphServiceHealthIssuePost) GetDescriptionOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphServiceHealthIssuePost) SetDescription(v AnyOfmicrosoftGraphItemBody)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphServiceHealthIssuePost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphServiceHealthIssuePost) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphServiceHealthIssuePost) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPostType

`func (o *MicrosoftGraphServiceHealthIssuePost) GetPostType() AnyOfmicrosoftGraphPostType`

GetPostType returns the PostType field if non-nil, zero value otherwise.

### GetPostTypeOk

`func (o *MicrosoftGraphServiceHealthIssuePost) GetPostTypeOk() (*AnyOfmicrosoftGraphPostType, bool)`

GetPostTypeOk returns a tuple with the PostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostType

`func (o *MicrosoftGraphServiceHealthIssuePost) SetPostType(v AnyOfmicrosoftGraphPostType)`

SetPostType sets PostType field to given value.

### HasPostType

`func (o *MicrosoftGraphServiceHealthIssuePost) HasPostType() bool`

HasPostType returns a boolean if a field has been set.

### SetPostTypeNil

`func (o *MicrosoftGraphServiceHealthIssuePost) SetPostTypeNil(b bool)`

 SetPostTypeNil sets the value for PostType to be an explicit nil

### UnsetPostType
`func (o *MicrosoftGraphServiceHealthIssuePost) UnsetPostType()`

UnsetPostType ensures that no value is present for PostType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


