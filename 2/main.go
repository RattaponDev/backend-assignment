package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	firstInput := input
	input = strings.ReplaceAll(input, "=", "")
	result := []int{}
	if(string(input[0])=="R" ||string(input[0])=="="){
		result = append(result,0)
	}else{
		count := 0
		for _, char := range input {
			if(string(char) == "R"){
				break
			}
			count++
		}
		result = append(result,count)
	}
	
	for i, char := range input {
		num := result[len(result)-1]
			switch char {
			case 'L':
				if(i+1<len(input)&&string(input[i+1])=="L"){
						result = append(result,num-1)
				}else{
					result = append(result,0)
				}
				
			case 'R':
				if(i+1<len(input)&&string(input[i+1])=="L"){
						count := 0
						for j := i+1 ;j<len(input) ;j++ {
							if(string(input[j]) == "R" || string(input[j]) == "="){
								break
							}
							count++
						}
						if(count==num){
							result = append(result,num+1)
						}else{
							result = append(result,count)
						}
				}else{
					result = append(result,num+1)
				}
			default:
				result = append(result,num)
			}
        
    }
	LastResult :=""
	for i,char := range firstInput{
		if(string(char)=="="){
			result = append(result[:i+1], result[i:]...) 
			result[i] = result[i]
		}
	}
	for _,item := range result{
		LastResult += strconv.Itoa(item) 
	}
	fmt.Println("result: ",LastResult) 
}
