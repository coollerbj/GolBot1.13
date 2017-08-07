package components

import (
	"encoding/base64"
)

func deobfuscate(Data string) string {
	var ClearText string
	for i := 0; i < len(Data); i++ {
		ClearText += string(int(Data[i]) - 1)
	}
	return ClearText
}

func obfuscate(Data string) string {
	var ObfuscateText string
	for i := 0; i < len(Data); i++ {
		ObfuscateText += string(int(Data[i]) + 1)
	}
	return ObfuscateText
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

