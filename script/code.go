package main

import (
	"bufio"
	"fmt"
	"github.com/micro/go-micro/debug/log"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// 根据code的定义生成到proto

func main() {
	codeDir := "./comm on/logic_error/code"
	protoFile := "./proto/platform/platform-ecode.proto"

	_ = protoFile

	files, err := ioutil.ReadDir(codeDir)
	if err != nil {
		log.Fatal(err)
	}

	putStr := ""

	for _, f := range files {
		file := codeDir + "/" + f.Name()
		content := getFileContent(file)
		codes := getCodeLine(content)
		doc := getFileDoc(content)
		putStr = putStr + "\n	// form " + file + "\n	" + doc + "\n" + codes
	}

	putText(protoFile, putStr)
}

func getFileDoc(text []string) string {
	for _, v := range text {
		v = strings.Trim(v, " ")
		has := strings.Index(v, "//")

		if has == 0 {
			return v
		}
	}

	return ""
}

type protoText struct {
	start string
	end   string
}

func putText(protoFile string, str string) {
	bytes, err := ioutil.ReadFile(protoFile)
	if err != nil {
		fmt.Println("error : %s", err)
		return
	}

	text := string(bytes)

	startSplit := "// @stack(code)"
	endSplit := "// @end(code)"

	arr := strings.Split(text, startSplit)
	arr2 := strings.Split(arr[1], endSplit)

	proto := protoText{
		start: arr[0],
		end:   arr2[1],
	}

	writeString := []byte(proto.start + startSplit + "\n" + str + "\n	" + endSplit + proto.end)
	_ = ioutil.WriteFile(protoFile, writeString, 0666)
}

func getCodeLine(text []string) string {
	var line string

	for _, v := range text {
		v = strings.Trim(v, " ")
		v = strings.Trim(v, "\t")
		has := strings.Index(v, "//")

		if has == 0 {
			continue
		}

		has = strings.Index(v, "=")

		if has == -1 {
			continue
		}

		has = strings.Index(v, "uint32")
		if has == -1 {
			log.Errorf("%s 状态码定义必须是uint32", v)
			continue
		}

		v = strings.Replace(v, "uint32", "", -1)
		v = strings.Replace(v, "_", "", -1)
		v = delete_extra_space(v)

		line += "\t"
		has = strings.Index(v, "//")
		if has == -1 {
			log.Errorf("%s 状态码定义必须要注释", v)
			line = line + v + ";\n"
		} else {
			arr := strings.Split(v, "//")
			line = line + arr[0] + "; // " + arr[1] + "\n"
		}
	}

	return line
}

func getFileContent(filename string) []string {
	var nameList []string
	r, _ := os.Open(filename)
	defer r.Close()
	s := bufio.NewScanner(r)
	for s.Scan() { // 循环直到文件结束
		nameList = append(nameList, s.Text()) // 这个 line 就是每一行的文本了，string 类型
	}

	return nameList
}

// 删除字符串中的多余空格，有多个空格时，仅保留一个空格
func delete_extra_space(s string) string {
	s1 := strings.Replace(s, "	", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}
