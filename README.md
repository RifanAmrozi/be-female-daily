# BE-FEMALE-DAILY

Backend service for Female Daily flash sale event app built with Go and Gin framework.

---

## 📦 Project Structure
```
---
BE-FEMALE-DAILY/
├── controller/ # HTTP handler functions
│ └── brand.go
├── data/ # Temporary seeded/mock data
│ └── seed20250721.go
├── model/ # Data models
│ └── brand.go
├── route/ # Route definitions
│ └── route.go
├── main.go # Entry point
├── go.mod
├── go.sum
└── .gitignore
---
```

## 🚀 Getting Started

### 1. Install Go
```bash
brew install go     # for Mac
```

### 2. Clone the repo
```bash
git clone https://github.com/yourusername/be-female-daily.git
cd be-female-daily
```

### 3. Run the project

```bash
go run main.go
```
API will be available at: http://localhost:8080/api/v1/

### ✅ Git Commit Convention
https://www.conventionalcommits.org/

Use Conventional Commits example:
feat: add brand list endpoint
fix: correct seed data typo
chore: update dependencies

### Pull Request Flow
#### Create new branch:
```
git checkout -b feat/brand-endpoint
```

#### Commit changes:
```
git add .
git commit -m "feat: add GET /brands endpoint"
```

#### Push and create PR:
```
git push origin feat/brand-endpoint
```

#### Open pull request on GitHub, assign reviewer.


### Rebase before merge
```
git checkout feat/brand-endpoint
git fetch origin
git rebase origin/main
# resolve any conflicts, then:
git push --force
```
