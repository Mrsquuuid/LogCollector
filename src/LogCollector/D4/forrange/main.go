package main

import "fmt"

type student struct {
	Name string
	Age int
}

func pase_student() {
	//m := make(map[string]*student, 16)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 从stus中依次取出学生放到m中
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}
	//
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.Name)
	//}

	for _, stu := range stus {
		stu.Age = stu.Age+10
	}

	for _, stu := range stus{
		fmt.Println(stu.Age)
	}

}
func main(){
	pase_student()
}
