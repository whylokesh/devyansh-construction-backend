# User API Demo Payloads

## 1. Signup (Register)
**Endpoint:** `POST /api/users/signup`

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "securepassword123",
  "role": "admin"
}
```

## 2. Login
**Endpoint:** `POST /api/users/login`

```json
{
  "email": "john.doe@example.com",
  "password": "securepassword123"
}
```

## CURL Commands

### Signup
```bash
curl -X POST http://localhost:8080/api/users/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "securepassword123",
    "role": "admin"
  }'
```

### Login
```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword123"
  }'
```
