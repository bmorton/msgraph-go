# MicrosoftGraphIdentityUserFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**UserFlowType** | Pointer to [**AnyOfmicrosoftGraphUserFlowType**](anyOf&lt;microsoft.graph.userFlowType&gt;.md) |  | [optional] 
**UserFlowTypeVersion** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphIdentityUserFlow

`func NewMicrosoftGraphIdentityUserFlow() *MicrosoftGraphIdentityUserFlow`

NewMicrosoftGraphIdentityUserFlow instantiates a new MicrosoftGraphIdentityUserFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityUserFlowWithDefaults

`func NewMicrosoftGraphIdentityUserFlowWithDefaults() *MicrosoftGraphIdentityUserFlow`

NewMicrosoftGraphIdentityUserFlowWithDefaults instantiates a new MicrosoftGraphIdentityUserFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentityUserFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentityUserFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentityUserFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentityUserFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserFlowType

`func (o *MicrosoftGraphIdentityUserFlow) GetUserFlowType() AnyOfmicrosoftGraphUserFlowType`

GetUserFlowType returns the UserFlowType field if non-nil, zero value otherwise.

### GetUserFlowTypeOk

`func (o *MicrosoftGraphIdentityUserFlow) GetUserFlowTypeOk() (*AnyOfmicrosoftGraphUserFlowType, bool)`

GetUserFlowTypeOk returns a tuple with the UserFlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowType

`func (o *MicrosoftGraphIdentityUserFlow) SetUserFlowType(v AnyOfmicrosoftGraphUserFlowType)`

SetUserFlowType sets UserFlowType field to given value.

### HasUserFlowType

`func (o *MicrosoftGraphIdentityUserFlow) HasUserFlowType() bool`

HasUserFlowType returns a boolean if a field has been set.

### SetUserFlowTypeNil

`func (o *MicrosoftGraphIdentityUserFlow) SetUserFlowTypeNil(b bool)`

 SetUserFlowTypeNil sets the value for UserFlowType to be an explicit nil

### UnsetUserFlowType
`func (o *MicrosoftGraphIdentityUserFlow) UnsetUserFlowType()`

UnsetUserFlowType ensures that no value is present for UserFlowType, not even an explicit nil
### GetUserFlowTypeVersion

`func (o *MicrosoftGraphIdentityUserFlow) GetUserFlowTypeVersion() AnyOfnumberstringstring`

GetUserFlowTypeVersion returns the UserFlowTypeVersion field if non-nil, zero value otherwise.

### GetUserFlowTypeVersionOk

`func (o *MicrosoftGraphIdentityUserFlow) GetUserFlowTypeVersionOk() (*AnyOfnumberstringstring, bool)`

GetUserFlowTypeVersionOk returns a tuple with the UserFlowTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowTypeVersion

`func (o *MicrosoftGraphIdentityUserFlow) SetUserFlowTypeVersion(v AnyOfnumberstringstring)`

SetUserFlowTypeVersion sets UserFlowTypeVersion field to given value.

### HasUserFlowTypeVersion

`func (o *MicrosoftGraphIdentityUserFlow) HasUserFlowTypeVersion() bool`

HasUserFlowTypeVersion returns a boolean if a field has been set.

### SetUserFlowTypeVersionNil

`func (o *MicrosoftGraphIdentityUserFlow) SetUserFlowTypeVersionNil(b bool)`

 SetUserFlowTypeVersionNil sets the value for UserFlowTypeVersion to be an explicit nil

### UnsetUserFlowTypeVersion
`func (o *MicrosoftGraphIdentityUserFlow) UnsetUserFlowTypeVersion()`

UnsetUserFlowTypeVersion ensures that no value is present for UserFlowTypeVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


