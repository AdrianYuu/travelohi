# Travelohi

## Project Description

Travelohi is a comprehensive full-stack travel booking platform that enables users to search and book flights and hotels. The platform features real-time multiplayer gaming capabilities for earning rewards, secure payment processing, digital wallet management, and a loyalty point system with promotional offers.

## Features

- **Flight Booking**: Search, filter, and book flights with seat selection and detailed flight information
- **Hotel Booking**: Browse hotels with detailed information, room types, facilities, and real-time availability
- **User Authentication**: Secure JWT-based authentication with OTP email verification and personal security questions
- **Payment System**: Secure payment processing with multiple credit card support and transaction management
- **Digital Wallet**: Balance tracking, point redemption, and transaction history
- **Real-time Gaming**: Multiplayer games via WebSocket for earning rewards and promotional codes
- **Ratings & Reviews**: User reviews and ratings for flights and hotels with rating types
- **Loyalty Program**: Point system, promotional code management, and user promo tracking
- **Search History**: Track and manage previous flight and hotel searches
- **Admin Broadcasting**: Real-time announcements and promotional broadcasts to users
- **Airport Transit Information**: Detailed airport and transit information for travelers
- **User Profile Management**: Personal information, security questions, and account settings

## Technology Stack

### Backend

- **Language**: Go 1.21.4
- **Framework**: Gin v1.9.1
- **Database**: PostgreSQL with GORM ORM v1.25.5
- **Authentication**: JWT (golang-jwt v5.2.0)
- **API Documentation**: Swagger/OpenAPI
- **Email Service**: Gomail v2.0.0
- **Encryption**: golang.org/x/crypto
- **CORS**: gin-contrib/cors v1.5.0
- **Database Driver**: PostgreSQL driver v1.5.4

### Frontend

- **Language**: TypeScript 5.2.2
- **Framework**: React 18.2.0
- **Build Tool**: Vite 5.0.8
- **Styling**: Emotion v11.11.3, SASS v1.69.7
- **HTTP Client**: Axios v1.6.5
- **Real-time Communication**: Socket.IO Client v4.8.1
- **Routing**: React Router DOM v6.21.2
- **Authentication**: Firebase v10.7.2
- **Form Validation**: React Google reCAPTCHA v3.1.0
- **Linting**: ESLint v8.55.0

### WebSocket Server

- **Runtime**: Node.js (LTS)
- **Framework**: Express v5.2.1
- **Real-time Communication**: Socket.IO v4.8.1
- **Environment Management**: Dotenv v17.2.3

### AI/ML Service

- **Language**: Python
- **Framework**: Flask v3.0.0