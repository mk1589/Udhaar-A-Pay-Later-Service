
# 💳 Udhaar-A-Pay-Later-Service

A simple terminal-based **Pay-Later** application written in Go, allowing users to make purchases on credit from merchants and repay later — with real-time credit validation, discount handling, and detailed reporting.

---

## 🛠️ Features

- 🧍 Add users with credit limits
- 🏬 Add merchants with custom discounts
- 💸 Record transactions with credit checks
- 💰 Handle paybacks (partial/full)
- 📊 Generate reports:
  - Total discount received from merchants
  - Outstanding dues per user
  - Users at credit limit
  - Total dues system-wide

---

## 📦 Project Structure
![Udhaar Flowchart](./assets/image.png)


---

## 🧑‍💻 Setup & Run Locally

### ✅ Prerequisites
- [Go installed](https://go.dev/doc/install) (v1.18 or higher)

### 📥 Clone the Repository
```bash
$ git clone https://github.com/mk1589/Udhaar-A-Pay-Later-Service.git
```

```bash
$ cd Udhaar-A-Pay-Later-Service
```
```bash
$ go mod init udhaar
```
```bash
$ go mod tidy
```
```bash
$ go run main.go
```

## 📘 Example CLI Session
### Add Users
```bash
new user user1 u1@users.com 300
new user user2 u2@users.com 400
```
### Add Merchants
```bash

new merchant m1 0.5%
new merchant m2 1.5%
```
### Transactions
```bash

new txn user1 m2 300        # Success
new txn user2 m1 500        # Rejected (exceeds credit limit)
```
### Payback
```bash
payback user1 100
```
### Reporting
```bash

report discount m2          # Total discount earned from m2
report dues user1           # Remaining dues for user1
report users-at-credit-limit
report total-dues
```


---

## ❓ Notes & Tips

- ✅ Inputs are **case-sensitive** and must follow the **exact syntax** shown in examples.
- 🧮 Merchant discount can be written as `1.5%` or simply `1.5`.
- 🧠 The application uses **in-memory storage** — no database setup is required.
- 🔚 To exit the CLI, type one of the following:
  - `exit`
  - `quit`
  - `stop`




