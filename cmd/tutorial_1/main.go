package main
import (
    "fmt"
    "unicode/utf8"
    "errors"
)

func main() {
//    var inputVar string = "Bien truyen vao"
   var intArr [3]int = [3]int{0,1,2}
   intArr[2] = 111
   fmt.Println(intArr[1])
   fmt.Println(intArr[2])
   fmt.Println(intArr[0])

   var make []int = make([]int, 5)
      fmt.Println(make);

   var myMap2 = map[string]int{"Adam" : 12, "Levi" : 36}
      fmt.Println(myMap2["Adam"])

   for  name := range myMap2 {
      fmt.Printf("Name %v", name)
   }
   for  i,v := range myMap2 {
         fmt.Printf("Index: %v, Value: %v\n", i, v)
      }
    var string1 = "Phương Duyên"
    var indexed = string1[0]
    fmt.Println(indexed)

     var strSlice = []string{"D", "i", "n", "l", "o"}
     var catStr = ""
     for i := range strSlice {
         fmt.Printf("Index: %v\n", i)
           catStr += strSlice[i]
     }
     fmt.Printf(catStr)


//    run(inputVar)
}

func run(inputVar string) {
   fmt.Println("Hello World!");
    fmt.Println(inputVar);
      var intNum int = 1246
      intNum = intNum + 1
      fmt.Println(intNum)

      var floatNum float64 = 1234.5
      var intNum2 int = int(floatNum) + intNum
       fmt.Println(intNum2)

       var myString string = "Dinlo"
       fmt.Println(myString)
       fmt.Println(utf8.RuneCountInString(myString));

      varName := "Day la 1 cach khai bao bien khac"
          fmt.Println(varName)

      const myConst string = "Luon phai khai bao"
         fmt.Println(myConst)

    var var1 int
    var var2 int = 200
    var cong, chia, err = hasReturn(var1, var2)
    if err!=nil {
        fmt.Printf(err.Error())
    }else {
        fmt.Printf("cong = %v chia = %v", cong, chia)
    }

 }


func hasReturn(var1 int, var2 int)(int, int, error) {
    var err error
    if var1 == 0 {
        err = errors.New("Khong chia duoc cho 0")
        return 0, 0, err
    }
    var cong int = var1 + var2
    var chia int = var2 / var1
    return cong, chia, err
}