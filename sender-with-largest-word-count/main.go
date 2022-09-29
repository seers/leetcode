package main

import (
	"fmt"
	"strings"
)

func main() {
	messages := []string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"}
	senders := []string{"Alice", "userTwo", "userThree", "Alice"}
	fmt.Println(largestWordCount(messages, senders))
	//Alice

	messages = []string{"How is leetcode for everyone", "Leetcode is useful for practice"}
	senders = []string{"Bob", "Charlie"}
	fmt.Println(largestWordCount(messages, senders))
	//Charlie
}

// 构造姓名，发言次数map
func largestWordCount(messages []string, senders []string) string {
	namecount := make(map[string]int)
	for k, v := range messages {
		namecount[senders[k]] += len(strings.Fields(v))
	}
	fmt.Print(namecount)
	return ""
}
