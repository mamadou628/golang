package main

import (
	"fmt"
)

func main() {

	tableau := []string{"python", "Golang", "Java", "Javascript", "C", "C++", "C",
		"cours de mathétique", "cours de chimie", "cous d'histoire "}
	multiresearch(tableau, "cours de chimie", "\n", 1, "C++")

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
			fmt.Println("element found :\n ", listtrouve)
		} else {
			fmt.Println("The elements are found :")
			for _, val := range listtrouve {
				fmt.Println(val)
			}

		}
	}
}
