# MicrosoftGraphTeamworkActivityTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**AnyOfmicrosoftGraphTeamworkActivityTopicSource**](anyOf&lt;microsoft.graph.teamworkActivityTopicSource&gt;.md) | Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom text, use text. | [optional] 
**Value** | Pointer to **string** | The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the vaule is text, this must be a plain text value. | [optional] 
**WebUrl** | Pointer to **NullableString** | The link the user clicks when they select the notification. Optional when source is entityUrl; required when source is text. | [optional] 

## Methods

### NewMicrosoftGraphTeamworkActivityTopic

`func NewMicrosoftGraphTeamworkActivityTopic() *MicrosoftGraphTeamworkActivityTopic`

NewMicrosoftGraphTeamworkActivityTopic instantiates a new MicrosoftGraphTeamworkActivityTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamworkActivityTopicWithDefaults

`func NewMicrosoftGraphTeamworkActivityTopicWithDefaults() *MicrosoftGraphTeamworkActivityTopic`

NewMicrosoftGraphTeamworkActivityTopicWithDefaults instantiates a new MicrosoftGraphTeamworkActivityTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *MicrosoftGraphTeamworkActivityTopic) GetSource() AnyOfmicrosoftGraphTeamworkActivityTopicSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MicrosoftGraphTeamworkActivityTopic) GetSourceOk() (*AnyOfmicrosoftGraphTeamworkActivityTopicSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MicrosoftGraphTeamworkActivityTopic) SetSource(v AnyOfmicrosoftGraphTeamworkActivityTopicSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MicrosoftGraphTeamworkActivityTopic) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MicrosoftGraphTeamworkActivityTopic) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MicrosoftGraphTeamworkActivityTopic) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphTeamworkActivityTopic) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphTeamworkActivityTopic) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphTeamworkActivityTopic) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphTeamworkActivityTopic) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWebUrl

`func (o *MicrosoftGraphTeamworkActivityTopic) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphTeamworkActivityTopic) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphTeamworkActivityTopic) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphTeamworkActivityTopic) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphTeamworkActivityTopic) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphTeamworkActivityTopic) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


