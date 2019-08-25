package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//Decode("10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet")
	//test := "ZEMS)wyq-8Q!B.fu)Bd7,#)B{,UL=WWcKPds0'!q[m,dfhj4?U]BUDITyp,c09z4,wkZ&>Cp91{IJC&e\nVnGFhW(.[g*2l2{ES2.+!eY9Qjd+%mgFePV.qI{2k8tEBY.1\"wH@<CPf5C.a+I.x6_h,590   jYVB_jVuh,88^,EF}mIu62&\"?G6'ozhF,c]pp4S+e\nj/*\"e>1N8.Np+dq _f<X/46BOE7jSe\"Alj_@mQI9cLqJo+O]_c9<[{+f4jYGM8gN\n+{.Kej*'SQso.}ML!JVd-G](]?1E>4,$+yP0msVvI8)d_{/}n:,l3O2x  CGBjTZ(w^gaO=_MD2OS(nZwx,G5><H}2\")te}2/Y^!Wf.ux[6L3l/RREEacG%?qS-+RvLt>rt#)nfP._1Cm%3_Pua.ebvp<ZzRcbW)W\nbH   !@72-2y0mP_:Y#jIN*w437vx6CQoeJT2,Bq.\n(HAx{nkNk\n7*^NI,4nw#DDZpJAyPE*Z."
	//test := "7 O$eegv4'0\"xRu:Dz2Wq1?}?#s9'c8#^EN5ApNxf7a&Yp3pn%LcX8RF]vA*6P"
	test := "40 l^v^l)>wAlMe[5x1^nxy'm'D<rhB7hxVN7ss3eN&1.yWbkqy6GW=ZsEV,$j%oMu),@^.C''bsexO&'.sJ-wm))F@\n:zfY*>;}NtbXmTuQHcXkLIN+S:6u{-abi@(iZ7+:^4e9]&*jZkj>zl!-rO+pxqL[IV^\":.V]z!4N{59'&k9Z<0$xj9PEg1p_9Tbg%VF[y'.r@aAN\"NuMPR_6  G]lj{h<^x  D&[+NA^)LPoE=?hCE   Bbx8T_^7O6!xte-o\"V(\n3')/93)"
	Decode(test)

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
	fmt.Printf("Welcome to Decode()\ns=%s\n", s)
	idx := strings.Index(s, " ")
	//fmt.Printf("idx=%d\n", idx)
	//fmt.Printf("s[:idx]=%v\n", s[:idx])
	n, _ := strconv.Atoi(s[:idx])
	//fmt.Println(n)

	dome := s[idx+1:]
	//s1 := dome[idx+1:]

	for x := 0; x < n; x++ {
		fmt.Printf("===============================================================================================\n")
		fmt.Printf("Decode() pass#:%d/%d\n", x, n-1)

		// Rotate each space-seperated substring to the left by n characters // TODO : put this in a function?
		fmt.Printf("Rotate each space-seperated substring to the left by %d characters.\n", n)

		s1 := dome
		fmt.Printf("s1=%v\n", s1)
		tmp := strings.Split(s1, " ")
		//fmt.Printf("tmp=%s\n", tmp)
		s2 := ""
		for i := 0; i < len(tmp); i++ {
			tmp[i] = rotateStringleft(n, tmp[i])
			s2 += tmp[i]
			if i < len(tmp)-1 {
				s2 = s2 + " "
			}
		}
		fmt.Printf("s2=%s\n\n", s2)

		// now remove spaces (keep track of where they were)
		fmt.Printf("Remove the spaces (keeping track of where they were)\n")
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
		fmt.Printf("s3=%v\n\n", s3)

		// now rotate left by n
		fmt.Printf("Now, rotate left %d times\n", n)
		s4 := rotateStringleft(n, s3)
		fmt.Printf("s4=%s\n\n", s4)

		// Place the spaces back in their original positions
		fmt.Printf("Put the spaces back in\n")

		// q? If there are no spaces, what do we do?
		if len(spaces) == 0 {
			dome = s4
			continue // we are done with this iteration
		}

		s5 := ""
		spacePtr := 0
		charPtr := 0

		for i := 0; i < len(dome); i++ {
			if spaces[spacePtr] == i {
				s5 += " "
				spacePtr++
				// if we have added the last space, just append the remaining characters and break out
				if spacePtr == len(spaces) {
					s5 += s4[charPtr:]
					break
				}
			} else {
				s5 += string(s4[charPtr])
				charPtr++
			}
		}
		fmt.Printf("s5=%s\n", s5)
		dome = s5
	}

	return dome
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

// TODO : create this fn
// putSpacesBack will put spaces back in a string based on an array of space indexes and a string
//func putSpacesBack(spaces int[], s string) string {

//}

func rotateStringleft(n int, s string) string {
	if s == "" {
		return ""
	}
	rs := ""
	loops := 0

	if n > len(s) {
		loops = n % len(s)
	} else {
		loops = n
	}

	if loops > 0 {
		rs = s[loops:] + s[:loops]
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
