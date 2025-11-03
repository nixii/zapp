# zapp
A password manager meant to be secure, local, yet allows you to connect devices together to sync your information.
Additionally, you may store secrets than just your passwords -- including your 2FA recovery codes.

## IMPORTANT
This is very much a work in progress!
> This means this is NOT meant to be used currently... please don't...

## SECURITY NOTES
1. This is a personal project and isn't currently meant for actual use. It is very possible it has security vulnerabilities. This is more meant to see what is possible to improve password managers for the future.

2. This uses a self-signed certificate. It is NOT meant to be used by other computers. It just has the certificate to make the TLS part happy, as encryption is important. It is assumed that, if you put your master password into the server, you are the one accessing it. Maybe not the safest, will probably be changed at some point.