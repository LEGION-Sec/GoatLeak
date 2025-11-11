# GoatLeak 🐐
Enhanced Gitleaks with severity levels, exception management &amp; CI/CD quality gates for enterprise security teams.

 ▄████  ▒█████   ▄▄▄     ▄▄▄█████▓ ██▓    ▓█████ ▄▄▄       ██ ▄█▀
 ██▒ ▀█▒▒██▒  ██▒▒████▄   ▓  ██▒ ▓▒▓██▒    ▓█   ▀▒████▄     ██▄█▒ 
▒██░▄▄▄░▒██░  ██▒▒██  ▀█▄ ▒ ▓██░ ▒░▒██░    ▒███  ▒██  ▀█▄  ▓███▄░ 
░▓█  ██▓▒██   ██░░██▄▄▄▄██░ ▓██▓ ░ ▒██░    ▒▓█  ▄░██▄▄▄▄██ ▓██ █▄ 
░▒▓███▀▒░ ████▓▒░ ▓█   ▓██▒ ▒██▒ ░ ░██████▒░▒████▒▓█   ▓██▒▒██▒ █▄
 ░▒   ▒ ░ ▒░▒░▒░  ▒▒   ▓▒█░ ▒ ░░   ░ ▒░▓  ░░░ ▒░ ░▒▒   ▓▒█░▒ ▒▒ ▓▒
  ░   ░   ░ ▒ ▒░   ▒   ▒▒ ░   ░    ░ ░ ▒  ░ ░ ░  ░ ▒   ▒▒ ░░ ░▒ ▒░
░ ░   ░ ░ ░ ░ ▒    ░   ▒    ░        ░ ░      ░    ░   ▒   ░ ░░ ░ 
      ░     ░ ░        ░  ░            ░  ░   ░  ░     ░  ░░  ░   
                                                                  
		⠀⠀⠀⠀⠀⠀⠀⢀⣠⡤⣤⡀⠀⠀⠀⠀⠀⠀⢀⡠⡤⣄⣀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⢀⡴⠊⢉⣑⣽⣕⡈⢢⠀⠀⠀⠀⡰⠉⣢⣿⣃⡉⠉⠢⡀⠀⠀⠀⠀
		⠀⠀⠒⠶⠥⠒⠉⠁⠀⠀⢱⠒⠈⡆⠀⠀⢠⠃⠒⡜⠀⠀⠀⠉⠒⠨⠵⠖⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡧⠬⢽⠀⠀⡼⠭⢬⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢛⣙⠚⡆⠀⡗⣊⣹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣸⡒⠫⢇⡸⠌⠒⣼⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⢀⠤⠤⢤⣒⠉⠁⡏⠀⠙⡦⠤⠤⢴⠞⠁⢸⠌⠉⢒⣤⠤⠤⡄⠀⠀⠀
		⠀⠀⠀⠘⢄⡀⠣⠤⣙⡿⣿⣕⡄⢸⠀⠀⡜⢀⡮⣽⠿⣛⠡⠜⢀⡠⠏⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠉⠁⠀⠀⠈⡛⢿⢸⠀⠀⡇⡼⠛⡇⠀⠀⠈⠉⠁⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⢸⢸⠀⠀⢃⡇⡰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⡇⠆⠀⡜⣠⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹⡔⠧⠴⠃⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣗⠤⠤⣺⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣏⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣻⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀

Creator: LEGION-Sec (shusig2603@gmail.com)


> **Don't Let Secrets Escape the Herd!** 🐐

Enterprise-grade secret detection tool built for modern development workflows. GoatLeak finds secrets like a GOAT (Greatest Of All Time)!

---

## 🚀 Features

- **🎯 Severity-Based Detection** - Critical/High/Medium/Low levels with risk scoring
- **📝 Exception Management** - JSON-based exception system. [File based and not line based]
- **🚦 Quality Gates** - CI/CD pipeline enforcement with customizable thresholds
- **⚙️ Config-Only Operation** - No default rules, explicit configuration required [No Internet Dependency]
- **📊 Multiple Output Formats** - JSON, CSV, SARIF, JUnit with severity integration
- **🔍 Advanced Scanning** - Git history, file system, and real-time protection, Remote Repository Scans
- **🎨 Enterprise Ready** - Built for security teams and development workflows

---

## Overview

GoatLeak is an enterprise-focused secret detection tool that extends the powerful Gitleaks engine with critical enterprise features. While Gitleaks [https://github.com/gitleaks/gitleaks] provides excellent secret detection capabilities, GoatLeak adds the workflow and governance features that security teams need for production environments.


### Why GoatLeak?

| Feature | Gitleaks | GoatLeak |
|---------|----------|----------|
| Severity Levels | ❌ Limited | ✅ Full support (Critical/High/Medium/Low) |
| Exception Management | ❌ Basic | ✅ JSON-based with audit trails |
| Quality Gates | ❌ Manual | ✅ CI/CD pipeline integration |
| Default Configs | ✅ Yes | ❌ Explicit configuration only |
| Risk Scoring | ❌ No | ✅ Built-in risk assessment |


---

## Quick Start

### CLI
Simply download the executable and run it on your system (LINUX and MacOS) with supported command. (Most of the gitleaks command will work here as well)

### Basic Usage

```bash
# Build the binary
go build -o goatleak .

# Simple scan
./goatleak detect --config goatleak-config.toml --source=. -f json -r output.json   (use --help for other report types)

# Scan with quality gates
./goatleak detect --config goatleak-config.toml --source=. --quality-gate Critical=0 --quality-gate High=5 (if scan report has 0 critical and <5 high, it will pass else fail)

# Scan with exceptions
./goatleak detect --config goatleak-config.toml --source=. --exceptions approved-exceptions.json

# Scan remote repository
./goatleak detect --config goatleak-config.toml --source="https://github.com/username/repo.git"

# Scan Specific branch
./goatleak detect --config goatleak-config.toml --source="https://github.com/username/repo.git" --branch="branch-name"

