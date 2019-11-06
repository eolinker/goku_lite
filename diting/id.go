package diting

import "github.com/eolinker/goku-api-gateway/diting/internal"

var (
	idCreate = internal.NewIDCreate()
)

//GetID getID
func GetID() uint64 {
	return idCreate.Next()
}
