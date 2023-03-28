package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("samples\n")
	resp := calculateV1(buildMonkeysSample())
	fmt.Printf("sample 1 solution: %d\n", resp)
	if resp != 10605 {
		fmt.Printf("test nr 1 failed\n")
		return
	}

	resp = calculateV2(buildMonkeysSample())
	fmt.Printf("sample 2 solution: %d\n", resp)
	if resp != 2713310158 {
		fmt.Printf("test nr 2 failed\n")
		return
	}

	fmt.Printf("inputs\n")
	resp = calculateV1(buildMonkeysInput())
	fmt.Printf("input 1 solution: %d\n", resp)

	resp = calculateV2(buildMonkeysInput())
	fmt.Printf("input 2 solution: %d\n", resp)

}

type monkey struct {
	items     []int
	operation func(old int) int
	test      int
	testTrue  int
	testFalse int
	business  int
}

func (m *monkey) addItem(item int) {
	m.items = append(m.items, item)
}

func calculateV1(monkeys []*monkey) int {
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.business++
				item = m.operation(item) / 3
				if item%m.test == 0 {
					monkeys[m.testTrue].addItem(item)
				} else {
					monkeys[m.testFalse].addItem(item)
				}
			}
			m.items = []int{}
		}
	}

	mb := []int{}
	for _, m := range monkeys {
		mb = append(mb, m.business)
	}

	sort.Ints(mb)

	return mb[len(mb)-1] * mb[len(mb)-2]
}

// chinese remainder theorem magic
func calculateV2(monkeys []*monkey) int {
	tests := []int{}
	for _, m := range monkeys {
		tests = append(tests, m.test)
	}
	supermodulo := 1
	for _, t := range tests {
		supermodulo *= t
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				m.business++
				item = m.operation(item)
				item %= supermodulo
				if item%m.test == 0 {
					monkeys[m.testTrue].addItem(item)
				} else {
					monkeys[m.testFalse].addItem(item)
				}
			}
			m.items = []int{}
		}
	}

	mb := []int{}
	for _, m := range monkeys {
		fmt.Printf("business: %d\n", m.business)
		mb = append(mb, m.business)
	}

	sort.Ints(mb)

	return mb[len(mb)-1] * mb[len(mb)-2]
}

func buildMonkeysSample() []*monkey {
	monkeys := []*monkey{
		{
			items:     []int{79, 98},
			operation: func(old int) int { return old * 19 },
			test:      23,
			testTrue:  2,
			testFalse: 3,
		}, {
			items:     []int{54, 65, 75, 74},
			operation: func(old int) int { return old + 6 },
			test:      19,
			testTrue:  2,
			testFalse: 0,
		}, {
			items:     []int{79, 60, 97},
			operation: func(old int) int { return old * old },
			test:      13,
			testTrue:  1,
			testFalse: 3,
		}, {
			items:     []int{74},
			operation: func(old int) int { return old + 3 },
			test:      17,
			testTrue:  0,
			testFalse: 1,
		},
	}
	return monkeys
}

func buildMonkeysInput() []*monkey {
	monkeys := []*monkey{
		{
			items:     []int{52, 78, 79, 63, 51, 94},
			operation: func(old int) int { return old * 13 },
			test:      5,
			testTrue:  1,
			testFalse: 6,
		}, {
			items:     []int{77, 94, 70, 83, 53},
			operation: func(old int) int { return old + 3 },
			test:      7,
			testTrue:  5,
			testFalse: 3,
		}, {
			items:     []int{98, 50, 76},
			operation: func(old int) int { return old * old },
			test:      13,
			testTrue:  0,
			testFalse: 6,
		}, {
			items:     []int{92, 91, 61, 75, 99, 63, 84, 69},
			operation: func(old int) int { return old + 5 },
			test:      11,
			testTrue:  5,
			testFalse: 7,
		}, {
			items:     []int{51, 53, 83, 52},
			operation: func(old int) int { return old + 7 },
			test:      3,
			testTrue:  2,
			testFalse: 0,
		}, {
			items:     []int{76, 76},
			operation: func(old int) int { return old + 4 },
			test:      2,
			testTrue:  4,
			testFalse: 7,
		}, {
			items:     []int{75, 59, 93, 69, 76, 96, 65},
			operation: func(old int) int { return old * 19 },
			test:      17,
			testTrue:  1,
			testFalse: 3,
		}, {
			items:     []int{89},
			operation: func(old int) int { return old + 2 },
			test:      19,
			testTrue:  2,
			testFalse: 4,
		},
	}
	return monkeys
}
