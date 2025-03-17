package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))

}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$ukaMULnlz6msDp8xCwmO.OOA1PHgrosqEUUcu6VaRcr62kWAx4pr2", "12341"))
}
