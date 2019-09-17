package utils

import (
  "fmt"
  "golang.org/x/crypto/ssh"
  "net"
  "time"
  // "io/ioutil"
)

func Connect(user, password, host, key string, port int, cipherList []string) (*ssh.Session, error) {
  var (
    auth     []ssh.AuthMethod
    addr     string
    clientConfig *ssh.ClientConfig
    client    *ssh.Client
    config    ssh.Config
    session   *ssh.Session
    err     error
  )
  // get auth method
  auth = make([]ssh.AuthMethod, 0)
  if key == "" {
    auth = append(auth, ssh.Password(password))
  } else {
    // pemBytes, err := ioutil.ReadFile(key)
    // if err != nil {
    //   return nil, err
    // }
    // fmt.Println(key)

    pemBytes := []byte(key)
    var signer ssh.Signer
    if password == "" {
      signer, err = ssh.ParsePrivateKey(pemBytes)
    } else {
			// 使用私钥解析密码
      signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password))
    }
    if err != nil {
      return nil, err
    }
    auth = append(auth, ssh.PublicKeys(signer))
  }
 
  if len(cipherList) == 0 {
    config = ssh.Config{
      Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
    }
  } else {
    config = ssh.Config{
      Ciphers: cipherList,
    }
  }
  clientConfig = &ssh.ClientConfig{
    User:  user,
    Auth:  auth,
    Timeout: 30 * time.Second,
    Config: config,
    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
      return nil
    },
  }
 
  // connet to ssh
  addr = fmt.Sprintf("%s:%d", host, port)
 
  if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
    return nil, err
  }
 
  // create session
  if session, err = client.NewSession(); err != nil {
    return nil, err
  }
 
  modes := ssh.TerminalModes{
    ssh.ECHO:     0,   // disable echoing
    ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
    ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
  }
 
  if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
    return nil, err
  }
 
  return session, nil
}