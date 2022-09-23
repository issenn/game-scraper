package mapstructure


import (
	"fmt"
	"encoding"
	"reflect"

	"github.com/mitchellh/mapstructure"

	mapstructureTypes "github.com/issenn/game-scraper/pkg/mapstructure/types"
)


// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshalerHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		result := reflect.New(t).Interface()
		unmarshaler, ok := result.(encoding.TextUnmarshaler)
		if !ok {
			return data, nil
		}
		if err := unmarshaler.UnmarshalText([]byte(data.(string))); err != nil {
			return nil, err
		}
		return result, nil
	}
}

func CustomUnmarshalerHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		// get the destination object address if it is not passed by reference
		if to.CanAddr() {
			to = to.Addr()
		}
		// If the destination implements the unmarshaling interface
		unmarshaler, ok := to.Interface().(mapstructureTypes.CustomUnmarshaler)
		if !ok {
			return from.Interface(), nil
		}
		// If it is nil and a pointer, create and assign the target value first
		if to.IsNil() && to.Type().Kind() == reflect.Ptr {
			to.Set(reflect.New(to.Type().Elem()))
			unmarshaler = to.Interface().(mapstructureTypes.CustomUnmarshaler)
		}
		// Call the custom unmarshaling method
		cont, err := unmarshaler.UnmarshalCustom(from.Interface())
		if cont {
			// Continue with the decoding stack
			return from.Interface(), err
		}
		// Decoding finalized
		return to.Interface(), err
	}
}

func UnmarshalerHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		// Todo
		if to.CanAddr() {
			to = to.Addr()
		}
		unmarshaler, ok := to.Interface().(mapstructureTypes.Unmarshaler)
		if !ok {
			return from.Interface(), nil
		}
		if to.IsNil() && to.Type().Kind() == reflect.Ptr {
			to.Set(reflect.New(to.Type().Elem()))
			unmarshaler = to.Interface().(mapstructureTypes.Unmarshaler)
		}
		if err := unmarshaler.UnmarshalMapStructure(from.Type(), from.Interface()); err != nil {
			return nil, err
		}
		return to.Interface(), nil
	}
}

func TestUnmarshalerHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		testPrint(from, to)
		return from.Interface(), nil
	}
}

func testPrint(from reflect.Value, to reflect.Value) {
	fmt.Printf("================\n\n")
	outputKind := getKind(to)
	fmt.Printf("outputKind: %v \n\n", outputKind)
	fmt.Printf("Value interface:\n  %#v\n =>\n  %#v \n\n", from.Interface(), to.Interface())
	fmt.Printf("Kind:\n  %v\n  =>\n  %v \n\n", from.Kind(), to.Kind())
	fmt.Printf("Type:\n  %v\n  =>\n  %v \n\n", from.Type(), to.Type())
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}
