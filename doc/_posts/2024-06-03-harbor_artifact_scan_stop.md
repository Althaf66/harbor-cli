---
title: harbor artifact scan stop
date: 2024-06-03 22:01:02 +0530
categories: [harbor,artifact,scan,stop]
tags: [harbor,artifact,scan,stop]
permalink: harbor/artifact/scan/stop
---
## harbor artifact scan stop

Stop a scan of an artifact

### Synopsis

Stop a scan of an artifact in Harbor Repository

```bash
harbor artifact scan stop [flags]
```

### Examples

```bash
harbor artifact scan stop <project>/<repository>/<reference>
```

### Options

```bash
  -h, --help   help for stop
```

### Options inherited from parent commands

```bash
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/user/.harbor/config.yaml")
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output
```

### SEE ALSO

* [harbor artifact scan](harbor)	 - Scan an artifact

