package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

var (
	MyKey []byte
	MyCertificate tls.Certificate
)

func Init() error {

	var err error
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	cert := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			CommonName: "localhost",
			Organization: []string{"In-memory TLS password manager"},
		},
		NotBefore: time.Now().Add(-time.Hour),
		NotAfter: time.Now().Add(365 * 24 * time.Hour),
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames: []string{"localhost", "127.0.0.1"},
	}

	certSigned, err := x509.CreateCertificate(rand.Reader, &cert, &cert, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certSigned})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	c, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return err
	}

	MyKey = keyPEM
	MyCertificate = c

	return nil
}