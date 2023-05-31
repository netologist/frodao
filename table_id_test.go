package frodao

import (
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/netologist/frodao/tableid"
	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	intID := ID[tableid.Int]{}
	stringID := ID[tableid.String]{}
	uuidID := ID[tableid.UUID]{}

	require.Nil(t, intID.Scan(int64(1)))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(1))

	require.Nil(t, intID.Scan(int32(2)))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(2))

	require.Nil(t, intID.Scan(int16(3)))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(3))

	require.Nil(t, intID.Scan(int(4)))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(4))

	require.Nil(t, intID.Scan([]byte(strconv.Itoa(5))))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(5))

	sliceValue := []byte(strconv.Itoa(6))
	require.Nil(t, intID.Scan(sliceValue[:]))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(6))

	require.Nil(t, intID.Scan(7))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(7))

	require.Nil(t, stringID.Scan("lorem-ipsum"))
	require.True(t, stringID.Exist)
	require.Equal(t, stringID.IDValue, "lorem-ipsum")

	require.Nil(t, stringID.Scan([]byte("ipsum-lorem")))

	require.Equal(t, stringID.IDValue, "ipsum-lorem")
	require.True(t, stringID.Exist)

	require.Nil(t, uuidID.Scan(uuid.New()))
	require.True(t, intID.Exist)
	require.Equal(t, intID.IDValue, int64(7))
}
