package accrual

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/accrual", new(Accrual))
}

var (
	_ modules.Module = &Accrual{}
)

var src = rand.NewSource(time.Now().UnixNano())
var count uint32

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type Accrual struct{}

type NewGoods struct {
	Match      string  `json:"match"`
	Reward     float64 `json:"reward"`
	RewardType string  `json:"reward_type"`
}

func (*Accrual) NewAccrual() *Accrual {
	return &Accrual{}
}

func (*Accrual) Generator() []byte {
	goods := NewGoods{}

	//Генерация рандомного товара
	goods.RewardType = "%"
	goods.Reward = float64(rand.Intn(30))
	str := RandStringBytesMaskImprSrcSB(7)
	goods.Match = str + " " + str[rand.Intn(5):6]

	//Возвращаем готовый JSON для вставки в тест
	jsonData, err := json.MarshalIndent(&goods, "", " ")
	if err != nil {
		fmt.Errorf("error JSON marshal %w", err)
	}
	return jsonData
}

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
