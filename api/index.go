package handler

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Result struct {
	Data struct {
		Skill []interface{} `json:"skill"`
	} `json:"data"`
}

type ErrorResult struct {
	Msg string `json:"msg"`
}

func (d Result) ToJson() []byte {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return b
}

func (d ErrorResult) ToJson() []byte {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return b
}

func AngelConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	angelConfig, err := http.Get("https://res.17roco.qq.com/conf/Angel.config")
	if err != nil {
		w.WriteHeader(500)
		errRes := ErrorResult{}
		errRes.Msg = err.Error()
		w.Write(errRes.ToJson())
		return
	}
	defer angelConfig.Body.Close()
	buf, err := ioutil.ReadAll(angelConfig.Body)
	if err != nil {
		w.WriteHeader(500)
		errRes := ErrorResult{}
		errRes.Msg = err.Error()
		w.Write(errRes.ToJson())
		return
	}
	// 删除前 7 个字节
	buf = buf[7:]
	// 解压
	result := DoZlibUnCompress(buf)
	result1 := matchSpiritSkillDes(string(result))
	w.WriteHeader(200)
	d := Result{}
	d.Data.Skill = matchSpiritSkill(result1)
	w.Write(d.ToJson())
}

func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

func matchSpiritSkillDes(str string) []string {
	reg1 := regexp.MustCompile(`<SpiritSkillDes.*/>`)
	if reg1 == nil {
		return nil
	}
	return reg1.FindAllString(str, -1)
}

func matchSpiritSkill(strarr []string) []interface{} {
	reg1 := regexp.MustCompile(`id="(\d+)"`)
	reg2 := regexp.MustCompile(`name="(.*?)"`)
	if reg1 == nil || reg2 == nil {
		return nil
	}
	var result []interface{}
	for _, v := range strarr {
		temp := make(map[string]string)
		temp["id"] = reg1.FindStringSubmatch(v)[1]
		temp["name"] = reg2.FindStringSubmatch(v)[1]
		result = append(result, temp)
	}
	return result
}
