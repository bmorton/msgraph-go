# InferenceClassificationOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassifyAs** | Pointer to [**AnyOfmicrosoftGraphInferenceClassificationType**](anyOf&lt;microsoft.graph.inferenceClassificationType&gt;.md) | Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other. | [optional] 
**SenderEmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The email address information of the sender for whom the override is created. | [optional] 

## Methods

### NewInferenceClassificationOverride

`func NewInferenceClassificationOverride() *InferenceClassificationOverride`

NewInferenceClassificationOverride instantiates a new InferenceClassificationOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceClassificationOverrideWithDefaults

`func NewInferenceClassificationOverrideWithDefaults() *InferenceClassificationOverride`

NewInferenceClassificationOverrideWithDefaults instantiates a new InferenceClassificationOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifyAs

`func (o *InferenceClassificationOverride) GetClassifyAs() AnyOfmicrosoftGraphInferenceClassificationType`

GetClassifyAs returns the ClassifyAs field if non-nil, zero value otherwise.

### GetClassifyAsOk

`func (o *InferenceClassificationOverride) GetClassifyAsOk() (*AnyOfmicrosoftGraphInferenceClassificationType, bool)`

GetClassifyAsOk returns a tuple with the ClassifyAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifyAs

`func (o *InferenceClassificationOverride) SetClassifyAs(v AnyOfmicrosoftGraphInferenceClassificationType)`

SetClassifyAs sets ClassifyAs field to given value.

### HasClassifyAs

`func (o *InferenceClassificationOverride) HasClassifyAs() bool`

HasClassifyAs returns a boolean if a field has been set.

### SetClassifyAsNil

`func (o *InferenceClassificationOverride) SetClassifyAsNil(b bool)`

 SetClassifyAsNil sets the value for ClassifyAs to be an explicit nil

### UnsetClassifyAs
`func (o *InferenceClassificationOverride) UnsetClassifyAs()`

UnsetClassifyAs ensures that no value is present for ClassifyAs, not even an explicit nil
### GetSenderEmailAddress

`func (o *InferenceClassificationOverride) GetSenderEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetSenderEmailAddress returns the SenderEmailAddress field if non-nil, zero value otherwise.

### GetSenderEmailAddressOk

`func (o *InferenceClassificationOverride) GetSenderEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEmailAddress

`func (o *InferenceClassificationOverride) SetSenderEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetSenderEmailAddress sets SenderEmailAddress field to given value.

### HasSenderEmailAddress

`func (o *InferenceClassificationOverride) HasSenderEmailAddress() bool`

HasSenderEmailAddress returns a boolean if a field has been set.

### SetSenderEmailAddressNil

`func (o *InferenceClassificationOverride) SetSenderEmailAddressNil(b bool)`

 SetSenderEmailAddressNil sets the value for SenderEmailAddress to be an explicit nil

### UnsetSenderEmailAddress
`func (o *InferenceClassificationOverride) UnsetSenderEmailAddress()`

UnsetSenderEmailAddress ensures that no value is present for SenderEmailAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


