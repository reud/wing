package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	TemplateCppFileName   string
	TemplateShellFileName string
	Number                int
)

func BinaryReplace(b []byte, from string, to string) []byte {
	s := string(b)
	return []byte(strings.Replace(s, from, to, -1))
}

func GetBinaryFromFile(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return b
}

func WriteFile(path string, b []byte) error {
	return ioutil.WriteFile(path, b, 0666)
}

func WriteContestCMakeList(contest string, number string, cnt int) {
	//os.O_RDWRを渡しているので、同時に読み込みも可能
	file, err := os.OpenFile("CMakeLists.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, fmt.Sprintf("# Contest %+v %+v", contest, number)) //書き込み
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(file, fmt.Sprintf("add_executable(%+v-%+v_%+v %+v/%+v/%+v.cpp)", contest, number, string(rune('a'+i)), contest, number, string(rune('a'+i)))) //書き込み
	}
}

func WriteCodeFiles(contest string, number string, cnt int, template []byte) {
	for i := 0; i < cnt; i++ {
		fn := fmt.Sprintf("%+v/%+v/%+v.cpp", contest, number, string(rune('a'+i)))
		if err := ioutil.WriteFile(fn, template, 0666); err != nil {
			fmt.Printf("%+v", err)
		}
	}
}

func CreateDirectory(directory string, contest string) error {
	return os.MkdirAll(fmt.Sprintf("%+v/%+v", directory, contest), 0777)
}
