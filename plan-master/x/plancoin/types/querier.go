package types

import (
	posmining "github.com/plan-crypto/node/x/posmining/types"
	structure "github.com/plan-crypto/node/x/structure/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QueryProfile = "profile"
)

// Profile response
type ProfileResolve struct {
	Owner     sdk.AccAddress             `json:"owner"`
	Balance   sdk.Int                    `json:"balance"`
	Posmining posmining.PosminingResolve `json:"posmining"`
	Structure structure.Structure        `json:"structure"`
}

func (r ProfileResolve) String() string {
	return r.Balance.String()
}

// EncodeResp defines a tx encoding response.
type EncodeResp struct {
	Tx string `json:"tx" yaml:"tx"`
}

type DecodeReq struct {
	Tx string `json:"tx" yaml:"tx"`
}
