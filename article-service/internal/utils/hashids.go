package utils

import (
	"github.com/speps/go-hashids"
)

//salt 盐值
const salt = "XpaxfQKo9"

//Encode 混淆
func EncodeID(data int) string {
	hd, _ := hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: 8,
		Salt:      salt,
	})
	e, _ := hd.Encode([]int{data})
	return e
}

//Decode 还原混淆
func DecodeID(data string) (int, error) {
	//todo 异常处理,需要捕获panic
	hd, _ := hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: 8,
		Salt:      salt,
	})
	e, err := hd.DecodeWithError(data)
	return e[0], err
}
