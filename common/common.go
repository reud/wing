package common

import (
	"io/ioutil"
	"strings"
)

func BinaryReplace(b []byte, contestName string) []byte {
	s := string(b)
	return []byte(strings.Replace(s, "###CONTEST_NAME###", contestName, -1))
}

func WriteFile(path string, b []byte) error {
	return ioutil.WriteFile(path, b, 0666)
}
