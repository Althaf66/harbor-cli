---
title: harbor artifact scan
date: 2024-06-03 22:01:02 +0530
categories: [harbor,artifact,scan]
tags: [harbor,artifact,scan]
permalink: harbor/artifact/scan
---
## harbor artifact scan

Scan an artifact

### Synopsis

Scan an artifact in Harbor Repository

### Examples

```bash
harbor artifact scan start <project>/<repository>/<reference>
```

### Options

```bash
  -h, --help   help for scan
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor artifact]()	 - Manage artifacts
* [harbor artifact scan start](scan/start)	 - Start a scan of an artifact
* [harbor artifact scan stop](scan/stop)	 - Stop a scan of an artifact

