package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "aabcb"
	fmt.Println(beautySum(s))

	// s = "aabcbaa"
	// fmt.Println(beautySum(s))

	// s = "njuvwehovhmnwfypqqzzguivrvmfuxdsbebskjfwzmjjqrrqwgdkibdkjhrrmjdehakrmjqtwwtqmyfzkqvyzbjggtmuaadmbwqynqosylguacktujyvxsoufzmltvttdupsuujbwrnfbbmqwyshsiytguxfuipcdutjoaracfxqogfsdhaczietafvkdlzzsyxycbemyedyrwkjzvudzgngixyftbldeyoadukxwsazunrrwafyvayfrbeqkzztdqnrgrgxmfkxxprzvshbsotcqffoteucenrnfxlphzrpllvuidrihuluvinpt"
	// fmt.Println(beautySum(s))
}

// 用切片保存，排序后最大值减去最小值
func beautySum(s string) int {
	substrs := make(chan string)

	go func() {
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				substrs <- s[i:j]
			}
		}
		close(substrs)
	}()

	sum := 0
	for substr := range substrs {
		fmt.Println(substr)
		arr := []int{}
		for _, strrune := range substr {
			arr = append(arr, strings.Count(substr, string(strrune)))
		}
		sort.Ints(arr)
		sum += arr[len(arr)-1] - arr[0]
	}

	return sum
}
