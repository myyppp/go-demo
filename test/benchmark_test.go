package test_test

import (
	"fmt"
	"strconv"
	"testing"
)

// 基准测试以 Benchmark 开头
// go test -benchmem -run=^$ -bench ^BenchmarkSprintf$ demo/test
// -v 详细信息
// -run 接受正则表达式，指定为 ^$ 排除单元测试
// -bench 接受正则表达式
// -benchtime=3s 更改测试执行的最短时间
// -benchmem 提供每次操作分配内存的次数，以及总共分配内存的字节数

// Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkSprintf$ demo/test

// goos: windows
// goarch: amd64
// pkg: demo/test
// cpu: Intel(R) Core(TM) i5-9600KF CPU @ 3.70GHz
// BenchmarkSprintf
// BenchmarkSprintf-6
// 22067950                53.95 ns/op            2 B/op          1 allocs/op
// PASS
// ok      demo/test       1.337s

// 22067950（循环中代码被执行的次数）  53.95 ns/op（每次操作消耗的纳秒数）  2 B/op（每次操作分配的字节数）  1 allocs/op（每次操作从堆上分配内存的次数）
// PASS
// ok      demo/test（被执行的代码文件的名字）   1.337s（基准测试总共消耗的时间）

func BenchmarkSprintf(b *testing.B) {
	number := 10

	// 重置计时器
	// 保证在测试代码前的初始化代码，不会干扰计时器的结果
	b.ResetTimer()

	// 将所有基准测试的代码都放入循环里，并且循环使用b.N的值
	// 默认持续1s内，反复测试
	// 每次调用测试函数时，都会增加b.N的值，第一次调用时，b.N的值为1
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat provides performance numbers for the
// strconv.FormatInt function.
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa provides performance numbers for the
// strconv.Itoa function.
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
