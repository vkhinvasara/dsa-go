package main
import "fmt"

func main(){
	result := contains_duplicate([]int{1,2,2,3,4,5,6,7});
	fmt.Print(result);
}

func contains_duplicate(nums []int) bool{
	hashmap := make(map[int]int);

	for _, num := range nums{
		_, exists := hashmap[num];
		if exists{
			return true;
		}
		hashmap[num]++;
	}
	return false;
}