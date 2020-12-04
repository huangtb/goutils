package main

import (
	"fmt"
	"testing"
)


func TestpassByValue(t *testing.T)  {

	u := User{
		UserName:"zhangsan",
		Age:11,
	}

	fmt.Printf("值:%+v ,地址:%p \n",u,&u)
	passByValue(u)

}

func TestpassByPointer(t *testing.T)  {

	u := &User{
		UserName:"zhangsan",
		Age:11,
	}

	fmt.Printf("值:%+v ,地址:%p \n",u,&u)
	passByPoniter(u)

}

func TestpassByPointer2(t *testing.T)  {

	u := User{
		UserName:"zhangsan",
		Age:11,
	}

	fmt.Printf("值:%+v ,地址:%p \n",u,&u)
	passByPoniter(u)

}
