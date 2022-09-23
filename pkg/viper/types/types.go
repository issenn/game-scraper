package types


import (
	"github.com/spf13/viper"
)


type Unmarshaler interface {
	UnmarshalViper(*viper.Viper, interface{}) error
}
