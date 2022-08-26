package main

import "fmt"

func main() {
	animals := make(map[string]string)
	animals["cat"] = "🐈"
	animals["dog"] = "🐕"

	fmt.Println(animals)//map[cat:🐈 dog:🐕]

	fruits := map[string]string{
		"apple":  "🍎",
		"banana": "🍌",
	}
	fmt.Println(fruits)//ap[apple:🍎 banana:🍌]
	delete(fruits, "banana")
	fmt.Println(fruits)//map[apple:🍎]

	if _, ok := animals["gorilla"]; !ok {
		animals["gorilla"] = "🦍"
	}

	fmt.Println(animals)//map[cat:🐈 dog:🐕 gorilla:🦍]

}
