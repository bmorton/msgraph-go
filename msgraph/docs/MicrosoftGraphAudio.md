# MicrosoftGraphAudio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to **NullableString** | The title of the album for this audio file. | [optional] 
**AlbumArtist** | Pointer to **NullableString** | The artist named on the album for the audio file. | [optional] 
**Artist** | Pointer to **NullableString** | The performing artist for the audio file. | [optional] 
**Bitrate** | Pointer to **NullableInt64** | Bitrate expressed in kbps. | [optional] 
**Composers** | Pointer to **NullableString** | The name of the composer of the audio file. | [optional] 
**Copyright** | Pointer to **NullableString** | Copyright information for the audio file. | [optional] 
**Disc** | Pointer to **NullableInt32** | The number of the disc this audio file came from. | [optional] 
**DiscCount** | Pointer to **NullableInt32** | The total number of discs in this album. | [optional] 
**Duration** | Pointer to **NullableInt64** | Duration of the audio file, expressed in milliseconds | [optional] 
**Genre** | Pointer to **NullableString** | The genre of this audio file. | [optional] 
**HasDrm** | Pointer to **NullableBool** | Indicates if the file is protected with digital rights management. | [optional] 
**IsVariableBitrate** | Pointer to **NullableBool** | Indicates if the file is encoded with a variable bitrate. | [optional] 
**Title** | Pointer to **NullableString** | The title of the audio file. | [optional] 
**Track** | Pointer to **NullableInt32** | The number of the track on the original disc for this audio file. | [optional] 
**TrackCount** | Pointer to **NullableInt32** | The total number of tracks on the original disc for this audio file. | [optional] 
**Year** | Pointer to **NullableInt32** | The year the audio file was recorded. | [optional] 

## Methods

### NewMicrosoftGraphAudio

`func NewMicrosoftGraphAudio() *MicrosoftGraphAudio`

NewMicrosoftGraphAudio instantiates a new MicrosoftGraphAudio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAudioWithDefaults

`func NewMicrosoftGraphAudioWithDefaults() *MicrosoftGraphAudio`

NewMicrosoftGraphAudioWithDefaults instantiates a new MicrosoftGraphAudio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *MicrosoftGraphAudio) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *MicrosoftGraphAudio) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *MicrosoftGraphAudio) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *MicrosoftGraphAudio) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *MicrosoftGraphAudio) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *MicrosoftGraphAudio) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetAlbumArtist

`func (o *MicrosoftGraphAudio) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *MicrosoftGraphAudio) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *MicrosoftGraphAudio) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *MicrosoftGraphAudio) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtistNil

`func (o *MicrosoftGraphAudio) SetAlbumArtistNil(b bool)`

 SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil

### UnsetAlbumArtist
`func (o *MicrosoftGraphAudio) UnsetAlbumArtist()`

UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
### GetArtist

`func (o *MicrosoftGraphAudio) GetArtist() string`

GetArtist returns the Artist field if non-nil, zero value otherwise.

### GetArtistOk

`func (o *MicrosoftGraphAudio) GetArtistOk() (*string, bool)`

GetArtistOk returns a tuple with the Artist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtist

`func (o *MicrosoftGraphAudio) SetArtist(v string)`

SetArtist sets Artist field to given value.

### HasArtist

`func (o *MicrosoftGraphAudio) HasArtist() bool`

HasArtist returns a boolean if a field has been set.

### SetArtistNil

`func (o *MicrosoftGraphAudio) SetArtistNil(b bool)`

 SetArtistNil sets the value for Artist to be an explicit nil

### UnsetArtist
`func (o *MicrosoftGraphAudio) UnsetArtist()`

UnsetArtist ensures that no value is present for Artist, not even an explicit nil
### GetBitrate

`func (o *MicrosoftGraphAudio) GetBitrate() int64`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MicrosoftGraphAudio) GetBitrateOk() (*int64, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *MicrosoftGraphAudio) SetBitrate(v int64)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *MicrosoftGraphAudio) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *MicrosoftGraphAudio) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *MicrosoftGraphAudio) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetComposers

`func (o *MicrosoftGraphAudio) GetComposers() string`

GetComposers returns the Composers field if non-nil, zero value otherwise.

### GetComposersOk

`func (o *MicrosoftGraphAudio) GetComposersOk() (*string, bool)`

GetComposersOk returns a tuple with the Composers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposers

`func (o *MicrosoftGraphAudio) SetComposers(v string)`

SetComposers sets Composers field to given value.

### HasComposers

`func (o *MicrosoftGraphAudio) HasComposers() bool`

HasComposers returns a boolean if a field has been set.

### SetComposersNil

`func (o *MicrosoftGraphAudio) SetComposersNil(b bool)`

 SetComposersNil sets the value for Composers to be an explicit nil

### UnsetComposers
`func (o *MicrosoftGraphAudio) UnsetComposers()`

UnsetComposers ensures that no value is present for Composers, not even an explicit nil
### GetCopyright

`func (o *MicrosoftGraphAudio) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *MicrosoftGraphAudio) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *MicrosoftGraphAudio) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *MicrosoftGraphAudio) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### SetCopyrightNil

`func (o *MicrosoftGraphAudio) SetCopyrightNil(b bool)`

 SetCopyrightNil sets the value for Copyright to be an explicit nil

### UnsetCopyright
`func (o *MicrosoftGraphAudio) UnsetCopyright()`

UnsetCopyright ensures that no value is present for Copyright, not even an explicit nil
### GetDisc

`func (o *MicrosoftGraphAudio) GetDisc() int32`

GetDisc returns the Disc field if non-nil, zero value otherwise.

### GetDiscOk

`func (o *MicrosoftGraphAudio) GetDiscOk() (*int32, bool)`

GetDiscOk returns a tuple with the Disc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisc

`func (o *MicrosoftGraphAudio) SetDisc(v int32)`

SetDisc sets Disc field to given value.

### HasDisc

`func (o *MicrosoftGraphAudio) HasDisc() bool`

HasDisc returns a boolean if a field has been set.

### SetDiscNil

`func (o *MicrosoftGraphAudio) SetDiscNil(b bool)`

 SetDiscNil sets the value for Disc to be an explicit nil

### UnsetDisc
`func (o *MicrosoftGraphAudio) UnsetDisc()`

UnsetDisc ensures that no value is present for Disc, not even an explicit nil
### GetDiscCount

`func (o *MicrosoftGraphAudio) GetDiscCount() int32`

GetDiscCount returns the DiscCount field if non-nil, zero value otherwise.

### GetDiscCountOk

`func (o *MicrosoftGraphAudio) GetDiscCountOk() (*int32, bool)`

GetDiscCountOk returns a tuple with the DiscCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscCount

`func (o *MicrosoftGraphAudio) SetDiscCount(v int32)`

SetDiscCount sets DiscCount field to given value.

### HasDiscCount

`func (o *MicrosoftGraphAudio) HasDiscCount() bool`

HasDiscCount returns a boolean if a field has been set.

### SetDiscCountNil

`func (o *MicrosoftGraphAudio) SetDiscCountNil(b bool)`

 SetDiscCountNil sets the value for DiscCount to be an explicit nil

### UnsetDiscCount
`func (o *MicrosoftGraphAudio) UnsetDiscCount()`

UnsetDiscCount ensures that no value is present for DiscCount, not even an explicit nil
### GetDuration

`func (o *MicrosoftGraphAudio) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MicrosoftGraphAudio) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MicrosoftGraphAudio) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MicrosoftGraphAudio) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *MicrosoftGraphAudio) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *MicrosoftGraphAudio) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetGenre

`func (o *MicrosoftGraphAudio) GetGenre() string`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *MicrosoftGraphAudio) GetGenreOk() (*string, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *MicrosoftGraphAudio) SetGenre(v string)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *MicrosoftGraphAudio) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### SetGenreNil

`func (o *MicrosoftGraphAudio) SetGenreNil(b bool)`

 SetGenreNil sets the value for Genre to be an explicit nil

### UnsetGenre
`func (o *MicrosoftGraphAudio) UnsetGenre()`

UnsetGenre ensures that no value is present for Genre, not even an explicit nil
### GetHasDrm

`func (o *MicrosoftGraphAudio) GetHasDrm() bool`

GetHasDrm returns the HasDrm field if non-nil, zero value otherwise.

### GetHasDrmOk

`func (o *MicrosoftGraphAudio) GetHasDrmOk() (*bool, bool)`

GetHasDrmOk returns a tuple with the HasDrm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDrm

`func (o *MicrosoftGraphAudio) SetHasDrm(v bool)`

SetHasDrm sets HasDrm field to given value.

### HasHasDrm

`func (o *MicrosoftGraphAudio) HasHasDrm() bool`

HasHasDrm returns a boolean if a field has been set.

### SetHasDrmNil

`func (o *MicrosoftGraphAudio) SetHasDrmNil(b bool)`

 SetHasDrmNil sets the value for HasDrm to be an explicit nil

### UnsetHasDrm
`func (o *MicrosoftGraphAudio) UnsetHasDrm()`

UnsetHasDrm ensures that no value is present for HasDrm, not even an explicit nil
### GetIsVariableBitrate

`func (o *MicrosoftGraphAudio) GetIsVariableBitrate() bool`

GetIsVariableBitrate returns the IsVariableBitrate field if non-nil, zero value otherwise.

### GetIsVariableBitrateOk

`func (o *MicrosoftGraphAudio) GetIsVariableBitrateOk() (*bool, bool)`

GetIsVariableBitrateOk returns a tuple with the IsVariableBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariableBitrate

`func (o *MicrosoftGraphAudio) SetIsVariableBitrate(v bool)`

SetIsVariableBitrate sets IsVariableBitrate field to given value.

### HasIsVariableBitrate

`func (o *MicrosoftGraphAudio) HasIsVariableBitrate() bool`

HasIsVariableBitrate returns a boolean if a field has been set.

### SetIsVariableBitrateNil

`func (o *MicrosoftGraphAudio) SetIsVariableBitrateNil(b bool)`

 SetIsVariableBitrateNil sets the value for IsVariableBitrate to be an explicit nil

### UnsetIsVariableBitrate
`func (o *MicrosoftGraphAudio) UnsetIsVariableBitrate()`

UnsetIsVariableBitrate ensures that no value is present for IsVariableBitrate, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphAudio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphAudio) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphAudio) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphAudio) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphAudio) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphAudio) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTrack

`func (o *MicrosoftGraphAudio) GetTrack() int32`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *MicrosoftGraphAudio) GetTrackOk() (*int32, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *MicrosoftGraphAudio) SetTrack(v int32)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *MicrosoftGraphAudio) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### SetTrackNil

`func (o *MicrosoftGraphAudio) SetTrackNil(b bool)`

 SetTrackNil sets the value for Track to be an explicit nil

### UnsetTrack
`func (o *MicrosoftGraphAudio) UnsetTrack()`

UnsetTrack ensures that no value is present for Track, not even an explicit nil
### GetTrackCount

`func (o *MicrosoftGraphAudio) GetTrackCount() int32`

GetTrackCount returns the TrackCount field if non-nil, zero value otherwise.

### GetTrackCountOk

`func (o *MicrosoftGraphAudio) GetTrackCountOk() (*int32, bool)`

GetTrackCountOk returns a tuple with the TrackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackCount

`func (o *MicrosoftGraphAudio) SetTrackCount(v int32)`

SetTrackCount sets TrackCount field to given value.

### HasTrackCount

`func (o *MicrosoftGraphAudio) HasTrackCount() bool`

HasTrackCount returns a boolean if a field has been set.

### SetTrackCountNil

`func (o *MicrosoftGraphAudio) SetTrackCountNil(b bool)`

 SetTrackCountNil sets the value for TrackCount to be an explicit nil

### UnsetTrackCount
`func (o *MicrosoftGraphAudio) UnsetTrackCount()`

UnsetTrackCount ensures that no value is present for TrackCount, not even an explicit nil
### GetYear

`func (o *MicrosoftGraphAudio) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MicrosoftGraphAudio) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MicrosoftGraphAudio) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MicrosoftGraphAudio) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *MicrosoftGraphAudio) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MicrosoftGraphAudio) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


