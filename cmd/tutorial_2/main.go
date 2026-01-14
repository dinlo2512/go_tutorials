package main
import (
    "fmt"
)

type gasEngine struct{
    msg int
    gallons int
}

type electricEngine struct {
    kmph int
    kwph int
}
type engine interface {
    engineLeft() int
}

func (e gasEngine) engineLeft() int {
    return e.msg * e.gallons
}

func (e electricEngine) engineLeft() int {
    return e.kmph * e.kwph
}

func canMakeIt(myGasEngine engine)string {
    if myGasEngine.engineLeft() > 200 {
        return "Can make it";
    }else {
        return "Can not make it"
    }
}

//Khai báo kiểu dữ liệu
func main() {
    var myGasEngine gasEngine = gasEngine{msg: 23, gallons: 42}
    fmt.Println(myGasEngine.msg, myGasEngine.gallons)
    fmt.Println(canMakeIt(myGasEngine))

    //Pointer
    var p *int = new(int)
    var i int

    fmt.Printf("\np = %v", p)
    fmt.Printf("\np pointer = %v", *p)

    // gán giá trị i vào con trỏ p
    p = &i
    *p = 1
    fmt.Printf("\np = %v", p)
    fmt.Printf("\ni = %v", i)

}