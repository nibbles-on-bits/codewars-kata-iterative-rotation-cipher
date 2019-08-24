package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	Decode("10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet")
	/*quote := `If you wish to make an apple pie from scratch, you must first invent the universe.`
	//solution := `10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet`


	fmt.Println()
	fmt.Println("result:")
	fmt.Println(Encode(10, quote)) //== solution //true

	encode_examples := [...]map[int]string{
		{10: "If you wish to make an apple pie from scratch, you must first invent the universe."},
		{14: "True evil is a mundane bureaucracy."},
		{22: "There is nothing more atrociously cruel than an adored child."},
		{36: "As I was going up the stair\nI met a man who wasn't there!\nHe wasn't there again today,\nOh how I wish he'd go away!"},
		{29: "I avoid that bleak first hour of the working day during which my still sluggish senses and body make every chore a penance.\nI find that in arriving later, the work which I do perform is of a much higher quality."},
	}

	for _, z := range encode_examples {
		for k, v := range z {
			//Encode(k, v)
			fmt.Println(Encode(k, v))
		}
	}*/

	/* Step-by-step breakdown:
	   Step 1 — remove all spaces:
	   `Ifyouwishtomakeanapplepiefromscratch,youmustfirstinventtheuniverse.`

	   Step 2 — shift the order of string characters to the right by 10:
	   `euniverse.Ifyouwishtomakeanapplepiefromscratch,youmustfirstinventth`

	   Step 3 — place the spaces back in their original positions:
	   `eu niv erse .I fyou wi shtom ake anap plepiefr oms crat ch,yo umustf irs tinventth`

	   Step 4 — shift the order of characters for each space-separated substring to the right by 10:
	   `eu vni seer .I oufy wi shtom eak apan frplepie som atcr ch,yo ustfum sir htinventt`

	   Repeat the steps 9 more times before returning the string with `10 ` prepended.
	*/

	//fmt.Println(rotateStringRight(10, "plepiefr"))
}

func Decode(s string) string {
	// parse off the pre-pended number
	fmt.Printf("Welcome to Decode() s=%s\n", s)
	idx := strings.Index(s, " ")
	fmt.Printf("idx=%d\n", idx)
	fmt.Printf("s[:idx]=%v\n", s[:idx])
	n, _ := strconv.Atoi(s[:idx])
	fmt.Println(n)

	s1 := s[idx:]
	fmt.Printf("s1=%v\n", s1)

	s2 := rotateStringleft(n, s1)
	fmt.Printf("s2=%v\n", s2)

	// now remove spaces (keep track of where they were)
	s3 := ""          // spaces removed will get stored here
	spaces := []int{} // indicies of  spaces in s2 get stored here
	for i := 0; i < len(s2); i++ {
		if s2[i] == ' ' {
			//fmt.Println(i)
			spaces = append(spaces, i)
		} else {
			s3 += string(s2[i])
		}
	}

	fmt.Printf("spaces[] = %v\n", spaces)
	fmt.Printf("s3=%v\n", s3)

	return ""
}

func Encode(n int, s string) string {
	ret := s
	for i := 0; i < n; i++ {
		ret = doapass(n, ret)
	}

	return strconv.Itoa(n) + " " + ret
}

func doapass(n int, s string) string {
	ret := ""
	process := s
	spaces := []int{}

	// Step 1 - remove all spaces keeping track of where they were at
	for i := 0; i < len(process); i++ {
		if process[i] == ' ' {
			//fmt.Println(i)
			spaces = append(spaces, i)
		} else {
			ret += string(process[i])
		}
	}

	// Step 2 - Shift the order of the string characters to the right by n
	ret = ret[len(ret)-n:] + ret[:len(ret)-n]

	// Step 3 - Place the spaces back in their original positions
	step3string := ""
	spacePtr := 0
	charPtr := 0

	for i := 0; i < len(process); i++ {
		if spaces[spacePtr] == i {
			step3string += " "
			spacePtr++
			// if we have added the last space, just append the remaining characters and break out
			if spacePtr == len(spaces) {
				step3string += ret[charPtr:]
				break
			}
		} else {
			step3string += string(ret[charPtr])
			charPtr++
		}
		//fmt.Println(step3string)
	}

	//fmt.Println(step3string)

	// Step 4 - now shift each space seperated substring to the right n times
	//tmp := strings.Fields(step3string) // use a different function
	tmp := strings.Split(step3string, " ")
	//fmt.Println(tmp)
	//fmt.Println(len(tmp))
	step4string := ""
	for i := 0; i < len(tmp); i++ {
		tmp[i] = rotateStringRight(n, tmp[i])
		step4string = step4string + tmp[i]
		if i < len(tmp)-1 {
			step4string = step4string + " "
		}
	}

	//fmt.Println("step4string")
	//fmt.Println(step4string)
	process = step4string

	return process
}

func rotateStringleft(n int, s string) string {
	rs := ""
	loops := 0

	if n > len(s) {
		loops = n % len(s)
	} else {
		loops = n
	}

	if loops > 0 {
		rs = s[n:] + s[:n]
	} else {
		rs = s
	}

	return rs
}

func rotateStringRight(n int, s string) string {
	rs := ""
	//fmt.Printf("rotateStringRight() called  n=%d  s=%s\n", n, s)

	loops := 0
	if n > len(s) {
		loops = n % len(s)
	} else {
		loops = n
	}

	pos := len(s) - loops

	if loops > 0 {
		rs = s[pos:] + s[:pos]
	} else {
		return s
	}

	return rs
}
