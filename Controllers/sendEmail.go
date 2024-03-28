package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	models "sendemail/Models"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName string = "Users"
var colName string = "Accounts"
var collection *mongo.Collection
var gOTP string

func init() {
	godotenv.Load()
	connectionString := os.Getenv("connectionString")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)

	}
	collection = client.Database(dbName).Collection(colName)
}

// Send Email
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var maildata models.Email
	err := json.NewDecoder(r.Body).Decode(&maildata)
	if err != nil {
		log.Fatal(err)

	}
	auth := smtp.PlainAuth("", maildata.FromEmail, maildata.AppPassword, maildata.Host)
	addr := maildata.Host + ":" + maildata.Port
	gOTP = strconv.Itoa(GenerateOtp())
	message := "Subject: Hello User\nYour OTP is " + gOTP
	msg := []byte(message)
	err = smtp.SendMail(addr, auth, maildata.FromEmail, []string{maildata.ToEmail}, msg)
	if err != nil {
		log.Fatal(err)

	}
	optdata := models.OTPModel{UserEmail: maildata.ToEmail, DeviceToken: maildata.DeviceToken, OTP: gOTP}
	SaveOTPWithEmail(optdata)
	json.NewEncoder(w).Encode("OTP IS SEND ON YOUR EMAIL")

}

// Generate OTP
func GenerateOtp() int {
	rand.New(rand.NewSource(time.Now().UnixMicro()))
	otp := rand.Intn(900000) + 100000

	return otp
}

//Save OTP database

func SaveOTPWithEmail(otpdata models.OTPModel) {
	filter := bson.M{"useremail": otpdata.UserEmail}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		print(err)
	}
	inserted, err := collection.InsertOne(context.Background(), otpdata)
	if err != nil {
		log.Fatal(err)

	}
	idString := inserted.InsertedID.(primitive.ObjectID).Hex()
	var dbOTP models.OTPModel = GetOtpFromDb(idString)
	gOTP = dbOTP.OTP
	fmt.Println("OTP SAVED : ", inserted.InsertedID)
}

// When verify opt later
func LaterOTP(user string) string {
	filter := bson.M{"useremail": user}
	var mongodb models.OTPModel
	collection.FindOne(context.Background(), filter).Decode(&mongodb)
	gOTP = mongodb.OTP
	return gOTP
}

// Verify OTP
func VerifyOtp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userotp models.OTPModel
	json.NewDecoder(r.Body).Decode(&userotp)
	print(userotp.OTP)
	if LaterOTP(userotp.UserEmail) == userotp.OTP {
		json.NewEncoder(w).Encode("TRUE")
		return
	} else {
		json.NewEncoder(w).Encode("FALSE")
		return
	}
}
func GetOtpFromDb(otpid string) models.OTPModel {
	id, err := primitive.ObjectIDFromHex(otpid)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var dbotp models.OTPModel
	collection.FindOne(context.Background(), filter).Decode(&dbotp)
	return dbotp
}
