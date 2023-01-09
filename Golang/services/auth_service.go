package services

import (
	"GOLANG/config"
	"GOLANG/dto"
	"GOLANG/models"
	"GOLANG/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LogIn(*dto.LogInReq) (*dto.TokenResponse, error)
	ForgotPassword(*dto.ForgotPasswordReq) (*dto.ForgotPasswordRes, error)
	Locations() ([]*dto.LocationRes, error)
	CreateAttendance(int, int) (*dto.AttendanceRes, error)
	CreateCheckout(int, int) (*dto.AttendanceRes, error)
	Logs(string) ([]*dto.LogsRes, error)
	Profile(int) (*dto.ProfileRes, error)
}

type authService struct {
	userRepository repositories.UserRepository
	appConfig      config.AppConfig
}

type AuthSConfig struct {
	UserRepository repositories.UserRepository
	AppConfig      config.AppConfig
}

func NewAuthService(c *AuthSConfig) AuthService {
	return &authService{
		userRepository: c.UserRepository,
		appConfig:      c.AppConfig,
	}
}

type idTokenClaims struct {
	jwt.RegisteredClaims
	User *models.User `json:"user"`
}

func (a *authService) generateJWTToken(user *models.User) (*dto.TokenResponse, error) {
	var idExp = a.appConfig.JWTExpireInMinutes * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp
	timeExpire := jwt.NumericDate{Time: time.Unix(tokenExp, 0)}
	timeNow := jwt.NumericDate{Time: time.Now()}

	claims := &idTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    a.appConfig.AppName,
			IssuedAt:  &timeNow,
			ExpiresAt: &timeExpire,
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.appConfig.JWTSecret)

	if err != nil {
		return new(dto.TokenResponse), err
	}
	return &dto.TokenResponse{IDToken: tokenString}, nil
}

func (a *authService) LogIn(req *dto.LogInReq) (*dto.TokenResponse, error) {
	user, err := a.userRepository.MatchingCredential(req.Username)

	errNotMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if errNotMatch != nil || user == nil {
		return nil, err
	}

	token, err := a.generateJWTToken(user)

	return token, err
}

func (a *authService) ForgotPassword(fpr *dto.ForgotPasswordReq) (*dto.ForgotPasswordRes, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(fpr.Password), 10)
	fpr.Password = string(bytes)
	updateInfo := &models.User{
		KTPID:    fpr.KTPID,
		Password: fpr.Password,
	}

	user, err := a.userRepository.ForgotPassword(updateInfo)
	if err != nil {
		return new(dto.ForgotPasswordRes), err
	}

	return new(dto.ForgotPasswordRes).FromForgotPassword(user), nil
}

func (a *authService) Locations() ([]*dto.LocationRes, error) {
	locations, err := a.userRepository.Location()

	// todo handle error

	var locationList []*dto.LocationRes

	for i := range locations {
		locationList = append(locationList, new(dto.LocationRes).FromLocation(locations[i]))
	}

	return locationList, err
}

func (a *authService) CreateAttendance(UserID int, LocationID int) (*dto.AttendanceRes, error) {
	attendanceInfo := &models.Attendance{
		UserID:     UserID,
		LocationID: LocationID,
		CheckInOut: "In",
	}

	res, err := a.userRepository.CreateAttendance(attendanceInfo)
	if err != nil {
		return new(dto.AttendanceRes), err
	}

	return new(dto.AttendanceRes).FromAttendance(res), nil
}

func (a *authService) CreateCheckout(UserID int, LocationID int) (*dto.AttendanceRes, error) {
	attendanceInfo := &models.Attendance{
		UserID:     UserID,
		LocationID: LocationID,
		CheckInOut: "Out",
	}

	res, err := a.userRepository.CreateCheckout(attendanceInfo)
	if err != nil {
		return new(dto.AttendanceRes), err
	}

	return new(dto.AttendanceRes).FromAttendance(res), nil
}

func (a *authService) Logs(filterBy string) ([]*dto.LogsRes, error) {
	logs, err := a.userRepository.Logs(filterBy)

	// todo handle error

	var logList []*dto.LogsRes

	for i := range logs {
		logList = append(logList, new(dto.LogsRes).FromLogs(logs[i]))
	}

	return logList, err
}

func (a *authService) Profile(id int) (*dto.ProfileRes, error) {
	user, err := a.userRepository.Profile(id)
	if err != nil {
		return new(dto.ProfileRes), err
	}

	return new(dto.ProfileRes).FromUser(user), nil
}
