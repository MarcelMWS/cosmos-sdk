package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "repo.mwaysolutions.com/blockscape/gaia-yubihsm/types"
)

func TestValidateGenesis(t *testing.T) {

	fp := InitialFeePool()
	require.Nil(t, fp.ValidateGenesis())

	fp2 := FeePool{CommunityPool: sdk.DecCoins{{"stake", sdk.NewDec(-1)}}}
	require.NotNil(t, fp2.ValidateGenesis())

}
