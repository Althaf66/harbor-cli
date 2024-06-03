---
title: harbor artifact scan start
date: 2024-06-03 22:01:02 +0530
categories: [harbor,artifact,scan,start]
tags: [harbor,artifact,scan,start]
permalink: harbor/artifact/scan/start
---
## harbor artifact scan start

Start a scan of an artifact

### Synopsis

Start a scan of an artifact in Harbor Repository

```bash
harbor artifact scan start [flags]
```

### Examples

```bash
harbor artifact scan start <project>/<repository>/<reference>
```

### Options

```bash
  -h, --help   help for start
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor artifact scan](harbor)	 - Scan an artifact

