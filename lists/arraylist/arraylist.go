package arraylist

type ArrayList struct {
	Size int
	List []int
}

// func init() *ArrayList {
// 	return &ArrayList
// }

func (a *ArrayList) Add(val int) int {
	if a.Size == 0 {
		arraylist := make([]int, 2)
		a.List = append(a.List, arraylist...)
	}
	if a.Size >= len(a.List) {
		arraylist := make([]int, a.Size*2)
		a.List = append(a.List, arraylist...)
	}
	a.List[a.Size] = val
	a.Size += 1
	return a.Size
}

func (a *ArrayList) Get(index int) int {
	return a.List[index]
}

func (a *ArrayList) GetFirst() int {
	return a.List[0]
}

func (a *ArrayList) GetLast() int {
	return a.List[a.Size-1]
}

func (a *ArrayList) Delete(index int) int {
	if (index >= a.Size) {
		panic("The index out of size of the arraylist")
	}
	for i:=index; i<a.Size-1; i++ {
		a.List[i] = a.List[i+1]
	}
	a.List[a.Size-1] = 0
	a.Size = a.Size-1

	return index
}