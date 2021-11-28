package bytespkg

import (
	"fmt"
)

func BasicBytes() {
	s := "this is a string"
	bbs := []byte(s)
	bs := []byte{71, 111}
	fmt.Printf("%s\n", bs)
	fmt.Printf("%s\n", bbs)
	fmt.Printf("%d", bbs)
}
