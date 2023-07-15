# aws-operations-poc
Creating and managing AWS resources using `aws-sdk-go`

# RUN
* Create a `secrets.go` file in `aws-operations-poc/aws` directory and provider your credential variables
    ```go
  package aws

  const (
  AWS_ACCESS_KEY_ID     = "y0UrAcc3$$1D"
  AWS_SECRET_ACCESS_KEY = "y0Ur$3CR3tK3y"
  )
  ```