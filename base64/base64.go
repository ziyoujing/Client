package base64

import "encoding/base64"

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}
func EncodeRaw(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func DecodeRaw(str string) []byte {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil
	}
	return data
}
