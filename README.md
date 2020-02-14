# SesLambda

Below is the brief explanation:

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This instructions file
├── application                 <-- Source code for a lambda function
│   ├── main.go                 <-- Lambda function code
│   └── main_test.go            <-- Unit tests
└── template.yml                <-- deploy template
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Packaging and deployment

Build before packaging:
```bash
make build
```

The command will package and deploy your application to AWS.
To package this application, run the following:

```bash
sam package \
--template-file tempalte.yml \
--s3-bucket BUCKET_NAME \
--output-template-file packaged.yml
--profile DEPLOY_USER
```

```bash
sam deploy \
--template-file packaged.yml \
--stack-name stg-email-bounce \
--capabilities CAPABILITY_IAM \
--profile DEPLOY_USER
```

Example is below.
BUCKET_NAME is stg-email-bounce-bucket-emailbouncebucket-1fkmbcmm3yz51,
DEPLOY_USER is us-east-1-user, and STACK_NAME is stg-email-bounce

```bash
sam package \
--template-file template.yml \
--s3-bucket stg-email-bounce-bucket-emailbouncebucket-1fkmbcmm3yz51 \
--output-template-file packaged.yml \
--profile us-east-1-user
```

```bash
sam deploy \
--template-file packaged.yml  \
--stack-name stg-email-bounce \
--capabilities CAPABILITY_IAM \
--profile us-east-1-user
```

### Testing

Run the following command to run our tests:

```shell
go test -v ./application/
```