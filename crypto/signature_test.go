package crypto

import "testing"
import "fmt"
import "encoding/hex"

func Test_GenerateKey(t *testing.T) {
	//t.Log(GenerateKey())
	x,y := GenerateKey()
	fmt.Println(hex.EncodeToString(x))
    fmt.Println(hex.EncodeToString(y))
	
}



