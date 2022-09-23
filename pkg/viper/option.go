package viper


import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)


func WithOption(name string, data interface{}) viper.DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		var dataVal reflect.Value
		isNil := data == nil
		if !isNil {
			dataVal = reflect.ValueOf(data)
			if dataVal.Kind() == reflect.Ptr && dataVal.IsNil() {
				isNil = true
			} else if dataVal.IsValid() {
				switch v := reflect.Indirect(dataVal); v.Kind() {
				case reflect.Chan,
					reflect.Func,
					reflect.Interface,
					reflect.Map,
					reflect.Ptr,
					reflect.Slice:
						isNil = v.IsValid() && v.IsNil()
				}
			}
		}

		var val reflect.Value = reflect.Indirect(reflect.ValueOf(c))
		nameVal := val.FieldByName(name)
		if !nameVal.IsValid() {
			return
		}

		if isNil {
			if !nameVal.IsZero() && nameVal.CanSet() {
				nilValue := reflect.Indirect(reflect.New(nameVal.Type()))
				nameVal.Set(nilValue)
			}
			return
		}

		if !dataVal.IsValid() {
			return
		}

		if dataVal.Kind() == reflect.Ptr && dataVal.Type().Elem() == nameVal.Type() {
			dataVal = reflect.Indirect(dataVal)
		}

		if dataVal.Kind() != reflect.Func && dataVal.Type() != nameVal.Type() {
			return
		}

		if nameVal.CanSet() {
			nameVal.Set(dataVal)
		}
	}
}

func Squash(squash bool) viper.DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		c.Squash = squash
	}
}

func Metadata(metadata *mapstructure.Metadata) viper.DecoderConfigOption {
	return func(c *mapstructure.DecoderConfig) {
		c.Metadata = metadata
	}
}
