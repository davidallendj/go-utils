package util

import (
	"encoding/base64"
	"math/rand"
	"os/exec"
	"runtime"
	"strings"

	"github.com/golang-jwt/jwt"
)

func QuoteArrayStrings(arr []string) []string {
	for i, v := range arr {
		arr[i] = "\"" + v + "\""
	}
	return arr
}

func OpenUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func GetCommit() string {
	bytes, err := exec.Command("git", "rev --parse HEAD").Output()
	if err != nil {
		return ""
	}
	return string(bytes)
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeJwt(encoded string) ([][]byte, error) {
	// split the string into 3 segments and decode
	segments := strings.Split(encoded, ".")
	decoded := [][]byte{}
	for _, segment := range segments {
		bytes, _ := jwt.DecodeSegment(segment)
		decoded = append(decoded, bytes)
	}
	return decoded, nil
}

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
