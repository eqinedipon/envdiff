# envdiff

> Compare `.env` files across environments and highlight missing or mismatched keys.

---

## Installation

```bash
go install github.com/yourusername/envdiff@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/envdiff.git
cd envdiff
go build -o envdiff .
```

---

## Usage

```bash
envdiff [flags] <file1> <file2>
```

### Example

```bash
envdiff .env.development .env.production
```

**Sample output:**

```
Missing in .env.production:
  - DATABASE_URL
  - REDIS_HOST

Mismatched keys:
  ~ LOG_LEVEL: "debug" → "info"

✔ All other keys match.
```

### Flags

| Flag | Description |
|------|-------------|
| `--keys-only` | Only compare key names, ignore values |
| `--quiet` | Suppress output, exit code reflects result |
| `--json` | Output results as JSON |

---

## Why envdiff?

Managing multiple environment files is error-prone. `envdiff` makes it easy to catch configuration drift before it causes issues in staging or production.

---

## Contributing

Pull requests are welcome. Please open an issue first to discuss any significant changes.

---

## License

[MIT](LICENSE) © 2024 yourusername