package main

import "fmt"
func main(){
	var workArray [10]uint8
var a,b uint8
for i:=0;i<10;i++{
    fmt.Scan(&a)
    workArray[i]=a
}
for i:=0;i<3;i++{
    fmt.Scan(&a)
    fmt.Scan(&b)
    workArray[a],workArray[b]=workArray[b],workArray[a]
}
for i:=0;i<10;i++{
    fmt.Printf("%d ",workArray[i])
}
}