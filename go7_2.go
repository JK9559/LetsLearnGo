package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Sv struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Svs []Sv `json:"servers"`
}

func getJson2Struct() {
	var s Serverslice
	str := `{"svs":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	// 能被赋值的字段一定一定是可导出字段（即首字母大写）
	// 如果json的key为Foo，那么其查找对应字段过程如下
	// 1. 先查找tag含有Foo的 可导出的struct字段
	// 2. 在查找字段名是Foo的 可导出字段
	// 3. 最后查找类似 FOO FoO的除了首字母大写外，其他大小写不敏感的导出字段
	// 同时JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略，这样的一个好处是：
	// 当你接收到一个很大的JSON数据结构而你却只想获取其中的部分数据的时候，你只需将
	// 你想要的数据对应的字段名大写，即可轻松解决这个问题。
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s.Svs[0].ServerName)
	fmt.Println(s.Svs[0].ServerIP)
	fmt.Println("=======")
	fmt.Println(s.Svs[1].ServerName)
	fmt.Println(s.Svs[1].ServerIP)
}

func getJson2Inter() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	// b := []byte(`{"Name":"Wednesday","Age":6,"Parents":[{"aca":"Gomez"},{"abx":"Morticia"}]}`)
	//b := []byte(`{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func transfer2Json() {
	var s Serverslice
	s.Svs = append(s.Svs, Sv{"Shanghai", "127.0.0.1"})
	s.Svs = append(s.Svs, Sv{"Beijing", "127.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

type Dada struct {
	/*
			字段的tag是"-"，那么这个字段不会输出到JSON
			tag中带有自定义名称，那么这个自定义名称会出现在JSON的字段名中，例如上面例子中serverName
			tag中如果带有"omitempty"选项，那么如果该字段值为空，就不会输出到JSON串中
			如果字段类型是bool, string, int, int64等，而tag中带有",string"选项，
		    那么这个字段在输出到JSON的时候会把该字段对应的值转换成JSON字符串
	*/
	ID          int    `json:"-"`
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	ServerIP    string `json:"serverIP,omitempty"`
}

func testSomething() {
	s := Dada{
		ID:          3,
		ServerName:  `Go "1.1"`,
		ServerName2: `Go "1.0"`,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}

func go7_2() {
	testSomething()
}
