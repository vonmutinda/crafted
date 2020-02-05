// package main

// import (  
// 	"github.com/vonmutinda/crafted/api" 
	
// )

// func main(){
// 	api.Run()
// }

package main

import (
"fmt"
)

func main(){
	numbers := []int{9090,}
	for i := 9991; i <10010; i++{
		numbers = append(numbers,i)
	}
	fmt.Println(numbers)
}