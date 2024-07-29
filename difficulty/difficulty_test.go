package difficulty

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHexStringDiffToDiifStruct(t *testing.T) {
	targetHex := "0x000000002bfaffc2f2c92abfbde3da69454e75d23cbfce328fb6a22c24f80d3c"

	aDiff := HexDifficultyToDifficulty(targetHex)
	jstr, _ := json.MarshalIndent(&aDiff, "", "    ")
	fmt.Printf("decode diff :%s\n", string(jstr))
}

func TestNewDiffStructs(t *testing.T) {
	diffNumber := int64(5000000000)

	newDiff := NewDifficultyMap(diffNumber)
	newjstr, _ := json.MarshalIndent(&newDiff, "", "    ")
	fmt.Printf("new diff :%s\n", string(newjstr))

	achanDiff := HexDifficultyToDifficulty("0x000000000000000000000000000000000000000000000000000000012a05f200")
	achanDiffJstr, _ := json.MarshalIndent(&achanDiff, "", "    ")
	fmt.Printf("dec new diff :%s\n", string(achanDiffJstr))
}
