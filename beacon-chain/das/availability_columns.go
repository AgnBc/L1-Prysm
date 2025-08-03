package das

import (
	"context"

	"github.com/OffchainLabs/prysm/v6/consensus-types/blocks"
)

// DataColumnsVerifier enables LazilyPersistentStoreColumn to manage the verification process
// going from RODataColumn->VerifiedRODataColumn, while avoiding the decision of which individual verifications
// to run and in what order. Since LazilyPersistentStoreColumn always tries to verify and save data columns only when
// they are all available, the interface takes a slice of data column sidecars.
type DataColumnsVerifier interface {
	VerifiedRODataColumns(ctx context.Context, blk blocks.ROBlock, scs []blocks.RODataColumn) ([]blocks.VerifiedRODataColumn, error)
}
