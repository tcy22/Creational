package main

import "bytes"

// builder设计模式，主要解决多变参数传递问题
// xorm就是使用了builder设计模式
// 故事： 平时去面馆吃面，有各种味道的面条（牛肉味、肥肠味等）
// 有各种配料（香菜、葱、姜、辣椒等）
// 第一个客人：一碗牛肉面 加葱、姜
// 第二个客人：一碗牛肉面 加葱、香菜


type MyBuilder struct {
	NoodlesType string // 面条口味
	Coriander   bool   // 是否加香菜
	Onion       bool   // 是否加葱
	Ginger      bool   // 是否加姜
	Pepper      bool   // 是否加辣椒
}


//构造一个结构体
func NewMyBuilder(noodsType string) MyBuilder {
	return MyBuilder{NoodlesType: noodsType}
}

//加香菜
func (myBuilder MyBuilder) withCoriander() MyBuilder{
	myBuilder.Coriander = true
	return myBuilder
}

// 加葱
func (myBuilder MyBuilder) withOnion() MyBuilder {
	myBuilder.Onion = true
	return myBuilder
}


// 加姜
func (myBuilder MyBuilder) withGinger() MyBuilder {
	myBuilder.Ginger = true
	return myBuilder
}


// 加辣椒
func (myBuilder MyBuilder) withPepper() MyBuilder {
	myBuilder.Pepper = true
	return myBuilder
}


func (myBuilder MyBuilder) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("要一碗" + myBuilder.NoodlesType + "面，")
	if myBuilder.Coriander {
		buffer.WriteString("加香菜、")
	} else {
		buffer.WriteString("不加香菜、")
	}


	if myBuilder.Onion {
		buffer.WriteString("加葱、")
	} else {
		buffer.WriteString("不加葱、")
	}


	if myBuilder.Ginger {
		buffer.WriteString("加姜、")
	} else {
		buffer.WriteString("不加姜、")
	}


	if myBuilder.Pepper {
		buffer.WriteString("加辣椒、")
	} else {
		buffer.WriteString("不加辣椒、")
	}
	buffer.WriteString("\n")

	return buffer.String()
}


func main(){
	myBuilder := NewMyBuilder("西红柿")

	myBuilder.withCoriander()
	myBuilder.withOnion()
}

