package dynaport

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	ports, err := GetWithErr(3)
	require.NoError(t, err)
	require.Equal(t, 3, len(ports))

	for _, port := range ports {
		ln, err := net.ListenTCP("tcp", &net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: port,
		})
		require.NoError(t, err)
		ln.Close()
	}
}
