package index

import (
	"testing"

	"github.com/wutong-paas/restic/pkg/restic"
	"github.com/wutong-paas/restic/pkg/test"
)

func TestMergeIndex(t testing.TB, mi *MasterIndex) ([]*Index, int) {
	finalIndexes := mi.finalizeNotFinalIndexes()
	for _, idx := range finalIndexes {
		test.OK(t, idx.SetID(restic.NewRandomID()))
	}

	test.OK(t, mi.MergeFinalIndexes())
	return finalIndexes, len(mi.idx)
}
