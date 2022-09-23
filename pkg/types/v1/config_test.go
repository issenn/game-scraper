package v1_test


import (
	// "fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
)


func TestMapstructure_DecodeRecursiveStructToMapHookFunc(t *testing.T) {
	var f mapstructure.DecodeHookFunc = RecursiveStructToMapHookFunc()

	type c struct {
		TestFoo string
	}

	type b struct {
		Sub c
		TestKey string
	}

	type a struct {
		Sub b
		N string
	}

	testStruct := a{
		Sub: b{
			Sub: c{
				TestFoo: "TestBarVal",
			},
			TestKey: "testval",
		},
		N: "issenn",
	}

	testMap := map[string]interface{}{
		"Sub": map[string]interface{}{
			"Sub": map[string]interface{}{
				"TestFoo": "TestBarVal",
			},
			"TestKey": "testval",
		},
		"N": "issenn",
	}

	cases := []struct {
		name     string
		receiver interface{}
		input    interface{}
		expected interface{}
		err      bool
	}{
		{
			"map receiver",
			func() interface{} {
				var res map[string]interface{}
				return &res
			}(),
			testStruct,
			&testMap,
			false,
		},
		{
			"interface receiver",
			func() interface{} {
				var res interface{}
				return &res
			}(),
			testStruct,
			func() interface{} {
				var exp interface{} = testMap
				return &exp
			}(),
			false,
		},
		{
			"slice receiver errors",
			func() interface{} {
				var res []string
				return &res
			}(),
			testStruct,
			new([]string),
			true,
		},
		{
			"slice to slice - no change",
			func() interface{} {
				var res []string
				return &res
			}(),
			[]string{"a", "b"},
			&[]string{"a", "b"},
			false,
		},
		{
			"string to string - no change",
			func() interface{} {
				var res string
				return &res
			}(),
			"test",
			func() *string {
				s := "test"
				return &s
			}(),
			false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &mapstructure.DecoderConfig{
				DecodeHook: f,
				Result:     tc.receiver,
			}

			d, err := mapstructure.NewDecoder(cfg)
			if err != nil {
				t.Fatalf("unexpected err %#v", err)
			}

			err = d.Decode(tc.input)
			if tc.err != (err != nil) {
				t.Fatalf("expected err %#v", err)
			}

			// fmt.Printf("%v", tc.receiver)
			if !reflect.DeepEqual(tc.expected, tc.receiver) {
				t.Fatalf("expected %#v, got %#v",
					tc.expected, tc.receiver)
			}
		})

	}
}

func RecursiveStructToMapHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Value, t reflect.Value) (interface{}, error) {
		// fmt.Printf("================\n\n")
		// outputKind := getKind(t)
		// fmt.Printf("outputKind: %v \n\n", outputKind)
		// fmt.Printf("Value interface: %#v => %#v \n\n", f.Interface(), t.Interface())
		// fmt.Printf("Kind: %v => %v \n\n", f.Kind(), t.Kind())
		if f.Kind() != reflect.Struct {
			// fmt.Printf("Return 0 interface f: %#v \n\n", f.Interface())
			return f.Interface(), nil
		}

		// fmt.Printf("Type: %v => %v \n\n", f.Type(), t.Type())
		var i interface{} = struct{}{}
		// fmt.Printf("t type != i interface type ?: %v != %v ? %v \n\n", t.Type(), reflect.TypeOf(&i).Elem(), t.Type() != reflect.TypeOf(&i).Elem())
		if t.Type() != reflect.TypeOf(&i).Elem() {
			// fmt.Printf("Return 1 interface f: %#v \n\n", f.Interface())
			return f.Interface(), nil
		}

		// fmt.Printf("if t type == i interface type\n\n")
		m := make(map[string]interface{})
		// fmt.Printf("m type: %v \n\n", reflect.ValueOf(m).Type())
		t.Set(reflect.ValueOf(m))
		// fmt.Printf("t type: %v \n\n", t.Type())
		// fmt.Printf("t Kind: %v \n\n", t.Kind())
		// fmt.Printf("t interface: %#v \n\n", t.Interface())

		// fmt.Printf("Return 2 interface f: %#v \n\n", f.Interface())
		return f.Interface(), nil
	}
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
