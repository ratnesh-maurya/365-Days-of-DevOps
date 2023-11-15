provider "aws" {
  region     = "ap-south-1"
  access_key = "your-access-key"
  secret_key = "your-secret-key"
}

# Create AWS Instance

resource "aws_instance" "myinstance1" {
  
  ami           = "ami-02e94b011299ef128"
  instance_type = "t3.medium"
  key_name      = "your-key-name"
  count			= 2

  tags = {
    Name = "my_Server"
  }
}