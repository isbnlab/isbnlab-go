# Author

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**AlternativeNames** | Pointer to **[]string** |  | [optional] 
**DeathDate** | Pointer to **NullableTime** |  | [optional] 
**BirthDate** | Pointer to **NullableTime** |  | [optional] 
**Bio** | Pointer to **NullableString** |  | [optional] 
**Photos** | Pointer to [**[]AuthorPhoto**](AuthorPhoto.md) |  | [optional] 

## Methods

### NewAuthor

`func NewAuthor() *Author`

NewAuthor instantiates a new Author object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorWithDefaults

`func NewAuthorWithDefaults() *Author`

NewAuthorWithDefaults instantiates a new Author object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Author) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Author) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Author) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Author) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Author) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Author) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlug

`func (o *Author) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Author) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Author) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Author) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *Author) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *Author) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetAlternativeNames

`func (o *Author) GetAlternativeNames() []string`

GetAlternativeNames returns the AlternativeNames field if non-nil, zero value otherwise.

### GetAlternativeNamesOk

`func (o *Author) GetAlternativeNamesOk() (*[]string, bool)`

GetAlternativeNamesOk returns a tuple with the AlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeNames

`func (o *Author) SetAlternativeNames(v []string)`

SetAlternativeNames sets AlternativeNames field to given value.

### HasAlternativeNames

`func (o *Author) HasAlternativeNames() bool`

HasAlternativeNames returns a boolean if a field has been set.

### SetAlternativeNamesNil

`func (o *Author) SetAlternativeNamesNil(b bool)`

 SetAlternativeNamesNil sets the value for AlternativeNames to be an explicit nil

### UnsetAlternativeNames
`func (o *Author) UnsetAlternativeNames()`

UnsetAlternativeNames ensures that no value is present for AlternativeNames, not even an explicit nil
### GetDeathDate

`func (o *Author) GetDeathDate() time.Time`

GetDeathDate returns the DeathDate field if non-nil, zero value otherwise.

### GetDeathDateOk

`func (o *Author) GetDeathDateOk() (*time.Time, bool)`

GetDeathDateOk returns a tuple with the DeathDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathDate

`func (o *Author) SetDeathDate(v time.Time)`

SetDeathDate sets DeathDate field to given value.

### HasDeathDate

`func (o *Author) HasDeathDate() bool`

HasDeathDate returns a boolean if a field has been set.

### SetDeathDateNil

`func (o *Author) SetDeathDateNil(b bool)`

 SetDeathDateNil sets the value for DeathDate to be an explicit nil

### UnsetDeathDate
`func (o *Author) UnsetDeathDate()`

UnsetDeathDate ensures that no value is present for DeathDate, not even an explicit nil
### GetBirthDate

`func (o *Author) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *Author) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *Author) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *Author) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *Author) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *Author) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBio

`func (o *Author) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *Author) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *Author) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *Author) HasBio() bool`

HasBio returns a boolean if a field has been set.

### SetBioNil

`func (o *Author) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *Author) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetPhotos

`func (o *Author) GetPhotos() []AuthorPhoto`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *Author) GetPhotosOk() (*[]AuthorPhoto, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *Author) SetPhotos(v []AuthorPhoto)`

SetPhotos sets Photos field to given value.

### HasPhotos

`func (o *Author) HasPhotos() bool`

HasPhotos returns a boolean if a field has been set.

### SetPhotosNil

`func (o *Author) SetPhotosNil(b bool)`

 SetPhotosNil sets the value for Photos to be an explicit nil

### UnsetPhotos
`func (o *Author) UnsetPhotos()`

UnsetPhotos ensures that no value is present for Photos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


