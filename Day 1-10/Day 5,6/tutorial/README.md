# Access here 
[Nanoc website deployment link](https://bit.ly/3sDtxXH)

# Nanoc Website Deployment to S3

This repository contains the code and configuration for deploying a Nanoc website to an AWS S3 bucket using GitHub Actions. The deployment is triggered on each push to the `main` branch.

## Prerequisites

Before using this repository, make sure you have the following:

1. **AWS Account and S3 Bucket:**
 - Create an AWS account if you don't have one.
 - Set up an S3 bucket to host the Nanoc website.

2. **IAM Credentials:**
 - Generate IAM credentials with the necessary permissions. Refer to the IAM policy in the repository for required permissions.

3. **GitHub Repository:**
 - Have a GitHub repository containing the Nanoc website source code.

## GitHub Actions Workflow

The GitHub Actions workflow defined in this repository performs the following tasks:

1. **Checkout Repository:**
 - Checks out the source code from the repository.

2. **Set up Ruby and Nanoc:**
 - Installs Ruby, Bundler, Nanoc, and other required gems.

3. **Build Nanoc Website:**
 - Navigates to the `tutorial` directory and compiles the Nanoc website.

4. **Set AWS Credentials:**
 - Configures AWS credentials for the GitHub Actions workflow using the provided secrets.

5. **Push to S3:**
 - Syncs the compiled Nanoc website in the `tutorial/output/` directory to the specified S3 bucket.

### Workflow Code

```yaml
name: Nanoc Compile and Upload to S3

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Repository
        uses: actions/checkout@v2

      - name: Set up Ruby and Nanoc
        run: |
          sudo apt-get update
          sudo apt-get install -y ruby-full build-essential zlib1g-dev
          sudo gem install bundler
          sudo gem install nanoc
          sudo gem install adsf
          sudo gem install kramdown

      - name: Build Nanoc Website
        run: |
          ls
          cd tutorial && nanoc

      - name: Set AWS credentials
        uses: aws-actions/configure-aws-credentials@v1
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ap-south-1

      - name: Push to S3
        run: aws s3 sync tutorial/output/ s3://docsite-github-action
```

S3 Bucket Policy
----------------

This is the S3 bucket policy:

```{
    "Version": "2008-10-17",
    "Id": "PolicyForPublicWebsiteContent",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::docsite-github-action/*"
        }
    ]
}
```

### Instructions

-   Replace placeholder values in the S3 bucket policy as needed for your specific use case.
-   Apply this S3 bucket policy to the `docsite-github-action` S3 bucket.
-   Ensure that the policy is configured securely for controlled access to the specified resources.

IAM User Policy
---------------

Add the following IAM user policy to the user created for GitHub Actions:


```{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AccessToGetBucketLocation",
            "Effect": "Allow",
            "Action": [
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::*"
            ]
        },
        {
            "Sid": "AccessToWebsiteBuckets",
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:PutObjectAcl",
                "s3:GetObject",
                "s3:ListBucket",
                "s3:DeleteObject"
            ],
            "Resource": [
                "arn:aws:s3:::docsite-github-action",
                "arn:aws:s3:::docsite-github-action/*"
            ]
        }
    ]
}
```

Adding Secrets to GitHub Repository
===================================

1.  Navigate to Settings:

    -   In your GitHub repository, go to the Settings tab.
2.  Access Secrets:

    -   Click on Secrets in the left sidebar.
3.  Add AWS_ACCESS_KEY_ID:

    -   Click on New secret.
    -   Enter `AWS_ACCESS_KEY_ID` as the secret name.
    -   Paste your AWS Access Key ID as the value.
    -   Click on Add secret.
4.  Add AWS_SECRET_ACCESS_KEY:

    -   Again, click on New secret.
    -   Enter `AWS_SECRET_ACCESS_KEY` as the secret name.
    -   Paste your AWS Secret Access Key as the value.
    -   Click on Add secret.

Now, you've securely added your AWS credentials as secrets to your GitHub repository.
# Give a star to repo 
<img src="https://github.com/ratnesh-maurya.png" alt="Your Name" width="100" height="100" style="border-radius: 50%">
