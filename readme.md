# giti

`giti` is a minimal Git-like version control system built in Go.  
It focuses on implementing the core features of Git, especially those needed to interact with GitHub.

> ⚙️ This project is not meant to replace Git — it's a learning tool to deeply understand how Git works under the hood.

## ✨ Features

- Initialize a repository (`giti init`)
- Add files to an index (`giti add`)
- Create commits (`giti commit`)
- View commit logs (`giti log`)
- Inspect index structure
- Basic GitHub-style SHA1 object storage

## 🧠 Why?

To build a strong mental model of:

- Git’s object model (blobs, trees, commits)
- Hashing with SHA1
- How Git tracks file content, not file names
- How `.git/index` and `.git/objects` work internally

## 📦 Installation

```bash
git clone https://github.com/programming-is-so-funny/giti.git
cd giti
go build -o giti
```

## 🚀 Usage

```bash
./giti init             # Initialize a new repository
./giti add <file>       # Add a file to the staging area
./giti commit -m "..."  # Commit staged files
./giti log              # Show commit logs
```

## 🏗 Project Structure

```bash
giti/
├─- cmd/        # CLI entry points (init, add, commit, etc.)
├── core/       # Core Git logic (index, object store, SHA1)
├── .git/       # Simulated Git repository (created at runtime)
└── main.go     # Entry point for the binary
```

Made with ❤️ for learning Git internals.