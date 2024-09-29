package ssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// RSA公钥私钥产生
func GenRsaKey(bits int, pwd string) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}

	// 用密码加密 the pem
	if pwd != "" {
		block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return err
		}
	}

	err = pem.Encode(file, block)
	// 生成到内存可以用 private := pem.EncodeToMemory(block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	// 用密码加密 the pem
	if pwd != "" {
		block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return err
		}
	}

	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	// 生成到内存可以用 public := pem.EncodeToMemory(block)
	if err != nil {
		return err
	}

	return nil
}
