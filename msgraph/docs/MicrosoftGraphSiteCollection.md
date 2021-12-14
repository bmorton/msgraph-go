# MicrosoftGraphSiteCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataLocationCode** | Pointer to **NullableString** | The geographic region code for where this site collection resides. Read-only. | [optional] 
**Hostname** | Pointer to **NullableString** | The hostname for the site collection. Read-only. | [optional] 
**Root** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | If present, indicates that this is a root site collection in SharePoint. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphSiteCollection

`func NewMicrosoftGraphSiteCollection() *MicrosoftGraphSiteCollection`

NewMicrosoftGraphSiteCollection instantiates a new MicrosoftGraphSiteCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSiteCollectionWithDefaults

`func NewMicrosoftGraphSiteCollectionWithDefaults() *MicrosoftGraphSiteCollection`

NewMicrosoftGraphSiteCollectionWithDefaults instantiates a new MicrosoftGraphSiteCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataLocationCode

`func (o *MicrosoftGraphSiteCollection) GetDataLocationCode() string`

GetDataLocationCode returns the DataLocationCode field if non-nil, zero value otherwise.

### GetDataLocationCodeOk

`func (o *MicrosoftGraphSiteCollection) GetDataLocationCodeOk() (*string, bool)`

GetDataLocationCodeOk returns a tuple with the DataLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLocationCode

`func (o *MicrosoftGraphSiteCollection) SetDataLocationCode(v string)`

SetDataLocationCode sets DataLocationCode field to given value.

### HasDataLocationCode

`func (o *MicrosoftGraphSiteCollection) HasDataLocationCode() bool`

HasDataLocationCode returns a boolean if a field has been set.

### SetDataLocationCodeNil

`func (o *MicrosoftGraphSiteCollection) SetDataLocationCodeNil(b bool)`

 SetDataLocationCodeNil sets the value for DataLocationCode to be an explicit nil

### UnsetDataLocationCode
`func (o *MicrosoftGraphSiteCollection) UnsetDataLocationCode()`

UnsetDataLocationCode ensures that no value is present for DataLocationCode, not even an explicit nil
### GetHostname

`func (o *MicrosoftGraphSiteCollection) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *MicrosoftGraphSiteCollection) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *MicrosoftGraphSiteCollection) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *MicrosoftGraphSiteCollection) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *MicrosoftGraphSiteCollection) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *MicrosoftGraphSiteCollection) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetRoot

`func (o *MicrosoftGraphSiteCollection) GetRoot() AnyOfobject`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *MicrosoftGraphSiteCollection) GetRootOk() (*AnyOfobject, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *MicrosoftGraphSiteCollection) SetRoot(v AnyOfobject)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *MicrosoftGraphSiteCollection) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### SetRootNil

`func (o *MicrosoftGraphSiteCollection) SetRootNil(b bool)`

 SetRootNil sets the value for Root to be an explicit nil

### UnsetRoot
`func (o *MicrosoftGraphSiteCollection) UnsetRoot()`

UnsetRoot ensures that no value is present for Root, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


