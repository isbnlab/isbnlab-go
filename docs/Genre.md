# Genre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BisacCode** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EditionCount** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewGenre

`func NewGenre() *Genre`

NewGenre instantiates a new Genre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenreWithDefaults

`func NewGenreWithDefaults() *Genre`

NewGenreWithDefaults instantiates a new Genre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBisacCode

`func (o *Genre) GetBisacCode() string`

GetBisacCode returns the BisacCode field if non-nil, zero value otherwise.

### GetBisacCodeOk

`func (o *Genre) GetBisacCodeOk() (*string, bool)`

GetBisacCodeOk returns a tuple with the BisacCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBisacCode

`func (o *Genre) SetBisacCode(v string)`

SetBisacCode sets BisacCode field to given value.

### HasBisacCode

`func (o *Genre) HasBisacCode() bool`

HasBisacCode returns a boolean if a field has been set.

### SetBisacCodeNil

`func (o *Genre) SetBisacCodeNil(b bool)`

 SetBisacCodeNil sets the value for BisacCode to be an explicit nil

### UnsetBisacCode
`func (o *Genre) UnsetBisacCode()`

UnsetBisacCode ensures that no value is present for BisacCode, not even an explicit nil
### GetName

`func (o *Genre) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Genre) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Genre) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Genre) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Genre) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Genre) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEditionCount

`func (o *Genre) GetEditionCount() int64`

GetEditionCount returns the EditionCount field if non-nil, zero value otherwise.

### GetEditionCountOk

`func (o *Genre) GetEditionCountOk() (*int64, bool)`

GetEditionCountOk returns a tuple with the EditionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditionCount

`func (o *Genre) SetEditionCount(v int64)`

SetEditionCount sets EditionCount field to given value.

### HasEditionCount

`func (o *Genre) HasEditionCount() bool`

HasEditionCount returns a boolean if a field has been set.

### SetEditionCountNil

`func (o *Genre) SetEditionCountNil(b bool)`

 SetEditionCountNil sets the value for EditionCount to be an explicit nil

### UnsetEditionCount
`func (o *Genre) UnsetEditionCount()`

UnsetEditionCount ensures that no value is present for EditionCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


