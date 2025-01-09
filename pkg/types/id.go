package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	uuid "github.com/gofrs/uuid"
)

// ID is a interface for uuid's
type ID interface {
	Bytes() []byte
	String() string
	// Format implements fmt.Formatter for ID values.
	Format(f fmt.State, c rune)
	UnmarshalJSON(data []byte) error
	IsNil() bool
	LessThen(other ID) bool
	Set(bytes []byte) error
	As(a ID)
}

//IDProvider can be used to inject id's for testing
type IDProvider interface {
	FromString(a ID, s string) error
	FromStringOrNil(a ID, s string)
	NewRandom(a ID) error
}

//DefaultIDProvider is a provider of ID's using the ID type
var DefaultIDProvider IDProvider = defaultIDProvider{}

type defaultIDProvider struct {
}

// FromStringOrNil creates a ID from string, if the string is not an ID it will returns a Nil one
func (p defaultIDProvider) FromStringOrNil(a ID, s string) {
	id := uuid.FromStringOrNil(s)
	a.Set(id.Bytes())
}

// FromString creates a new ID from a string
func (p defaultIDProvider) FromString(a ID, s string) error {
	parsed, err := uuid.FromString(s)
	if err == nil {
		a.Set(parsed.Bytes())
	}
	return err
}

// NewRandom creates a new RandomV4 id and assigns it to the type passed in
func (p defaultIDProvider) NewRandom(a ID) error {
	u, err := uuid.NewV4()
	if err == nil {
		a.Set(u.Bytes())
	}
	return err
}

var _ ID = (*GenID)(nil)

// GenID is a generic id it provides all the functionality for derived ids
type GenID struct {
	uuid.UUID
}

// IDType is the runtime type of UUID
var IDType = reflect.TypeOf(NilID)

// NilID is special form of UUID that is specified to have all
// 128 bits set to zero.
var NilID = GenID{}

// UnmarshalJSON sets *u to a copy of data
// we create this function to allow to change the type of errors thrown by the original unmarshaller.
func (i *GenID) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &i.UUID)
	if err != nil {
		return &json.UnmarshalTypeError{Type: IDType, Value: string(data)}
	}
	return nil
}

// IsNil check if id is the NilID
func (i GenID) IsNil() bool {
	return i == NilID
}

// LessThan compare this instance to another.
func (i GenID) LessThen(other ID) bool {
	return bytes.Compare(i.Bytes(), other.Bytes()) < 0
}

// Set assigns the value of the bytes to an ID
func (i *GenID) Set(bytes []byte) error {
	u, err := uuid.FromBytes(bytes)
	if err == nil {
		i.UUID = u
	}
	return err
}

//As assigns equivalent valies to a different ID type
func (i *GenID) As(a ID) {
	a.Set(i.Bytes())
}

// ProvideIDProvider provides a idprovider
func ProvideIDProvider() *defaultIDProvider {
	return &defaultIDProvider{}
}

// Must is a wrapper which assures that a assignation did not return an error
// it will panic if they do.
// This is meant to assure that hardcoded ID values are valid
// usage Must(&RoutingIDGlobal, "cb98a312-26b5-4d90-96f3-af0f40f63291")
func Must(a ID, s string) {
	err := DefaultIDProvider.FromString(a, s)
	if err != nil {
		panic(err)
	}
}
