package wallet

import "github.com/sevndata/blockchain_box/coin_flags"

type Wallet struct {
	Coin                    coin_flags.CoinFlag // 币种标识
	CoinName                string              // 币种名称
	WalletAddress           string              // 钱包地址
	AliasAddress            string              // 别名地址，针对有多重转换的地址适配预留字段
	AliasAddressAppendZeroX string              // 别名地址，针对有多重转换的地址适配预留字段 补充0x开头
	PublicKey               string              // 公钥
	PrivateKey              string              // 私钥
	Mnemonic                string              // 助记词
}
