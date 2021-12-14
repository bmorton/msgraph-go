# Initiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorType** | Pointer to [**AnyOfmicrosoftGraphInitiatorType**](anyOf&lt;microsoft.graph.initiatorType&gt;.md) | Type of initiator. Possible values are: user, application, system, unknownFutureValue. | [optional] 

## Methods

### NewInitiator

`func NewInitiator() *Initiator`

NewInitiator instantiates a new Initiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiatorWithDefaults

`func NewInitiatorWithDefaults() *Initiator`

NewInitiatorWithDefaults instantiates a new Initiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorType

`func (o *Initiator) GetInitiatorType() AnyOfmicrosoftGraphInitiatorType`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *Initiator) GetInitiatorTypeOk() (*AnyOfmicrosoftGraphInitiatorType, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *Initiator) SetInitiatorType(v AnyOfmicrosoftGraphInitiatorType)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *Initiator) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### SetInitiatorTypeNil

`func (o *Initiator) SetInitiatorTypeNil(b bool)`

 SetInitiatorTypeNil sets the value for InitiatorType to be an explicit nil

### UnsetInitiatorType
`func (o *Initiator) UnsetInitiatorType()`

UnsetInitiatorType ensures that no value is present for InitiatorType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


