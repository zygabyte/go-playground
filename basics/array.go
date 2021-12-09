package main

import "fmt"

func main() {
	a := [...]int{10, 40, 29}
	var b []int
	fmt.Println("a is")
	fmt.Println(a)
	if b == nil {
		fmt.Println("not initialized")
	}

	fmt.Println("____POINTERS____")
	var c [3]*string
	d := [...]*string{new(string), new(string), new(string)}
	*d[0] = "jane"
	*d[1] = "james"
	*d[2] = "grace"

	c = d

	fmt.Println("c is ")
	fmt.Println(c[2])

	*c[2] = "joy"

	fmt.Println("new d is ")
	fmt.Println(d[2])

	e := make([]string, 3, 5)
	f := []int{3, 4, 5, 3, 22, 54}
	g := f[1:4]
	h := f[1:4:4]

	fmt.Println("new e and f is ")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	g = append(g, 60)
	h = append(h, 90)
	fmt.Println(g)
	fmt.Printf("h is -> %v\n", h)
	fmt.Printf("f is -> %v\n", f)

	for i, v := range f {
		fmt.Printf("index %d and value %d.. address from range and slice %X %X \n", i, v, &v, &f[i])
	}

	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])
	}

	ii := [5]string{"first", "", "", "", ""}

	fmt.Printf("ii is %v \n", ii)

	aa := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	fmt.Printf("map is %v", aa)

}
