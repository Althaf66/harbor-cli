---
title: harbor artifact tags delete
date: 2024-06-03 22:01:02 +0530
categories: [harbor,artifact,tags,delete]
tags: [harbor,artifact,tags,delete]
permalink: harbor/artifact/tags/delete
---
## harbor artifact tags delete

Delete a tag of an artifact

```bash
harbor artifact tags delete [flags]
```

### Examples

```bash
harbor artifact tags delete <project>/<repository>/<reference> <tag>
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

* [harbor artifact tags](harbor)	 - Manage tags of an artifact

