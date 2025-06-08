### Chaos Data Comparator

A utility to compare outputs from [ProjectDiscovery's Chaos](https://chaos.projectdiscovery.io/) program across different days. It helps identify **added**, **removed**, and optionally **unchanged** URLs for each program.

---

## 📦 Features

- Compares two `chaos-output` folders line-by-line.
- Organizes comparison results by program name.
- Generates:
  - `added.txt`
  - `removed.txt`
  - `unchanged.txt` (optional)
- Supports **verbose output**, **custom output directory**, and **filtering unchanged entries**.

---

## Installation

```
go install github.com/computerauditor/compare-chaos@latest
```

Can Build a go binary using : 
```
go build -o compare-chaos compare-chaos.go
```
Move to /usr/local/bin [OPTIONAL]

```
mv /go/bin/compare-chaos /bin
```

OR
```
mv /go/bin/compare-chaos /usr/local/bin
```

## 📂 Folder Structure Example

```

chaos-output-2025-06-07/
├── mastercard.nl.txt
└── adobe.com.txt

chaos-output-2025-06-08/
├── mastercard.nl.txt
└── adobe.com.txt

results/
├── mastercard.nl/
│   ├── added.txt
│   ├── removed.txt
│   └── unchanged.txt
└── adobe.com/
├── added.txt
├── removed.txt
└── unchanged.txt

````

---

## 🚀 Usage

```bash
go run compare-chaos.go [options]
````
EXAMPLE:
```bash
./compare-chaos -n chaos-data/chaos-output-2025-06-08 -p chaos-data/chaos-output-2025-06-07/ -output output/results -nu
````

### 📋 Flags

| Flag(s)                      | Description                                                  |
| ---------------------------- | ------------------------------------------------------------ |
| `-n`, `--new`, `--today`     | Path to today’s `chaos-output` folder                        |
| `-p`, `--old`, `--yesterday` | Path to yesterday’s `chaos-output` folder                    |
| `-o`, `--output`             | Output directory for comparison results (default: `results`) |
| `-v`, `--verbose`            | Enable verbose output with stats                             |
| `--nu`, `--no-unchanged`     | Skip writing `unchanged.txt` files                           |
| `-h`, `--help`               | Show usage help                                              |

---

### ✅ Example

```bash
go run compare-chaos.go \
  -n chaos-output-2025-06-08 \
  -p chaos-output-2025-06-07 \
  -o results \
  -v \
  --nu
```

> Compares the latest `chaos-output` with the previous day and stores added/removed entries per program in the `results/` folder. Skips writing unchanged URLs.

---

## 🛠 Requirements

* Go 1.18+
* Two valid folders containing `.txt` files from the `chaoser` tool or ProjectDiscovery Chaos outputs.

---

## 📄 License

MIT License — use freely, credit appreciated. Contributions welcome.

---

## 🤝 Acknowledgements

* [ProjectDiscovery](https://projectdiscovery.io/)
* Inspired by the [chaoser](https://github.com/computerauditor/chaoser) automation tool

```

---

Let me know if you’d like:
- Badge support (Go version, License, etc.)
- GIF demo of usage
- CLI release versioning instructions

Happy reconning! 🕵️‍♂️
```
