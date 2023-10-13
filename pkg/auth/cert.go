package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"time"

	"software.sslmate.com/src/go-pkcs12"
)

// PEM block types
const (
	certificateType = "CERTIFICATE"
	privateKeyType  = "PRIVATE KEY"
)

var (
	ErrPEMInvalid = errors.New("failed to parse PEM")
)

type Certificate struct {
	caCert       *x509.Certificate
	caPrivateKey crypto.PrivateKey
}

func NewCertificate(opts ...func(*Certificate)) *Certificate {
	caPrivateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	c := &Certificate{
		caCert: &x509.Certificate{
			SerialNumber: big.NewInt(0),
			Subject: pkix.Name{
				CommonName: "0.0.0.0",
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().AddDate(10, 0, 0),
			KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			BasicConstraintsValid: true,
			IsCA:                  true,
		},
		caPrivateKey: caPrivateKey,
	}

	for _, f := range opts {
		f(c)
	}

	return c
}

func WithCACert(caCert *x509.Certificate) func(*Certificate) {
	return func(c *Certificate) {
		c.caCert = caCert
	}
}

func WithCAPrivateKey(caPrivateKey crypto.PrivateKey) func(*Certificate) {
	return func(c *Certificate) {
		c.caPrivateKey = caPrivateKey
	}
}

func (c *Certificate) Generate(commonName string) (certPEM []byte, privPEM []byte, err error) {
	serialNumber, err := generateSerialNumber()
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, c.caCert, priv.Public(), c.caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	certPEM = pem.EncodeToMemory(&pem.Block{Type: certificateType, Bytes: derBytes})

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	privPEM = pem.EncodeToMemory(&pem.Block{Type: privateKeyType, Bytes: privBytes})

	return certPEM, privPEM, nil
}

// generateSerialNumber generate serial number
func generateSerialNumber() (*big.Int, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)

	return rand.Int(rand.Reader, serialNumberLimit)
}

func NewTLSConfig(pfxData []byte) (*tls.Config, error) {
	caCertPEM, clientCertPEM, clientKeyPEM, err := ParsePKCS12(pfxData)
	if err != nil {
		return nil, err
	}

	// Import trusted certificates from caCertPEM.
	// Alternatively, manually add CA certificates to
	// default openssl CA bundle.
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caCertPEM)

	// Import client certificate/key pair
	cert, err := tls.X509KeyPair(clientCertPEM, clientKeyPEM)
	if err != nil {
		return nil, err
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}, nil
}

func ParsePKCS12(pfxData []byte) (caCertPEM, clientCertPEM, clientKeyPEM []byte, err error) {
	privateKey, certificate, caCerts, err := pkcs12.DecodeChain(pfxData, "")
	if err != nil {
		return nil, nil, nil, err
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, nil, nil, err
	}

	caCertPEM = pem.EncodeToMemory(&pem.Block{Type: certificateType, Bytes: caCerts[0].Raw})
	clientCertPEM = pem.EncodeToMemory(&pem.Block{Type: certificateType, Bytes: certificate.Raw})
	clientKeyPEM = pem.EncodeToMemory(&pem.Block{Type: privateKeyType, Bytes: privBytes})

	return
}

func ParseCertPEM(certPEM []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, ErrPEMInvalid
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func ParsePrivatePEM(privPEM []byte) (any, error) {
	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, ErrPEMInvalid
	}

	return x509.ParsePKCS8PrivateKey(block.Bytes)
}
