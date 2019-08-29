
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var DebugLevel int = 0

func main() {

	//Decode("10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet")
	//test := "ZEMS)wyq-8Q!B.fu)Bd7,#)B{,UL=WWcKPds0'!q[m,dfhj4?U]BUDITyp,c09z4,wkZ&>Cp91{IJC&e\nVnGFhW(.[g*2l2{ES2.+!eY9Qjd+%mgFePV.qI{2k8tEBY.1\"wH@<CPf5C.a+I.x6_h,590   jYVB_jVuh,88^,EF}mIu62&\"?G6'ozhF,c]pp4S+e\nj/*\"e>1N8.Np+dq _f<X/46BOE7jSe\"Alj_@mQI9cLqJo+O]_c9<[{+f4jYGM8gN\n+{.Kej*'SQso.}ML!JVd-G](]?1E>4,$+yP0msVvI8)d_{/}n:,l3O2x  CGBjTZ(w^gaO=_MD2OS(nZwx,G5><H}2\")te}2/Y^!Wf.ux[6L3l/RREEacG%?qS-+RvLt>rt#)nfP._1Cm%3_Pua.ebvp<ZzRcbW)W\nbH   !@72-2y0mP_:Y#jIN*w437vx6CQoeJT2,Bq.\n(HAx{nkNk\n7*^NI,4nw#DDZpJAyPE*Z."
	//test := "7 O$eegv4'0\"xRu:Dz2Wq1?}?#s9'c8#^EN5ApNxf7a&Yp3pn%LcX8RF]vA*6P"
	//test := "40 l^v^l)>wAlMe[5x1^nxy'm'D<rhB7hxVN7ss3eN&1.yWbkqy6GW=ZsEV,$j%oMu),@^.C''bsexO&'.sJ-wm))F@\n:zfY*>;}NtbXmTuQHcXkLIN+S:6u{-abi@(iZ7+:^4e9]&*jZkj>zl!-rO+pxqL[IV^\":.V]z!4N{59'&k9Z<0$xj9PEg1p_9Tbg%VF[y'.r@aAN\"NuMPR_6  G]lj{h<^x  D&[+NA^)LPoE=?hCE   Bbx8T_^7O6!xte-o\"V(\n3')/93)"
	//test := "10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet"
	//test := "74 K99?2g-\"1EvI>C+:NwFeT\"nMNk1 kd:^XZI0;4-}I,\n6k')),&tH,>Y!rmj.Gb  =egR(   .<b*VXflJNv<^3Qe)&hRF6tlv/Z9jP7mL]__blH*j3)2/dRh4>/>9z4!60uY\nYE35Ie^4'K,exlL7=nf&SWXD2Mk}Q*U2:fMgX6'HJ?5>;W)fSusd$jAW: kn(7As C;J<;t\nFGX;9\",E$cr.f&[(T15v1#! bOVd.E#\"y{rWFY )D]O\n++&Gd8z'/]YT0V6>[Cc*6.kD}F,:v$sYG1J>#Y\nDpz/.TAMI43zP?SWY\"QSTwiZFr.5M^8#]&y( .V<jH_.Dr&n[dhL-6xZUfY@]I#^6'b.DULwyR  eg*Tg-''o-H ,xy%F   <zGLt*[Hy3'.(?CxhdNSU%n8Q^<FnDEK@<\"Z(Ij<bnU5K"
	//Decode(test)

	//Decode("10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet")
	Decode2("10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet")

	//test := "l1NkL%.bX#($o9.#   q<R;S2iMpP;#}k"
	//Encode(7, "0bOg&gWIfAN.y@M)R4P)SAZ3^aDuHKOCw")
	//Encode(8, "vd-5E2L@mI$   @3mdvW[iF$?g>k  uOiMk/Cfd")

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

func Decode2(s string) string {
	// given the length of the string, and the location of the spaces
	ret := ""
	fmt.Printf("Decode2() s = >>>%v<<<\n", s)
	idx := strings.Index(s, " ")
	n, _ := strconv.Atoi(s[:idx])
	fmt.Printf("n = %v\n", n)
	dome := s[idx+1:] // This will get changed at the end of each cycle

	spaces := []int{}    // indexes of spaces in the original string
	nonspaces := []int{} // indexes of non spaces in the orinal string
	//template := []int{}  // This is the template  should be the length of dome
	template := make([]int, len(dome), len(dome))

	for x := 0; x < len(template); x++ {
		template[x] = x
	}

	///////////// begin main loop
	for cycles := 0; cycles < n; cycles++ {

		for x := 0; x < len(dome); x++ {
			if dome[x] == ' ' {
				spaces = append(spaces, x)
			} else {
				nonspaces = append(nonspaces, x)
			}

		}

		fmt.Printf("spaces    : %v\n", spaces)
		fmt.Printf("nonspaces : %v\n", nonspaces)

		// now, find the groups, just dump it to the screen
		groups := [][]int{}
		group := []int{}

		//group = append(group, nonspaces[0]) // start here

		for x := 0; x < len(nonspaces); x++ {

			group = append(group, nonspaces[x])

			// if the next one is a gap, or if we are done then write this group out
			if x == len(nonspaces)-1 || nonspaces[x+1]-nonspaces[x] > 1 {
				groups = append(groups, group)
				group = []int{} // reset
			}
		}

		fmt.Printf("groups : %v\n", groups)

		// now iterate thru each group, rotating to the left n times
		fmt.Printf("Rotating each group left %d times\n", n)

		for x := 0; x < len(groups); x++ {
			if len(groups[x]) == 1 {
				continue
			}
			loops := n % len(groups[x])
			if loops > 0 {
				groups[x] = append(groups[x][loops:], groups[x][0:loops]...)
			}
		}

		fmt.Printf("groups : %v\n", groups)

		fmt.Println("Put the spaces back in and rotate that to the left n times")
		// build template here
		spaceCnt := 0
		groupCnt := 0
		tmplatePtr := 0
		//TODO : STOPPED HERE.... not the below isn't working properly yet.... we are trying to rebuild template
		cycles := len(groups) + len(spaces)
		for x := 0; x < cycles; x++ {
			if spaceCnt < len(spaces) && spaces[spaceCnt] == tmplatePtr {
				template[tmplatePtr] = tmplatePtr
				fmt.Printf("%d.", tmplatePtr)
				tmplatePtr++
				spaceCnt++
			} else {
				if groupCnt < len(groups) {
					for y := 0; y < len(groups[groupCnt]); y++ {
						template[tmplatePtr] = groups[groupCnt][y]
						tmplatePtr++
						fmt.Printf("%d.", groups[groupCnt][y])
						//fmt.Printf("template : %v\n", template)
					}
					groupCnt++
				}
			}
		}

		fmt.Printf("\ntemplate : %v\n", template)
		// now rotate template to the left n times
		fmt.Printf("and rotate left %d times\n", n)
		template = append(template[n:], template[0:n]...)
		fmt.Println(template)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------")
		//now, rebuild dome
		s := ""
		for x := 0; x < len(template); x++ {
			s += string(dome[template[x]])
		}
		dome = s
	}

	return ret
}

func Decode(s string) string {
	// parse off the pre-pended number
	fmt.Printf("Welcome to Decode()\ns=%s\n", s)
	idx := strings.Index(s, " ")
	//fmt.Printf("idx=%d\n", idx)
	//fmt.Printf("s[:idx]=%v\n", s[:idx])
	n, _ := strconv.Atoi(s[:idx])
	//fmt.Println(n)

	s1 := ""
	s2 := ""
	s3 := ""
	s4 := ""
	s5 := ""

	dome := s[idx+1:]
	//s1 := dome[idx+1:]

	for x := 0; x < n; x++ {
		fmt.Printf("===============================================================================================\n")
		fmt.Printf("Decode() pass#:%d/%d\n", x, n-1)

		// Rotate each space-seperated substring to the left by n characters // TODO : put this in a function?
		fmt.Printf("Rotate each space-seperated substring to the left by %d characters.\n", n)

		s1 = dome
		fmt.Printf("s1=%v\n", s1)
		tmp := strings.Split(s1, " ")
		//fmt.Printf("tmp=%s\n", tmp)
		s2 = ""
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
		s3 = ""           // spaces removed will get stored here
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
		s4 = rotateStringleft(n, s3)
		fmt.Printf("s4=%s\n\n", s4)

		// Place the spaces back in their original positions
		fmt.Printf("Put the spaces back in\n")

		// q? If there are no spaces, what do we do?
		if len(spaces) == 0 {
			dome = s4
			continue // we are done with this iteration
		}

		s5 = ""
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
	fmt.Printf("Welcome to Encode()\ns=>>>%s<<<\nn=%d\n", s, n)
	//fmt.Printf("Welcome to Encode()\ns=>>>%s<<<\n", s)
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
	if len(spaces) > 0 {
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

	} else {
		step3string = ret

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
	if len(s) == 0 {
		return s
	}
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
