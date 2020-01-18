package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

//SftpClient sftp客户端
type SftpClient struct {
	client    *sftp.Client
	passSize  int64
	totalSize int64
	finish    chan error
}

//NewSftpClient new sftp客户端
func NewSftpClient(client *ssh.Client) (*SftpClient, error) {
	sc, err := SftpConnect(client)
	if err != nil {
		return nil, err
	}
	return &SftpClient{
		client:    sc,
		passSize:  0,
		totalSize: 0,
		finish:    make(chan error),
	}, nil
}

//ScpCopy scp复制
func (s *SftpClient) ScpCopy(localFilePath, remoteDir string) error {
	var (
		err error
	)

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		s.finish <- err
		return err
	}
	defer srcFile.Close()
	fInfo, err := srcFile.Stat()
	if err != nil {
		s.finish <- err
		return err
	}
	s.totalSize = fInfo.Size()

	var remoteFileName = path.Base(localFilePath)
	tmpFile := path.Join(remoteDir, fmt.Sprintf("%s%s", remoteFileName, ".tmp"))

	dstFile, err := s.client.Create(tmpFile)
	if err != nil {
		s.finish <- err
		return err
	}
	defer dstFile.Close()
	fmt.Println(dstFile.ReadFrom(srcFile))

	_, err = dstFile.ReadFrom(srcFile)
	if err != nil {
		s.finish <- err
		return err
	}
	s.client.Rename(tmpFile, path.Join(remoteDir, remoteFileName))
	s.finish <- nil

	return nil
}

//GetProcess GetProcess
func (s *SftpClient) GetProcess() string {
	if s.totalSize == 0 {
		return "0.00"
	}
	return fmt.Sprintf("%.2f", float64(s.passSize)*100/float64(s.totalSize))
}

//CheckPathIsExisted CheckPathIsExisted
func (s *SftpClient) CheckPathIsExisted(path string) error {
	_, err := s.client.Stat(path)
	return err
}

//Finish Finish
func (s *SftpClient) Finish() error {
	//period := time.Duration(5) * time.Second
	//t := time.NewTicker(period)
	//for {
	//	select {
	//	case <-t.C:
	//		{
	//			fmt.Println(s.GetProcess())
	//		}
	//	case err := <-s.finish:
	//		{
	//			fmt.Println("100.00")
	//			return err
	//		}
	//	}
	//}
	return <-s.finish
}

//Close Close
func (s *SftpClient) Close() error {
	close(s.finish)
	return s.client.Close()
}

//SftpConnect SftpConnect
func SftpConnect(client *ssh.Client) (*sftp.Client, error) {
	var (
		sftpClient *sftp.Client
		err        error
	)
	// create sftp client
	if sftpClient, err = sftp.NewClient(client); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
