# MicrosoftGraphApplicationServicePrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphApplication**](anyOf&lt;microsoft.graph.application&gt;.md) |  | [optional] 
**ServicePrincipal** | Pointer to [**AnyOfmicrosoftGraphServicePrincipal**](anyOf&lt;microsoft.graph.servicePrincipal&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphApplicationServicePrincipal

`func NewMicrosoftGraphApplicationServicePrincipal() *MicrosoftGraphApplicationServicePrincipal`

NewMicrosoftGraphApplicationServicePrincipal instantiates a new MicrosoftGraphApplicationServicePrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphApplicationServicePrincipalWithDefaults

`func NewMicrosoftGraphApplicationServicePrincipalWithDefaults() *MicrosoftGraphApplicationServicePrincipal`

NewMicrosoftGraphApplicationServicePrincipalWithDefaults instantiates a new MicrosoftGraphApplicationServicePrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *MicrosoftGraphApplicationServicePrincipal) GetApplication() AnyOfmicrosoftGraphApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MicrosoftGraphApplicationServicePrincipal) GetApplicationOk() (*AnyOfmicrosoftGraphApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MicrosoftGraphApplicationServicePrincipal) SetApplication(v AnyOfmicrosoftGraphApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MicrosoftGraphApplicationServicePrincipal) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MicrosoftGraphApplicationServicePrincipal) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MicrosoftGraphApplicationServicePrincipal) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetServicePrincipal

`func (o *MicrosoftGraphApplicationServicePrincipal) GetServicePrincipal() AnyOfmicrosoftGraphServicePrincipal`

GetServicePrincipal returns the ServicePrincipal field if non-nil, zero value otherwise.

### GetServicePrincipalOk

`func (o *MicrosoftGraphApplicationServicePrincipal) GetServicePrincipalOk() (*AnyOfmicrosoftGraphServicePrincipal, bool)`

GetServicePrincipalOk returns a tuple with the ServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipal

`func (o *MicrosoftGraphApplicationServicePrincipal) SetServicePrincipal(v AnyOfmicrosoftGraphServicePrincipal)`

SetServicePrincipal sets ServicePrincipal field to given value.

### HasServicePrincipal

`func (o *MicrosoftGraphApplicationServicePrincipal) HasServicePrincipal() bool`

HasServicePrincipal returns a boolean if a field has been set.

### SetServicePrincipalNil

`func (o *MicrosoftGraphApplicationServicePrincipal) SetServicePrincipalNil(b bool)`

 SetServicePrincipalNil sets the value for ServicePrincipal to be an explicit nil

### UnsetServicePrincipal
`func (o *MicrosoftGraphApplicationServicePrincipal) UnsetServicePrincipal()`

UnsetServicePrincipal ensures that no value is present for ServicePrincipal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


