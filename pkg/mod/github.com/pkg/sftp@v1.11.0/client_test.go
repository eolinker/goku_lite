package sftp

import (
	"errors"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/kr/fs"
)

// assert that *Client implements fs.FileSystem
var _ fs.FileSystem = new(Client)

// assert that *File implements io.ReadWriteCloser
var _ io.ReadWriteCloser = new(File)

func TestNormaliseError(t *testing.T) {
	var (
		ok         = &StatusError{Code: sshFxOk}
		eof        = &StatusError{Code: sshFxEOF}
		fail       = &StatusError{Code: sshFxFailure}
		noSuchFile = &StatusError{Code: sshFxNoSuchFile}
		foo        = errors.New("foo")
	)

	var tests = []struct {
		desc string
		err  error
		want error
	}{
		{
			desc: "nil error",
		},
		{
			desc: "not *StatusError",
			err:  foo,
			want: foo,
		},
		{
			desc: "*StatusError with ssh_FX_EOF",
			err:  eof,
			want: io.EOF,
		},
		{
			desc: "*StatusError with ssh_FX_NO_SUCH_FILE",
			err:  noSuchFile,
			want: os.ErrNotExist,
		},
		{
			desc: "*StatusError with ssh_FX_OK",
			err:  ok,
		},
		{
			desc: "*StatusError with ssh_FX_FAILURE",
			err:  fail,
			want: fail,
		},
	}

	for _, tt := range tests {
		got := normaliseError(tt.err)
		if got != tt.want {
			t.Errorf("normaliseError(%#v), test %q\n- want: %#v\n-  got: %#v",
				tt.err, tt.desc, tt.want, got)
		}
	}
}

var flagsTests = []struct {
	flags int
	want  uint32
}{
	{os.O_RDONLY, sshFxfRead},
	{os.O_WRONLY, sshFxfWrite},
	{os.O_RDWR, sshFxfRead | sshFxfWrite},
	{os.O_RDWR | os.O_CREATE | os.O_TRUNC, sshFxfRead | sshFxfWrite | sshFxfCreat | sshFxfTrunc},
	{os.O_WRONLY | os.O_APPEND, sshFxfWrite | sshFxfAppend},
}

func TestFlags(t *testing.T) {
	for i, tt := range flagsTests {
		got := flags(tt.flags)
		if got != tt.want {
			t.Errorf("test %v: flags(%x): want: %x, got: %x", i, tt.flags, tt.want, got)
		}
	}
}

func TestUnmarshalStatus(t *testing.T) {
	requestID := uint32(1)

	id := marshalUint32([]byte{}, requestID)
	idCode := marshalUint32(id, sshFxFailure)
	idCodeMsg := marshalString(idCode, "err msg")
	idCodeMsgLang := marshalString(idCodeMsg, "lang tag")

	var tests = []struct {
		desc   string
		reqID  uint32
		status []byte
		want   error
	}{
		{
			desc:   "well-formed status",
			reqID:  1,
			status: idCodeMsgLang,
			want: &StatusError{
				Code: sshFxFailure,
				msg:  "err msg",
				lang: "lang tag",
			},
		},
		{
			desc:   "missing error message and language tag",
			reqID:  1,
			status: idCode,
			want: &StatusError{
				Code: sshFxFailure,
			},
		},
		{
			desc:   "missing language tag",
			reqID:  1,
			status: idCodeMsg,
			want: &StatusError{
				Code: sshFxFailure,
				msg:  "err msg",
			},
		},
		{
			desc:   "request identifier mismatch",
			reqID:  2,
			status: idCodeMsgLang,
			want:   &unexpectedIDErr{2, requestID},
		},
	}

	for _, tt := range tests {
		got := unmarshalStatus(tt.reqID, tt.status)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("unmarshalStatus(%v, %v), test %q\n- want: %#v\n-  got: %#v",
				requestID, tt.status, tt.desc, tt.want, got)
		}
	}
}

type packetSizeTest struct {
	size  int
	valid bool
}

var maxPacketCheckedTests = []packetSizeTest{
	{size: 0, valid: false},
	{size: 1, valid: true},
	{size: 32768, valid: true},
	{size: 32769, valid: false},
}

var maxPacketUncheckedTests = []packetSizeTest{
	{size: 0, valid: false},
	{size: 1, valid: true},
	{size: 32768, valid: true},
	{size: 32769, valid: true},
}

func TestMaxPacketChecked(t *testing.T) {
	for _, tt := range maxPacketCheckedTests {
		testMaxPacketOption(t, MaxPacketChecked(tt.size), tt)
	}
}

func TestMaxPacketUnchecked(t *testing.T) {
	for _, tt := range maxPacketUncheckedTests {
		testMaxPacketOption(t, MaxPacketUnchecked(tt.size), tt)
	}
}

func TestMaxPacket(t *testing.T) {
	for _, tt := range maxPacketCheckedTests {
		testMaxPacketOption(t, MaxPacket(tt.size), tt)
	}
}

func testMaxPacketOption(t *testing.T, o ClientOption, tt packetSizeTest) {
	var c Client

	err := o(&c)
	if (err == nil) != tt.valid {
		t.Errorf("MaxPacketChecked(%v)\n- want: %v\n- got: %v", tt.size, tt.valid, err == nil)
	}
	if c.maxPacket != tt.size && tt.valid {
		t.Errorf("MaxPacketChecked(%v)\n- want: %v\n- got: %v", tt.size, tt.size, c.maxPacket)
	}
}

func testFstatOption(t *testing.T, o ClientOption, value bool) {
	var c Client

	err := o(&c)
	if err == nil && c.useFstat != value {
		t.Errorf("UseFStat(%v)\n- want: %v\n- got: %v", value, value, c.useFstat)
	}
}

func TestUseFstatChecked(t *testing.T) {
	testFstatOption(t, UseFstat(true), true)
	testFstatOption(t, UseFstat(false), false)
}
