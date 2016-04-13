# gwn-oauth

https://aaronparecki.com/2012/07/29/2/oauth2-simplified

## Certs
```bash
openssl ecparam -genkey -name secp521r1 -noout -out private.pem 
openssl ec -in private.pem -pubout -out public.pem
```