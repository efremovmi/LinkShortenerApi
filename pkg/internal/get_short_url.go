package internal

import "github.com/speps/go-hashids"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = 10
)

func GetHash(longURL string) string {
	hd := hashids.NewData()
	hd.MinLength = length
	hd.Alphabet = alphabet
	hd.Salt = longURL
	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{1, 2, 3, 4, 5})
	return id
}

func GetShortUrl(longURL string) string {
	ind := 0
	str := "0000000000"
	shift := length * 2
	end := ind + shift
	for ind < len(longURL) {
		if end > len(longURL) {
			end = len(longURL)
			if ind == len(longURL) {
				break
			}
		}

		str = SumTwoStrings(GetHash(longURL[ind:end]), str)
		ind += shift
		end = ind + shift

	}

	res := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		res[i] = str[i] % 122
	}

	return GetHash(string(res))
}

func SumTwoStrings(left, right string) string {
	res := make([]byte, len(left))
	for i, _ := range left {
		res[i] = left[i] + right[i]
	}
	return string(res)
}
