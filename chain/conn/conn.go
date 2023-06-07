package conn

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type Conn struct {
	Rpc            string
	Api            *gsrpc.SubstrateAPI
	Metadata       *types.Metadata
	GenesisHash    types.Hash
	RuntimeVersion *types.RuntimeVersion
}

func NewConn(url string) (*Conn, error) {
	api, err := gsrpc.NewSubstrateAPI(url)
	if err != nil {
		return nil, err
	}

	metadata, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}

	runtimeVersion, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		Rpc:            url,
		Api:            api,
		Metadata:       metadata,
		GenesisHash:    genesisHash,
		RuntimeVersion: runtimeVersion,
	}

	return conn, nil
}
