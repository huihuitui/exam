package main

func work() {
	//找不同批次：
	//维护多个连续数组，新批次如果不在里面就自成一个连续数组，在里面直接结果加1
	//这样可以快速判断是否是新批次，只用维护数组的边界
	//新船的结果等于   上一搜船的结果+自己的批次结果 - 上一批多包含的批次结果
	//  OVO!
	// cpu烧了 就这么多了
}
