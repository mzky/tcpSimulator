package common

import (
	"github.com/mzky/tls"
	"path"
)

func Generate(fp string) error {
	var ca tls.CACert
	ca.Cert, ca.Key, _ = tls.GenerateRoot()
	ipArray, _ := GetLocalIPList()
	c, k, _ := ca.GenerateServer(ipArray)

	_ = tls.WritePEM(path.Join(fp, "server.pem"), c)
	_ = tls.WritePEM(path.Join(fp, "server.key"), k)

	_, err := tls.CertificateInfo("server.pem")

	return err
}
