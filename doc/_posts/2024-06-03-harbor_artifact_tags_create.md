---
title: harbor artifact tags create
date: 2024-06-03 22:01:02 +0530
categories: [harbor,artifact,tags,create]
tags: [harbor,artifact,tags,create]
permalink: harbor/artifact/tags/create
---
## harbor artifact tags create

Create a tag of an artifact

```bash
harbor artifact tags create [flags]
```

### Examples

```bash
harbor artifact tags create <project>/<repository>/<reference> <tag>
```

### Options

```bash
  -h, --help   help for create
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor artifact tags](harbor)	 - Manage tags of an artifact

