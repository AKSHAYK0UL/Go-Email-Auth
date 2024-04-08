Email Authentication

This project implements a simple email verification system using Go and MongoDB. Users can sign up, receive a verification code via email, verify the code, and login to their accounts.

Features:

Sign Up: Users can sign up by providing their email address. A verification code is sent to the provided email address for verification.

Code Verification: Users can verify the code received via email to complete the sign-up process.

Login: Registered users can log in to their accounts using their email and password.

RESTful API: The system provides RESTful APIs for signing up, verifying codes and logging in.

Project Structure:

models: Contains the data models for the application, including Email and User.

helpers: Contains helper functions for sending emails, generating verification codes, saving codes to the database, and user authentication.

controllers: Contains the HTTP request handlers for different endpoints, including signing up, verifying codes and logging in.

router: Configures the HTTP routes for the application using the Gorilla Mux router.

.env: Environment file for storing configuration variables such as MongoDB connection string.

main.go: Entry point of the application, initializes the server and starts listening for incoming requests.

Usage:

Set up a MongoDB instance and provide the connection string in the .env file.

Install dependencies using go mod tidy.

Run the application using go run main.go.

Use the provided RESTful APIs to interact with the system.

API Endpoints:

POST /signup: Sign up a new user by providing their email address. Returns a success message and sends a verification code to the provided email.

POST /verify: Verify the code received via email to complete the sign-up process.

POST /reset: Initiate the password reset process by providing the email address and new password. Sends a verification code to the provided email.

POST /resetverify: Verify the code received via email to reset the password. If the code is correct, the new password is saved.

POST /login: Log in to an existing user account by providing email and password. Returns a success message upon successful login.

"This project is currently under development, so please be aware that it may undergo further changes and updates"
