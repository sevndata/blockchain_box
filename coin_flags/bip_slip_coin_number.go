package coin_flags

// BIP and SLIP format Coin Number
/*
	SLIP-44: https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	BIP-44: https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
*/
const (
	BIP_SLIP_CoinNumber_BitCoin         = 0
	BIP_SLIP_CoinNumber_BitCoinTestNet  = 1
	BIP_SLIP_CoinNumber_LiteCoin        = 2
	BIP_SLIP_CoinNumber_DogeCoin        = 3
	BIP_SLIP_CoinNumber_Ethereum        = 60
	BIP_SLIP_CoinNumber_EthereumClassic = 61
	BIP_SLIP_CoinNumber_Zilliqa         = 313
	BIP_SLIP_CoinNumber_Filecoin        = 461
	BIP_SLIP_CoinNumber_Solana          = 501
	BIP_SLIP_CoinNumber_Aleo            = 683
	BIP_SLIP_CoinNumber_Bells           = 1
	BIP_SLIP_CoinNumber_Pep             = 1 //?=pepe 3434
	BIP_SLIP_CoinNumber_Aus             = 1 //not found aus slip
	BIP_SLIP_CoinNumber_DogmCoin        = 1 //not found dogm slip
	BIP_SLIP_CoinNumber_Earthcoin       = 1 // ?=earths 406
	BIP_SLIP_CoinNumber_Dingocoin       = 1 //not found dingo slip
	BIP_SLIP_CoinNumber_Luckycoin       = 1 //not found lky slip
	BIP_SLIP_CoinNumber_Junkcoin        = 1 //not found jkc slip
	BIP_SLIP_CoinNumber_Craftcoin       = 1 //not found crc slip
	BIP_SLIP_CoinNumber_Myriad          = 1 //not found xmy slip

)

// GetBIPSLIPCoinNumber
/*
	en: get bip-44 coinNumber;
	zh-CN: 获取BIP-44 或者 SLIP44 规范的 币种编号;
	@return [☑][type: int64] en: coin number ;zh-CN: 比重编号;
*/
func (cf CoinFlag) GetBIPSLIPCoinNumber() uint32 {
	switch cf {
	case CoinFlagBTC:
		return BIP_SLIP_CoinNumber_BitCoin
	case CoinFlagLTC:
		return BIP_SLIP_CoinNumber_LiteCoin
	case CoinFlagDOGE:
		return BIP_SLIP_CoinNumber_DogeCoin
	case CoinFlagETC:
		return BIP_SLIP_CoinNumber_EthereumClassic
	case CoinFlagETHW:
		return BIP_SLIP_CoinNumber_Ethereum
	case CoinFlagZIL:
		return BIP_SLIP_CoinNumber_Zilliqa
	case CoinFlagETH:
		return BIP_SLIP_CoinNumber_Ethereum
	case CoinFlagSOL:
		return BIP_SLIP_CoinNumber_Solana
	case CoinFlagFIL:
		return BIP_SLIP_CoinNumber_Filecoin
	case CoinFlagAleo:
		return BIP_SLIP_CoinNumber_Aleo
	case CoinFlagBEL:
		return BIP_SLIP_CoinNumber_Bells
	case CoinFlagPEP:
		return BIP_SLIP_CoinNumber_Pep
	case CoinFlagAUS:
		return BIP_SLIP_CoinNumber_Aus
	case CoinFlagDOGM:
		return BIP_SLIP_CoinNumber_DogmCoin
	case CoinFlagEAC:
		return BIP_SLIP_CoinNumber_Earthcoin
	case CoinFlagDINGO:
		return BIP_SLIP_CoinNumber_Dingocoin
	case CoinFlagLKY:
		return BIP_SLIP_CoinNumber_Luckycoin
	case CoinFlagJKC:
		return BIP_SLIP_CoinNumber_Junkcoin
	case CoinFlagCRC:
		return BIP_SLIP_CoinNumber_Craftcoin
	case CoinFlagXMY:
		return BIP_SLIP_CoinNumber_Myriad
	default:
		return BIP_SLIP_CoinNumber_BitCoinTestNet
	}
}
