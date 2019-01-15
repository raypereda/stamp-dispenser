package stamps_test

import (
	"fmt"

	"github.com/raypereda/stamps"
)

func Example_dispense1() {
	stampDenominations := []int{90, 30, 24, 10, 6, 2, 1}
	d, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}
	stamps, err := d.Dispense1(18)
	if err != nil {
		fmt.Println("Can't give exacts stamps given the demoninations")
		return
	}
	fmt.Println("d =", d)
	fmt.Println("d.Dispense(18) = ", stamps)

	for _, p := range stampDenominations {
		s, err := d.Dispense1(p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("d.Dispense(%d) = %d\n", p, s)
	}

	// Output: d = &{[90 30 24 10 6 2 1] map[] 0}
	// d.Dispense(18) =  3
	// d.Dispense(90) = 1
	// d.Dispense(30) = 1
	// d.Dispense(24) = 1
	// d.Dispense(10) = 1
	// d.Dispense(6) = 1
	// d.Dispense(2) = 1
	// d.Dispense(1) = 1
}

func Example_dispense2() {
	stampDenominations := []int{90, 30, 24, 10, 6, 2, 1}
	d, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}

	// for _, p := range stampDenominations {
	for p := 1; p < 200; p++ {
		s1, err := d.Dispense1(p)
		if err != nil {
			fmt.Println(err)
		}
		s2, err := d.Dispense2(p)
		if err != nil {
			fmt.Println(err)
		}
		if s1 != s2 {
			fmt.Printf("d.Dispense1(%d) = %d but d.Dispense2(%d) = %d\n", p, s1, p, s2)
		}
	}
	// Output: d.Dispense1(34) = 3 but d.Dispense2(34) = 2
	// d.Dispense1(35) = 4 but d.Dispense2(35) = 3
	// d.Dispense1(44) = 4 but d.Dispense2(44) = 3
	// d.Dispense1(45) = 5 but d.Dispense2(45) = 4
	// d.Dispense1(48) = 4 but d.Dispense2(48) = 2
	// d.Dispense1(49) = 5 but d.Dispense2(49) = 3
	// d.Dispense1(58) = 4 but d.Dispense2(58) = 3
	// d.Dispense1(59) = 5 but d.Dispense2(59) = 4
	// d.Dispense1(64) = 4 but d.Dispense2(64) = 3
	// d.Dispense1(65) = 5 but d.Dispense2(65) = 4
	// d.Dispense1(72) = 4 but d.Dispense2(72) = 3
	// d.Dispense1(73) = 5 but d.Dispense2(73) = 4
	// d.Dispense1(74) = 5 but d.Dispense2(74) = 4
	// d.Dispense1(75) = 6 but d.Dispense2(75) = 5
	// d.Dispense1(78) = 5 but d.Dispense2(78) = 3
	// d.Dispense1(79) = 6 but d.Dispense2(79) = 4
	// d.Dispense1(82) = 5 but d.Dispense2(82) = 4
	// d.Dispense1(83) = 6 but d.Dispense2(83) = 5
	// d.Dispense1(88) = 5 but d.Dispense2(88) = 4
	// d.Dispense1(89) = 6 but d.Dispense2(89) = 5
	// d.Dispense1(124) = 4 but d.Dispense2(124) = 3
	// d.Dispense1(125) = 5 but d.Dispense2(125) = 4
	// d.Dispense1(134) = 5 but d.Dispense2(134) = 4
	// d.Dispense1(135) = 6 but d.Dispense2(135) = 5
	// d.Dispense1(138) = 5 but d.Dispense2(138) = 3
	// d.Dispense1(139) = 6 but d.Dispense2(139) = 4
	// d.Dispense1(148) = 5 but d.Dispense2(148) = 4
	// d.Dispense1(149) = 6 but d.Dispense2(149) = 5
	// d.Dispense1(154) = 5 but d.Dispense2(154) = 4
	// d.Dispense1(155) = 6 but d.Dispense2(155) = 5
	// d.Dispense1(162) = 5 but d.Dispense2(162) = 4
	// d.Dispense1(163) = 6 but d.Dispense2(163) = 5
	// d.Dispense1(164) = 6 but d.Dispense2(164) = 5
	// d.Dispense1(165) = 7 but d.Dispense2(165) = 6
	// d.Dispense1(168) = 6 but d.Dispense2(168) = 4
	// d.Dispense1(169) = 7 but d.Dispense2(169) = 5
	// d.Dispense1(172) = 6 but d.Dispense2(172) = 5
	// d.Dispense1(173) = 7 but d.Dispense2(173) = 6
	// d.Dispense1(178) = 6 but d.Dispense2(178) = 5
	// d.Dispense1(179) = 7 but d.Dispense2(179) = 6
}

func Example_dispense3() {
	stampDenominations := []int{90, 30, 24, 10, 6, 2, 1}
	d2, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}
	d3, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}

	for p := 1; p < 1000; p++ {
		s2, err := d2.Dispense2(p)
		if err != nil {
			fmt.Println(err)
		}
		s3, err := d3.Dispense3(p)
		if err != nil {
			fmt.Println(err)
		}
		if s2 != s3 {
			fmt.Printf("d2.Dispense2(%d) = %d but d3.Dispense3(%d) = %d\n", p, s2, p, s3)
		}
	}
	// Output:
}

func Example_dispense1a() {
	stampDenominations := []int{20, 10, 5, 2, 1} // greedy is a optimal in this case
	d1, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}
	d2, err := stamps.New(stampDenominations)
	if err != nil {
		fmt.Println(err)
	}

	for p := 1; p < 1000; p++ {
		s1, err := d1.Dispense1(p)
		if err != nil {
			fmt.Println(err)
		}
		s2, err := d2.Dispense2(p)
		if err != nil {
			fmt.Println(err)
		}
		if s1 != s2 {
			fmt.Printf("d1.Dispense1(%d) = %d but d1.Dispense2(%d) = %d\n", p, s1, p, s2)
		}
	}
	// Output:
}
