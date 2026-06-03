package services

type NoopMedicineRatingUpdater struct{}

func (NoopMedicineRatingUpdater) UpdateAvgRating(medicineID uint, avgRating float64) error {
	return nil
}
