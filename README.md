# CertificateGenerator

The **CertificateGenerator** CertificateGenerator is a project developed in Go that facilitates the generation of PEM certificates and their conversion to the PFX format. It offers a simple and efficient solution for generating and converting SSL/TLS certificates.

## Features

- Self-signed PEM certificate generation.
- Conversion of PEM certificates to the PFX format.

## How to Use

1. Clone the repository:

```bash
git clone https://github.com/OtavioCoelho/CertificateGenerator.git
```

2. Navigate to the project directory:

```
cd CertificateGenerator
```

3. Execute o arquivo main.go:

```
go run main.go
```

4. Provide the requested data to generate the certificate.
5. Access the `output` folder to check the generated certificate. This will initiate the generation of a self-signed PEM certificate and its conversion to the PFX format. After successful execution, the resulting certificate will be available in the project directory.
6. 
```
cd output
```

## Requirements

Make sure you have Go installed on your machine. If you don't have it yet, you can download and install it from the [official Go website.](https://golang.org/).

## Contribuindo

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

This README was generated via chat GPT.