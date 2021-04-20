package model

import (
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

type User struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"_"`
	UpdatedAt       time.Time `json:"_"`
	Email           string    `json:"email"`
	passwordHash    string    `json:"-"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Role            int       `json:"role"`
}

func (u *User) Register(conn *sql.DB) error {
	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("password must be at least 4 characters long.")
	}

	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}

	if len(u.Email) < 4 {
		return fmt.Errorf("email must be at least 4 characters long.")
	}

	u.Email = strings.ToLower(u.Email)
	row := conn.QueryRow("SELECT id from user_account WHERE email = ?", u.Email)
	userLookup := User{}
	err := row.Scan(&userLookup)
	if err != sql.ErrNoRows {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email already exists" + err.Error())
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.passwordHash = string(pwdHash)

	//Todo: think about the role logic
	u.Role = 0
	now := time.Now()
	_, err = conn.Exec("INSERT INTO user_account (created_at, updated_at, email, password, role) VALUES(?,?,?,?,?)", now, now, u.Email, u.passwordHash, u.Role)

	return err
}

// GetAuthToken returns the auth token to be used
func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["role"] = u.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}

// IsAuthenticated checks to make sure password is correct
func (u *User) IsAuthenticated(conn *sql.DB) error {
	row := conn.QueryRow("SELECT id, password, role from user_account WHERE email = ?", u.Email)
	err := row.Scan(&u.ID, &u.passwordHash, &u.Role)
	if err == sql.ErrNoRows {
		fmt.Println("User with email not found")
		return fmt.Errorf("Invalid login credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.passwordHash), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("Invalid login credentials")
	}

	return nil
}

func GetMyUSer(conn *sql.DB, id int) (error, User) {
	u := User{}
	row := conn.QueryRow("SELECT id, email , role from user_account WHERE id = ?", id)
	err := row.Scan(&u.ID, &u.Email, &u.Role)
	if err == sql.ErrNoRows {
		return fmt.Errorf("Invalid login credentials"), u
	}
	return err, u
}

func IsTokenValid(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// fmt.Printf("Parsing: %v \n", token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Token signing method is not valid: %v", token.Header["alg"])
		}

		return tokenSecret, nil
	})

	if err != nil {
		fmt.Printf("Err %v \n", err)
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims)
		userID := claims["user_id"]
		return true, fmt.Sprint(userID)
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		fmt.Println(err)
		return false, "uuid.UUID{}"
	}
}
