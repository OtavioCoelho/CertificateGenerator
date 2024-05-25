package Services

import (
	"CertificateGenerator/Models"
	"CertificateGenerator/Utilities"
	"CertificateGenerator/Validates"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"software.sslmate.com/src/go-pkcs12"
)

func GetInfo() (*Models.CertificateInfo, error) {
	var certificationInfo Models.CertificateInfo

	if err := Utilities.PromptAndScan("Enter Country (e.g., US): ", &certificationInfo.Country); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter Organization (e.g., My Company): ", &certificationInfo.Organization); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter Organizational Unit (e.g., IT): ", &certificationInfo.OrganizationalUnit); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter Common Name (e.g., My Company Certificate): ", &certificationInfo.CommonName); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter URL (e.g., example.com): ", &certificationInfo.Url); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter Years Validating (e.g., 2): ", &certificationInfo.YearsValidate); err != nil {
		return nil, err
	}
	if err := Utilities.PromptAndScan("Enter Password for PFX file: ", &certificationInfo.PfxPassword); err != nil {
		return nil, err
	}

	return &certificationInfo, nil
}

func CreateTemplate(certInfo Models.CertificateInfo) *x509.Certificate {
	//Verify year
	year := Validates.InputInt(certInfo.YearsValidate)
	if year <= 0 {
		year = 1
	}

	// Prepare the certificate template
	certificateTemplate := x509.Certificate{
		SerialNumber:       Utilities.GenerateNumericUUID(),
		SignatureAlgorithm: x509.SHA512WithRSA,
		Subject: pkix.Name{
			Country:            Validates.InputArrayString(certInfo.Country),
			Organization:       Validates.InputArrayString(certInfo.Organization),
			OrganizationalUnit: Validates.InputArrayString(certInfo.OrganizationalUnit),
			CommonName:         certInfo.CommonName,
		},
		URIs:                  Validates.ValidateInputArrayUrl(certInfo.Url),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(year, 0, 0), // 1 Year default
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	return &certificateTemplate
}

func Generate(template *x509.Certificate, pathOutput string) ([]byte, *rsa.PrivateKey, error) {
	keys, err := generateKeys()
	if err != nil {
		return nil, nil, err
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &keys.PublicKey, keys)
	if err != nil {
		fmt.Printf("Failed to create certificate: %s\n", err)
		return nil, nil, err
	}

	err = Utilities.CheckDiretory(pathOutput)
	if err != nil {
		return nil, nil, err
	}

	certOut, err := os.Create(".\\" + pathOutput + "\\certificate.pem")
	if err != nil {
		fmt.Printf("Failed to open cert.pem for writing: %s\n", err)
		return nil, nil, err
	}
	defer func(certOut *os.File) {
		err := certOut.Close()
		if err != nil {
			return
		}
	}(certOut)

	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Certificate save  in .\\" + pathOutput + "\\certificate.pem")

	keyOut, err := os.Create(".\\" + pathOutput + "\\privateKey.pem")
	if err != nil {
		return nil, nil, err
	}
	defer func(keyOut *os.File) {
		err := keyOut.Close()
		if err != nil {
			return
		}
	}(keyOut)

	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(keys)})
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Private key save  in .\\" + pathOutput + "\\privateKey.pem")

	return certDER, keys, nil
}

func generateKeys() (*rsa.PrivateKey, error) {
	keys, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Printf("Failed to generate private key: %s\n", err)
		return nil, err
	}
	return keys, nil
}

func ConvertPfx(certDER []byte, keys *rsa.PrivateKey, pfxPassword string, pathOutput string) {
	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		fmt.Printf("Failed to parse certificate: %s\n", err)
		return
	}

	// Converter para .pfx
	pfxData, err := pkcs12.Legacy.Encode(keys, cert, nil, pfxPassword) //OK
	if err != nil {
		fmt.Printf("Failed to create PFX data: %s\n", err)
		return
	}

	// Salvar o arquivo .pfx
	err = SavePfxToFile(".\\"+pathOutput+"\\certificate.pfx", pfxData)
	if err != nil {
		fmt.Printf("Failed to save pfx: %s\n", err)
	}

	fmt.Println("Certificate and private key save in  .\\" + pathOutput + "\\certificate.pfx")
}

func SavePfxToFile(filename string, pfxData []byte) error {
	pfxOut, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to open certificate.pfx for writing: %s\n", err)
		return err
	}
	_, err = pfxOut.Write(pfxData)
	if err != nil {
		return err
	}

	err = pfxOut.Close()
	if err != nil {
		return err
	}

	return nil
}
