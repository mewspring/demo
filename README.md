# demo

The aim of this repository is to make it trivial to parse and analyze the `*.dmo` file format of DevilutionX, as used for input event demo recording and replay.

To facilitate this effort, lexers and parses are automatically generated from a [BNF grammar for the demo file format](dmo/demo.tm).

## Installation

```bash
go install github.com/mewspring/demo/cmd/demo_dump@master
```

## Usage

```
demo_dump FILE.dmo
```
