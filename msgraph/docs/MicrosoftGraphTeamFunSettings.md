# MicrosoftGraphTeamFunSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCustomMemes** | Pointer to **NullableBool** | If set to true, enables users to include custom memes. | [optional] 
**AllowGiphy** | Pointer to **NullableBool** | If set to true, enables Giphy use. | [optional] 
**AllowStickersAndMemes** | Pointer to **NullableBool** | If set to true, enables users to include stickers and memes. | [optional] 
**GiphyContentRating** | Pointer to [**AnyOfmicrosoftGraphGiphyRatingType**](anyOf&lt;microsoft.graph.giphyRatingType&gt;.md) | Giphy content rating. Possible values are: moderate, strict. | [optional] 

## Methods

### NewMicrosoftGraphTeamFunSettings

`func NewMicrosoftGraphTeamFunSettings() *MicrosoftGraphTeamFunSettings`

NewMicrosoftGraphTeamFunSettings instantiates a new MicrosoftGraphTeamFunSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTeamFunSettingsWithDefaults

`func NewMicrosoftGraphTeamFunSettingsWithDefaults() *MicrosoftGraphTeamFunSettings`

NewMicrosoftGraphTeamFunSettingsWithDefaults instantiates a new MicrosoftGraphTeamFunSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCustomMemes

`func (o *MicrosoftGraphTeamFunSettings) GetAllowCustomMemes() bool`

GetAllowCustomMemes returns the AllowCustomMemes field if non-nil, zero value otherwise.

### GetAllowCustomMemesOk

`func (o *MicrosoftGraphTeamFunSettings) GetAllowCustomMemesOk() (*bool, bool)`

GetAllowCustomMemesOk returns a tuple with the AllowCustomMemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomMemes

`func (o *MicrosoftGraphTeamFunSettings) SetAllowCustomMemes(v bool)`

SetAllowCustomMemes sets AllowCustomMemes field to given value.

### HasAllowCustomMemes

`func (o *MicrosoftGraphTeamFunSettings) HasAllowCustomMemes() bool`

HasAllowCustomMemes returns a boolean if a field has been set.

### SetAllowCustomMemesNil

`func (o *MicrosoftGraphTeamFunSettings) SetAllowCustomMemesNil(b bool)`

 SetAllowCustomMemesNil sets the value for AllowCustomMemes to be an explicit nil

### UnsetAllowCustomMemes
`func (o *MicrosoftGraphTeamFunSettings) UnsetAllowCustomMemes()`

UnsetAllowCustomMemes ensures that no value is present for AllowCustomMemes, not even an explicit nil
### GetAllowGiphy

`func (o *MicrosoftGraphTeamFunSettings) GetAllowGiphy() bool`

GetAllowGiphy returns the AllowGiphy field if non-nil, zero value otherwise.

### GetAllowGiphyOk

`func (o *MicrosoftGraphTeamFunSettings) GetAllowGiphyOk() (*bool, bool)`

GetAllowGiphyOk returns a tuple with the AllowGiphy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGiphy

`func (o *MicrosoftGraphTeamFunSettings) SetAllowGiphy(v bool)`

SetAllowGiphy sets AllowGiphy field to given value.

### HasAllowGiphy

`func (o *MicrosoftGraphTeamFunSettings) HasAllowGiphy() bool`

HasAllowGiphy returns a boolean if a field has been set.

### SetAllowGiphyNil

`func (o *MicrosoftGraphTeamFunSettings) SetAllowGiphyNil(b bool)`

 SetAllowGiphyNil sets the value for AllowGiphy to be an explicit nil

### UnsetAllowGiphy
`func (o *MicrosoftGraphTeamFunSettings) UnsetAllowGiphy()`

UnsetAllowGiphy ensures that no value is present for AllowGiphy, not even an explicit nil
### GetAllowStickersAndMemes

`func (o *MicrosoftGraphTeamFunSettings) GetAllowStickersAndMemes() bool`

GetAllowStickersAndMemes returns the AllowStickersAndMemes field if non-nil, zero value otherwise.

### GetAllowStickersAndMemesOk

`func (o *MicrosoftGraphTeamFunSettings) GetAllowStickersAndMemesOk() (*bool, bool)`

GetAllowStickersAndMemesOk returns a tuple with the AllowStickersAndMemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStickersAndMemes

`func (o *MicrosoftGraphTeamFunSettings) SetAllowStickersAndMemes(v bool)`

SetAllowStickersAndMemes sets AllowStickersAndMemes field to given value.

### HasAllowStickersAndMemes

`func (o *MicrosoftGraphTeamFunSettings) HasAllowStickersAndMemes() bool`

HasAllowStickersAndMemes returns a boolean if a field has been set.

### SetAllowStickersAndMemesNil

`func (o *MicrosoftGraphTeamFunSettings) SetAllowStickersAndMemesNil(b bool)`

 SetAllowStickersAndMemesNil sets the value for AllowStickersAndMemes to be an explicit nil

### UnsetAllowStickersAndMemes
`func (o *MicrosoftGraphTeamFunSettings) UnsetAllowStickersAndMemes()`

UnsetAllowStickersAndMemes ensures that no value is present for AllowStickersAndMemes, not even an explicit nil
### GetGiphyContentRating

`func (o *MicrosoftGraphTeamFunSettings) GetGiphyContentRating() AnyOfmicrosoftGraphGiphyRatingType`

GetGiphyContentRating returns the GiphyContentRating field if non-nil, zero value otherwise.

### GetGiphyContentRatingOk

`func (o *MicrosoftGraphTeamFunSettings) GetGiphyContentRatingOk() (*AnyOfmicrosoftGraphGiphyRatingType, bool)`

GetGiphyContentRatingOk returns a tuple with the GiphyContentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiphyContentRating

`func (o *MicrosoftGraphTeamFunSettings) SetGiphyContentRating(v AnyOfmicrosoftGraphGiphyRatingType)`

SetGiphyContentRating sets GiphyContentRating field to given value.

### HasGiphyContentRating

`func (o *MicrosoftGraphTeamFunSettings) HasGiphyContentRating() bool`

HasGiphyContentRating returns a boolean if a field has been set.

### SetGiphyContentRatingNil

`func (o *MicrosoftGraphTeamFunSettings) SetGiphyContentRatingNil(b bool)`

 SetGiphyContentRatingNil sets the value for GiphyContentRating to be an explicit nil

### UnsetGiphyContentRating
`func (o *MicrosoftGraphTeamFunSettings) UnsetGiphyContentRating()`

UnsetGiphyContentRating ensures that no value is present for GiphyContentRating, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


