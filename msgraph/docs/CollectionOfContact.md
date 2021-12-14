# CollectionOfContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**[]MicrosoftGraphContact**](MicrosoftGraphContact.md) |  | [optional] 
**OdataNextLink** | Pointer to **string** |  | [optional] 

## Methods

### NewCollectionOfContact

`func NewCollectionOfContact() *CollectionOfContact`

NewCollectionOfContact instantiates a new CollectionOfContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionOfContactWithDefaults

`func NewCollectionOfContactWithDefaults() *CollectionOfContact`

NewCollectionOfContactWithDefaults instantiates a new CollectionOfContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CollectionOfContact) GetValue() []MicrosoftGraphContact`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CollectionOfContact) GetValueOk() (*[]MicrosoftGraphContact, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CollectionOfContact) SetValue(v []MicrosoftGraphContact)`

SetValue sets Value field to given value.

### HasValue

`func (o *CollectionOfContact) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOdataNextLink

`func (o *CollectionOfContact) GetOdataNextLink() string`

GetOdataNextLink returns the OdataNextLink field if non-nil, zero value otherwise.

### GetOdataNextLinkOk

`func (o *CollectionOfContact) GetOdataNextLinkOk() (*string, bool)`

GetOdataNextLinkOk returns a tuple with the OdataNextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdataNextLink

`func (o *CollectionOfContact) SetOdataNextLink(v string)`

SetOdataNextLink sets OdataNextLink field to given value.

### HasOdataNextLink

`func (o *CollectionOfContact) HasOdataNextLink() bool`

HasOdataNextLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


