package c


import "fmt"

func ResErr(m string) error{
	return fmt.Errorf("%s",m)
}

func ResErrAndStr(m string)(string,error){
	return m, fmt.Errorf("%s",m)
}

func errPattern1() {
	if err := ResErr("test"); err != nil {//want "abbreviated notation"
		fmt.Printf("%v",err)
	}
}

func okPattern1() {
	err := ResErr("test")
	if err != nil {
		fmt.Printf("%v",err)
	}
}

func okPattern2() {
	err := ResErr("test")
	if err != nil {
		fmt.Printf("%v",err)
	}
}

func okPattern3() {
	err := ResErr("test")
	if err != nil {
		fmt.Printf("%v",err)
	}
}

func errPattern2() {
	_,err := ResErrAndStr("test")
	if err != nil {
		fmt.Printf("%v",err)
	}
}
