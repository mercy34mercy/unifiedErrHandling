package a

import "fmt"

func ResErr(m string) error{
	return fmt.Errorf("%s",m)
}

func ResErrAndStr(m string)(string,error){
	return m, fmt.Errorf("%s",m)
}

func errPattern1() {
	if err := ResErr("test"); err != nil { //want "abbreviated notation"
		fmt.Printf("%v",err)
	}
}

func okPattern1() {
	err := ResErr("test")
	if err != nil { //want "Separated notation"
		fmt.Printf("%v",err)
	}
}

func errPattern2() {
	if _,err := ResErrAndStr("test"); err != nil { //want "abbreviated notation"
		fmt.Printf("%v",err)
	}
}


func errPattern3() {
	if _,hogehoge := ResErrAndStr("test"); hogehoge != nil { //want "abbreviated notation"
		fmt.Printf("%v",hogehoge)
	}
}

func errPattern4() {
	if hogehoge := ResErr("test"); hogehoge != nil { //want "abbreviated notation"
		fmt.Printf("%v",hogehoge)
	}
}

func ok1() {
	err := 1
	if err != 0{
		fmt.Println(err)
	}
}

func ok2() {
	if err := 0; err != 1{
		fmt.Println(err)
	}
}

