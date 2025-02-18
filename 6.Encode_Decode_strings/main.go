package main

import (
	"fmt"
	"strconv"
)

func main(){
	split_string := []string{"Hello", "neetcode", "love", "5you"};
	encoded_string := encode(split_string);
	decoded_string := decode(encoded_string);

	fmt.Print(decoded_string);

}

func encode(strings []string)string{
	var res string;
	for _, s := range strings{
		res = res + strconv.Itoa(len(s)) + "$" +s;
	}
	return res;
}

func decode(s string)[]string{
	var res []string;
	i:= 0;
	for i < len(s){
		j := i;
		for s[j] != '$'{
			j++;
		}
		length, _ := strconv.Atoi(s[i:j]);
		res = append(res, s[j+1:j+1+length]);
		i = j+1+length;
	}
	return res;
}