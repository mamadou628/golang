package main

import (
	"fmt"
)

func main() {

	tableau := []string{"python", "Golang", "Java", "Javascript", "C", "C++", ""}
	multiresearch(tableau, "Golang", 1, "C++")

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
	if listtrouve != nil {
		if len(listtrouve) == 1 {
			fmt.Println("element found : ", listtrouve)
		} else {
			fmt.Println("The elements found : \n", listtrouve, "\n")

		}
	}
}
