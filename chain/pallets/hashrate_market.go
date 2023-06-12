package pallets

import (
	"MatrixAI-Client/chain"
	"MatrixAI-Client/chain/events"
	"MatrixAI-Client/chain/pattern"
	"MatrixAI-Client/hardwareinfo"
	"MatrixAI-Client/logs"
	"MatrixAI-Client/utils"
	"encoding/json"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
	"time"
)

type WrapperMatrix struct {
	*chain.InfoChain
}

func (chain WrapperMatrix) AddMachine(hardwareInfo hardwareinfo.HardwareInfo) (string, error) {

	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	uuid, err := utils.ParseUUID(string(hardwareInfo.MachineUUID))
	//uuid, err := utils.ParseUUID("E39911FB-03C7-A00A-B29E-50EBF6B66205")
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error parsing uuid: %v", err))
	}

	var machineUUID pattern.MachineUUID
	for i := 0; i < len(machineUUID); i++ {
		machineUUID[i] = types.U8(uuid[i])
	}

	jsonData, err := json.Marshal(hardwareInfo)
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error marshaling the struct to JSON: %v", err))
	}

	key, err := types.CreateStorageKey(chain.Conn.Metadata, pattern.SYSTEM, pattern.ACCOUNT, chain.Wallet.KeyringPair.PublicKey)
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error creating storage key: %v", err))
	}

	_, err = chain.Conn.Api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error getting storage latest: %v", err))
	}

	call, err := types.NewCall(chain.Conn.Metadata, pattern.TX_HASHRATE_MARKET_REGISTER, machineUUID, types.NewBytes(jsonData))
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error creating call: %v", err))
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
		return txhash, errors.New(fmt.Sprintf("Error signing the extrinsic: %v", err))
	}

	sub, err := chain.Conn.Api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.New(fmt.Sprintf("Error submitting extrinsic: %v", err))
	}

	defer sub.Unsubscribe()

	timeout := time.NewTimer(time.Second * time.Duration(12))
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				logs.Result(fmt.Sprintf("------------------ 交易完成 ------------------ : %#x", status.AsInBlock))

				tEvents := events.EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := chain.Conn.Api.RPC.State.GetStorageRaw(chain.Conn.KeyEvents, status.AsInBlock)
				if err != nil {
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}

				err = types.EventRecordsRaw(*h).DecodeEventRecords(chain.Conn.Metadata, &tEvents)
				if err != nil {
					return txhash, errors.New(fmt.Sprintf("DecodeEventRecords error : %v", err))
				}

				for _, e := range tEvents.HashrateMarket_MachineAdded {
					if codec.Eq(e.Id, machineUUID) {
						logs.Result("add machine bingo!!!")
						return txhash, nil
					}
				}
				return txhash, errors.New(pattern.ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[WatchExtrinsic]")
		case <-timeout.C:
			return txhash, errors.New(pattern.ERR_Timeout)
		}
	}
}

func NewMatrixWrapper(info *chain.InfoChain) *WrapperMatrix {
	return &WrapperMatrix{info}
}
