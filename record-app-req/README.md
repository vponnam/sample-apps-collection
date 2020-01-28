## This is a sample application that will help to analyze http request headers and body
This sample application will log all the request headers and body that the application receives from an external client in the application logs as well as responding the same to the web request.

### Deploying to cloud foundry
git clone && cd record-app-req
cf push

## Using the app to inspect a client request
The application will basically capture/log any request information(headers and body) received by the app.  

For example: Sending a simple web request to the app using curl will log as below:

> `curl https://app-url`

Tip: `cf logs simple-request-viewer --recent`

```
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT GET / HTTP/1.1
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT Host: REDACTED
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT Accept: */*
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT B3: c81bc69ac5346b6b-c81bc69ac5346b6b
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT User-Agent: curl/7.47.0
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-B3-Spanid: c81bc69ac5346b6b
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-B3-Traceid: c81bc69ac5346b6b
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Cf-Applicationid: fa85827b-1e5c-4268-a1d2-394322a39455
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Cf-Instanceid: cb28f9e2-d4f5-4b85-62c7-900d
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Cf-Instanceindex: 0
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Forwarded-For: 10.193.78.6, 10.193.78.250
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Forwarded-Proto: http
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Request-Start: 1580252645567
   2020-01-28T18:04:05.57-0500 [APP/PROC/WEB/0] OUT X-Vcap-Request-Id: f89ea965-5d91-4100-47b8-4084
```

## Use cases where this type of analysis is helpful
- You as as operator or developer wants to capture the exact request that was sent by external client for troubleshooting purposes.
- Wanted to check if all the expected request headers are made it to the app after passing through intermediate network hops/load-balancers.
- mTLS related triage is a good example where ideally if the mTLS is successful between an external client and PCF, the app is PCF should see the client_certificate in [XFCC header](https://docs.cloudfoundry.org/concepts/http-routing.html#forward-client-cert) for further validation purposes.


## mTLS troubleshooting:
Step 1:  
Push this sample application to your foundation.

Step 2:  
Initiate a client side mTLS call to your PCF app, below is an example.

```
curl https://app-url --cert=cert.pem --key=key.pem
```

Above curl will work as long the provided certs are issue by a well known CA, else you have 2 options:
1. Adding you internal CA+intermediate chain to bosh trusted certs section as documented [here](https://docs.pivotal.io/pivotalcf/2-4/customizing/trusted-certificates.html).
2. Pass the ca that issued your certs as below.
```
curl https://app-url --cert=cert.pem --key=key.pem --cacert ca.pem
```

As long as everything else at the load-balancing layer is configured properly, the app in PCF will see the contents of public cert(XFCC header) that was sent by the client during mTLS handshake.
