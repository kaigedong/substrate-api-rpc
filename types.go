package substrate

import (
	"github.com/kaigedong/scale.go/source"
	"github.com/kaigedong/scale.go/types"
)

func RegCustomTypes(sourceCode []byte) {
	types.RegCustomTypes(source.LoadTypeRegistry(sourceCode))
}
