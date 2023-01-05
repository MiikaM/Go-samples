# Description

This project is a **VERY** simple aws lambda function.

# Try it out

***NOTE!*** This assumes that you have Golang installed on your device

You can try the project by cloning the project and moving to "aws-lambda" directory and building the main.go file by running *"go build main.go"* command in the directory. After this you can either use [aws cli](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-awscli.html) to deploy the function or you can package the main.exe file to a .zip file and uploading it to your aws lambda function.

***Note!*** Be sure to change the aws function handler to "**main**" since by default it is **"hello"**. The **main** handler points to the main function of the application.

This was done with the help of AWS lambda documentation:  

- [Handler](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [Getting started](https://docs.aws.amazon.com/lambda/latest/dg/getting-started.html)