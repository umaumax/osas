package osas

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

func StringsToListStrings(items []string) string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, item := range items {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%q", item)
	}
	buf.WriteString("}")
	return buf.String()
}
func IntsToListInts(items []int) string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, item := range items {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", item)
	}
	buf.WriteString("}")
	return buf.String()
}

func MakeListAddPrefixFunc(prefix string) func([]string) string {
	return func(input []string) string {
		if input == nil {
			return ""
		}
		return prefix + " " + StringsToListStrings(input)
	}
}
func MakeIntListAddPrefixFunc(prefix string) func([]int) string {
	return func(input []int) string {
		if input == nil {
			return ""
		}
		return prefix + " " + IntsToListInts(input)
	}
}
func MakeAddPrefixFunc(prefix string) func(string) string {
	return func(input string) string {
		if input == "" {
			return ""
		}
		return prefix + " " + fmt.Sprintf("%q", input)
	}
}
func MakeAddPrefixNotQoutedFunc(prefix string) func(string) string {
	return func(input string) string {
		if input == "" {
			return ""
		}
		return prefix + " " + input
	}
}
func MakeBoolAddPrefixFunc(prefix string) func(bool) string {
	return func(input bool) string {
		return prefix + " " + fmt.Sprintf("%t", input)
	}
}
func MakeIntAddPrefixFunc(prefix string) func(int) string {
	return func(input int) string {
		if input == 0 {
			return ""
		}
		return prefix + " " + fmt.Sprintf("%d", input)
	}
}

var Verbose bool

func RunOSAScript(t *template.Template, data interface{}) (result Result, err error) {
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return
	}
	if Verbose {
		fmt.Println(buf.String())
	}
	cmd := exec.Command("osascript", "-e", buf.String())
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("osascript %s\n%s", err, string(bytes))
		return
	}
	//	AppleScriptからの返り値に改行が含まれている
	str := strings.TrimSpace(string(bytes))
	index := strings.Index(str, ":")
	if index < 0 {
		err = fmt.Errorf("unexpected result:[%s]", str)
		return
	}
	result = make(Result)
	key := str[:index]
	value := str[index+1:]
	result[key] = value
	return
}
