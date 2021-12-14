# MicrosoftGraphInferenceClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Overrides** | Pointer to [**[]MicrosoftGraphInferenceClassificationOverride**](MicrosoftGraphInferenceClassificationOverride.md) | A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphInferenceClassification

`func NewMicrosoftGraphInferenceClassification() *MicrosoftGraphInferenceClassification`

NewMicrosoftGraphInferenceClassification instantiates a new MicrosoftGraphInferenceClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphInferenceClassificationWithDefaults

`func NewMicrosoftGraphInferenceClassificationWithDefaults() *MicrosoftGraphInferenceClassification`

NewMicrosoftGraphInferenceClassificationWithDefaults instantiates a new MicrosoftGraphInferenceClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphInferenceClassification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphInferenceClassification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphInferenceClassification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphInferenceClassification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOverrides

`func (o *MicrosoftGraphInferenceClassification) GetOverrides() []MicrosoftGraphInferenceClassificationOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *MicrosoftGraphInferenceClassification) GetOverridesOk() (*[]MicrosoftGraphInferenceClassificationOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *MicrosoftGraphInferenceClassification) SetOverrides(v []MicrosoftGraphInferenceClassificationOverride)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *MicrosoftGraphInferenceClassification) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


