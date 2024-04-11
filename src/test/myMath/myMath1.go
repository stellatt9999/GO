package mathClass
import "fmt"

func Add(x,y int) int {
	return x + y
}

//打印九九乘法表
func printMultiTable() {
	var n, m int
	for n = 1; n <= 9; n++ {
		for m = 1; m <= n; m++ {
			fmt.Printf("%dx%d=%d ", m, n, m*n)
		}
		fmt.Println("")
	}
}