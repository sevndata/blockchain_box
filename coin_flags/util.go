package coin_flags

func GetCoinFlagByCoinName(name string) CoinFlag {
	switch name {
	case "BTC", "BitCoin", "Bitcoin":
		return CoinFlagBTC
	case "LTC", "LiteCoin", "Litecoin":
		return CoinFlagLTC
	case "DOGE", "DogeCoin", "Dogecoin":
		return CoinFlagDOGE
	case "ETC", "Ethereum Classic", "EthereumClassic":
		return CoinFlagETC
	case "ETHW", "EthereumPoW":
		return CoinFlagETHW
	case "ZIL", "Zilliqa":
		return CoinFlagZIL
	case "OCTA", "OctaSpace":
		return CoinFlagOCTA
	case "META", "MetaChain":
		return CoinFlagMETA
	case "CAU", "Canxium":
		return CoinFlagCAU
	case "DNX", "Dynexcoin":
		return CoinFlagDNX
	case "ETH", "Ethereum":
		return CoinFlagETH
	case "SOL", "Solana":
		return CoinFlagSOL
	case "FIL", "FileCoin", "Filecoin":
		return CoinFlagFIL
	case "ALEO", "Aleo", "aleo":
		return CoinFlagAleo
	case "BEL", "BELLS", "Bel", "Bells":
		return CoinFlagBEL
	case "PEP", "pep", "Pep", "Pepecoin", "PepeCoin":
		return CoinFlagPEP
	case "AUS", "AUS-Cash", "Australiacash", "Australia Cash", "AustraliaCash":
		return CoinFlagAUS
	case "Earthcoin", "EAC", "Earth", "earth", "EarthCoin", "eac":
		return CoinFlagEAC
	case "Dogmcoin", "DogmCoin", "dogmCoin", "dogmcoin", "DOGMCOIN", "DOGM", "Dogm", "dogm":
		return CoinFlagDOGM
	case "Dingocoin", "DingoCoin", "dingoCoin", "dingocoin", "DINGOCOIN", "DINGO", "Dingo", "dingo":
		return CoinFlagDINGO
	case "Luckycoin", "LuckyCoin", "luckyCoin", "luckycoin", "LUCKYCOIN", "LKY", "Lky", "lky":
		return CoinFlagLKY
	case "Junkcoin", "JunkCoin", "junkCoin", "junkcoin", "JUNKCOIN", "JKC", "Jkc", "jkc":
		return CoinFlagJKC
	case "Craftcoin", "CraftCoin", "craftCoin", "craftcoin", "CRAFTCOIN", "CRC", "Crc", "crc":
		return CoinFlagCRC
	case "Myriadcoin", "MyriadCoin", "myriadCoin", "myriadcoin", "MYRIADCOIN", "XMY", "Xmy", "xmy", "MYRIAD", "Myriad", "myriad":
		return CoinFlagXMY
	case "Shibacoin", "ShibaCoin", "shibaCoin", "shibacoin", "SHIBACOIN", "Shibainucoin", "ShibainuCoin", "shibainuCoin", "shibainucoin", "SHIBAINUCOIN", "SHIC", "shic", "Shic":
		return CoinFlagSHIC
	case "BeersCoin", "Beerscoin", "beersCoin", "beerscoin", "BEERSCOIN", "BRC", "Brc", "brc":
		return CoinFlagBRC
	case "FlinCoin", "Flincoin", "flinCoin", "flincoin", "FLINCOIN", "FLIN", "Flin", "flin":
		return CoinFlagFLIN
	case "BonkCoin", "Bonkcoin", "bonkCoin", "bonkcoin", "BONKCOIN", "BONC", "Bonc", "bonc":
		return CoinFlagBONC
	case "DogecoinEV", "Dogecoinev", "dogecoinEV", "dogecoinev", "DOGECOINEV", "DEV", "Dev", "dev":
		return CoinFlagDEV
	default:
		return CoinFlagNone
	}
}
