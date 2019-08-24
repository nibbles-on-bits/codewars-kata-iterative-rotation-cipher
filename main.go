package main

import (
	"fmt"
	"strings"
)

func main() {

	quote := `If you wish to make an apple pie from scratch, you must first invent the universe.`
	//solution := `10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet`
	IRCEncode(10, quote) //== solution //true

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

	fmt.Println(rotateStringRight(132345, "something"))
}

func IRCEncode(n int, s string) string {
	ret := ""

	//sArr := []rune(s)
	// remove the spaces but remember the position
	spaces := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			fmt.Println(i)
			spaces = append(spaces, i)
		} else {
			ret += string(s[i])
		}
	}

	fmt.Println(spaces)
	fmt.Println(ret)

	// shift the order of the string characters to the right by n
	ret = ret[len(ret)-n:] + ret[:len(ret)-n]
	fmt.Println("after step 2")
	fmt.Println(ret)
	fmt.Println()

	step3string := ""
	spacePtr := 0
	charPtr := 0

	for i := 0; i < len(s); i++ {
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
		fmt.Println(step3string)
	}

	fmt.Println(step3string)

	// now shift each space seperated substring to the right n times
	tmp := strings.Fields(step3string)
	fmt.Println(tmp)
	fmt.Println(len(tmp))
	for i := 0; i < len(tmp); i++ {
		rotateStringRight(n, tmp[i])
	}

	return ret
}

func rotateStringRight(n int, s string) string {
	rs := ""
	fmt.Printf("rotateStringRight() called  n=%d  s=%s\n", n, s)

	loops := 0
	if n > len(s) {
		loops = len(s) % n
	} else {
		loops = n
	}

	rs = s[loops-1:] + s[:loops-1]

	return rs
}
