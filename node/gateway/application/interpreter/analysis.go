package interpreter

import (
	"bytes"
	"strings"
)

var (
	start = []byte("{{")
	end   = []byte("}}")

	restfulStart = []byte("{")
	restfulEnd   = []byte("}")
	restfulHead  = []byte(Restful)
	restfulSp    = []byte("/:")
	restfulSpE   = "/?"
)

//Gen gen
func Gen(tpl string, encode string) Interpreter {
	if encode == "origin" {
		exe := make(_Executor, 1, 1)
		exe[0] = new(_OrgReader)
		return exe
	}
	tpl = strings.TrimSpace(tpl)
	i, err := Parse(tpl)
	if err != nil {
		exe := make(_Executor, 1, 1)
		exe[0] = new(_OrgReader)
		return exe
	}
	return i
}

//Parse 编译
func Parse(tpl string) (Interpreter, error) {

	exe := make(_Executor, 0, 10)

	data := []byte(tpl)
	for {
		pre, key, sub := parse(data)

		if pre != nil {
			exe = append(exe, _NotReader(string(pre)))
		}

		if key != nil {
			r, e := genReader(key)
			if e != nil {
				return nil, e
			}
			exe = append(exe, r)
		}
		if sub == nil {
			break
		}
		data = sub

	}
	return exe, nil
}

func parse(data []byte) (pre, key, sub []byte) {

	firstStart := bytes.Index(data, start)
	if firstStart == -1 {
		return data, nil, nil
	}
	firstEnd := bytes.Index(data[firstStart:], end)
	if firstEnd == -1 {
		return data, nil, nil
	}
	firstEnd += firstStart
	pre = data[:firstStart]
	key = data[firstStart+2 : firstEnd]
	if firstEnd+2 < len(data) {
		sub = data[firstEnd+2:]
	}
	return
}

//GenPath genPath
func GenPath(path string) Interpreter {
	interpreter, e := ParsePath(path)
	if e != nil {
		exe := make(_Executor, 1, 1)
		exe[0] = _NotReader(path)
		return exe
	}
	return interpreter
}

//ParsePath 解析路径
func ParsePath(path string) (Interpreter, error) {

	exe := make(_Executor, 0, 10)

	data := []byte(path)
	for {
		pre, key, sub := parse(data)

		//fmt.Println("pre:", string(pre), "\tkey:", string(key), "\tsub:", string(sub))
		if pre != nil {
			restfulReader, err := parsePath(pre)
			if err != nil {
				return nil, err
			}
			exe = append(exe, restfulReader...)
		}

		if key != nil {
			r, e := genReader(key)
			if e != nil {
				return nil, e
			}
			exe = append(exe, r)
		}
		if sub == nil {
			break
		}
		data = sub

	}
	return exe, nil

}

func parsePath(line []byte) ([]Reader, error) {
	readers := make([]Reader, 0, 10)
	data := line
	for {
		pre, key, sub := parsePathDo(data)

		//fmt.Println("path:\tpre:", string(pre), "\tkey:", string(key), "\tsub:", string(sub))
		if pre != nil {

			restfulReader, err := parseRestful(pre)
			if err != nil {
				return nil, err
			}
			readers = append(readers, restfulReader...)
		}

		if key != nil {
			r, e := genResfult(restfulHead, key)
			if e != nil {
				return nil, e
			}
			readers = append(readers, r)
		}
		if sub == nil {
			break
		}
		data = sub

	}
	return readers, nil
}

func parsePathDo(line []byte) (pre, key, sub []byte) {

	firstStart := bytes.Index(line, restfulStart)
	if firstStart == -1 {
		return line, nil, nil
	}
	firstEnd := bytes.Index(line[firstStart:], restfulEnd)
	if firstEnd == -1 {
		return line, nil, nil
	}
	firstEnd += firstStart
	pre = line[:firstStart]
	key = line[firstStart+1 : firstEnd]
	if firstEnd+1 < len(line) {
		sub = line[firstEnd+1:]
	}
	return
}
func parseRestful(line []byte) ([]Reader, error) {
	readers := make([]Reader, 0, 10)
	data := line
	for {
		pre, key, sub := parseRestfulDo(data)

		//fmt.Println("restful:\tpre:", string(pre), "\tkey:", string(key), "\tsub:", string(sub))
		if pre != nil {
			readers = append(readers, _NotReader(pre))
		}

		if key != nil {
			r, e := genResfult(restfulHead, key)
			if e != nil {
				return nil, e
			}
			readers = append(readers, r)
		}
		if sub == nil {
			break
		}
		data = sub

	}
	return readers, nil
}
func parseRestfulDo(line []byte) (pre, key, sub []byte) {

	firstStart := bytes.Index(line, restfulSp)
	if firstStart == -1 {
		return line, nil, nil
	}
	firstEnd := bytes.IndexAny(line[firstStart+2:], restfulSpE)
	if firstEnd == -1 {
		pre = line[:firstStart+1]
		key = line[firstStart+2:]
		sub = nil
		return
	}

	firstEnd += firstStart + 2
	pre = line[:firstStart+1]
	key = line[firstStart+2 : firstEnd]

	if firstEnd+1 < len(line) {
		sub = line[firstEnd:]
	}
	return
}
