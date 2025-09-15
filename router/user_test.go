package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest-api/mocks"
	"strings"
	"testing"
)

func TestSignUp_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	//mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("sdfsf", nil)
	mockUserService.EXPECT().
		AddUser(gomock.Any()).
		Return(nil)

	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/signup", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.SignUp(ctx)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	mockUserService.EXPECT().
		ValidateUser(gomock.Any()).
		Return(nil)
	mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("sdfsf", nil)

	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logged In Successfully")
}

func TestLogin_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("sdfsf", nil)
	mockUserService.EXPECT().
		ValidateUser(gomock.Any()).
		Return(nil)

	body := `{"password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Logged In Successfully")
}

func TestLogin_Unauthorised(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	mockUserService.EXPECT().
		ValidateUser(gomock.Any()).
		Return(nil)
	mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("", errors.New("Unauthorised"))

	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}

func TestSignUp_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	//mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("sdfsf", nil)
	mockUserService.EXPECT().
		AddUser(gomock.Any()).
		Return(nil)

	body := `{"password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/signup", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.SignUp(ctx)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignUp_Unauthorised(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Arrange
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUtil := mocks.NewMockUtil(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	//mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("sdfsf", nil)

	mockUserService.EXPECT().
		AddUser(gomock.Any()).
		Return(errors.New("unauthorised"))

	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/signup", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.SignUp(ctx)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
func TestLogin_FailureTokenValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockUtil := mocks.NewMockUtil(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)

	mockUserService.EXPECT().
		ValidateUser(gomock.Any()).
		Return(nil)
	mockUtil.EXPECT().TokenGeneration(gomock.Any(), gomock.Any()).Return("", errors.New("unauthorised"))

	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}
func TestLogin_FailureValidationFails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	mockUtil := mocks.NewMockUtil(ctrl)
	handler := NewUserHandler(mockUserService, mockUtil)
	mockUserService.EXPECT().
		ValidateUser(gomock.Any()).
		Return(errors.New("Error"))
	body := `{"email":"test@example.com","password":"secret"}`

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/login", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.Login(ctx)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
}
