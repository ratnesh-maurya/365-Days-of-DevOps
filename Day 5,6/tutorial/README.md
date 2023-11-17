# Access here 
[Nanoc website deployl link](https://bit.ly/3sDtxXH)
  
# Nanoc Website Deployment to S3

This repository contains the code and configuration for deploying a Nanoc website to an AWS S3 bucket using GitHub Actions. The deployment is triggered on each push to the `main` branch.

## Prerequisites

Before using this repository, make sure you have the following:

1. An AWS account with an S3 bucket created for hosting the website.
2. AWS IAM credentials with the necessary permissions (see IAM policy in the repository).
3. GitHub repository containing the Nanoc website source code.

## GitHub Actions Workflow

The GitHub Actions workflow defined in this repository performs the following tasks:

1. **Checkout Repository**: Checks out the source code from the repository.

2. **Set up Ruby and Nanoc**: Installs Ruby, Bundler, Nanoc, and other required gems.

3. **Build Nanoc Website**: Navigates to the `tutorial` directory and compiles the Nanoc website.

4. **Set AWS credentials**: Configures AWS credentials for the GitHub Actions workflow using the provided secrets.

5. **Push to S3**: Syncs the compiled Nanoc website in the `tutorial/output/` directory to the specified S3 bucket.

### Workflow Code

```yaml
name: Nanoc Compile and Upload to S3

on:
  push:
    branches:
      - main

jobs:
  build:
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
# Nanoc Website Deployment to AWS S3

This guide provides instructions for deploying a Nanoc website to Amazon Simple Storage Service (S3) using GitHub Actions.

## Prerequisites

1. GitHub account
2. AWS account with S3 bucket
3. Nanoc website project
4. Basic knowledge of Markdown and JSON

## Steps

1. **Create an AWS IAM User and Access Key:**

   a. In the AWS Management Console, navigate to the IAM service.

   b. Click "Users" and then click "Add user."

   c. Provide a user name and select "Programmatic access" as the access type.

   d. Attach the "AmazonS3FullAccess" IAM policy to the user.

   e. Click "Next: Permissions" and then click "Next: Tags."

   f. Click "Next: Review" and then click "Create user."

   g. Note the Access Key ID and Secret Access Key displayed on the screen. Save these credentials securely.

2. **Create an S3 Bucket:**

   a. In the AWS Management Console, navigate to the S3 service.

   b. Click "Create bucket" and provide a bucket name.

   c. Select the appropriate region for your website.

   d. Click "Create" to create the bucket.

3. **Create a GitHub Repository:**

   a. Create a new GitHub repository to store your Nanoc website project.

   b. Clone the repository to your local machine.

4. **Configure GitHub Actions Workflow:**

   a. Create a `.github/workflows/deploy.yml` file in the root directory of your repository.

   b. Add the following YAML code to the file, replacing the placeholder values with your specific details:

```yaml
name: Deploy Nanoc Website to S3

on:
  push:
    branches:
      - main

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout repository
        uses: actions/checkout@v3

      - name: Set up Node.js
        uses: actions/setup-node@v3
        with:
          node-version: 16

      - name: Install dependencies
        run: npm install

      - name: Build Nanoc website
        run: npm run build

      - name: Configure AWS credentials
        env:
          AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
          AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        run: |
          echo "AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID" >> $GITHUB_ENV
          echo "AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY" >> $GITHUB_ENV

      - name: Deploy to S3
        uses: actions/upload-artifact@v2
        with:
          name: website-build
          path: ./public

      - name: Deploy to S3 bucket
        uses: s3-actions/s3-sync@v2
        with:
          bucket: "your-s3-bucket-name"
          region: "your-s3-bucket-region"
          src: ./website-build/*
          acl: "public-read"
```


5. **Add Secrets to GitHub Repository:**

   a. In your GitHub repository, navigate to the "Settings" tab.

   b. Click "Secrets" and then click "New secret."

   c. Enter "AWS_ACCESS_KEY_ID" as the secret name and paste your AWS Access Key ID as the value.

   d. Click "Add secret" and repeat the process for "AWS_SECRET_ACCESS_KEY," using your AWS Secret Access Key as the value.

6. **Commit and Push Changes:**

   a. Commit the changes you made to the workflow file and the secrets to your local repository.

   b. Push the changes to the main branch of your GitHub repository.

## Deployment Verification

Once you push the changes to the main branch, the GitHub Actions workflow will be triggered, building the Nanoc website and deploying it to the specified S3 bucket. You can verify the deployment by visiting the S3 bucket URL. The website content should be publicly accessible.
