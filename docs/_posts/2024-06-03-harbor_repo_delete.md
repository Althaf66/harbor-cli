---
title: harbor repo delete
date: 2024-06-03 22:01:02 +0530
categories: [harbor,repo,delete]
tags: [harbor,repo,delete]
permalink: harbor/repo/delete
---
## harbor repo delete

Delete a repository

### Synopsis

Delete a repository within a project in Harbor

```bash
harbor repo delete [flags]
```

### Examples

```bash
  harbor repository delete [project_name]/[repository_name]
```

### Options

```bash
  -h, --help   help for delete
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor repo]()	 - Manage repositories

