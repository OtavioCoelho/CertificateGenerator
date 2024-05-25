package main

import (
	"CertificateGenerator/Services"
)

func main() {
	certInfo, err := Services.GetInfo()
	if err != nil {
		println(err.Error())
		return
	}
	certInfo.Output = "output"

	certificateTemplate := Services.CreateTemplate(*certInfo)

	certDER, keys, err := Services.Generate(certificateTemplate, certInfo.Output)
	if err != nil {
		println(err.Error())
		return
	}
	Services.ConvertPfx(certDER, keys, certInfo.PfxPassword, certInfo.Output)
}
