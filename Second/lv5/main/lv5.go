package main

import "fmt"

// 定义链表元素的数据
type conder struct {
	name          string
	school_number string
	phone_number  string
	code_line     int
	birth_date    string
}

// 定义链表元素
type Student_data struct {
	data  conder
	point *Student_data
}

var header_data Student_data

// 定义末端指针副本
var pointlast *Student_data

// 定义向链表末尾增加数据的函数
func adddata(newdata Student_data) {
	(*pointlast).point = &newdata
	pointlast = &newdata
}

// 定义遍历函数
func throughout() []Student_data {
	var result = []Student_data{}
	for point := header_data.point; (*point).point != nil; point = (*point).point {
		result = append(result, *point)
	}
	return result
}

// 定义按代码行数统计的函数
// symbol:可输入">""<""==",用以统计代码行数code_line大于、小于或等于形参codeline的数据
// 返回由数据构成的切片
func statistics(codeline int, symbol string) []Student_data {
	//fmt.Print("Statistics is running...\n")
	var result = []Student_data{}
	if symbol == ">" {
		for point := header_data.point; (*point).point != nil; point = (*point).point {
			//fmt.Printf("The type of point is%T\n", point)
			//fmt.Printf("The type of point = (*point).point is%T\n", (*point).point)
			if (*point).data.code_line > codeline {
				//fmt.Print("Statistics is running>...\n")
				//fmt.Printf("%v\n", *point)
				result = append(result, *point)
			}
		}
	} else if symbol == "<" {
		for point := header_data.point; (*point).point != nil; point = (*point).point {
			if (*point).data.code_line < codeline {
				result = append(result, *point)
			}
		}
	} else if symbol == "==" {
		for point := header_data.point; (*point).point != nil; point = (*point).point {
			if (*point).data.code_line == codeline {
				result = append(result, *point)
			}
		}
	} else {
		fmt.Printf("The symbol import symbol error!")
	}
	//fmt.Printf("Statistics is end...\n")
	return result
}

// 定义查找函数
// value:查找值 attribute:要查找的属性
func find(value interface{}, attribute string) []Student_data {
	//fmt.Print("Statistics is running...\n")
	var result = []Student_data{}
	if attribute == "name" {
		for point := header_data.point; (*point).point != nil; point = (*point).point {
			if (*point).data.name == value {
				result = append(result, *point)
			}
		}
	} else if attribute == "school_number" {
		for point := header_data.point; (*point).point != nil; point = (*point).point {
			if (*point).data.school_number == value {
				result = append(result, *point)
			}
		}
	}
	return result
}

// 定义插入函数
// 在学号为schoolnumber_before的元素前插入object元素
func insertb(object Student_data, schoolnumber_before string) []Student_data {
	//fmt.Print("Statistics is running...\n")
	var result = []Student_data{}
	var pointbefore *Student_data
	var pointlater *Student_data
	point := header_data.point
	for ((*point).data.school_number != schoolnumber_before) && ((*point).point != nil) {
		//fmt.Printf("(*point).data.school_number : \"%v\" %T\n", (*point).data.school_number, (*point).data.school_number)
		//fmt.Printf("schoolnumber_before : \"%v\" %T\n", schoolnumber_before, schoolnumber_before)
		pointbefore = point
		point = (*point).point
	}
	pointlater = (*pointbefore).point
	if pointbefore != pointlast {
		//fmt.Printf("Before %v\n", pointlater.data.name)
		pointlater = (*pointbefore).point
		(*pointbefore).point = &object
		object.point = pointlater
	} else {
		fmt.Printf("There is no %v schoolnumber!\n", schoolnumber_before)
	}
	return result
}

// 在学号为schoolnumber_before的元素后插入object元素
func insertl(object Student_data, schoolnumber_before string) []Student_data {
	//fmt.Print("Statistics is running...\n")
	var result = []Student_data{}
	var pointbefore *Student_data
	var pointlater *Student_data
	point := header_data.point
	for ((*point).data.school_number != schoolnumber_before) && ((*point).point != nil) {
		//fmt.Printf("(*point).data.school_number : \"%v\" %T\n", (*point).data.school_number, (*point).data.school_number)
		//fmt.Printf("schoolnumber_before : \"%v\" %T\n", schoolnumber_before, schoolnumber_before)
		pointbefore = point
		point = (*point).point
	}
	pointbefore = (*pointbefore).point
	pointlater = (*pointbefore).point
	if pointbefore != pointlast {
		//fmt.Printf("After %v\n", pointbefore.data.name)
		pointlater = (*pointbefore).point
		(*pointbefore).point = &object
		object.point = pointlater
	} else if pointbefore == pointlast {
		//保证pointlast是最后一个元素的指针
		fmt.Printf("After %v\n", pointbefore.data.name)
		adddata(object)
	} else {
		fmt.Printf("There is no %v schoolnumber!\n", schoolnumber_before)
	}
	return result
}

// 定义按学号删除元素的函数
func delete(schoolnumber string) {
	for point := header_data.point; (*point).point != nil; point = (*point).point {
		var pointdel *Student_data
		//var pointpre *Student_data
		if (*((*point).point)).data.school_number == schoolnumber {
			//pointpre = (*point).point
			//保证pointlast是最后一个元素的指针
			if (*point).point == pointlast {
				pointlast = point
			}
			pointdel = (*((*point).point)).point
			(*((*point).point)).point = new(Student_data)
			(*point).point = pointdel
		}
	}

}

func main() {

	//定义链表首个元素为空
	header_data = Student_data{
		data:  conder{},
		point: new(Student_data),
	}

	//初始化末端指针副本
	pointlast = &header_data

	//定义链表元素
	student1 := Student_data{
		data: conder{
			name:          "秦天毅",
			school_number: "2022210401",
			phone_number:  "13649795411",
			code_line:     5114,
			birth_date:    "2004-4-7",
		},
		point: new(Student_data),
	}

	student2 := Student_data{
		data: conder{
			name:          "吕科良",
			school_number: "2022210433",
			phone_number:  "13627756896",
			code_line:     5114,
			birth_date:    "2004-4-5",
		},
		point: new(Student_data),
	}

	student3 := Student_data{
		data: conder{
			name:          "唐香茗",
			school_number: "2022210444",
			phone_number:  "13627756789",
			code_line:     4781,
			birth_date:    "2003-9-9",
		},
		point: new(Student_data),
	}

	student4 := Student_data{
		data: conder{
			name:          "牟炫霖",
			school_number: "2022210412",
			phone_number:  "13568863462",
			code_line:     6510,
			birth_date:    "2003-2-3",
		},
		point: new(Student_data),
	}

	student5 := Student_data{
		data: conder{
			name:          "陆承涛",
			school_number: "2022210434",
			phone_number:  "13700851974",
			code_line:     6510,
			birth_date:    "2003-9-10",
		},
		point: new(Student_data),
	}

	//fmt.Printf("%T", *pointnow.point)

	student1.data.code_line = 5223

	//向链表末端添加元素
	adddata(student1)
	adddata(student2)
	adddata(student3)
	adddata(student5)

	//fmt.Printf("code_line>5000的学生有：%v\n", statistics(5114, "=="))
	//fmt.Printf("查找学号为\"2022210433\"的学生：%v\n", find("2022210433", "school_number"))
	fmt.Printf("Throughout 1 :%v\n\n", throughout())
	insertl(student4, "2022210444")
	fmt.Printf("Throughout 2 :%v\n\n", throughout())
	delete("2022210433")
	fmt.Printf("Del 吕科良 :%v\n\n", throughout())
}
