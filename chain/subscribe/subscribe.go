package subscribe

import (
	"MatrixAI-Client/chain"
	"MatrixAI-Client/chain/events"
	"MatrixAI-Client/logs"
	"MatrixAI-Client/utils"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

type WrapperSubscribe struct {
	*chain.InfoChain
}

func (chain *WrapperSubscribe) SubscribeEvents() {

	//// ---------- 协程处理区块 ----------
	//newHeadsChan := make(chan types.Hash, 10)
	//quitChan := make(chan struct{})
	//
	//analysisEvent := events.NewAnalysisWrapper(chain.InfoChain)
	//go analysisEvent.AnalysisEvents(newHeadsChan, quitChan)
	//// ---------- 协程处理区块 ----------

	//// ---------- 订阅区块 ----------
	//sub, err := chain.Conn.Api.RPC.Chain.SubscribeNewHeads()
	//if err != nil {
	//	_ = fmt.Errorf("subscribe error: %v", err)
	//	return
	//}
	//defer sub.Unsubscribe()
	//// ---------- 订阅区块 ----------

	//count := 0
	//
	//for {
	//	head := <-sub.Chan()
	//	fmt.Printf("Chain is at block: #%v\n", head.Number)
	//
	//	parentHash, _ := codec.EncodeToHex(head.ParentHash)
	//	fmt.Printf("Chain is at block ParentHash: #%v\n", parentHash)
	//	count++
	//
	//	if count == 5 {
	//		fmt.Printf("Subscribe done")
	//		return
	//	}
	//}

	//	// ---------- 协程处理区块 ----------
	//Loop:
	//	for {
	//		select {
	//		case header := <-sub.Chan():
	//			fmt.Println("Received new block:", header.Number)
	//			parentHash, _ := codec.EncodeToHex(header.ParentHash)
	//			fmt.Println("ParentHash:", parentHash)
	//			newHeadsChan <- header.ParentHash
	//		case <-sub.Err():
	//			_ = fmt.Errorf("subscribe error: %v", sub.Err())
	//			break Loop
	//		case <-quitChan:
	//			fmt.Println("Unsubscribing from new block")
	//			sub.Unsubscribe()
	//			break Loop
	//		}
	//	}
	//	// ---------- 协程处理区块 ----------

	sub, err := chain.Conn.Api.RPC.State.SubscribeStorageRaw([]types.StorageKey{chain.Conn.KeyEvents})
	if err != nil {
		logs.Error(fmt.Sprintf("SubscribeStorageRaw error: %v", err))
		return
	}
	defer sub.Unsubscribe()

	for {
		set := <-sub.Chan()

		logs.Normal(fmt.Sprintf("---------- block hash ---------- %v", set.Block.Hex()))

		// inner loop for the changes within one of those notifications
		for _, chng := range set.Changes {
			// 判断chng.StorageKey是否为chain.Conn.KeyEvents
			if !codec.Eq(chng.StorageKey, chain.Conn.KeyEvents) || !chng.HasStorageData {
				// skip, we are only interested in tEvents with content
				continue
			}

			// Decode the event records
			tEvents := events.EventRecords{}
			err = types.EventRecordsRaw(chng.StorageData).DecodeEventRecords(chain.Conn.Metadata, &tEvents)
			if err != nil {
				logs.Error(fmt.Sprintf("DecodeEventRecords error block hash : %v", set.Block.Hex()))
				logs.Error(fmt.Sprintf("DecodeEventRecords error : %v", err))
				continue
			}

			// Show what we are busy with

			for _, e := range tEvents.Oss_Authorize {
				logs.Normal(fmt.Sprintf("Oss:Authorize:: (phase=%#v)", e.Phase))
				account, err := utils.EncodePublicKeyAsCessAccount(e.Acc[:])
				if err != nil {
					return
				}
				operator, err := utils.EncodePublicKeyAsCessAccount(e.Operator[:])
				if err != nil {
					return
				}
				logs.Normal(fmt.Sprintf("acc : %v, operator : %v", account, operator))

				if utils.AreStorageKeysEqual(chain.Wallet.KeyringPair.PublicKey, e.Acc[:]) {
					logs.Result("bingo!!!")
					return
				}
			}

			//for _, e := range tEvents.Balances_Endowed {
			//	fmt.Printf("\tBalances:Endowed:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#x, %v\n", e.Who, e.Balance)
			//}
			//for _, e := range tEvents.Balances_DustLost {
			//	fmt.Printf("\tBalances:DustLost:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#x, %v\n", e.Who, e.Balance)
			//}
			//for _, e := range tEvents.Balances_Transfer {
			//	fmt.Printf("\tBalances:Transfer:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v, %v, %v\n", e.From, e.To, e.Value)
			//}
			//for _, e := range tEvents.Balances_BalanceSet {
			//	fmt.Printf("\tBalances:BalanceSet:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v, %v, %v\n", e.Who, e.Free, e.Reserved)
			//}
			//for _, e := range tEvents.Balances_Deposit {
			//	fmt.Printf("\tBalances:Deposit:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v, %v\n", e.Who, e.Balance)
			//}
			//for _, e := range tEvents.Grandpa_NewAuthorities {
			//	fmt.Printf("\tGrandpa:NewAuthorities:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.NewAuthorities)
			//}
			//for _, e := range tEvents.Grandpa_Paused {
			//	fmt.Printf("\tGrandpa:Paused:: (phase=%#v)\n", e.Phase)
			//}
			//for _, e := range tEvents.Grandpa_Resumed {
			//	fmt.Printf("\tGrandpa:Resumed:: (phase=%#v)\n", e.Phase)
			//}
			//for _, e := range tEvents.ImOnline_HeartbeatReceived {
			//	fmt.Printf("\tImOnline:HeartbeatReceived:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#x\n", e.AuthorityID)
			//}
			//for _, e := range tEvents.ImOnline_AllGood {
			//	fmt.Printf("\tImOnline:AllGood:: (phase=%#v)\n", e.Phase)
			//}
			//for _, e := range tEvents.ImOnline_SomeOffline {
			//	fmt.Printf("\tImOnline:SomeOffline:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.IdentificationTuples)
			//}
			//for _, e := range tEvents.Indices_IndexAssigned {
			//	fmt.Printf("\tIndices:IndexAssigned:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#x%v\n", e.AccountID, e.AccountIndex)
			//}
			//for _, e := range tEvents.Indices_IndexFreed {
			//	fmt.Printf("\tIndices:IndexFreed:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.AccountIndex)
			//}
			//for _, e := range tEvents.Offences_Offence {
			//	fmt.Printf("\tOffences:Offence:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v%v\n", e.Kind, e.OpaqueTimeSlot)
			//}
			//for _, e := range tEvents.Session_NewSession {
			//	fmt.Printf("\tSession:NewSession:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.SessionIndex)
			//}
			//for _, e := range tEvents.Staking_OldSlashingReportDiscarded {
			//	fmt.Printf("\tStaking:OldSlashingReportDiscarded:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.SessionIndex)
			//}
			//for _, e := range tEvents.System_ExtrinsicSuccess {
			//	fmt.Printf("\tSystem:ExtrinsicSuccess:: (phase=%#v)\n", e.Phase)
			//}
			//for _, e := range tEvents.System_ExtrinsicFailed {
			//	fmt.Printf("\tSystem:ErtrinsicFailed:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%v\n", e.DispatchError)
			//}
			//for _, e := range tEvents.System_CodeUpdated {
			//	fmt.Printf("\tSystem:CodeUpdated:: (phase=%#v)\n", e.Phase)
			//}
			//for _, e := range tEvents.System_NewAccount {
			//	fmt.Printf("\tSystem:NewAccount:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#x\n", e.Who)
			//}
			//for _, e := range tEvents.System_KilledAccount {
			//	fmt.Printf("\tSystem:KilledAccount:: (phase=%#v)\n", e.Phase)
			//	fmt.Printf("\t\t%#X\n", e.Who)
			//}
		}
	}
}

func NewSubscribeWrapper(info *chain.InfoChain) *WrapperSubscribe {
	return &WrapperSubscribe{info}
}
