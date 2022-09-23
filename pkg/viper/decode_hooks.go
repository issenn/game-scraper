package viper


import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	viperTypes "github.com/issenn/game-scraper/pkg/viper/types"
)


func ViperUnmarshalerHookFunc(v *viper.Viper) mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		if to.CanAddr() {
			to = to.Addr()
		}
		unmarshaler, ok := to.Interface().(viperTypes.Unmarshaler)
		if !ok {
			return from.Interface(), nil
		}
		if to.IsNil() && to.Type().Kind() == reflect.Ptr {
			to.Set(reflect.New(to.Type().Elem()))
			unmarshaler = to.Interface().(viperTypes.Unmarshaler)
		}
		if err := unmarshaler.UnmarshalViper(v, from.Interface()); err != nil {
			return nil, err
		}
		return from.Interface(), nil
	}
}
