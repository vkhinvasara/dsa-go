package main

import "fmt"

func main(){
	result := topKFrequent([]int{1,2,2,3,3,3}, 2)
	fmt.Print(result)
}

func topKFrequent(nums []int, k int) []int{
	count := make(map[int]int);
	freq := make([][]int, len(nums)+1);

	for _, n := range nums{
		count[n]++;
	}
	for i, v := range count{
		freq[v]= append(freq[v], i);
	}

	var res []int;

	for i:=len(freq)-1; i>=0; i--{
		for _, n := range freq[i]{
			res = append(res, n);
			if len(res) == k{
				return res;
			}
		}
	}
	return []int{};
}