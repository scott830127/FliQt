# FliQt (Backend)

This is a Golang-based HR system backend built with the Gin framework. It demonstrates clean backend architecture and functionality suitable for technical interviews.

---

### 🗓️ Development Started: **2025-05-03**

---

## ✅ Features Implemented

- ✅ RESTful API using Gin
- ✅ GORM with MySQL: auto-migration + initial seed data
- ✅ Redis integration (Cache + Distributed Locking)
- ✅ Leave Application API
- ✅ Leave Type Query API
- ✅ Dependency Injection via Google Wire
- ✅ Docker support (Dockerfile + Docker Compose)
- ✅ Makefile for easy local development

---

## 🗃️ Database Models

- **Employee**: Basic profile, contact info, salary, etc.
- **LeaveType**: Leave category (e.g., Annual, Sick)
- **LeaveQuota**: Tracks entitlement vs. usage per leave type
- **LeaveRecord**: Employee leave history

---

## 🚀 Getting Started

### 📦 Prerequisites

- Go 1.23.0 (darwin/arm64)
- Docker & Docker Compose

### 🔧 Build and Run

```bash
make docker
# Then start services
docker compose up -d
```

### 🛠 Development Mode
```bash
make start
```

### 🔁 Regenerate Wire DI
```bash
make wire
```

---

## 💡 Sample API Usage

### `POST /fliqt/leave`
```json
{
  "employeeID": 1,
  "leaveTypeID": 1,
  "startTime": "2025-05-10T09:00:00Z",
  "endTime": "2025-05-10T18:00:00Z",
  "hours": 8,
  "reason": "Personal leave"
}
```

### `GET /fliqt/leave-types`
```json
{
  "leaveTypes": [
    {"code": "ANNUAL", "name": "特休", "description": "年假"},
    {"code": "SICK", "name": "病假", "description": "生病請假"}
  ]
}
```

---

## 📁 Folder Structure Overview
```
.
├── cmd/
│   └── fliqt/main.go
├── deploy/
│   └── config/config.local.toml
├── internals/
│   └── app/
│       ├── api/
│       ├── config/
│       ├── entity/
│       ├── repository/
│       ├── service/
│       ├── router/
│       └── wire.go
├── pkg/
│   ├── gormx/
│   └── redisx/
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 🛠️ TODOs & Recommendations

### 📘 Documentation & Testing
- [ ] API Docs / Swagger UI
- [ ] Unit Tests (API / Repository)

### 📋 Additional Features
- [ ] Employee: Create / Update / Get
- [ ] Leave Approval Flow (Pending → Approve/Reject)
- [ ] Employee Leave History API
- [ ] Leave Quota Validation & Rules

### 🔒 Quality Enhancements
- [ ] Unified Error Handling Format
- [ ] Input Validation & Role-based Authorization
- [ ] Pagination & Filtering Support

---