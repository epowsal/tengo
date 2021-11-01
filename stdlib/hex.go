package stdlib

import (
	"encoding/hex"

	"github.com/epowsal/tengo"
)

var hexModule = map[string]tengo.Object{
	"encode": &tengo.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &tengo.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
