package main

import "fmt"

func main() {
	// roles := make(map[bool]string, 3)
	
	// roles[true] = "Admin"
	// roles[false] = "User"
	// roles[true] = "owner"

	// roles[true] = "Akram"

	// fmt.Println(roles[true])

	// for {
	// 	fmt.Println("111")
	// 	break
	// }

    // isLoop := true
	// 	for isLoop {
	// 		fmt.Println("222")
	// 		isLoop = false
	// 	}

		
	// 	for i :=0; i< 5; i++ {
	// 		if i == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("iter:", i)

	// 	}

	sl := []string{"b","j","t"}
	idx := 0

		for idx < len(sl){
			fmt.Println("Slice index is", idx, "value is", sl[idx])
			idx++
		}


		for v, i := range sl {
			fmt.Println("Slice index is", i, "and value is", v)
		}


		str := "Hmmm 6"

		for t,f := range str {
			fmt.Println("Test t", t, "test f", f)
		}
	 


nums :=[]int{44,45,66,54}
func sum(in ...int) (result int){
	fmt.Println("in", in)
	for _,val := range in {
		result +=val
	}
	return
}


fmt.Println()




}
