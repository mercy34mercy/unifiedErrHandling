package b

import "fmt"

func ResErr(m string) error{
	return fmt.Errorf("%s",m)
}

func ResErrAndStr(m string)(string,error){
	return m, fmt.Errorf("%s",m)
}

func errPattern1() {
	if err := ResErr("test"); err != nil {
		fmt.Printf("%v",err)
	}
}

func okPattern1() {
	err := ResErr("test")
	if err != nil { //want "separated notation"
		fmt.Printf("%v",err)
	}
}

func errPattern2() {
	if _,err := ResErrAndStr("test"); err != nil {
		fmt.Printf("%v",err)
	}
}