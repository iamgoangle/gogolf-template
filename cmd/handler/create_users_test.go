package handler_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/iamgoangle/gogolf-template/cmd/domain"
	"github.com/iamgoangle/gogolf-template/cmd/handler"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	handlerMock "github.com/iamgoangle/gogolf-template/cmd/handler/mock"
)

var _ = Describe("CreateUsers", Label("library"), func() {
	var (
		mockCtrl    *gomock.Controller
		mockService *handlerMock.MockUserService
		mockAppPort = "1234"
		mockServer  *echo.Echo

		h *handler.Handler
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockServer = echo.New()
		mockServer.Validator = &handler.MockValidator{
			Validator: validator.New(),
		}
		mockService = handlerMock.NewMockUserService(mockCtrl)

		h = &handler.Handler{
			Server:      mockServer,
			Port:        mockAppPort,
			UserService: mockService,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	When("user sends a valid request payload", func() {
		It("should create a new user", func() {
			mockUser := &domain.User{
				Name:  "unit-test",
				Email: "unit-test@unit.com",
			}

			mockService.EXPECT().Create(gomock.Any()).Return(mockUser, nil).Times(1)

			req := httptest.NewRequest(http.MethodPost, "/v1/users", nil)
			rec := httptest.NewRecorder()
			c := mockServer.NewContext(req, rec)

			err := h.CreateUserHandler(c)

			Expect(err).Should(BeNil())
		})
	})
})
