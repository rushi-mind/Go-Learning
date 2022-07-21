/*
	Problem: 		Max Words Count
	Date: 			30-06-2022
	Last Modified: 	30-06-2022
	Author: 		Rushiraj Parekh
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	str := "Aliquam fringilla nunc et placerat auctor. Vivamus rhoncus nulla vel lacus placerat, id tincidunt dolor mattis. Etiam ex elit, posuere a venenatis vel, vulputate ut orci. In ligula ex, finibus a augue non, laoreet accumsan magna. Praesent ac mollis elit. Curabitur tincidunt risus elementum velit interdum vestibulum. Donec libero tellus, pulvinar ut turpis rhoncus, sollicitudin vestibulum ipsum. Nulla eu quam nisi. Morbi placerat congue nibh, in tincidunt tellus ullamcorper sed. Aenean a est efficitur, suscipit enim eu, ornare nunc. Etiam malesuada tortor eu posuere vehicula. In suscipit aliquet urna a ornare. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Morbi molestie nunc felis, a congue justo lobortis sit amet. Vivamus varius arcu neque. Nam eu magna lobortis, ultrices libero a, pellentesque tortor. Donec massa dolor, rhoncus ac mi sit amet, eleifend pulvinar velit. Aliquam dignissim elit ligula, a pellentesque diam dapibus vitae. Mauris posuere lobortis dui ornare porttitor. Pellentesque elit turpis, posuere ac tellus id, tempor facilisis nulla. Phasellus gravida tempus sapien et consectetur. Phasellus tempus in ante quis pellentesque. Nunc tincidunt, libero dapibus pulvinar efficitur, ex risus varius elit, in vehicula mi lacus quis sapien. Cras eros massa, vulputate in gravida sit amet, vestibulum quis tellus. Integer efficitur pharetra mauris nec dignissim.  Nulla eget odio justo. Suspendisse egestas turpis venenatis arcu convallis sagittis. Morbi a eros facilisis, accumsan magna nec, venenatis dui. Praesent placerat sagittis vulputate. Suspendisse at elementum ipsum. Phasellus venenatis risus ipsum, quis venenatis tortor egestas at. Suspendisse sit amet massa auctor, pulvinar est ac, suscipit nisl. Etiam interdum, libero in consequat mollis, eros est fermentum neque, vel tempor elit justo eu quam. Nullam quis efficitur velit, nec dignissim"

	temp := strings.Split(str, " ") // splitting by space and creating array of every words
	_map := make(map[string]int)    // defining empty map => string-int

	// creating map with different words count
	for _, k := range temp {
		// removing blank strings
		if k == "" {
			continue
		}
		// trimming string from both side with punctuations
		k = strings.TrimFunc(k, func(r rune) bool {
			return unicode.IsPunct(r)
		})
		_map[k]++
	}

	words := []string{} // defining empty array of string
	for key := range _map {
		words = append(words, key)
	}

	// sorting words array by count in descending order
	sort.SliceStable(words, func(i, j int) bool {
		return _map[words[i]] > _map[words[j]]
	})

	// Printing top 10 results
	for i, k := range words {
		fmt.Printf("%v =>  %v\n", _map[k], k)
		if i >= 9 {
			break
		}
	}
}
