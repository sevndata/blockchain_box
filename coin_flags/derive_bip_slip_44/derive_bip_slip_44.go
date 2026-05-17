package derive_bip_slip_44

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/sevndata/blockchain_box/coin_flags"
)

const (
	Derive_BIP_SLIP_Tag = uint32(44) // 默认 BIP执行标准为BIP-44 or SLIP-44
)

// DeriveBIP44PathByCoinFlag 根据 CoinFlag 获取币种的 BIP44 硬化路径
func DeriveBIP44PathByCoinFlag(coinFlag coin_flags.CoinFlag) []uint32 {
	// BIP44 的目的为 44，硬化处理后使用 | hdkeychain.HardenedKeyStart
	purpose := Derive_BIP_SLIP_Tag | hdkeychain.HardenedKeyStart

	// 获取币种对应的硬化编号
	coinType := coinFlag.GetBIPSLIPCoinNumber() | hdkeychain.HardenedKeyStart

	// 设置账户、变更和地址索引（通常为0）
	account := uint32(0) | hdkeychain.HardenedKeyStart
	change := uint32(0)       // 0 表示外部链，1 表示内部链
	addressIndex := uint32(0) // 第一个地址

	//完整的硬化路径
	aPath := []uint32{purpose, coinType, account, change, addressIndex}
	return aPath
}

// DeriveBIP44PathByCoinName 根据币种名称获取对应的 BIP44 硬化路径
func DeriveBIP44PathByCoinName(coinName string) []uint32 {
	// 根据 coinName 获取对应的 CoinFlag
	return DeriveBIP44PathByCoinFlag(coin_flags.GetCoinFlagByCoinName(coinName))
}
