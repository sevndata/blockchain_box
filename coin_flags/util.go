package coin_flags

func GetCoinFlagByCoinName(name string) CoinFlag {
	switch name {
	case "BitCoin", "Bitcoin", "bitcoin", "BTC", "btc", "Btc":
		return CoinFlagBTC
	case "LiteCoin", "Litecoin", "LTC", "Ltc", "ltc":
		return CoinFlagLTC
	case "DogeCoin", "Dogecoin", "dogecoin", "DOGE", "doge", "Doge":
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
	case "BELLS", "Bells", "bells", "BEL", "Bel", "bel":
		return CoinFlagBEL
	case "PEP", "pep", "Pep", "Pepecoin", "PepeCoin", "pepcoin":
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
	case "BBQCoin", "BBQcoin", "bbqCoin", "bbqcoin", "BBQCOIN", "BQC", "Bqc", "bqc":
		return CoinFlagBQC
	case "MarsCoin", "Marscoin", "marsCoin", "marscoin", "MARSCOIN", "MARS", "Mars", "mars":
		return CoinFlagMARS
	case "FlopCoin", "Flopcoin", "flopCoin", "flopcoin", "FLOPCOIN", "FLOP", "Flop", "flop":
		return CoinFlagFLOP
	case "NameCoin", "Namecoin", "nameCoin", "namecoin", "NAMECOIN", "NMC", "Nmc", "nmc":
		return CoinFlagNMC
	case "FractalBitcoin", "Fractalbitcoin", "fractalBitcoin", "fractalbitcoin", "FRACTALBITCOIN", "FB", "Fb", "fb":
		return CoinFlagFB
	case "Hathor", "hathor", "HATHOR", "HTR", "Htr", "htr":
		return CoinFlagHTR
	case "Elastos", "elastos", "ELASTOS", "ELA", "Ela", "ela":
		return CoinFlagELA
	case "BitcoinCash", "Bitcoincash", "bitcoinCash", "bitcoincash", "BITCOINCASH", "BCH", "Bch", "bch":
		return CoinFlagBCH
	case "LebowskisCoin", "Lebowskiscoin", "lebowskisCoin", "lebowskiscoin", "LEBOWSKISCOIN", "LBW", "Lbw", "lbw":
		return CoinFlagLBW
	case "Bit", "bit", "BIT", "B1T", "b1t", "B1t":
		return CoinFlagB1T
	case "TrumPOW", "Trumpow", "trumPOW", "trumpow", "TRUMPOW", "TRMP", "Trmp", "trmp":
		return CoinFlagTRMP
	case "TRON", "TRX":
		return CoinFlagTRX
	default:
		return CoinFlagNone
	}
}
