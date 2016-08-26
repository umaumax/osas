package osas

import (
	"regexp"
	"strconv"
)

type Result map[string]string

func (r Result) IsOK() (ok bool) {
	_, ok = r["OK"]
	return
}

func (r Result) IsCancel() (ok bool) {
	_, ok = r["CANCEL"]
	return
}

func (r Result) IsError() (ok bool) {
	_, ok = r["ERROR"]
	return
}

func (r Result) IsTimeOut() (ok bool) {
	_, ok = r["TIMEOUT"]
	return
}

func (r Result) GetOKText() string {
	return r["OK"]
}
func (r Result) GetCancelText() string {
	return r["CANCEL"]
}
func (r Result) GetErrorText() string {
	return r["ERROR"]
}
func (r Result) GetTimeOutText() string {
	return r["TIMEOUT"]
}

func (r Result) GetOKTexts() []string {
	return r.GetTexts("OK")
}

func (r Result) GetText(name string) string {
	return r[name]
}

func (r Result) GetTexts(name string) []string {
	text, ok := r[name]
	if !ok {
		return nil
	}
	return regexp.MustCompile("\r\n|\n\r|\n|\r").Split(text, -1)
}

func (r Result) GetRGB() (cr, cg, cb int) {
	rgb := r.GetOKTexts()
	if len(rgb) != 3 {
		return 0, 0, 0
	}
	cr, _ = strconv.Atoi(rgb[0])
	cg, _ = strconv.Atoi(rgb[1])
	cb, _ = strconv.Atoi(rgb[2])
	return
}
