package helpers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

func SendEmailToUser(maildata models.Email) {
	auth := smtp.PlainAuth("", maildata.FromEmail, maildata.AppPassword, maildata.Host)
	addr := maildata.Host + ":" + maildata.Port
	gOTP = strconv.Itoa(GenerateOtp())
	message := "Subject: Hello User\nYour Verification Code is " + gOTP
	msg := []byte(message)
	err := smtp.SendMail(addr, auth, maildata.FromEmail, []string{maildata.ToEmail}, msg)
	if err != nil {
		log.Fatal(err)

	}
	optdata := models.OTPModel{UserEmail: maildata.ToEmail, DeviceToken: maildata.DeviceToken, OTP: gOTP}
	SaveOTPWithEmail(optdata)
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

// check otp
func CheckOTP(userotp models.OTPModel) bool {
	if LaterOTP(userotp.UserEmail) == userotp.OTP {

		return true
	} else {

		return false
	}
}

//Get OTP from DB

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
