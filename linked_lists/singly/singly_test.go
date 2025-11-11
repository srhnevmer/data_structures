package singly

const template = "Test case number:"

// func TestTraversal(t *testing.T) {
// 	testCases := []struct {
// 		list
// 		want []int
// 	}{
// 		{list: fillList(), want: make([]int, 0)},
// 		{list: fillList(10), want: []int{10}},
// 		{list: fillList(10, 20, 30), want: []int{10, 20, 30}},
// 	}

// 	for i, tc := range testCases {
// 		t.Run(fmt.Sprintf("%s%d", template, i+1), func(t *testing.T) {
// 			if got := tc.list.traversal(); slices.Compare(tc.want, got) != 0 {
// 				t.Errorf("Expected result: %v got %v", tc.want, got)
// 			}
// 		})
// 	}
// }
