package chain

import (
	"MatrixAI-Client/chain/conn"
	"MatrixAI-Client/chain/wallet"
	"MatrixAI-Client/config"
)

// InfoChain 结构体用于存放conn\wallet信息
type InfoChain struct {
	Conn   *conn.Conn
	Wallet *polkadot_wallet.Wallet
}

// GetChainInfo 创建conn\wallet信息并返回ChainInfo
func GetChainInfo(cfg *config.Config) (*InfoChain, error) {
	newConn, err := conn.NewConn(cfg.ChainRPC)
	if err != nil {
		return nil, err
	}

	wallet, err := polkadot_wallet.InitWallet(cfg)
	if err != nil {
		return nil, err
	}

	chainInfo := &InfoChain{
		Conn:   newConn,
		Wallet: wallet,
	}

	return chainInfo, nil
}
