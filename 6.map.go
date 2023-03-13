package main

//map是一种key：value键值对的数据结构容器，map内部实现是哈希表（hash）
//map最重要的一点是通过key来快速检索，key类似于索引，指向数据的值
//map是引用类型的

//map的语法格式
//可以使用内建函数make也可以使用map关键字来定义map
/*
	声明变量，默认map是nil
	var map_variable map[key_data_type]value_data_type
	使用make函数
	map_variable = make(map[key_data_type]value_data_type
*/

func main() {
	m1 := make(map[string]string)

	m1["name"] = "tom"

	//遍历 一般使用for range循环进行map遍历 得到key和value的值

	m2 := map[string]string{
		"name": "kite",
		"age":  "20",
	}
	for s, s2 := range m2 {
		println(s, s2)
	}
	println(m1, m2, m1["name"])

}
