## pesudocode
#Execrise 1:
a function called addOne that receives the address of a number, goes to that address, and adds 1 to the real value -so the change sticks.

function addOne(num *int) {
    go to  the address
    get what is there add 1 to it 
    it back
}
    main:{
    creat x = 100
    call addOne and pass &x
    print x 
    }
```go
package main

import "fmt"

func addOne(num *int) {
    *num = *num + 1
}
func main() {
    x := 100
    addOne(&x)
    fmt.Println(x)
}
```

# pesudocode 2

function double(num *int){
    go to the address 
    get what is there
    multiply it by 2
    put it back
}
main{
    create x = 50
    call doulble and pass &x
    print x
}
```go
package main

import "fmt"

func double(num *int) {
    *num = *num * 2

}
func main() {
    x := 50
    double(&x)
    fmt.Println(x)
}
```

