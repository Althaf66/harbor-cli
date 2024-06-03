---
title: harbor repo info
date: 2024-06-03 22:01:02 +0530
categories: [harbor,repo,info]
tags: [harbor,repo,info]
permalink: harbor/repo/info
---
## harbor repo info

Get repository information

### Synopsis

Get information of a particular repository in a project

```bash
harbor repo info [flags]
```

### Examples

```bash
  harbor repo info <project_name>/<repo_name>
```

### Options

```bash
  -h, --help   help for info
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor repo]()	 - Manage repositories

