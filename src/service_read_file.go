package src

import (
	"github.com/json-iterator/go"
	"io/ioutil"
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
func MustGetMatchesFromMatchSamples(path string) *Matches {
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
func (matches *Matches) PrintCompareResult() {
	for _, v := range matches.Matches {
		Compare(v.PlayerA, v.PlayerB)
	}
}
