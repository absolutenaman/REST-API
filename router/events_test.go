package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest-api/mocks"
	"rest-api/models"
	"strings"
	"testing"
	"time"
)

func TestCreateEvent_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEventService := mocks.NewMockEventsService(ctrl)
	handler := NewEventsHandler(mockEventService)
	mockEventService.EXPECT().Save(gomock.Any()).Return()
	body := `{"name": "Naman Srivastava",
	"description": "Marriage",
	"location": "India",
	"dateTime": "2026-01-01T00:00:00Z"}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/events", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1Nzk3Mjg1NiwiaWQiOjJ9.UmNZ4Gbi6FpH1EJjclkna4DlwjO6W1iLRW1I_8eJz1U")
	handler.createEvent(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"event":{"id":0,"name":"Naman Srivastava","description":"Marriage","location":"India","dateTime":"2026-01-01T00:00:00Z","user":0},"message":"Event created succesfully"}`)
}
func TestCreateEvent_FailureBad(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEventService := mocks.NewMockEventsService(ctrl)
	handler := NewEventsHandler(mockEventService)
	//mockEventService.EXPECT().Save(gomock.Any()).Return()
	body := `{
	"description": "Marriage",
	"location": "India",
	"dateTime": "2026-01-01T00:00:00Z"}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/events", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1Nzk3Mjg1NiwiaWQiOjJ9.UmNZ4Gbi6FpH1EJjclkna4DlwjO6W1iLRW1I_8eJz1U")
	handler.createEvent(ctx)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `{"message":"Bad Request"}`)
}
func TestUpdateEvent_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEventService := mocks.NewMockEventsService(ctrl)
	handler := NewEventsHandler(mockEventService)
	mockTime := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
	mockEventService.EXPECT().GetAllEventsById(gomock.Any()).Return(
		models.Events{ID: 0, Name: "Naman Srivastava", Description: "Marriage",
			Location: "India", DateTime: mockTime, User: 0}, nil)
	body := `{"name": "Naman Srivastava",
	"description": "Marriage",
	"location": "India",
	"dateTime": "2030-01-01T00:00:00Z"}`
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("PUT", "/events/2", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5hbWFuQGdtYWlsLmNvbSIsImV4cCI6MTc1Nzk3Mjg1NiwiaWQiOjJ9.UmNZ4Gbi6FpH1EJjclkna4DlwjO6W1iLRW1I_8eJz1U")
	handler.updateEvent(ctx)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"event":{"id":2,"name":"Naman Srivastava","description":"Marriage","location":"India","dateTime":"2026-01-01T00:00:00Z","user":0},"message":"Event created succesfully"}`)
}
