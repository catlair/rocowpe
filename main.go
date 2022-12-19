package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func index() *gin.Engine {
	//静态文件路径
	const staticPath = `dist/`
	mime.AddExtensionType(".js", "application/javascript")
	r.Static("/assets", staticPath+"assets")
	r.StaticFile("/", staticPath+"index.html")
	r.StaticFile("/vue.svg", staticPath+"vue.svg")
	// 关键点【解决页面刷新404的问题】
	r.NoRoute(func(c *gin.Context) {
	})
	return r
}

func main() {
	index()

	r.Use(cors.Default())
	api := r.Group("/api")

	{
		api.GET("/angel/config", angelConfig)
	}

	r.Run()
}

func angelConfig(ctx *gin.Context) {
	// 请求 https://res.17roco.qq.com/conf/Angel.config
	angelConfig, err := http.Get("https://res.17roco.qq.com/conf/Angel.config")
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer angelConfig.Body.Close()
	buf, err := ioutil.ReadAll(angelConfig.Body)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 删除前 7 个字节
	buf = buf[7:]
	// 解压
	result := DoZlibUnCompress(buf)
	result1 := matchSpiritSkillDes(string(result))
	ctx.JSON(200, gin.H{
		"data": gin.H{
			"skill": matchSpiritSkill(result1),
		},
	})
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
