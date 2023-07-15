package tiansuo_snow

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	id1 := Generate_SnowflakeID_str()
	fmt.Println(id1)
	id2 := Generate_SnowflakeID_i64()
	fmt.Println(id2)
}
