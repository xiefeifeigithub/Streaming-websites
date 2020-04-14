package main

import (
	"fmt"
	"testing"
)

/*
command: go test -v
1. test文件下的每一个test case均必须以Test开头并且符合TestXxx形式，
否则go test会直接跳过测试不执行。
2. test case的入参为 t *testing.T(测试对错)或b *testing.B(测试性能)。
3. t.Errorf为打印错误信息，并且当前test case会被跳过。
4. t.SkipNow()为跳过当前test，并且直接按PASS处理继续下一个test。
5. Go的test不会保证多个TestXxx是顺序执行，但是通常会按顺序执行。
6. 使用t.Run来执行subtests可以做到控制test输出以及test的顺序。
7. 使用TestMain作为初始化test，并且使用m.Run()来调用其他tests可以完成
一些需要初始化操作的testing，比如数据库连接，文件打开，REST服务登录等
7.1 注：如果没有在TestMain中调用m.Run()则除了TestMain以外的其他tests都不会被执行。
*/
func testPrint(t *testing.T) {
	//t.SkipNow()
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("testPrint",testPrint) // 使用subtests来控制test执行的顺序
	t.Run("testPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins...")
	m.Run()
}

/*
Test之Benchmark
command: go test -bench=.
1. benchmark函数一般以Benchmark开头。
2. benchmark的case一般会跑b.N次，而且每次执行都会如此。
3. 在执行过程中会根据实际case的执行时间是否稳定会增加b.N的次数以达到稳态。
 */

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}

