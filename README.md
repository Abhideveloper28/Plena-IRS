# Plena-IRS
🚀 Access Key Management and Web3 Token Information Retrieval System
This project includes two Go-based microservices:

Access Key Management Service

Web3 Token Information Service

These services communicate asynchronously using Redis Pub/Sub, enabling secure, rate-limited access to mock Web3 token data via access keys.

🛠️ Technologies Used
Go (Golang)

Gin – HTTP Web Framework

GORM – ORM with SQLite

Redis – Pub/Sub for inter-service communication

SQLite – Local database for access key storage

UUID – For generating unique keys

✅ Prerequisites
Make sure the following are installed:

Go (go version)

Redis (redis-server)

Git

Install dependencies:

go mod tidy

▶️ How to Run
1. Start Redis
redis-server

2. Run Access Key Management Service
cd access-key-service
go run main.go

This service runs at: http://localhost:8080

3. Run Web3 Token Information Service
cd web3-token-info-service
go run main.go

This service runs at: http://localhost:8081

🔐 Access Key Management Service - API Endpoints
Method	Endpoint	Description
POST	/api/admin/key	Create a new access key
GET	/api/admin/keys	List all keys
PUT	/api/admin/key/:key	Update rate limit or expiry for a key
DELETE	/api/admin/key/:key	Delete a key
GET	/api/user/plan/:key	Get plan details for a key

🌐 Web3 Token Info Service - API Endpoint
Method	Endpoint	Description
GET	/api/token/:key	Fetch token info using access key

🔄 Redis Pub/Sub Events
When a new key is created, it is published to Redis so the Web3 service can maintain a local cache and enforce rate limits or access control.

✨ Features
Create, update, and delete access keys

Set rate limits and expiration times

Retrieve mock Web3 token information

Redis Pub/Sub for decoupled service communication

Clean, modular, and testable Go codebase

📦 Deliverables
✅ Two microservices
✅ REST APIs following standard practices
✅ Redis-based asynchronous communication
✅ Readable and scalable code
✅ Easy setup and testing instructions
