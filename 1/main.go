package main

import (
	"encoding/json"
	"fmt"
    "net/http"
)

func main() {
    response, err := http.Get("https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json")
	if err != nil {
		fmt.Printf("Error Fetch data: %v", err)
        return 
	}
	defer response.Body.Close()
	var data [][]int
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&data)
	if err != nil {
        fmt.Printf("Error decoding JSON: %v", err)
        return
	}
    nextArr := data[len(data)-1]
    for i := len(data) - 2; i >= 0; i--{
        var Arr []int
        for i,item := range data[i]{
            if(item+nextArr[i]>item+nextArr[i+1]){
                Arr =  append(Arr,item+nextArr[i])
            }else{
                Arr = append(Arr,item+nextArr[i+1])
            }
        }
        nextArr = Arr

        
    }
    fmt.Println("result:",nextArr[0])
}
