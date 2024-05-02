package auth

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/snowflakedb/gosnowflake"
)

func append_private_key_string(dsn *string, config gosnowflake.Config) string {
	var b bytes.Buffer
	b.WriteString(*dsn)
	b.WriteString(fmt.Sprintf("&authenticator=%s", config.Authenticator))
	b.WriteString(fmt.Sprintf("&private_key=%s", config.Token))
	return b.String()
}

func Setup_private_key(key_path string) *rsa.PrivateKey {
	data, err := os.ReadFile(key_path)
	block, _ := pem.Decode(data)
	if err != nil || block.Type != "PRIVATE KEY" {
		panic(err)
	}
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return privKey.(*rsa.PrivateKey)
}

func Generate_private_key_string(key string) string {
	rsa_key := Setup_private_key(key)
	tmp_bytes, _ := x509.MarshalPKCS8PrivateKey(rsa_key)
	privateKeyPKCS8 := base64.URLEncoding.EncodeToString(tmp_bytes)
	return privateKeyPKCS8
}
