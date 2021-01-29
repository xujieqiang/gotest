package main

import (
	std "fmt"
	"math/rand"
	_ "strings"


	//"github.com/gin-gonic/gin"

)



type sear struct{
	name string
	age int

}

var (
	ab =20
	bc =30
)





func main(){
	const (
		a1=97
		b1="abc"
	)

	/*
	fmt.Println("ddddda")
	engin:=gin.Default()
	engin.Handle("GET","/",func(c *gin.Context){
		c.Writer.Write([]byte("dfdfsdf"))
	})
	engin.Run(":85")
	*/
	var a int
	var b string
	var mp sear
	var cha int

	a=123
	b="lucy is my name"
	std.Println(a,"==",b)
	var stu sear
	stu.name="jack"
	stu.age = 192
	q:=sear{"zhangsan",35}

	var r int=rand.Intn(6)
	std.Println(stu,"-->",q)
	
	var tag bool
	tag=true
	std.Println(tag,"-->",r)

	mp,cha=Older(stu,q)
	std.Println(mp,"-->",cha)

	al:=123
	ar:=456

	std.Printf("%T,%b\n",al,ar)
	std.Printf("%o,%50d\n",al,ar)
	std.Println(ab+bc)
	std.Println(string(a1)+b1)
}

func Older(p1,p2 sear)(sear , int){
	if p1.age>p2.age {
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

