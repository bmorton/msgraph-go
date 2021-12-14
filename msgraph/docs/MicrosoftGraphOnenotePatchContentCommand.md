# MicrosoftGraphOnenotePatchContentCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AnyOfmicrosoftGraphOnenotePatchActionType**](anyOf&lt;microsoft.graph.onenotePatchActionType&gt;.md) | The action to perform on the target element. The possible values are: replace, append, delete, insert, or prepend. | [optional] 
**Content** | Pointer to **NullableString** | A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a &#39;Commands&#39; part. | [optional] 
**Position** | Pointer to [**AnyOfmicrosoftGraphOnenotePatchInsertPosition**](anyOf&lt;microsoft.graph.onenotePatchInsertPosition&gt;.md) | The location to add the supplied content, relative to the target element. The possible values are: after (default) or before. | [optional] 
**Target** | Pointer to **string** | The element to update. Must be the #&lt;data-id&gt; or the generated &lt;id&gt; of the element, or the body or title keyword. | [optional] 

## Methods

### NewMicrosoftGraphOnenotePatchContentCommand

`func NewMicrosoftGraphOnenotePatchContentCommand() *MicrosoftGraphOnenotePatchContentCommand`

NewMicrosoftGraphOnenotePatchContentCommand instantiates a new MicrosoftGraphOnenotePatchContentCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenotePatchContentCommandWithDefaults

`func NewMicrosoftGraphOnenotePatchContentCommandWithDefaults() *MicrosoftGraphOnenotePatchContentCommand`

NewMicrosoftGraphOnenotePatchContentCommandWithDefaults instantiates a new MicrosoftGraphOnenotePatchContentCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetAction() AnyOfmicrosoftGraphOnenotePatchActionType`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetActionOk() (*AnyOfmicrosoftGraphOnenotePatchActionType, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetAction(v AnyOfmicrosoftGraphOnenotePatchActionType)`

SetAction sets Action field to given value.

### HasAction

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *MicrosoftGraphOnenotePatchContentCommand) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphOnenotePatchContentCommand) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetPosition() AnyOfmicrosoftGraphOnenotePatchInsertPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetPositionOk() (*AnyOfmicrosoftGraphOnenotePatchInsertPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetPosition(v AnyOfmicrosoftGraphOnenotePatchInsertPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *MicrosoftGraphOnenotePatchContentCommand) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphOnenotePatchContentCommand) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphOnenotePatchContentCommand) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


