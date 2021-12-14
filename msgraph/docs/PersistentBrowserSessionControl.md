# PersistentBrowserSessionControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**AnyOfmicrosoftGraphPersistentBrowserSessionMode**](anyOf&lt;microsoft.graph.persistentBrowserSessionMode&gt;.md) | Possible values are: always, never. | [optional] 

## Methods

### NewPersistentBrowserSessionControl

`func NewPersistentBrowserSessionControl() *PersistentBrowserSessionControl`

NewPersistentBrowserSessionControl instantiates a new PersistentBrowserSessionControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentBrowserSessionControlWithDefaults

`func NewPersistentBrowserSessionControlWithDefaults() *PersistentBrowserSessionControl`

NewPersistentBrowserSessionControlWithDefaults instantiates a new PersistentBrowserSessionControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PersistentBrowserSessionControl) GetMode() AnyOfmicrosoftGraphPersistentBrowserSessionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PersistentBrowserSessionControl) GetModeOk() (*AnyOfmicrosoftGraphPersistentBrowserSessionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PersistentBrowserSessionControl) SetMode(v AnyOfmicrosoftGraphPersistentBrowserSessionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PersistentBrowserSessionControl) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *PersistentBrowserSessionControl) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *PersistentBrowserSessionControl) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


