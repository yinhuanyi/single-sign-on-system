/**
 * @Author：Robby
 * @Date：2022/3/1 15:13
 * @Function：
 **/

package main

import "fmt"

func main() {
	a := []string { "1", "2" }
	for i := range a {
		fmt.Println(i)
	}
}
