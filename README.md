<!-- markdownlint-disable first-line-h1 no-inline-html -->
<a href="https://terraform.io">
    <img src=".github/terraform_logo.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform AWS Provider

[![Forums][discuss-badge]][discuss]

[discuss-badge]: https://img.shields.io/badge/discuss-terraform--aws-623CE4.svg?style=flat
[discuss]: https://discuss.hashicorp.com/c/terraform-providers/tf-aws/

The [AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs) allows [Terraform](https://terraform.io) to manage [AWS](https://aws.amazon.com) resources.

- [Contributing guide](https://hashicorp.github.io/terraform-provider-aws/)
- [Quarterly development roadmap](ROADMAP.md)
- [FAQ](https://hashicorp.github.io/terraform-provider-aws/faq/)
- [Tutorials](https://learn.hashicorp.com/collections/terraform/aws-get-started)
- [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-aws/)
- [gitter](https://gitter.im/hashicorp-terraform/Lobby)
- [Google Groups](http://groups.google.com/group/terraform-tool)

_**Please note:** We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform AWS Provider, please responsibly disclose it by contacting us at security@hashicorp.com._



In IAM role you need to create eks policy 
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "eksadministrator",
            "Effect": "Allow",
            "Action": "eks:*",
            "Resource": "*"
        }
    ]
}

using this u can access iam role 

commands:
  - export AWS_ACCESS_KEY_ID=$TF_VAR_AWS_ACCESS_KEY_ID
  - export AWS_SECRET_ACCESS_KEY=$TF_VAR_AWS_SECRET_ACCESS_KEY

for installing terraform google it
