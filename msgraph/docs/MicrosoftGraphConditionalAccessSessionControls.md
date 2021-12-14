# MicrosoftGraphConditionalAccessSessionControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationEnforcedRestrictions** | Pointer to [**AnyOfmicrosoftGraphApplicationEnforcedRestrictionsSessionControl**](anyOf&lt;microsoft.graph.applicationEnforcedRestrictionsSessionControl&gt;.md) | Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control. | [optional] 
**CloudAppSecurity** | Pointer to [**AnyOfmicrosoftGraphCloudAppSecuritySessionControl**](anyOf&lt;microsoft.graph.cloudAppSecuritySessionControl&gt;.md) | Session control to apply cloud app security. | [optional] 
**PersistentBrowser** | Pointer to [**AnyOfmicrosoftGraphPersistentBrowserSessionControl**](anyOf&lt;microsoft.graph.persistentBrowserSessionControl&gt;.md) | Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly. | [optional] 
**SignInFrequency** | Pointer to [**AnyOfmicrosoftGraphSignInFrequencySessionControl**](anyOf&lt;microsoft.graph.signInFrequencySessionControl&gt;.md) | Session control to enforce signin frequency. | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessSessionControls

`func NewMicrosoftGraphConditionalAccessSessionControls() *MicrosoftGraphConditionalAccessSessionControls`

NewMicrosoftGraphConditionalAccessSessionControls instantiates a new MicrosoftGraphConditionalAccessSessionControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessSessionControlsWithDefaults

`func NewMicrosoftGraphConditionalAccessSessionControlsWithDefaults() *MicrosoftGraphConditionalAccessSessionControls`

NewMicrosoftGraphConditionalAccessSessionControlsWithDefaults instantiates a new MicrosoftGraphConditionalAccessSessionControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationEnforcedRestrictions

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetApplicationEnforcedRestrictions() AnyOfmicrosoftGraphApplicationEnforcedRestrictionsSessionControl`

GetApplicationEnforcedRestrictions returns the ApplicationEnforcedRestrictions field if non-nil, zero value otherwise.

### GetApplicationEnforcedRestrictionsOk

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetApplicationEnforcedRestrictionsOk() (*AnyOfmicrosoftGraphApplicationEnforcedRestrictionsSessionControl, bool)`

GetApplicationEnforcedRestrictionsOk returns a tuple with the ApplicationEnforcedRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnforcedRestrictions

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetApplicationEnforcedRestrictions(v AnyOfmicrosoftGraphApplicationEnforcedRestrictionsSessionControl)`

SetApplicationEnforcedRestrictions sets ApplicationEnforcedRestrictions field to given value.

### HasApplicationEnforcedRestrictions

`func (o *MicrosoftGraphConditionalAccessSessionControls) HasApplicationEnforcedRestrictions() bool`

HasApplicationEnforcedRestrictions returns a boolean if a field has been set.

### SetApplicationEnforcedRestrictionsNil

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetApplicationEnforcedRestrictionsNil(b bool)`

 SetApplicationEnforcedRestrictionsNil sets the value for ApplicationEnforcedRestrictions to be an explicit nil

### UnsetApplicationEnforcedRestrictions
`func (o *MicrosoftGraphConditionalAccessSessionControls) UnsetApplicationEnforcedRestrictions()`

UnsetApplicationEnforcedRestrictions ensures that no value is present for ApplicationEnforcedRestrictions, not even an explicit nil
### GetCloudAppSecurity

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetCloudAppSecurity() AnyOfmicrosoftGraphCloudAppSecuritySessionControl`

GetCloudAppSecurity returns the CloudAppSecurity field if non-nil, zero value otherwise.

### GetCloudAppSecurityOk

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetCloudAppSecurityOk() (*AnyOfmicrosoftGraphCloudAppSecuritySessionControl, bool)`

GetCloudAppSecurityOk returns a tuple with the CloudAppSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAppSecurity

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetCloudAppSecurity(v AnyOfmicrosoftGraphCloudAppSecuritySessionControl)`

SetCloudAppSecurity sets CloudAppSecurity field to given value.

### HasCloudAppSecurity

`func (o *MicrosoftGraphConditionalAccessSessionControls) HasCloudAppSecurity() bool`

HasCloudAppSecurity returns a boolean if a field has been set.

### SetCloudAppSecurityNil

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetCloudAppSecurityNil(b bool)`

 SetCloudAppSecurityNil sets the value for CloudAppSecurity to be an explicit nil

### UnsetCloudAppSecurity
`func (o *MicrosoftGraphConditionalAccessSessionControls) UnsetCloudAppSecurity()`

UnsetCloudAppSecurity ensures that no value is present for CloudAppSecurity, not even an explicit nil
### GetPersistentBrowser

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetPersistentBrowser() AnyOfmicrosoftGraphPersistentBrowserSessionControl`

GetPersistentBrowser returns the PersistentBrowser field if non-nil, zero value otherwise.

### GetPersistentBrowserOk

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetPersistentBrowserOk() (*AnyOfmicrosoftGraphPersistentBrowserSessionControl, bool)`

GetPersistentBrowserOk returns a tuple with the PersistentBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentBrowser

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetPersistentBrowser(v AnyOfmicrosoftGraphPersistentBrowserSessionControl)`

SetPersistentBrowser sets PersistentBrowser field to given value.

### HasPersistentBrowser

`func (o *MicrosoftGraphConditionalAccessSessionControls) HasPersistentBrowser() bool`

HasPersistentBrowser returns a boolean if a field has been set.

### SetPersistentBrowserNil

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetPersistentBrowserNil(b bool)`

 SetPersistentBrowserNil sets the value for PersistentBrowser to be an explicit nil

### UnsetPersistentBrowser
`func (o *MicrosoftGraphConditionalAccessSessionControls) UnsetPersistentBrowser()`

UnsetPersistentBrowser ensures that no value is present for PersistentBrowser, not even an explicit nil
### GetSignInFrequency

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetSignInFrequency() AnyOfmicrosoftGraphSignInFrequencySessionControl`

GetSignInFrequency returns the SignInFrequency field if non-nil, zero value otherwise.

### GetSignInFrequencyOk

`func (o *MicrosoftGraphConditionalAccessSessionControls) GetSignInFrequencyOk() (*AnyOfmicrosoftGraphSignInFrequencySessionControl, bool)`

GetSignInFrequencyOk returns a tuple with the SignInFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInFrequency

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetSignInFrequency(v AnyOfmicrosoftGraphSignInFrequencySessionControl)`

SetSignInFrequency sets SignInFrequency field to given value.

### HasSignInFrequency

`func (o *MicrosoftGraphConditionalAccessSessionControls) HasSignInFrequency() bool`

HasSignInFrequency returns a boolean if a field has been set.

### SetSignInFrequencyNil

`func (o *MicrosoftGraphConditionalAccessSessionControls) SetSignInFrequencyNil(b bool)`

 SetSignInFrequencyNil sets the value for SignInFrequency to be an explicit nil

### UnsetSignInFrequency
`func (o *MicrosoftGraphConditionalAccessSessionControls) UnsetSignInFrequency()`

UnsetSignInFrequency ensures that no value is present for SignInFrequency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


