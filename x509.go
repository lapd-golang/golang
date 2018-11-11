     template := &x509.Certificate {
               IsCA : true,
               BasicConstraintsValid : true,
               SubjectKeyId : []byte{1,2,3},
               SerialNumber : big.NewInt(1234),
               Subject : pkix.Name{
                         Country : []string{"Earth"},
                         Organization: []string{"Mother Nature"},
               },
               NotBefore : time.Now(),
               NotAfter : time.Now().AddDate(5,5,5),
               // see http://golang.org/pkg/crypto/x509/#KeyUsage
               ExtKeyUsage : []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
               KeyUsage : x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
      }

      // generate private key
      privatekey, err := rsa.GenerateKey(rand.Reader, 2048)

      if err != nil {
        fmt.Println(err)
      }

      publickey := &privatekey.PublicKey

      // create a self-signed certificate. template = parent
      var parent = template
      cert, err := x509.CreateCertificate(rand.Reader, template, parent, publickey,privatekey)

      if err != nil {
         fmt.Println(err)
      }
