# InferenceClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overrides** | Pointer to [**[]MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md) | A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable. | [optional] 

## Methods

### NewInferenceClassification

`func NewInferenceClassification() *InferenceClassification`

NewInferenceClassification instantiates a new InferenceClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceClassificationWithDefaults

`func NewInferenceClassificationWithDefaults() *InferenceClassification`

NewInferenceClassificationWithDefaults instantiates a new InferenceClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrides

`func (o *InferenceClassification) GetOverrides() []MicrosoftGraphInferenceClassificationOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InferenceClassification) GetOverridesOk() (*[]MicrosoftGraphInferenceClassificationOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InferenceClassification) SetOverrides(v []MicrosoftGraphInferenceClassificationOverride)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InferenceClassification) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


