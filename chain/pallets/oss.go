package pallets

import (
	"MatrixAI-Client/chain"
	"MatrixAI-Client/chain/pattern"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

type OssWrapper struct {
	*chain.InfoChain
}

func (chain *OssWrapper) Authorize(puk []byte) (string, error) {

	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	address, err := types.NewAccountID(puk)
	if err != nil {
		return txhash, err
	}

	fmt.Printf("------------------ 构建交易 ------------------\n")

	call, err := types.NewCall(chain.Conn.Metadata, pattern.TX_OSS_REGISTER, *address)
	if err != nil {
		return txhash, err
	}

	key, err := types.CreateStorageKey(chain.Conn.Metadata, pattern.SYSTEM, pattern.ACCOUNT, chain.Wallet.KeyringPair.PublicKey)
	if err != nil {
		return txhash, err
	}

	_, err = chain.Conn.Api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, err
	}

	options := types.SignatureOptions{
		BlockHash:          chain.Conn.GenesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        chain.Conn.GenesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        chain.Conn.RuntimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: chain.Conn.RuntimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)
	err = ext.Sign(chain.Wallet.KeyringPair, options)
	if err != nil {
		return txhash, err
	}

	hash, err := chain.Conn.Api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		return txhash, err
	}
	txhash, _ = codec.EncodeToHex(hash)

	fmt.Printf("------------------ 提交交易 ------------------\n%+v\n", txhash)
	return txhash, nil
}

func NewOssWrapper(info *chain.InfoChain) *OssWrapper {
	return &OssWrapper{info}
}
