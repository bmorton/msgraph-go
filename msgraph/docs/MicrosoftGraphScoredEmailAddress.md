# MicrosoftGraphScoredEmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** | The email address. | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**RelevanceScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships. | [optional] 
**SelectionLikelihood** | Pointer to [**AnyOfmicrosoftGraphSelectionLikelihoodInfo**](anyOf&lt;microsoft.graph.selectionLikelihoodInfo&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphScoredEmailAddress

`func NewMicrosoftGraphScoredEmailAddress() *MicrosoftGraphScoredEmailAddress`

NewMicrosoftGraphScoredEmailAddress instantiates a new MicrosoftGraphScoredEmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphScoredEmailAddressWithDefaults

`func NewMicrosoftGraphScoredEmailAddressWithDefaults() *MicrosoftGraphScoredEmailAddress`

NewMicrosoftGraphScoredEmailAddressWithDefaults instantiates a new MicrosoftGraphScoredEmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MicrosoftGraphScoredEmailAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphScoredEmailAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphScoredEmailAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphScoredEmailAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphScoredEmailAddress) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphScoredEmailAddress) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetItemId

`func (o *MicrosoftGraphScoredEmailAddress) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *MicrosoftGraphScoredEmailAddress) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *MicrosoftGraphScoredEmailAddress) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *MicrosoftGraphScoredEmailAddress) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *MicrosoftGraphScoredEmailAddress) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *MicrosoftGraphScoredEmailAddress) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) GetRelevanceScore() AnyOfnumberstringstring`

GetRelevanceScore returns the RelevanceScore field if non-nil, zero value otherwise.

### GetRelevanceScoreOk

`func (o *MicrosoftGraphScoredEmailAddress) GetRelevanceScoreOk() (*AnyOfnumberstringstring, bool)`

GetRelevanceScoreOk returns a tuple with the RelevanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) SetRelevanceScore(v AnyOfnumberstringstring)`

SetRelevanceScore sets RelevanceScore field to given value.

### HasRelevanceScore

`func (o *MicrosoftGraphScoredEmailAddress) HasRelevanceScore() bool`

HasRelevanceScore returns a boolean if a field has been set.

### SetRelevanceScoreNil

`func (o *MicrosoftGraphScoredEmailAddress) SetRelevanceScoreNil(b bool)`

 SetRelevanceScoreNil sets the value for RelevanceScore to be an explicit nil

### UnsetRelevanceScore
`func (o *MicrosoftGraphScoredEmailAddress) UnsetRelevanceScore()`

UnsetRelevanceScore ensures that no value is present for RelevanceScore, not even an explicit nil
### GetSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) GetSelectionLikelihood() AnyOfmicrosoftGraphSelectionLikelihoodInfo`

GetSelectionLikelihood returns the SelectionLikelihood field if non-nil, zero value otherwise.

### GetSelectionLikelihoodOk

`func (o *MicrosoftGraphScoredEmailAddress) GetSelectionLikelihoodOk() (*AnyOfmicrosoftGraphSelectionLikelihoodInfo, bool)`

GetSelectionLikelihoodOk returns a tuple with the SelectionLikelihood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) SetSelectionLikelihood(v AnyOfmicrosoftGraphSelectionLikelihoodInfo)`

SetSelectionLikelihood sets SelectionLikelihood field to given value.

### HasSelectionLikelihood

`func (o *MicrosoftGraphScoredEmailAddress) HasSelectionLikelihood() bool`

HasSelectionLikelihood returns a boolean if a field has been set.

### SetSelectionLikelihoodNil

`func (o *MicrosoftGraphScoredEmailAddress) SetSelectionLikelihoodNil(b bool)`

 SetSelectionLikelihoodNil sets the value for SelectionLikelihood to be an explicit nil

### UnsetSelectionLikelihood
`func (o *MicrosoftGraphScoredEmailAddress) UnsetSelectionLikelihood()`

UnsetSelectionLikelihood ensures that no value is present for SelectionLikelihood, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


