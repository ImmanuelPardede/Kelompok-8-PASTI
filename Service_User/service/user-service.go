package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iqbalsiagian17/User/dto"
	"github.com/iqbalsiagian17/User/model"
	"github.com/iqbalsiagian17/User/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) Register(userDTO dto.UserDTO) (*dto.UserResponseDTO, error) {
	// Hash password (use bcrypt or any other secure hashing algorithm)
	hashedPassword := hashPassword(userDTO.Password)

	// Create new user
	newUser := model.User{
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  hashedPassword,
		Telephone: userDTO.Telephone,
		Role:      "customer", // default role
	}

	// Save user to database
	newUser, err := us.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	// Return user response DTO
	userResponse := dto.UserResponseDTO{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Telephone: newUser.Telephone,
		Role:      newUser.Role,
	}

	return &userResponse, nil
}

func (us *UserService) Login(loginDTO dto.LoginDTO) (string, error) {
	// Find user by email
	user, err := us.userRepository.FindByEmail(loginDTO.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check password (use bcrypt or any other secure hashing algorithm)
	if !checkPassword(loginDTO.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	// Generate JWT token (use a JWT library)
	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) Logout(tokenString string) error {
	// Validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil // Mengambil kunci rahasia dari variabel lingkungan
	})
	if err != nil {
		return err
	}

	// Validasi token
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return errors.New("invalid token")
	}

	// Tambahkan logika untuk menambahkan token ke daftar token yang sudah kadaluwarsa
	// atau di-blacklist di sini, jika diperlukan

	return nil
}

func hashPassword(password string) string {
	// Hash password menggunakan MD5
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func checkPassword(password, hashedPassword string) bool {
	// Memeriksa kecocokan password dengan hashed password
	return hashedPassword == hashPassword(password)
}

func generateToken(user model.User) (string, error) {
	// Membuat payload JWT
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token akan berlaku selama 72 jam

	// Membuat token JWT dengan signing key yang diperoleh dari variabel lingkungan JWT_SECRET
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
