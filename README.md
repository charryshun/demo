# This is a CI/CD demo using github action workflow

## How to access the application?
Direct to url `http://f3a7ddec-us-south.lb.appdomain.cloud/YOUR_NAME` replacing the `YOUR_NAME` to a string to see the response.

## How to verify the CI/CD process?
1. Create a PR against `main` branch with your changes of the [greeting string](https://github.com/ywysuibian/demo/blob/main/main.go#L15)
2. Wait for the the workflow complete
3. Direct to the url again to see if the changes reflected in the reponse

## Sample output for PR #8
```
while true; do curl http://f3a7ddec-us-south.lb.appdomain.cloud/myname ; sleep 2; echo ; done
Greeting: Welcome {myname}!
2021-02-04 07:58:39
Greeting: Welcome {myname}!
2021-02-04 07:58:41
Greeting: Welcome {myname}!
2021-02-04 07:58:44
Greeting: Welcome {myname}!
2021-02-04 07:58:47
Greeting: Welcome {myname}!
2021-02-04 07:58:49
Greeting: Welcome {myname}!
2021-02-04 07:58:52
Greeting: Welcome {myname}!
2021-02-04 07:58:55
Greeting: Welcome to this page myname!
2021-02-04 07:58:57
Greeting: Welcome to this page myname!
2021-02-04 07:59:00
Greeting: Welcome to this page myname!
2021-02-04 07:59:03
Greeting: Welcome to this page myname!
2021-02-04 07:59:06
Greeting: Welcome to this page myname!
2021-02-04 07:59:08
```
