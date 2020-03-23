package chshare

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func loadCerts(certPaths []string) (*x509.CertPool,  error){
	certPool := x509.NewCertPool()
	for _, certPath := range certPaths {
		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			return nil, err
		}
		if !certPool.AppendCertsFromPEM(cert) {
			return nil, fmt.Errorf("error while parsing certificate %s", cert)
		}
	}
	return certPool, nil
}

func NewTLSConfig(certPEM string, keyPEM string, rootCAs *[]string, clientCAs *[]string) (*tls.Config, error) {
	var tlsConfig = new(tls.Config)

	if rootCAs != nil {
		rootCertPool, err := loadCerts(*rootCAs)
		if err != nil {
			return nil, err
		}
		tlsConfig.RootCAs = rootCertPool
	}

	if clientCAs != nil {
		clientCertPool, err := loadCerts(*clientCAs)
		if err != nil {
			return nil, err
		}
		tlsConfig.ClientCAs = clientCertPool
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	cert, err := tls.LoadX509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	tlsConfig.Certificates = []tls.Certificate{cert}
	tlsConfig.MinVersion = tls.VersionTLS12

	return tlsConfig, nil
}