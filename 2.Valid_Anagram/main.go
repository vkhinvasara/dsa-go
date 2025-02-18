package main
import "fmt"
func main(){
	result := is_anagram("ABBC","CBBB");
	fmt.Print(result)
}

func is_anagram(s1 string, s2 string)bool{
	if len(s1) != len(s2){
		return false
	}
	s1_map, s2_map := make(map[rune]int), make(map[rune]int)
	
	for i,value := range s1{
		s1_map[value]++
		s2_map[rune(s2[i])]++
	}
	
	for char, count := range s1_map{
		if s2_map[char] != count{
			return false
		}
	}
	return true
}