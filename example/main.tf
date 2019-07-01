terraform {
  required_version = "= 0.11.8"
}

provider "dynadot" {
  api_key = "your_api_key_here"
}

resource "dynadot_domain_nameservers" "testcom_ns" {
  domain = "test.com"
  nameservers = [
    "ns-1.test.com",
    "ns-2.test.org"
  ]
}