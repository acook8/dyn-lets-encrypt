REM This script would be called with the parameters <target domain> <record name> <record value> <zone id (optionally)>

REM this example then calls a go application forwarding all the arguments

%~dp0dynDns.exe %* cleanup