package main

import (
	"fmt"
	"sort"
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

type sc struct {
	sender string
	count  int
}
type scs []sc

func (s scs) Len() int {
	return len(s)
}

func (s scs) Less(i, j int) bool {
	if s[i].count == s[j].count {
		return s[i].sender < s[j].sender
	}
	return s[i].count < s[j].count
}

func (s scs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 构造姓名，发言次数map
func largestWordCount(messages []string, senders []string) string {
	sendercount := make(map[string]int)
	for k, v := range messages {
		sendercount[senders[k]] += len(strings.Fields(v))
	}
	scs := scs{}
	for k, v := range sendercount {
		scs = append(scs, sc{sender: k, count: v})
	}

	sort.Sort(scs)
	return scs[len(scs)-1].sender

}
