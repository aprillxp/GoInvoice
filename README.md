🚀 GoInvoice
GoInvoice is a REST API backend application built with Golang to help businesses process invoices efficiently.
It comes with powerful features like invoice management, secure authentication, email notifications, and payment processing.

✨ Key Features

✅ User Authentication
Register & login with password hashing (bcrypt)
Secure JWT-based authentication middleware

✅ Invoice Management
Full CRUD (Create, Read, Update, Delete) operations on invoices
Data validation with GORM and Gin

✅ Email Notifications
Send invoices to user emails via SendGrid
Plan: Automatically send emails after successful Stripe payment

✅ Stripe Payment Integration
Securely handle payments for invoices using Stripe

✅ PostgreSQL Database
Managed via GORM ORM
Auto migration support

✅ Docker & Railway Deployment
Easy deployment via Docker containers
Live deployment on Railway with CI/CD

🛠️ Tech Stack
Golang: Gin HTTP framework, GORM ORM
PostgreSQL: Relational database
SendGrid: For sending invoice emails
Stripe: For payment processing
JWT & Bcrypt: For secure user authentication
Docker: For containerization
Railway: As the deployment platform


📦 Getting Started (Local)
git clone https://github.com/<your-username>/goinvoice.git
cd goinvoice
go mod tidy
go run main.go

Make sure to set your environment variables (or use a .env file):
DATABASE_URL=postgres://<user>:<password>@<host>:<port>/<db>?sslmode=disable
SENDGRID_API_KEY=...
STRIPE_SECRET_KEY=...
JWT_SECRET=...
APP_ENV=...

🚀 Docker Usage
docker build -t goinvoice .
docker run -p 8080:8080 goinvoice

🎯 Roadmap
 CRUD Invoice API
 JWT Authentication
 SendGrid Email Integration
 Stripe Payment Integration
 Stripe webhook to automatically send email after payment success
 Frontend in React (handled by collaborator)

📝 License
MIT © 2025 Gavrila Ariendra
