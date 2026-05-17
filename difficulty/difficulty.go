package difficulty

import (
	"fmt"
	"github.com/sevndata/blockchain_box/coin_flags"
	"math/big"
	"strings"
)

// Difficulty 难度结构体
type Difficulty struct {
	Algorithm                     string
	Difficulty                    *big.Int
	DifficultyStringShort         string
	DifficultyStringShortFull     string
	DifficultyStringShort1024     string
	DifficultyStringShortFull1024 string
	DifficultyUnitShort           string
	DifficultyHexString           string
	DifficultyUnit                DiffUnit
	DifficultyUnitExtendStr       string     // 扩展拼接字段
	DifficultyValue               *big.Float // 存储计算后的难度值
}

// FormatUnit 单位计算公式 以 byte 为单位传入，根据数值不同换算成进制单位MB、GB等
// fSed 小数点保留位数
func (diff *Difficulty) FormatUnit(fSed int) {
	baseDifficulty := new(big.Float).SetInt(diff.Difficulty)
	k := big.NewFloat(1000)
	m := new(big.Float).Mul(k, k)
	g := new(big.Float).Mul(m, k)
	t := new(big.Float).Mul(g, k)
	p := new(big.Float).Mul(t, k)
	e := new(big.Float).Mul(p, k)
	z := new(big.Float).Mul(e, k)
	y := new(big.Float).Mul(z, k)

	k1024 := big.NewFloat(1024)
	m1024 := new(big.Float).Mul(k, k1024)
	g1024 := new(big.Float).Mul(m, k1024)
	t1024 := new(big.Float).Mul(g, k1024)
	p1024 := new(big.Float).Mul(t, k1024)
	e1024 := new(big.Float).Mul(p, k1024)
	z1024 := new(big.Float).Mul(e, k1024)
	y1024 := new(big.Float).Mul(z, k1024)

	format := fmt.Sprintf("%%.%df", fSed)

	switch {
	case baseDifficulty.Cmp(y) >= 0:
		diff.DifficultyUnit = DiffUnitYotta
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, y))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, y1024))

	case baseDifficulty.Cmp(z) >= 0:
		diff.DifficultyUnit = DiffUnitZetta
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, z))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, z1024))

	case baseDifficulty.Cmp(e) >= 0:
		diff.DifficultyUnit = DiffUnitExa
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, e))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, e1024))

	case baseDifficulty.Cmp(p) >= 0:
		diff.DifficultyUnit = DiffUnitPeta
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, p))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, p1024))

	case baseDifficulty.Cmp(t) >= 0:
		diff.DifficultyUnit = DiffUnitTera
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, t))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, t1024))

	case baseDifficulty.Cmp(g) >= 0:
		diff.DifficultyUnit = DiffUnitGiga
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, g))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, g1024))

	case baseDifficulty.Cmp(m) >= 0:
		diff.DifficultyUnit = DiffUnitMega
		diff.DifficultyStringShort = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, m))
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, m1024))

	case baseDifficulty.Cmp(k) >= 0:
		diff.DifficultyUnit = DiffUnitKilo
		diff.DifficultyStringShort1024 = fmt.Sprintf(format, new(big.Float).Quo(baseDifficulty, k1024))
	default:
		diff.DifficultyUnit = DiffUnitNone
		diff.DifficultyStringShort = fmt.Sprintf(format, baseDifficulty)
	}

	diff.DifficultyStringShortFull = fmt.Sprintf("%s%s", diff.DifficultyStringShort, diff.DifficultyUnit.Shortname())
	diff.DifficultyStringShortFull1024 = fmt.Sprintf("%s%s", diff.DifficultyStringShort1024, diff.DifficultyUnit.Shortname())
	diff.DifficultyUnitShort = diff.DifficultyUnit.Shortname()
}

func (diff *Difficulty) CalculateSpecificDifficulty() {
	switch diff.Algorithm {
	case coin_flags.CoinFlagBTC.PowAlgorithm():
		diff.DifficultyValue = new(big.Float).SetInt(TargetHexToDiffBySha256d(diff.DifficultyHexString))

	case coin_flags.CoinFlagLTC.PowAlgorithm():
		diff.DifficultyValue = new(big.Float).SetInt(TargetHexToDiffByScrypt(diff.DifficultyHexString))

	case coin_flags.CoinFlagETC.PowAlgorithm(), coin_flags.CoinFlagETHW.PowAlgorithm():
		diff.DifficultyValue = new(big.Float).Quo(new(big.Float).SetInt(pow256Ethash), new(big.Float).SetInt(diff.Difficulty))

	case coin_flags.CoinFlagDNX.PowAlgorithm():
		// 需要根据 DynexSolve 的具体难度计算方式进行实现
		diff.DifficultyValue = new(big.Float).SetInt(diff.Difficulty)

	case coin_flags.CoinFlagAleo.PowAlgorithm():
		// zkSNARK 不使用工作量证明，难度为零
		diff.DifficultyValue = big.NewFloat(0)

	default:
		// 未知算法，设置为零
		diff.DifficultyValue = big.NewFloat(0)
	}
}

func HexDifficultyToDifficulty(hexDifficulty string) map[string]*Difficulty {
	targetHex := strings.TrimPrefix(hexDifficulty, "0x")
	target := new(big.Int)
	target.SetString(targetHex, 16)

	diffMap := map[string]*Difficulty{}

	for _, coin := range coin_flags.SupportedCoinsMap {
		var difficulty *big.Int
		switch coin.PowAlgorithm() {
		case "SHA256d":
			difficulty = TargetHexToDiffBySha256d(hexDifficulty)
		case "Scrypt":
			difficulty = TargetHexToDiffByScrypt(hexDifficulty)
		case "Ethash", "EtcHash":
			difficulty = TargetHexToDiff(hexDifficulty)
		default:
			difficulty = target
		}

		diff := &Difficulty{
			Algorithm:           coin.PowAlgorithm(),
			Difficulty:          difficulty,
			DifficultyHexString: hexDifficulty,
		}
		diff.CalculateSpecificDifficulty()
		diff.FormatUnit(3)
		diffMap[coin.PowAlgorithm()] = diff
	}

	return diffMap
}

func NewDifficultyMap(diffNumber int64) map[string]*Difficulty {
	diffMap := map[string]*Difficulty{}

	for _, coin := range coin_flags.SupportedCoinsMap {
		var hexStr string
		switch coin.PowAlgorithm() {
		case "SHA256d":
			hexStr = GetTargetHexBySha256d(diffNumber)
		case "Scrypt":
			hexStr = GetTargetHexByScrypt(diffNumber)
		case "Ethash", "EtcHash":
			hexStr = GetTargetHex(diffNumber)
		default:
			hexStr = new(big.Int).SetInt64(diffNumber).Text(16)
			if len(hexStr) < 66 {
				hexStr = "0x" + strings.Repeat("0", 66-len(hexStr)) + hexStr[2:]
			}
		}

		difficulty := new(big.Int).SetInt64(diffNumber)
		diff := &Difficulty{
			Algorithm:           coin.PowAlgorithm(),
			Difficulty:          difficulty,
			DifficultyHexString: hexStr,
		}
		diff.CalculateSpecificDifficulty()
		diff.FormatUnit(3)
		diffMap[coin.PowAlgorithm()] = diff
	}

	return diffMap
}
