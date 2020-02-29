package sftp

import (
	"io"
	"syscall"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrFxCode(t *testing.T) {
	ider := sshFxpStatusPacket{ID: 1}
	table := []struct {
		err error
		fx  fxerr
	}{
		{err: errors.New("random error"), fx: ErrSSHFxFailure},
		{err: syscall.EBADF, fx: ErrSSHFxFailure},
		{err: syscall.ENOENT, fx: ErrSSHFxNoSuchFile},
		{err: syscall.EPERM, fx: ErrSSHFxPermissionDenied},
		{err: io.EOF, fx: ErrSSHFxEOF},
	}
	for _, tt := range table {
		statusErr := statusFromError(ider, tt.err).StatusError
		assert.Equal(t, statusErr.FxCode(), tt.fx)
	}
}

func TestSupportedExtensions(t *testing.T) {
	for _, supportedExtension := range supportedSFTPExtensions {
		_, err := getSupportedExtensionByName(supportedExtension.Name)
		assert.NoError(t, err)
	}
	_, err := getSupportedExtensionByName("invalid@example.com")
	assert.Error(t, err)
}

func TestExtensions(t *testing.T) {
	var supportedExtensions []string
	for _, supportedExtension := range supportedSFTPExtensions {
		supportedExtensions = append(supportedExtensions, supportedExtension.Name)
	}

	testSFTPExtensions := []string{"hardlink@openssh.com"}
	expectedSFTPExtensions := []sshExtensionPair{
		{"hardlink@openssh.com", "1"},
	}
	err := SetSFTPExtensions(testSFTPExtensions...)
	assert.NoError(t, err)
	assert.Equal(t, expectedSFTPExtensions, sftpExtensions)

	invalidSFTPExtensions := []string{"invalid@example.com"}
	err = SetSFTPExtensions(invalidSFTPExtensions...)
	assert.Error(t, err)
	assert.Equal(t, expectedSFTPExtensions, sftpExtensions)

	emptySFTPExtensions := []string{}
	expectedSFTPExtensions = []sshExtensionPair{}
	err = SetSFTPExtensions(emptySFTPExtensions...)
	assert.NoError(t, err)
	assert.Equal(t, expectedSFTPExtensions, sftpExtensions)

	// if we only have an invalid extension nothing will be modified.
	invalidSFTPExtensions = []string{
		"hardlink@openssh.com",
		"invalid@example.com",
	}
	err = SetSFTPExtensions(invalidSFTPExtensions...)
	assert.Error(t, err)
	assert.Equal(t, expectedSFTPExtensions, sftpExtensions)

	err = SetSFTPExtensions(supportedExtensions...)
	assert.Equal(t, supportedSFTPExtensions, sftpExtensions)
}
