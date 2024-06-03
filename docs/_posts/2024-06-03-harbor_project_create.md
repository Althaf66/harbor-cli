---
title: harbor project create
date: 2024-06-03 22:01:02 +0530
categories: [harbor,project,create]
tags: [harbor,project,create]
permalink: harbor/project/create
---
## harbor project create

create project

```bash
harbor project create [flags]
```

### Options

```bash
  -h, --help                   help for create
      --name string            Name of the project
      --proxy-cache            Whether the project is a proxy cache project
      --public                 Project is public or private (default true)
      --registry-id string     ID of referenced registry when creating the proxy cache project
      --storage-limit string   Storage quota of the project (default "-1")
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor project]()	 - Manage projects and assign resources to them

