/*
Package blockchain_node_install blockchain Node install tools
*/
package blockchain_node_install

import (
	"errors"
	"github.com/sevndata/blockchain_box/coin_flags"
)

// InstallBlockChainNodeWithCoinFlag 通过传CoinFlag 安装区块链节点
func InstallBlockChainNodeWithCoinFlag(coinFlag coin_flags.CoinFlag) error {
	var err error

	if coinFlag == coin_flags.CoinFlagNone {
		err = errors.New("no support coin")
		return err
	}

	return nil
}

// InstallBlockChainNodeWithCoinName 通过传Coin Name 安装对应的区块链节点
func InstallBlockChainNodeWithCoinName(coinName string) error {
	aCoinFlag := coin_flags.GetCoinFlagByCoinName(coinName)

	return InstallBlockChainNodeWithCoinFlag(aCoinFlag)
}
