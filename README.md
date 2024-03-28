Description:

The SendEmail project is a simple Go application designed to facilitate the sending of one-time passwords (OTP) via email. It leverages SMTP for email delivery and MongoDB for storing OTP data. This project provides a straightforward solution for integrating OTP-based authentication or verification mechanisms into web applications.

Key Features:

OTP Email Sending: Users can trigger the generation and sending of OTP emails to their registered email addresses.

OTP Verification: Users can verify the OTP provided via email against the OTP stored in the database.

MongoDB Integration: OTP data, including user email, OTP, and device token, is stored in MongoDB, providing persistence and scalability.

Technology Stack:

Go: The project is developed using the Go programming language, known for its simplicity, concurrency support, and efficiency.

MongoDB: MongoDB is used as the NoSQL database to store OTP-related data. It offers flexibility and scalability for handling large volumes of data.

SMTP (Simple Mail Transfer Protocol): SMTP is utilized for sending OTP emails to users' email addresses. It's a widely used protocol for email transmission over the internet.

Usage:

User Signup: Users trigger the signup process by providing their email address. Upon signup, an OTP is generated and sent to the provided email address for verification.

OTP Verification: After receiving the OTP via email, users enter the OTP into the application for verification. The application compares the provided OTP with the OTP stored in the database to authenticate the user.

Project Structure:

Controllers: Contains the HTTP request handlers responsible for processing user requests, such as sending OTP emails and verifying OTPs.

Models: Defines the data models used in the application, including the Email and OTPModel structs.
