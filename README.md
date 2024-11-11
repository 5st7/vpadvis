# vpadvis
vpadvis is a command-line tool designed to aggregate and display recommendations from the Vertical Pod Autoscaler (VPA) in Kubernetes.

## Installation
```bash
go install github.com/5st7/vpadvis@latest

```
Ensure that your $GOPATH/bin is included in your system's PATH to execute the binary directly.


## Usage
To display recommendations in plaintext format:
```bash
vpadvis recommend --namespace default --format plaintext

```

To display recommendations in Markdown format:

```bash
vpadvis recommend --namespace default --format markdown
```
