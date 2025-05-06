package util

import "github.com/tiendc/go-deepcopy"

func Copy(dest any, src any) error {
	return deepcopy.Copy(dest, src, deepcopy.IgnoreNonCopyableTypes(true))
}
