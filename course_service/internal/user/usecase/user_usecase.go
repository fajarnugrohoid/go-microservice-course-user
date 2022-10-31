package usecase

import (
	"course/internal/domain"
	"errors"
	"golang.org/x/net/context"
)

type UserGetter interface {
	GetByID(ctx context.Context, ID int) (domain.User, error)
}

type UserUsecase struct {
	repo UserGetter
}

func NewUserUsecase(repo UserGetter) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uu UserUsecase) GetUserByID(ctx context.Context, userID int) (domain.User, error) {
	if userID <= 0 {
		return domain.User{}, errors.New("Invalid User")
	}

	return uu.repo.GetByID(ctx, userID)
}

/*
func (uu UserUsecase) Register(c *gin.Context) {
	var userRequest UserRequest
	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	user, err := domain.NewUser(userRequest.Email, userRequest.Name, userRequest.Password, userRequest.NoHp)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	if err := uu.db.Create(&user).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "error when create user",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, map[string]string{
			"message": "error when generate token",
		})
		return
	}
	c.JSON(200, map[string]string{
		"token": token,
	})
}

type UserRequest struct {
	Name     string
	Email    string
	Password string
	NoHp     string
}

type UserLoginRequest struct {
	Email    string
	Password string
}


func (uu UserUsecase) Login(c *gin.Context) {
	var userRequest UserLoginRequest
	if err := c.ShouldBind(&userRequest); err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	var user domain.User
	err := uu.db.Where("email = ?", userRequest.Email).Take(&user).Error
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid username or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid username or password",
		})
		return
	}

	token, err := user.GenerateJWT()
	if err != nil {
		c.JSON(500, map[string]string{
			"message": "error when generate token",
		})
		return
	}
	c.JSON(200, map[string]string{
		"token": token,
	})
} */
