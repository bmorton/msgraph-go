# MicrosoftGraphWebsite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** | The URL of the website. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name of the web site. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphWebsiteType**](anyOf&lt;microsoft.graph.websiteType&gt;.md) | The possible values are: other, home, work, blog, profile. | [optional] 

## Methods

### NewMicrosoftGraphWebsite

`func NewMicrosoftGraphWebsite() *MicrosoftGraphWebsite`

NewMicrosoftGraphWebsite instantiates a new MicrosoftGraphWebsite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWebsiteWithDefaults

`func NewMicrosoftGraphWebsiteWithDefaults() *MicrosoftGraphWebsite`

NewMicrosoftGraphWebsiteWithDefaults instantiates a new MicrosoftGraphWebsite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MicrosoftGraphWebsite) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphWebsite) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphWebsite) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphWebsite) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphWebsite) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphWebsite) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphWebsite) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWebsite) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphWebsite) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphWebsite) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphWebsite) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphWebsite) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetType

`func (o *MicrosoftGraphWebsite) GetType() AnyOfmicrosoftGraphWebsiteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphWebsite) GetTypeOk() (*AnyOfmicrosoftGraphWebsiteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphWebsite) SetType(v AnyOfmicrosoftGraphWebsiteType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphWebsite) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphWebsite) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphWebsite) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


