package BOE
import (
    "fmt"
    "testing"
)
var (
    boe = BoeGetInstance()
)

func TestValidateSign(t *testing.T) {
    var (
        hash = make([]byte, 32)
        r    = make([]byte, 32)
        s    = make([]byte, 32)
    )
    var v byte

    x,y,err := boe.ValidateSign(hash, r, s, v)
    if err == nil {
        //fmt.Printf("len(x)=%d\n", len(x))
        for i:=0; i < len(x); i++ {
            fmt.Printf("x[%d]=%02x\n", i, x[i])
        }
        for i:=0; i < len(y); i++ {
            fmt.Printf("y[%d]=%02x\n", i, y[i])
        }
    }
}

func TestHWSign(t *testing.T) {
    var (
        hash = make([]byte, 32)
    )
    for i:=0; i < 32; i++ {
        hash[i] = byte(i)
    }

    result,err := boe.HWSign(hash)
    if err == nil {
        //fmt.Printf("len(x)=%d\n", len(x))
        for i:=0; i < 32; i++ {
            fmt.Printf("signval[%d]=%02x\n",i,result.val[i])
        }
    }
}

func TestNewEvent(t *testing.T) {
    var ver = boe.GetHWVersion()
    fmt.Printf("hwversion = %02x\n", ver)
}
