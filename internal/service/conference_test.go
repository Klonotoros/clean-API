package service

import (
	"clean-API/internal/dto"
	"clean-API/internal/model"
	"clean-API/internal/repository"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"time"
)

var _ = Describe("ConferenceService", func() {

	var (
		conferenceService  ConferenceService
		conferenceRepoCtrl *gomock.Controller
		conferenceRepoMock *repository.MockConferenceRepository
		mockConferences    []model.Conference
		validConference    model.Conference
	)

	BeforeEach(func() {
		conferenceRepoCtrl = gomock.NewController(GinkgoT())
		conferenceRepoMock = repository.NewMockConferenceRepository(conferenceRepoCtrl)
		conferenceService = newConferenceService(conferenceRepoMock, dto.Config{})

		mockConferences = []model.Conference{
			{ID: 1, Name: "Conference 1", Description: "Description 1", Location: "Location 1", DateTime: time.Now(), UserID: 1},
			{ID: 2, Name: "Conference 2", Description: "Description 2", Location: "Location 2", DateTime: time.Now().Add(10 * time.Hour), UserID: 2},
		}
		validConference = model.Conference{
			Name:        "Test Conference",
			Description: "Test Description",
			Location:    "Test Location",
			DateTime:    time.Now(),
			UserID:      1,
		}

	})

	AfterEach(func() {
		conferenceRepoCtrl.Finish()
	})

	Describe("GetAllConferences", func() {
		Context("when GetAllConferences is called", func() {
			It("should return all conferences", func() {
				// given
				conferenceRepoMock.EXPECT().GetAllConferences().Return(mockConferences, nil)

				// when
				conferences, err := conferenceService.GetAllConferences()

				// then
				Expect(err).To(BeNil())
				Expect(conferences).To(Equal(mockConferences))
			})

			It("should return an error if GetAllConferences fails", func() {
				// given
				conferenceRepoMock.EXPECT().GetAllConferences().Return(nil, errors.New("error"))

				// when
				conferences, err := conferenceService.GetAllConferences()
				// then
				Expect(err.Error()).To(Equal("error"))
				Expect(conferences).To(BeNil())
			})
		})
	})

	Describe("GetConferenceById", func() {
		Context("when GetConferenceById is called with a valid conference ID", func() {
			It("should return the corresponding conference", func() {
				// given

				conferenceRepoMock.EXPECT().GetConferenceByID(int64(1)).Return(mockConferences[0], nil)

				// when
				conference, err := conferenceService.GetConferenceById(int64(1))

				// then
				Expect(err).To(BeNil())
				Expect(conference).To(Equal(mockConferences[0]))
			})
		})

		Context("when GetConferenceById is called with an invalid conference ID", func() {
			It("should return an error", func() {
				// given

				conferenceRepoMock.EXPECT().GetConferenceByID(int64(100)).Return(model.Conference{}, errors.New("conference not found"))

				// when
				conference, err := conferenceService.GetConferenceById(int64(100))

				// then
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("conference not found"))
				Expect(conference).To(Equal(model.Conference{}))
			})
		})
	})

	Describe("Save", func() {
		Context("when Save is called with a valid conference", func() {
			It("should save the conference successfully", func() {
				// given

				conferenceRepoMock.EXPECT().Save(validConference).Return(validConference, nil)

				// when
				savedConference, err := conferenceService.Save(validConference)

				// then
				Expect(err).To(BeNil())
				Expect(savedConference).To(Equal(validConference))
			})
		})

		Context("when Save is called with an invalid conference", func() {
			It("should return an error", func() {
				// given
				invalidConference := model.Conference{}
				conferenceRepoMock.EXPECT().Save(invalidConference).Return(model.Conference{}, errors.New("validation error"))

				// when
				savedConference, err := conferenceService.Save(invalidConference)

				// then
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(Equal("validation error"))
				Expect(savedConference).To(Equal(model.Conference{}))
			})
		})
	})

	Describe("Update", func() {
		Context("when Update is called with a valid conference", func() {
			It("should not return an error", func() {
				// given
				conferenceRepoMock.EXPECT().Update(validConference).Return(nil)
				// when
				err := conferenceService.Update(validConference)
				// then
				Expect(err).To(BeNil())

			})
		})

		Context("when Update is called with invalid conference", func() {
			It("should return an error", func() {
				// given
				invalidConference := model.Conference{}
				conferenceRepoMock.EXPECT().Update(invalidConference).Return(errors.New("can't update data"))
				// when
				err := conferenceService.Update(invalidConference)
				//then
				Expect(err.Error()).To(Equal("can't update data"))
			})
		})

	})

	Describe("Delete", func() {
		Context("when Delete is called with valid conference", func() {
			It("should not return an error", func() {
				// given
				conferenceRepoMock.EXPECT().Delete(validConference).Return(nil)
				// when
				err := conferenceService.Delete(validConference)
				// then
				Expect(err).To(BeNil())
			})
		})

		Context("when Delete is called with invalid conference", func() {
			It("should return an error", func() {
				// given
				invalidConference := model.Conference{}
				conferenceRepoMock.EXPECT().Delete(invalidConference).Return(errors.New("can't delete data"))
				// when
				err := conferenceService.Delete(invalidConference)
				// then
				Expect(err.Error()).To(Equal("can't delete data"))
			})
		})
	})
})
