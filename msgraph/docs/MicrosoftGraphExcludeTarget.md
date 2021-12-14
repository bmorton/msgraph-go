# MicrosoftGraphExcludeTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The object identifier of an Azure Active Directory user or group. | [optional] 
**TargetType** | Pointer to [**AnyOfmicrosoftGraphAuthenticationMethodTargetType**](anyOf&lt;microsoft.graph.authenticationMethodTargetType&gt;.md) | The type of the authentication method target. Possible values are: user, group, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphExcludeTarget

`func NewMicrosoftGraphExcludeTarget() *MicrosoftGraphExcludeTarget`

NewMicrosoftGraphExcludeTarget instantiates a new MicrosoftGraphExcludeTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExcludeTargetWithDefaults

`func NewMicrosoftGraphExcludeTargetWithDefaults() *MicrosoftGraphExcludeTarget`

NewMicrosoftGraphExcludeTargetWithDefaults instantiates a new MicrosoftGraphExcludeTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphExcludeTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphExcludeTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphExcludeTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphExcludeTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetType

`func (o *MicrosoftGraphExcludeTarget) GetTargetType() AnyOfmicrosoftGraphAuthenticationMethodTargetType`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *MicrosoftGraphExcludeTarget) GetTargetTypeOk() (*AnyOfmicrosoftGraphAuthenticationMethodTargetType, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *MicrosoftGraphExcludeTarget) SetTargetType(v AnyOfmicrosoftGraphAuthenticationMethodTargetType)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *MicrosoftGraphExcludeTarget) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *MicrosoftGraphExcludeTarget) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *MicrosoftGraphExcludeTarget) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


