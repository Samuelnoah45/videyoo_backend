package auth_controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockGraphQLClient is a mock of the GraphQL client
type MockGraphQLClient struct {
    mock.Mock
}

func (m *MockGraphQLClient) Query(ctx context.Context, query interface{}, variables map[string]interface{}) error {
    args := m.Called(ctx, query, variables)
    return args.Error(0)
}

// MockUtilService is a mock of the utility service
type MockUtilService struct {
    mock.Mock
}

func (m *MockUtilService) ComparePasswords(hashedPwd string, plainPwd string) bool {
    args := m.Called(hashedPwd, plainPwd)
    return args.Bool(0)
}

func (m *MockUtilService) HasuraAccessToken(user authModel.User) (string, error) {
    args := m.Called(user)
    return args.String(0), args.Error(1)
}

// SetupRouter sets up the Gin router for testing
func SetupRouter(mockGraphQLClient *MockGraphQLClient, mockUtilService *MockUtilService) *gin.Engine {
    r := gin.Default()
    r.POST("/login", func(ctx *gin.Context) {
        Login(ctx, mockGraphQLClient, mockUtilService)
    })
    return r
}

func TestLoginSuccess(t *testing.T) {
    mockGraphQLClient := new(MockGraphQLClient)
    mockUtilService := new(MockUtilService)

    router := SetupRouter(mockGraphQLClient, mockUtilService)

    // Mock the GraphQL query response
    mockGraphQLClient.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
        query := args.Get(1).(*struct {
            Users []struct {
                ID         string `json:"id"`
                First_name string `json:"first_name"`
                Last_name  string `json:"last_name"`
                Email      string `json:"email"`
                Password   string `json:"password"`
                User_roles []struct {
                    Role_name string `json:"role_name"`
                } `json:"user_roles"`
            } `graphql:"users(where: {email: {_eq: $email}})"`
        })
        query.Users = []struct {
            ID         string `json:"id"`
            First_name string `json:"first_name"`
            Last_name  string `json:"last_name"`
            Email      string `json:"email"`
            Password   string `json:"password"`
            User_roles []struct {
                Role_name string `json:"role_name"`
            } `json:"user_roles"`
        }{
            {
                ID:         "1",
                First_name: "John",
                Last_name:  "Doe",
                Email:      "john.doe@example.com",
                Password:   "$2y$12$WzPbXU1hcQK8EkcAXQv8POr4E.nGJiG6MwI8uw5G/h.lJx3ddvRjq", // hashed password
                User_roles: []struct {
                    Role_name string `json:"role_name"`
                }{
                    {Role_name: "user"},
                },
            },
        }
    }).Return(nil)

    // Mock the password comparison
    mockUtilService.On("ComparePasswords", mock.Anything, mock.Anything).Return(true)

    // Mock the token generation
    mockUtilService.On("HasuraAccessToken", mock.Anything).Return("mock-token", nil)

    // Create a request to send to the route
    payload := `{"input":{"email":"john.doe@example.com","password":"password"}}`
    req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(w, req)

    // Check the status code
    assert.Equal(t, http.StatusOK, w.Code)

    // Check the response body
    var responseBody map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, true, responseBody["success"])
    assert.Equal(t, "mock-token", responseBody["token"])
}

func TestLoginInvalidCredentials(t *testing.T) {
    mockGraphQLClient := new(MockGraphQLClient)
    mockUtilService := new(MockUtilService)

    router := SetupRouter(mockGraphQLClient, mockUtilService)

    // Mock the GraphQL query response
    mockGraphQLClient.On("Query", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
        query := args.Get(1).(*struct {
            Users []struct {
                ID         string `json:"id"`
                First_name string `json:"first_name"`
                Last_name  string `json:"last_name"`
                Email      string `json:"email"`
                Password   string `json:"password"`
                User_roles []struct {
                    Role_name string `json:"role_name"`
                } `json:"user_roles"`
            } `graphql:"users(where: {email: {_eq: $email}})"`
        })
        query.Users = []struct {
            ID         string `json:"id"`
            First_name string `json:"first_name"`
            Last_name  string `json:"last_name"`
            Email      string `json:"email"`
            Password   string `json:"password"`
            User_roles []struct {
                Role_name string `json:"role_name"`
            } `json:"user_roles"`
        }{}
    }).Return(nil)

    // Create a request to send to the route
    payload := `{"input":{"email":"invalid@example.com","password":"password"}}`
    req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(w, req)

    // Check the status code
    assert.Equal(t, http.StatusBadRequest, w.Code)

    // Check the response body
    var responseBody map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &responseBody)
    assert.NoError(t, err)
    assert.Equal(t, "There is no account with email address invalid@example.com", responseBody["message"])
}
