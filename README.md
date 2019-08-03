## terraform-provider-nfs
  
A Terraform provider for [NearlyFreeSpeech.net](https://nearlyfreespeech.net)'s API. This provider is still under development and highly unstable. Please feel free to make a pull request if you would like to contribute.

[Api Docs](https://members.nearlyfreespeech.net/wiki/API/Introduction)

You need the following information to use this provider:

- NFS ACCOUNT ID - can be found on the web UI on the accounts tab
- NFS_API_KEY - need to be manually requested via an assistance request
- NFS LOGIN - your username

# Example Usage in Terraform

To create a DNS record of type TXT for the domain "aareet.com", you could do the following:
<code>
resource "nfs_dns_record" "new-dns-record" {
    domain = "aareet.com"
    name = "tftest"
    type = "TXT"
    data = "I created this record using Terraform! Woo!"
}
<code>

# TODO
* Read existing DNS records
* Remove a DNS record