# testWarningHook

simple (and very dirty) webhook with warnings. This was only to get an overview of warnings on CLI/openshift console.


run ```go run main.go``` and then  ```curl -k https://localhost:8443 -H "Content-Type: application/json" -d @request.json```


you need your own ssl certs and to update the registration hook to include its signature so kubernetes api trusts it
