# MicrosoftGraphVisualInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribution** | Pointer to [**AnyOfmicrosoftGraphImageInfo**](anyOf&lt;microsoft.graph.imageInfo&gt;.md) | Optional. JSON object used to represent an icon which represents the application used to generate the activity | [optional] 
**BackgroundColor** | Pointer to **NullableString** | Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color | [optional] 
**Content** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI | [optional] 
**Description** | Pointer to **NullableString** | Optional. Longer text description of the user&#39;s unique activity (example: document name, first sentence, and/or metadata) | [optional] 
**DisplayText** | Pointer to **string** | Required. Short text description of the user&#39;s unique activity (for example, document name in cases where an activity refers to document creation) | [optional] 

## Methods

### NewMicrosoftGraphVisualInfo

`func NewMicrosoftGraphVisualInfo() *MicrosoftGraphVisualInfo`

NewMicrosoftGraphVisualInfo instantiates a new MicrosoftGraphVisualInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphVisualInfoWithDefaults

`func NewMicrosoftGraphVisualInfoWithDefaults() *MicrosoftGraphVisualInfo`

NewMicrosoftGraphVisualInfoWithDefaults instantiates a new MicrosoftGraphVisualInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribution

`func (o *MicrosoftGraphVisualInfo) GetAttribution() AnyOfmicrosoftGraphImageInfo`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *MicrosoftGraphVisualInfo) GetAttributionOk() (*AnyOfmicrosoftGraphImageInfo, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *MicrosoftGraphVisualInfo) SetAttribution(v AnyOfmicrosoftGraphImageInfo)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *MicrosoftGraphVisualInfo) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### SetAttributionNil

`func (o *MicrosoftGraphVisualInfo) SetAttributionNil(b bool)`

 SetAttributionNil sets the value for Attribution to be an explicit nil

### UnsetAttribution
`func (o *MicrosoftGraphVisualInfo) UnsetAttribution()`

UnsetAttribution ensures that no value is present for Attribution, not even an explicit nil
### GetBackgroundColor

`func (o *MicrosoftGraphVisualInfo) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *MicrosoftGraphVisualInfo) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *MicrosoftGraphVisualInfo) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *MicrosoftGraphVisualInfo) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *MicrosoftGraphVisualInfo) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *MicrosoftGraphVisualInfo) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphVisualInfo) GetContent() AnyOfobject`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphVisualInfo) GetContentOk() (*AnyOfobject, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphVisualInfo) SetContent(v AnyOfobject)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphVisualInfo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphVisualInfo) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphVisualInfo) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphVisualInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphVisualInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphVisualInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphVisualInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphVisualInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphVisualInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayText

`func (o *MicrosoftGraphVisualInfo) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *MicrosoftGraphVisualInfo) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *MicrosoftGraphVisualInfo) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *MicrosoftGraphVisualInfo) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


