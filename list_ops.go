package listops

type (
	binFunc   func(a, b int) int
	predFunc  func(x int) bool
	unaryFunc func(x int) int
	// IntList represents a slice of ints.
	IntList []int
)

// Foldl returns the result of fn in an accumulator
func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, number := range l {
		initial = fn(initial, number)
	}
	return initial
}

// Foldr returns the result of fn in an accumulator. It starts from the right
func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := len(l) - 1; i >= 0; i-- {
		initial = fn(l[i], initial)
	}
	return initial
}

// Filter returns a filtered result based on the result of the fn function
func (l IntList) Filter(fn predFunc) IntList {
	out := IntList{}
	for _, number := range l {
		if fn(number) {
			out = append(out, number)
		}
	}
	return out
}

// Length returns the total number of elements in the list
func (l IntList) Length() int {
	i := 0
	for range l {
		i++
	}
	return i
}

// Map returns the list of the results if applying fn on all items
func (l IntList) Map(fn unaryFunc) IntList {
	for i, number := range l {
		l[i] = fn(number)
	}
	return l
}

// Reverse returns the original list with the reserved order
func (l IntList) Reverse() IntList {
	out := IntList{}
	for i := len(l) - 1; i >= 0; i-- {
		out = append(out, l[i])
	}
	return out
}

// Append add all items in the second list to the end of the first list
func (l IntList) Append(this IntList) IntList {
	out := l
	for _, number := range this {
		out = append(out, number)
	}
	return out
}

// Concat returns a flattened list that combines all the items in all the lists
func (l IntList) Concat(args []IntList) IntList {
	for _, number := range args {
		if number.Length() == 0 {
			continue
		}
		l = append(l, number...)
	}
	return l
}
