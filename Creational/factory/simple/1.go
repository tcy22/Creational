package simple

import "fmt"

//产品，工厂，客户端
//简单工厂：根据条件产生不同功能的类
//工厂方法，抽象工厂

type Operater interface{
	Operate(int,int)int
}

type AddOperate struct{}

func (t *AddOperate)Operate(a int,b int)int {
	return a + b
}

type SubOperate struct{}

func (t *SubOperate)Operate(a int,b int)int {
	return a - b
}

type OperateFactory struct {}

func NewOperateFactory() *OperateFactory{
	return &OperateFactory{}
}

func (t *OperateFactory) CreateOperate(operatename string) Operater{
	switch operatename {
	case "+":
		return &AddOperate{}
	case "-":
		return &SubOperate{}
	default:
		panic("no this type")
		return nil
	}
}

func main(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	Operator := NewOperateFactory().CreateOperate("+")
	res := Operator.Operate(1,2)
	fmt.Println(res)
}

