package sftp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestPflags(t *testing.T) {
	pflags := newFileOpenFlags(sshFxfRead | sshFxfWrite | sshFxfAppend)
	assert.True(t, pflags.Read)
	assert.True(t, pflags.Write)
	assert.True(t, pflags.Append)
	assert.False(t, pflags.Creat)
	assert.False(t, pflags.Trunc)
	assert.False(t, pflags.Excl)
}

func TestRequestAflags(t *testing.T) {
	aflags := newFileAttrFlags(
		sshFileXferAttrSize | sshFileXferAttrUIDGID)
	assert.True(t, aflags.Size)
	assert.True(t, aflags.UidGid)
	assert.False(t, aflags.Acmodtime)
	assert.False(t, aflags.Permissions)
}

func TestRequestAttributes(t *testing.T) {
	// UID/GID
	fa := FileStat{UID: 1, GID: 2}
	fl := uint32(sshFileXferAttrUIDGID)
	at := []byte{}
	at = marshalUint32(at, 1)
	at = marshalUint32(at, 2)
	testFs, _ := getFileStat(fl, at)
	assert.Equal(t, fa, *testFs)
	// Size and Mode
	fa = FileStat{Mode: 700, Size: 99}
	fl = uint32(sshFileXferAttrSize | sshFileXferAttrPermissions)
	at = []byte{}
	at = marshalUint64(at, 99)
	at = marshalUint32(at, 700)
	testFs, _ = getFileStat(fl, at)
	assert.Equal(t, fa, *testFs)
	// FileMode
	assert.True(t, testFs.FileMode().IsRegular())
	assert.False(t, testFs.FileMode().IsDir())
	assert.Equal(t, testFs.FileMode().Perm(), os.FileMode(700).Perm())
}

func TestRequestAttributesEmpty(t *testing.T) {
	fs, b := getFileStat(sshFileXferAttrAll, nil)
	assert.Equal(t, &FileStat{
		Extended: []StatExtended{},
	}, fs)
	assert.Empty(t, b)
}
