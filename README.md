
# Gitto - AI-Powered Git Commits

Gitto is a CLI tool that generates smart and meaningful Git commit messages using AI. It simplifies your Git workflow by analyzing code changes and automatically suggesting clear and descriptive commit messages.

## 🚀 Features
- **AI-Generated Commit Messages** – Let AI analyze your code changes and generate consistent, meaningful commit messages.
- **Easy Configuration** – Set up your API key for AI-based commit generation.
- **Seamless Git Integration** – Stage and commit changes directly from the CLI.

---

## 🛠️ Installation

### 1. Clone the repository:
```bash
git clone https://github.com/Hemsara/gitto.git
```

### 2. Navigate to the project directory:
```bash
cd gitto
```

### 3. Build the binary:
```bash
go build -o gitto
```

---

## 🌟 Usage

### 1. **Configure API Key**  
Set up the AI key for generating commit messages:
```bash
./gitto config --apikey YOUR_API_KEY
```

---

### 2. **Generate and Commit Changes**  
Let AI create a meaningful commit message and commit changes:
```bash
./gitto commit
```

Example output:
```bash
💡 Generated commit message:
Refactor user authentication flow and improve error handling.

✅ Commit? (y/N):
```

---

## 🎯 Commands

| Command              | Description                                           |
|--------------------- | ----------------------------------------------------- |
| `gitto config`       | Configure Gitto with your API key                      |
| `gitto commit`       | Generate and commit changes with AI-generated message   |
| `gitto help`         | Display help information                                |

---

## ✅ Example

```bash
# Configure API key
./gitto config --apikey "YOUR_API_KEY"

# Generate and commit changes
./gitto commit
```

---

## 💡 How It Works
1. Gitto checks if the current directory is a Git repository.
2. AI generates a commit message based on `git diff` output.
3. Prompts the user to confirm the generated message.
4. Stages and commits changes using the generated message.

---

## 📄 License
This project is licensed under the [MIT License](LICENSE).
