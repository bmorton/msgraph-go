# InlineObject6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyCredential** | Pointer to [**MicrosoftGraphKeyCredential**](MicrosoftGraphKeyCredential.md) |  | [optional] 
**PasswordCredential** | Pointer to [**AnyOfmicrosoftGraphPasswordCredential**](anyOf&lt;microsoft.graph.passwordCredential&gt;.md) |  | [optional] 
**Proof** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineObject6

`func NewInlineObject6() *InlineObject6`

NewInlineObject6 instantiates a new InlineObject6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject6WithDefaults

`func NewInlineObject6WithDefaults() *InlineObject6`

NewInlineObject6WithDefaults instantiates a new InlineObject6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyCredential

`func (o *InlineObject6) GetKeyCredential() MicrosoftGraphKeyCredential`

GetKeyCredential returns the KeyCredential field if non-nil, zero value otherwise.

### GetKeyCredentialOk

`func (o *InlineObject6) GetKeyCredentialOk() (*MicrosoftGraphKeyCredential, bool)`

GetKeyCredentialOk returns a tuple with the KeyCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCredential

`func (o *InlineObject6) SetKeyCredential(v MicrosoftGraphKeyCredential)`

SetKeyCredential sets KeyCredential field to given value.

### HasKeyCredential

`func (o *InlineObject6) HasKeyCredential() bool`

HasKeyCredential returns a boolean if a field has been set.

### GetPasswordCredential

`func (o *InlineObject6) GetPasswordCredential() AnyOfmicrosoftGraphPasswordCredential`

GetPasswordCredential returns the PasswordCredential field if non-nil, zero value otherwise.

### GetPasswordCredentialOk

`func (o *InlineObject6) GetPasswordCredentialOk() (*AnyOfmicrosoftGraphPasswordCredential, bool)`

GetPasswordCredentialOk returns a tuple with the PasswordCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCredential

`func (o *InlineObject6) SetPasswordCredential(v AnyOfmicrosoftGraphPasswordCredential)`

SetPasswordCredential sets PasswordCredential field to given value.

### HasPasswordCredential

`func (o *InlineObject6) HasPasswordCredential() bool`

HasPasswordCredential returns a boolean if a field has been set.

### SetPasswordCredentialNil

`func (o *InlineObject6) SetPasswordCredentialNil(b bool)`

 SetPasswordCredentialNil sets the value for PasswordCredential to be an explicit nil

### UnsetPasswordCredential
`func (o *InlineObject6) UnsetPasswordCredential()`

UnsetPasswordCredential ensures that no value is present for PasswordCredential, not even an explicit nil
### GetProof

`func (o *InlineObject6) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *InlineObject6) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *InlineObject6) SetProof(v string)`

SetProof sets Proof field to given value.

### HasProof

`func (o *InlineObject6) HasProof() bool`

HasProof returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


