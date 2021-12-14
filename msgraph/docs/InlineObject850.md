# InlineObject850

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to [**AnyOfmicrosoftGraphTeamworkActivityTopic**](anyOf&lt;microsoft.graph.teamworkActivityTopic&gt;.md) |  | [optional] 
**ActivityType** | Pointer to **NullableString** |  | [optional] 
**ChainId** | Pointer to **NullableInt64** |  | [optional] 
**PreviewText** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) |  | [optional] 
**TemplateParameters** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) |  | [optional] 
**Recipient** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### NewInlineObject850

`func NewInlineObject850() *InlineObject850`

NewInlineObject850 instantiates a new InlineObject850 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject850WithDefaults

`func NewInlineObject850WithDefaults() *InlineObject850`

NewInlineObject850WithDefaults instantiates a new InlineObject850 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *InlineObject850) GetTopic() AnyOfmicrosoftGraphTeamworkActivityTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *InlineObject850) GetTopicOk() (*AnyOfmicrosoftGraphTeamworkActivityTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *InlineObject850) SetTopic(v AnyOfmicrosoftGraphTeamworkActivityTopic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *InlineObject850) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *InlineObject850) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *InlineObject850) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetActivityType

`func (o *InlineObject850) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *InlineObject850) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *InlineObject850) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *InlineObject850) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### SetActivityTypeNil

`func (o *InlineObject850) SetActivityTypeNil(b bool)`

 SetActivityTypeNil sets the value for ActivityType to be an explicit nil

### UnsetActivityType
`func (o *InlineObject850) UnsetActivityType()`

UnsetActivityType ensures that no value is present for ActivityType, not even an explicit nil
### GetChainId

`func (o *InlineObject850) GetChainId() int64`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *InlineObject850) GetChainIdOk() (*int64, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *InlineObject850) SetChainId(v int64)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *InlineObject850) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### SetChainIdNil

`func (o *InlineObject850) SetChainIdNil(b bool)`

 SetChainIdNil sets the value for ChainId to be an explicit nil

### UnsetChainId
`func (o *InlineObject850) UnsetChainId()`

UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
### GetPreviewText

`func (o *InlineObject850) GetPreviewText() AnyOfmicrosoftGraphItemBody`

GetPreviewText returns the PreviewText field if non-nil, zero value otherwise.

### GetPreviewTextOk

`func (o *InlineObject850) GetPreviewTextOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetPreviewTextOk returns a tuple with the PreviewText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewText

`func (o *InlineObject850) SetPreviewText(v AnyOfmicrosoftGraphItemBody)`

SetPreviewText sets PreviewText field to given value.

### HasPreviewText

`func (o *InlineObject850) HasPreviewText() bool`

HasPreviewText returns a boolean if a field has been set.

### SetPreviewTextNil

`func (o *InlineObject850) SetPreviewTextNil(b bool)`

 SetPreviewTextNil sets the value for PreviewText to be an explicit nil

### UnsetPreviewText
`func (o *InlineObject850) UnsetPreviewText()`

UnsetPreviewText ensures that no value is present for PreviewText, not even an explicit nil
### GetTemplateParameters

`func (o *InlineObject850) GetTemplateParameters() []*AnyOfmicrosoftGraphKeyValuePair`

GetTemplateParameters returns the TemplateParameters field if non-nil, zero value otherwise.

### GetTemplateParametersOk

`func (o *InlineObject850) GetTemplateParametersOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetTemplateParametersOk returns a tuple with the TemplateParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameters

`func (o *InlineObject850) SetTemplateParameters(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetTemplateParameters sets TemplateParameters field to given value.

### HasTemplateParameters

`func (o *InlineObject850) HasTemplateParameters() bool`

HasTemplateParameters returns a boolean if a field has been set.

### GetRecipient

`func (o *InlineObject850) GetRecipient() AnyOfobject`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *InlineObject850) GetRecipientOk() (*AnyOfobject, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *InlineObject850) SetRecipient(v AnyOfobject)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *InlineObject850) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *InlineObject850) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *InlineObject850) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


