package src

import (
	"io/ioutil"
	"github.com/json-iterator/go"
	"fmt"
)

type Matches struct {
	Matches []*Match `json:"matches"`
}

type Match struct {
	PlayerA string `json:"alice"`
	PlayerB string `json:"bob"`
	Result  int    `json:"result"`
}

// 获取牌组（必然获取，否则抛出异常）
func MustGetMatchesFromMatchSamples(path string) (*Matches) {
	var file []byte
	var err error

	if file, err = ioutil.ReadFile(path); err != nil {
		panic("panic: " + err.Error())
	}

	matches := Matches{}
	if err := jsoniter.Unmarshal(file, &matches); err != nil {
		panic("panic: " + err.Error())
	}
	return &matches
}

// 打印牌组比较结果
func (matches *Matches)PrintCompareResult(checkResult bool){
	for _, v := range matches.Matches {
		playerA := v.PlayerA
		playerB := v.PlayerB

		if result := Compare(playerA, playerB); checkResult && result != v.Result {
			fmt.Printf("error result:%s, %s ,%d, should be %d\n", playerA, playerB, result, v.Result)
		} else {
			fmt.Printf("%s, %s ,%d\n", playerA, playerB, result)
		}
	}
	fmt.Printf("总计：%d 条\n", len(matches.Matches))
}
