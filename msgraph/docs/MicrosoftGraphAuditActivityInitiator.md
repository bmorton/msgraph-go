# MicrosoftGraphAuditActivityInitiator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AnyOfmicrosoftGraphAppIdentity**](anyOf&lt;microsoft.graph.appIdentity&gt;.md) | If the resource initiating the activity is an app, this property indicates all the app related information like appId, Name, servicePrincipalId, Name. | [optional] 
**User** | Pointer to [**AnyOfmicrosoftGraphUserIdentity**](anyOf&lt;microsoft.graph.userIdentity&gt;.md) | If the resource initiating the activity is a user, this property Indicates all the user related information like userId, Name, UserPrinicpalName. | [optional] 

## Methods

### NewMicrosoftGraphAuditActivityInitiator

`func NewMicrosoftGraphAuditActivityInitiator() *MicrosoftGraphAuditActivityInitiator`

NewMicrosoftGraphAuditActivityInitiator instantiates a new MicrosoftGraphAuditActivityInitiator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuditActivityInitiatorWithDefaults

`func NewMicrosoftGraphAuditActivityInitiatorWithDefaults() *MicrosoftGraphAuditActivityInitiator`

NewMicrosoftGraphAuditActivityInitiatorWithDefaults instantiates a new MicrosoftGraphAuditActivityInitiator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *MicrosoftGraphAuditActivityInitiator) GetApp() AnyOfmicrosoftGraphAppIdentity`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MicrosoftGraphAuditActivityInitiator) GetAppOk() (*AnyOfmicrosoftGraphAppIdentity, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *MicrosoftGraphAuditActivityInitiator) SetApp(v AnyOfmicrosoftGraphAppIdentity)`

SetApp sets App field to given value.

### HasApp

`func (o *MicrosoftGraphAuditActivityInitiator) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *MicrosoftGraphAuditActivityInitiator) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *MicrosoftGraphAuditActivityInitiator) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetUser

`func (o *MicrosoftGraphAuditActivityInitiator) GetUser() AnyOfmicrosoftGraphUserIdentity`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MicrosoftGraphAuditActivityInitiator) GetUserOk() (*AnyOfmicrosoftGraphUserIdentity, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MicrosoftGraphAuditActivityInitiator) SetUser(v AnyOfmicrosoftGraphUserIdentity)`

SetUser sets User field to given value.

### HasUser

`func (o *MicrosoftGraphAuditActivityInitiator) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *MicrosoftGraphAuditActivityInitiator) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *MicrosoftGraphAuditActivityInitiator) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


