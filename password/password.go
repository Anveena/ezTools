package password

import (
	"encoding/base64"
	"github.com/Anveena/ezTools/crypto"
)

func Encode(origStr string) (string, error) {
	return EncodeData([]byte(origStr))
}
func Decode(base64Str string) (string, error) {
	rs, err := DecodeData(base64Str)
	if err != nil {
		return "", err
	}
	return string(rs), nil
}
func EncodeData(origData []byte) (string, error) {
	if len(origData) == 0 {
		return "", nil
	}
	encData, err := crypto.EZEncrypt(origData, "this code may be not working", 9458)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encData), nil
}
func DecodeData(base64Str string) ([]byte, error) {
	if len(base64Str) == 0 {
		return []byte{}, nil
	}
	encData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	origData, err := crypto.EZDecrypt(encData, "this code may be not working", 9458)
	if err != nil {
		return nil, err
	}
	return origData, nil
}
