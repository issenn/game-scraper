package types


import (
	"reflect"
)


// Any type implementing this interface can eventually decode a
// map[string]interface{} into the type or just run some logic before the
// default decoder is applied, this is controller with the returned boolean.
// If it returns true the mapstructure continues with the default decoder
// logic after running the custom logic. If it returns false it assumes the
// decoding is completely done after this custom logic.
type CustomUnmarshaler interface {
	UnmarshalCustom(interface{}) (bool, error)
}

type Unmarshaler interface {
	// Todo
	UnmarshalMapStructure(from reflect.Type, data interface{}) error
}

// type UnmarshalerHookFunc func()
