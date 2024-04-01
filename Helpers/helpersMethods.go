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
var colName string = "Verification Code"
var colAccount string = "Account"
var collection *mongo.Collection
var collectionAccount *mongo.Collection

func init() {
	godotenv.Load()
	connectionString := os.Getenv("connectionString")
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)

	}
	collection = client.Database(dbName).Collection(colName)
	collectionAccount = client.Database(dbName).Collection(colAccount)
}

// check and send email
func SendEmailToUser(maildata models.Email) string {
	filteremail := bson.M{"email": maildata.ToEmail}
	filtername := bson.M{"name": maildata.UserName}
	var hasAccounte bson.M
	var hasAccountn bson.M

	collectionAccount.FindOne(context.Background(), filteremail).Decode(&hasAccounte)
	collectionAccount.FindOne(context.Background(), filtername).Decode(&hasAccountn)

	if hasAccounte == nil && hasAccountn == nil {
		auth := smtp.PlainAuth("", maildata.FromEmail, maildata.AppPassword, maildata.Host)
		addr := maildata.Host + ":" + maildata.Port
		vcode := strconv.Itoa(GenerateVcode()) // Generate Vcode
		message := "Subject: Hello User\nYour Verification Code is " + vcode
		msg := []byte(message)

		err := smtp.SendMail(addr, auth, maildata.FromEmail, []string{maildata.ToEmail}, msg)
		if err != nil {
			log.Fatal(err)
		}
		maildata.Vcode = vcode
		optdata := maildata
		SaveVcodeWithEmail(optdata)
		return "Vcode Send"
	}
	if hasAccounte != nil && hasAccountn != nil {
		return "Email & UserName Already Exist"
	} else if hasAccounte != nil {
		return "Email Already Exist"
	}
	return "UserName Already Exist"
}

// Generate Verification Code
func GenerateVcode() int {
	rand.New(rand.NewSource(time.Now().UnixMicro()))
	Vcode := rand.Intn(900000) + 100000

	return Vcode
}

// Save VerificationCode database
func SaveVcodeWithEmail(user models.Email) {
	filter := bson.M{"toemail": user.ToEmail}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		print(err)
	}
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Code SAVED : ", inserted.InsertedID)
}

// When verify Verification code later (within 5 min)
func LaterVcode(user string) models.Email {
	filter := bson.M{"toemail": user}
	var mongodb models.Email
	collection.FindOne(context.Background(), filter).Decode(&mongodb)
	return mongodb
}

// check Verification code
func CheckVCode(userVcode models.Email) bool {
	data := LaterVcode(userVcode.ToEmail)
	if data.Vcode == userVcode.Vcode {
		var UserData models.User
		UserData.Email = data.ToEmail
		UserData.Name = data.UserName
		UserData.Password = data.Password
		UserData.Islogin = true
		UserData.DeviceToken = data.DeviceToken
		UserData.HostName = data.Host
		CreateAndSaveUser(UserData)
		return true

	} else {

		return false
	}
}

// Get Verification code from DB
func GetVcodeFromDb(VcodeID string) models.Email {
	id, err := primitive.ObjectIDFromHex(VcodeID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	var dbvcode models.Email
	collection.FindOne(context.Background(), filter).Decode(&dbvcode)
	return dbvcode
}

// Create New User and Save in DB
func CreateAndSaveUser(user models.User) {
	inserted, err := collectionAccount.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE", inserted.InsertedID)

}

// Login
func Login(userEmail string, password string) bool {
	filter := bson.M{"email": userEmail, "password": password}
	var hasAccount bson.M
	collectionAccount.FindOne(context.Background(), filter).Decode(&hasAccount)
	if hasAccount != nil {
		return true
	} else {
		return false
	}

}
