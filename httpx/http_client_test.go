package httpx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func _TestNewClient(t *testing.T) {
	client := NewClient(UseProxy("socks5://127.0.0.1:1080"))
	ri, body, err := client.Get("https://www.google.com").ReadAllString()
	assert.NoError(t, err)
	fmt.Println(ri)
	fmt.Println(body)
}
