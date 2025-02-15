package main

import "fmt"

func main() {

	tableau := []string{"python", "Golang", "Java", "Javascript", "C", "C++", "C",
		"cours de mathétique", "cours de chimie", "cous d'histoire "}
	multiresearch(tableau, "\n", 1, "Mamadou")

}

func multiresearch(slice []string, para ...interface{}) {
	listtrouve := []string{}
	for _, val := range slice {

		for _, p := range para {
			if val == p {
				listtrouve = append(listtrouve, val)
			}
		}
	}
	// if listtrouve != nil {
	// 	if len(listtrouve) == 1 {
	// 		fmt.Println("element not found :\n ", listtrouve)
	// 	} else {
	// 		fmt.Println("The elements are found :")
	// 		for _, val := range listtrouve {
	// 			fmt.Println(val)
	// 		}

	// 	}
	// }
	if listtrouve != nil {
		fmt.Println("The elements are found :")
		for _, val := range listtrouve {
			fmt.Println(val)
		}
	} else {
		fmt.Println("The element are not found !")
	}

}
