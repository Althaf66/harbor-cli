---
title: harbor registry update
date: 2024-06-03 22:01:02 +0530
categories: [harbor,registry,update]
tags: [harbor,registry,update]
permalink: harbor/registry/update
---
## harbor registry update

update registry

```bash
harbor registry update [flags]
```

### Options

```bash
      --credential-access-key string      Access key, e.g. user name when credential type is 'basic'
      --credential-access-secret string   Access secret, e.g. password when credential type is 'basic'
      --credential-type string            Credential type, such as 'basic', 'oauth'
      --description string                Description of the registry
  -h, --help                              help for update
      --insecure                          Whether or not the certificate will be verified when Harbor tries to access the server (default true)
      --name string                       Name of the registry
      --type string                       Type of the registry
      --url string                        Registry endpoint URL
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor registry]()	 - Manage registries

