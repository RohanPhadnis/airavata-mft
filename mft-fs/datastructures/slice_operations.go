package datastructures

import (
	"github.com/jacobsa/fuse/fuseops"
)

func Remove(slice []fuseops.InodeID, i int) []fuseops.InodeID {
	if i+1 < len(slice) {
		return append(slice[:i], slice[i+1:]...)
	}
	return slice[:i]
}
