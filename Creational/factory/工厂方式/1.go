package 工厂方法

//定义一个用于创建对象的接口，让子类决定实例化哪个类，使一个类的实例化延迟道其子类。

type Operation struct {
	a float64
	b float64
}

type OperationI interface {
	GetResult() float64
	setA(float64)
	setB(float64)
}

func (op *Operation) SetA(a float64){
	op.a = a
}

func (op *Operation) SetB(b float64){
	op.b = b
}


type AddOperation struct {
	Operation
}

func (t *AddOperation) GetResult() float64{
	return t.a + t.b
}



type SubOperation struct {
	Operation
}

func (t *SubOperation) GetResult() float64{
	return t.a - t.b
}

type IFactory interface {
	CreteOperation() Operation
}

type AddFactory struct {}

func(t *AddFactory) CreateOperation() OperationI {
	return &(AddOperation{})
}


/*func main(){
	fac := &(AddFactory{})
	oper := fac.CreateOperation()
	oper.setA(1)
	oper.setA(2)
	fmt.Println(oper.GetResult())
}*/

