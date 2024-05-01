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

	currentTime := time.Now()
	formatTime := currentTime.Format("1504")
	user.SendAt = formatTime
	fmt.Println(user.SendAt)
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

// reset password
func ResetPasssword(maildata models.Email) bool {

	filter := bson.M{"email": maildata.ToEmail}
	var accountExist bson.M
	collectionAccount.FindOne(context.Background(), filter).Decode(&accountExist)

	fmt.Println(accountExist)
	if accountExist != nil {
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
		return true
	} else {
		return false
	}

}

// Check Vcode when reset password
func CheckResetVCode(userVcode models.Email) bool {
	data := LaterVcode(userVcode.ToEmail)
	if data.Vcode == userVcode.Vcode {
		var UserData models.User
		UserData.Email = data.ToEmail
		UserData.Password = data.Password
		UserData.Islogin = true
		UserData.DeviceToken = data.DeviceToken
		UpdateUserPassword(UserData)
		return true

	} else {

		return false
	}
}

// Update user password when resetting
func UpdateUserPassword(user models.User) {
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	inserted, err := collectionAccount.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE", inserted.UpsertedCount)

}

// Delete Vcode after 10 min
func DeleteVcode() {
	currentTime := time.Now().Format("1504")
	converttoINT, _ := strconv.ParseInt(currentTime, 10, 64)
	filter := bson.M{}
	curser, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)

	}

	for curser.Next(context.Background()) {
		var vcodeDbObject models.Email
		curser.Decode(&vcodeDbObject)

		vcodetime := vcodeDbObject.SendAt

		vcodeTimetoINT, _ := strconv.ParseInt(vcodetime, 10, 64)
		fmt.Println("Add time", converttoINT)
		fmt.Println("vcodeTime :", vcodeTimetoINT)
		if converttoINT-vcodeTimetoINT >= 10 {
			deleteFilter := bson.M{"toemail": vcodeDbObject.ToEmail}
			collection.DeleteOne(context.Background(), deleteFilter)
		}
	}

}
