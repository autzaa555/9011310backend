## Getting Started

1. Run Database

```bash
docker compose up -d
```
2. Run Backend

```bash
cd backend
# go mod tidy
# goose -dir ./migrations postgres postgres://devpool:123456789@localhost:5432/mydb up
go run cmd/app/main.go
```
3. Run Frontend

```bash
cd frontend
npm i
npm run dev 
```
# Project tree
```bash
.
├── README.md
├── docker-compose.yml
├── .gitignore
├── .env
├── go.mod
├── go.sum
├── #backend
│   ├── cmd
│   │   └──  app
│   │        └──  main.go
│   ├── data
│   │    └──  db
│   │
│   ├── internal
│   │   ├── constant
│   │   │   └──  item.go
│   │   ├── item
│   │   │   ├── controller.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── model
│   │      └── item.go
│   └── migrations
│       └── 20240921032817_create_item_table.sql
└── #frontend
    ├── README.md
    ├── .next    
    ├── app
    │   ├── component
    │   │   ├── button.tsx
    │   │   └── navbar.tsx
    │   ├── create
    │   │   └── page.tsx
    │   ├── edit
    │   │   └── id
    │   │       └── page.tsx
    │   ├── fonts 
    │   │   ├── GeistMonoVF.woff
    │   │   └── GeistVF.woff
    │   ├── table
    │   │   └── page.tsx
    │   ├── favicon.ico
    │   ├── globals.css
    │   ├── layout.tsx
    │   └── page.tsx
    ├── next-env.d.ts
    ├── next.config.mjs
    ├── node_modules
    ├── package-lock.json
    ├── package.json
    ├── postcss.config.mjs
    ├── tailwind.config.ts
    └── tsconfig.json
```
# 9011310backend
```bash
echo "# 9011310backend" >> README.md
git init
git add README.md
git add .
git status
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/autzaa555/9011310backend.git
git push -u origin main
```