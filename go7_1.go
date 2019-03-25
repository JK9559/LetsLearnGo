package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
	Desc    string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []svs    `xml:"server"`
}

type svs struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func getXML() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Println("error: %v\n", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error: %v\n", err)
		return
	}
	v := Recurlyservers{}
	/*
	 关于Unmarshal函数：
	 两个参数 第一个是XML数据流，第二个是存储的数据类型，目前支持struct
	 slice和string 解析时 先读取struct tag，如果没有 就读取对应的字段名 注意大小写敏感
	 如果某个struct中的tag定义中含有 ",attr" 那就表示为该节点的属性 <a name="avs" age="20"> 的name和age

	*/
	err1 := xml.Unmarshal(data, &v)
	if err1 != nil {
		fmt.Println("error: %v\n", err1)
		return
	}

	fmt.Println(v.XMLName)
	fmt.Println("===========")
	fmt.Println(v.Svs)
	fmt.Println("===========")
	fmt.Println(v.Version)
	fmt.Println("===========")
	fmt.Println(v.Desc)
	fmt.Println("===========")
	fmt.Println(v.Svs[0].XMLName)
	fmt.Println(v.Svs[0].ServerIP)
	fmt.Println(v.Svs[0].ServerName)
	fmt.Println("===========")
	fmt.Println(v.Svs[1].XMLName)
	fmt.Println(v.Svs[1].ServerIP)
	fmt.Println(v.Svs[1].ServerName)
	fmt.Println("===========")
}

func outXML() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, svs{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, svs{"Beijing_VPN", "127.0.0.2"})
	out, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error: %v\n", err)
	}
	// 添加xml头
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(out)
}

func go7_1() {
	outXML()
}
