package main

type Person struct {
	name     string
	age      int
	flg      bool
	interest []string
}

//构造方法
func NewPerson(name string, age int, flg bool, interest []string) *Person {
	return &Person{
		name:     name,
		age:      age,
		flg:      flg,
		interest: interest,
	}
}
func main() {
	p := NewPerson("张三", 20, true, []string{"篮球", "足球"})
	println(p.name)
}
