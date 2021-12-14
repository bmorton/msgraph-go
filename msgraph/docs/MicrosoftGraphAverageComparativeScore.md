# MicrosoftGraphAverageComparativeScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageScore** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Average score within specified basis. | [optional] 
**Basis** | Pointer to **NullableString** | Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes. | [optional] 

## Methods

### NewMicrosoftGraphAverageComparativeScore

`func NewMicrosoftGraphAverageComparativeScore() *MicrosoftGraphAverageComparativeScore`

NewMicrosoftGraphAverageComparativeScore instantiates a new MicrosoftGraphAverageComparativeScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAverageComparativeScoreWithDefaults

`func NewMicrosoftGraphAverageComparativeScoreWithDefaults() *MicrosoftGraphAverageComparativeScore`

NewMicrosoftGraphAverageComparativeScoreWithDefaults instantiates a new MicrosoftGraphAverageComparativeScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageScore

`func (o *MicrosoftGraphAverageComparativeScore) GetAverageScore() AnyOfnumberstringstring`

GetAverageScore returns the AverageScore field if non-nil, zero value otherwise.

### GetAverageScoreOk

`func (o *MicrosoftGraphAverageComparativeScore) GetAverageScoreOk() (*AnyOfnumberstringstring, bool)`

GetAverageScoreOk returns a tuple with the AverageScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageScore

`func (o *MicrosoftGraphAverageComparativeScore) SetAverageScore(v AnyOfnumberstringstring)`

SetAverageScore sets AverageScore field to given value.

### HasAverageScore

`func (o *MicrosoftGraphAverageComparativeScore) HasAverageScore() bool`

HasAverageScore returns a boolean if a field has been set.

### SetAverageScoreNil

`func (o *MicrosoftGraphAverageComparativeScore) SetAverageScoreNil(b bool)`

 SetAverageScoreNil sets the value for AverageScore to be an explicit nil

### UnsetAverageScore
`func (o *MicrosoftGraphAverageComparativeScore) UnsetAverageScore()`

UnsetAverageScore ensures that no value is present for AverageScore, not even an explicit nil
### GetBasis

`func (o *MicrosoftGraphAverageComparativeScore) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *MicrosoftGraphAverageComparativeScore) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *MicrosoftGraphAverageComparativeScore) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *MicrosoftGraphAverageComparativeScore) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### SetBasisNil

`func (o *MicrosoftGraphAverageComparativeScore) SetBasisNil(b bool)`

 SetBasisNil sets the value for Basis to be an explicit nil

### UnsetBasis
`func (o *MicrosoftGraphAverageComparativeScore) UnsetBasis()`

UnsetBasis ensures that no value is present for Basis, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


