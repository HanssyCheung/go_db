package main

//学习结构体之前，先搞清楚类型定义和类型别名

//类型定义的语法
//type NewType Type

//相当于定义了一个全新的类型，但类型别名并没有定义一个新的类型，而是使用一个别名来替换
//存在于代码中，编译后就不存在了
//因为和原类型是一致的，所以原类型有的方法，类型别名中也可以调用，但是如果是重新定义的一个类型，不可以调用之前的任何方法
