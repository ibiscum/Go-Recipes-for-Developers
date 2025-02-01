This is how you can create a self-signed certificate using openssl.

1. Create a private key

```
openssl ecparam -name secp521r1 -genkey -noout -out privatekey.pem
```

2. Extract the corresponding public key from the private key

```
openssl ec -in privatekey.pem -pubout > publickey.pem
```

3. Create a certificate signing request. Note the subject--it is not
   provigin a "common name", which is deprecated

```
openssl req -new -sha256  -subj "/C=US"  -key privatekey.pem -out my.csr 
```

4. Create a certificate by signing it using the private key. The
   config.ext provides SANs for localhost

```
openssl x509 -signkey privatekey.pem -in my.csr -extfile config.ext -req -days 3650 -out server.crt
```
