# MicrosoftGraphInferenceClassificationOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClassifyAs** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) | Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other. | [optional] 
**SenderEmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The email address information of the sender for whom the override is created. | [optional] 

## Methods

### NewMicrosoftGraphInferenceClassificationOverride

`func NewMicrosoftGraphInferenceClassificationOverride() *MicrosoftGraphInferenceClassificationOverride`

NewMicrosoftGraphInferenceClassificationOverride instantiates a new MicrosoftGraphInferenceClassificationOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInferenceClassificationOverrideWithDefaults

`func NewMicrosoftGraphInferenceClassificationOverrideWithDefaults() *MicrosoftGraphInferenceClassificationOverride`

NewMicrosoftGraphInferenceClassificationOverrideWithDefaults instantiates a new MicrosoftGraphInferenceClassificationOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphInferenceClassificationOverride) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphInferenceClassificationOverride) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphInferenceClassificationOverride) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAs() AnyOfmicrosoftGraphInferenceClassificationType`

GetClassifyAs returns the ClassifyAs field if non-nil, zero value otherwise.

### GetClassifyAsOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAsOk() (*AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetClassifyAsOk returns a tuple with the ClassifyAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) SetClassifyAs(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetClassifyAs sets ClassifyAs field to given value.

### HasClassifyAs

`func (o *MicrosoftGraphInferenceClassificationOverride) HasClassifyAs() bool`

HasClassifyAs returns a boolean if a field has been set.

### SetClassifyAsNil

`func (o *MicrosoftGraphInferenceClassificationOverride) SetClassifyAsNil(b bool)`

 SetClassifyAsNil sets the value for ClassifyAs to be an explicit nil

### UnsetClassifyAs
`func (o *MicrosoftGraphInferenceClassificationOverride) UnsetClassifyAs()`

UnsetClassifyAs ensures that no value is present for ClassifyAs, not even an explicit nil
### GetSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) SetSenderEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetSenderEmailAddress sets SenderEmailAddress field to given value.

### HasSenderEmailAddress

`func (o *MicrosoftGraphInferenceClassificationOverride) HasSenderEmailAddress() bool`

HasSenderEmailAddress returns a boolean if a field has been set.

### SetSenderEmailAddressNil

`func (o *MicrosoftGraphInferenceClassificationOverride) SetSenderEmailAddressNil(b bool)`

 SetSenderEmailAddressNil sets the value for SenderEmailAddress to be an explicit nil

### UnsetSenderEmailAddress
`func (o *MicrosoftGraphInferenceClassificationOverride) UnsetSenderEmailAddress()`

UnsetSenderEmailAddress ensures that no value is present for SenderEmailAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


