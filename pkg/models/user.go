package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rushabh2390/gousersmodule/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type CustomDate time.Time

var jwtKey = []byte("123456")

const customLayout = "2006-01-02"

// UnmarshalJSON parses the JSON date
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	t, err := time.Parse(`"`+customLayout+`"`, s)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}

// MarshalJSON formats the date for JSON
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(cd).Format(customLayout))), nil
}

// String method for easy printing
func (cd CustomDate) String() string {
	return time.Time(cd).Format(customLayout)
}

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares the hashed password with the plain password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;index"`
	Username    string `gorm:"unique;index"`
	Fullname    string
	Email       string `gorm:"unique"`
	Password    string
	DateOfBirth CustomDate `json:"date_of_birth"`
	IsSuperUser bool       `gorm:"default:false"`
	IsStaffUser bool       `gorm:"default:false"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
type JWTResponse struct {
	Token string `json:"token"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	password, err := HashPassword(u.Password)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	u.Password = password
	db.Create(&u)
	return u
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(Id int64) (*User, *gorm.DB, error) {
	var getUser User
	if err := db.Where("ID=?", Id).Find(&getUser).Error; err != nil {
		return nil, nil, err
	}
	return &getUser, db, nil
}

func DeleteUser(Id int64) (*User, error) {
	var user User
	if err := db.Where("ID=?", Id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) UpdateUser() (*User, error) {
	if err := db.Model(u).Where("id = ?", u.ID).Updates(u).Error; err != nil {
		return nil, err
	} // Return the updated user
	return u, nil
}

func (cred *Credentials) LoginUser() (*JWTResponse, error) {
	var user User
	if err := db.Model(&User{}).Where("username = ?", cred.Username).First(&user).Error; err != nil {
		if err := db.Model(&User{}).Where("email = ?", cred.Username).First(&user).Error; err != nil {
			return nil, errors.New("invalid username or password")
		}
	}
	if !CheckPasswordHash(cred.Password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Create JWT token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		jwtKey = []byte(env.JWTSecret)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &JWTResponse{Token: tokenString}, nil
}
