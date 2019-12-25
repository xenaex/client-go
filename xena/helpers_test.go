package xena

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDLen(t *testing.T) {
	id := ID("")
	fmt.Println(len(id), id)
	assert.Equal(t, maxIdLen, len(id))

	id = ID("dskjfhasdofhweifohwekfhsdklfhsdikhfsdhfsdkfhahsfksdhfasdifhaskjfhsduilfhdsj")
	fmt.Println(len(id), id)
	assert.Equal(t, maxIdLen, len(id))
}
