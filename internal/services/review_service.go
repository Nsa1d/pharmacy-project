package services

import (
	"math"
	"pharmacy-project/internal/apperrors"
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
	"strings"
)

type PurchaseValidator interface {
	UserPurchasedMedicine(userID, medicineID uint) (bool, error)
}

type MedicineRatingUpdater interface {
	UpdateAvgRating(medicineID uint, avgRating float64) error
}

type ReviewService interface {
	CreateReview(medicineID uint, req models.ReviewCreateRequest) (*models.Review, error)
	GetReviewsByMedicineID(medicineID uint) ([]models.Review, error)
	UpdateReview(id uint, req models.ReviewUpdateRequest) (*models.Review, error)
	DeleteReview(id uint) error
}

type reviewService struct {
	reviews       repository.ReviewRepository
	users         repository.UserRepository
	medicine      repository.MedicineRepository
	purchases     PurchaseValidator
	ratingUpdater MedicineRatingUpdater
}

func NewReviewService(
	reviews repository.ReviewRepository,
	users repository.UserRepository,
	medicine repository.MedicineRepository,
	purchases PurchaseValidator,
	ratingUpdater MedicineRatingUpdater,
) ReviewService {
	if ratingUpdater == nil {
		ratingUpdater = NoopMedicineRatingUpdater{}
	}

	return &reviewService{
		reviews:       reviews,
		users:         users,
		medicine:      medicine,
		purchases:     purchases,
		ratingUpdater: ratingUpdater,
	}
}

func (s *reviewService) CreateReview(medicineID uint, req models.ReviewCreateRequest) (*models.Review, error) {
	if err := s.validateReviewCreate(req); err != nil {
		return nil, err
	}

	user, err := s.users.GetByID(req.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}

	med, err := s.medicine.GetByID(medicineID)
	if err != nil {
		return nil, err
	}
	if med == nil {
		return nil, apperrors.ErrMedicineNotFound
	}

	ok, err := s.purchases.UserPurchasedMedicine(req.UserID, medicineID)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, apperrors.ErrReviewNotAllowed
	}

	review := &models.Review{
		UserID:     req.UserID,
		MedicineID: medicineID,
		Rating:     req.Rating,
		Text:       strings.TrimSpace(req.Text),
	}

	if err := s.reviews.Create(review); err != nil {
		return nil, err
	}

	if err := s.recalculateMedicineRating(medicineID); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) GetReviewsByMedicineID(medicineID uint) ([]models.Review, error) {
	med, err := s.medicine.GetByID(medicineID)
	if err != nil {
		return nil, err
	}
	if med == nil {
		return nil, apperrors.ErrMedicineNotFound
	}

	return s.reviews.GetByMedicineID(medicineID)
}

func (s *reviewService) UpdateReview(id uint, req models.ReviewUpdateRequest) (*models.Review, error) {
	review, err := s.reviews.GetByID(id)
	if err != nil {
		return nil, err
	}
	if review == nil {
		return nil, apperrors.ErrReviewNotFound
	}

	if req.Rating != nil {
		if *req.Rating < 1 || *req.Rating > 5 {
			return nil, apperrors.ErrInvalidReviewInput
		}
		review.Rating = *req.Rating
	}

	if req.Text != nil {
		text := strings.TrimSpace(*req.Text)
		if text == "" {
			return nil, apperrors.ErrInvalidReviewInput
		}
		review.Text = text
	}

	if err := s.reviews.Update(review); err != nil {
		return nil, err
	}

	if err := s.recalculateMedicineRating(review.MedicineID); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) DeleteReview(id uint) error {
	review, err := s.reviews.GetByID(id)
	if err != nil {
		return err
	}
	if review == nil {
		return apperrors.ErrReviewNotFound
	}

	if err := s.reviews.Delete(id); err != nil {
		return err
	}

	return s.recalculateMedicineRating(review.MedicineID)
}

func (s *reviewService) validateReviewCreate(req models.ReviewCreateRequest) error {
	if req.UserID == 0 {
		return apperrors.ErrInvalidReviewInput
	}
	if req.Rating < 1 || req.Rating > 5 {
		return apperrors.ErrInvalidReviewInput
	}
	if strings.TrimSpace(req.Text) == "" {
		return apperrors.ErrInvalidReviewInput
	}

	return nil
}

func (s *reviewService) recalculateMedicineRating(medicineID uint) error {
	reviews, err := s.reviews.GetByMedicineID(medicineID)
	if err != nil {
		return err
	}

	if len(reviews) == 0 {
		return s.ratingUpdater.UpdateAvgRating(medicineID, 0)
	}

	var sum int
	for _, review := range reviews {
		sum += review.Rating
	}

	avg := float64(sum) / float64(len(reviews))
	avg = math.Round(avg*10) / 10

	return s.ratingUpdater.UpdateAvgRating(medicineID, avg)
}
