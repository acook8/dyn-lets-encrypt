# dyn-lets-encrypt
Scripts to automate dns challenges in Dyn DNS for Certify The Web for Windows and Certbot for Linux

## Background
I built this  that would work allow interaction with the Dyn DNS API, and one that would work on both Windows and Linux. I could have done it in either Powershell or Python, but I wanted to avoid installing language runtimes on servers if possible. I had been looking for an exuse to learn Go so I decided this was a great opportunity. One advantage of Go is you can compile it for a variety of operating systems and architectures. And the runtime is included in the single executable created, so no languages have to be installed.

## Building
Before you build the application find the getToken function, and in it replace CUSTOMER_NAME, USER_NAME, and PASSWORD with the appropriate credentials to your Dyn account.

To build a Go application first install Go using instructions found here: 
After Go is installed build the code by running:
```
go build dynSetup.go
```
To build it for a specific operating system/architecture use environment variables. For example to build it for 64-bit Windows run:
```
GOOS=windows GOARCH=amd64 go build dynSetup.go
```
To learn more about this go here:

## Use
### Use with Certify the Web
Certify the Web is a Windows application that can get certificates from Let's Encrypt, update them automatically, and deploy them inside IIS. 

Inside Certify the Web create a new certificate, select the iis site, and add the domain you want. Currently this go application only accepts one domain in the cert so for example choose either *.github.com or github.com, but not both. 

In the Authorization tab choose dns-01 as the challenge type, and as for the DNS Update Method scroll to the top and choose (Use Custom Script). Type in the paths to the createDNS.bat file and cleanupDNS.bat file where it says Create Script Path and Delete Script Path. Make sure the two bat files and the compiled exe are inside the same directory on the computer/server. The Go application has a 10 second sleep in it already, but if you want a longer sleep add a Propagation Delay. Don't put in a DNS Zone ID, it isn't needed and will cause an error.

After you choose your Deployment Preferences go ahead and click test. It is always a good idea to test your settings before requesting a real certificate. Let's Encrypt limits how many errors can occur in a specific time frame, so if you have two many errors in a short amount of time you'll have to wait before requesting a certificate again. Once the test passes go ahead and request a certificate.

### Use with Certbot

## To-Do
* [ ] Finish ReadMe
* [ ] Create bash scripts for Certbot
* [ ] Update ReadMe with Certbot instructions
