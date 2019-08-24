package main

import (
	"fmt"
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
}

func IRCEncode(n int, s string) string {
	ret := ""

	//sArr := []rune(s)
	// remove the spaces but remember the position
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			fmt.Println(i)
		} else {
			ret += string(s[i])
		}
	}

	fmt.Println(ret)

	// shift the order of the string characters to the right by n
	ret = ret[len(ret)-n:] + ret[:len(ret)-n]
	//fmt.Println(ret[len(ret)-n:])
	//fmt.Println(ret[:len(ret)-n])
	//fmt.Println(ret[5:])

	return ret
}
