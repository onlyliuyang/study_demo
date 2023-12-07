package main

type set struct {
	elements map[interface{}]bool
}

type Set interface {
	Add(item interface{})
	Del(item interface{})
	Len() int
	GetItems() []interface{}
	In(item interface{}) bool
}

func New(items ...interface{}) Set {
	st := set{
		elements: make(map[interface{}]bool),
	}
	for _, i := range items {
		st.Add(i)
	}
	return &st
}

func (s *set) Add(item interface{}) {
	s.elements[item] = true
}

func (s *set) Del(item interface{}) {
	delete(s.elements, item)
}

func (s *set) Len() int {
	return len(s.elements)
}

func (s *set) GetItems() []interface{} {
	keys := make([]interface{}, 0)
	for k, _ := range s.elements {
		keys = append(keys, k)
	}
	return keys
}

func (s *set) In(item interface{}) bool {
	_, ok := s.elements[item]
	return ok
}

//func main() {
//	mySet := New(1, 2, 3, 4, 5, 6, 7)
//	fmt.Println(mySet.GetItems())
//	mySet.Add(9)
//	mySet.Add(10)
//	fmt.Println(mySet.GetItems())
//	fmt.Println(mySet.Len())
//	mySet.Del(3)
//	fmt.Println(mySet.Len())
//	fmt.Println(mySet.In(20))
//	fmt.Println(mySet.In(1))
//}

func main() {
	//set := []int{0, 2, 4, 6}
	//list := []string{}
	//for i := 0; i < len(set); i++ {
	//	if set[i] == 0 {
	//		continue
	//	}
	//	for j := 0; j < len(set); j++ {
	//		list = append(list, strconv.Itoa(set[i])+strconv.Itoa(set[j]))
	//	}
	//}
	//fmt.Println(list)
	//fmt.Println(len(list))
}
