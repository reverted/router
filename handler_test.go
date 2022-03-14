package router_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	"github.com/reverted/router"
	"github.com/reverted/router/mocks"
)

var _ = Describe("Handler", func() {

	var (
		err error
		req *http.Request
		rec *httptest.ResponseRecorder

		mockCtrl    *gomock.Controller
		mockRouter  *mocks.MockRouter
		mockHandler *mocks.MockHandler

		handler http.Handler
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRouter = mocks.NewMockRouter(mockCtrl)
		mockHandler = mocks.NewMockHandler(mockCtrl)

		handler = router.NewHandler(
			newLogger(),
			mockRouter,
			mockHandler,
		)
	})

	Describe("ServeHTTP", func() {
		BeforeEach(func() {
			req, err = http.NewRequest("GET", "http://localhost", nil)
			Expect(err).NotTo(HaveOccurred())

			rec = httptest.NewRecorder()
		})

		JustBeforeEach(func() {
			handler.ServeHTTP(rec, req)
		})

		Context("when the router fails", func() {
			BeforeEach(func() {
				mockRouter.EXPECT().Route(req).Return(errors.New("nope"))
			})

			It("responds with Not Found", func() {
				Expect(rec.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})

		Context("when the router succeeds", func() {
			BeforeEach(func() {
				mockRouter.EXPECT().Route(req).Return(nil)
			})

			Context("when it forwards the request to the handler", func() {
				BeforeEach(func() {
					mockHandler.EXPECT().ServeHTTP(rec, req)
				})

				It("succeeds", func() {
					Expect(rec.Result().StatusCode).To(Equal(http.StatusOK))
				})
			})
		})
	})
})

func newLogger() *logger {
	return &logger{}
}

type logger struct{}

func (self *logger) Error(args ...interface{}) {
	fmt.Fprintln(GinkgoWriter, args...)
}
