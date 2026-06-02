package apperrors

import (
	"errors"
	"net/http"
)

type Err struct {
	StatusCode int
	Msg        string
}

var (
	ErrUserNotFound       = errors.New("пользователь с таким ID не найден")
	ErrMedicineNotFound   = errors.New("лекарство с таким ID не найдено")
	ErrItemNotFound       = errors.New("лекарства с таким номером нет в корзине")
	ErrMedicineOutOfStock = errors.New("лекарства нет в наличии")
	ErrInsufficientStock  = errors.New("недостаточно товара на складе")
	ErrInvalidQuantity    = errors.New("количество товара должно быть положительным")
	ErrUserIDMismatch     = errors.New("user_id в URL и в теле запроса не совпадают")
	ErrCartEmpty          = errors.New("вы еще ничего не добавили в корзину")
	ErrItemAlreadyInCart  = errors.New("товар уже в корзине, измените количество")

	ErrEmptyRequiredFields     = errors.New("все обязательные поля должны быть заполнены")
	ErrAddressLengthInvalid    = errors.New("адрес доставки должен быть от 10 до 250 символов")
	ErrCommentLengthInvalid    = errors.New("комментарий должен содержать от 5 до 250 символов")
	ErrPromocodeNotFound       = errors.New("такой промокод не существует")
	ErrPromocodeInactive       = errors.New("промокод неактивен")
	ErrPromocodeExpired        = errors.New("срок действия промокода истек")
	ErrPromoUsageLimit         = errors.New("превышен лимит использования промокода")
	ErrPromoUserLimit          = errors.New("превышен лимит на пользователя для этого промокода")
	ErrDiscountTooHigh         = errors.New("скидка не может превышать или быть равной стоимости заказа")
	ErrOrdersNotFound          = errors.New("вы еще не делали заказов")
	ErrInvalidOrderStatus      = errors.New("статус заказа должен быть одним из: pending_payment, paid, canceled, shipped, completed")
	ErrInvalidStatusTransition = errors.New("недопустимый переход статуса заказа")

	ErrAmountMustBePositive = errors.New("сумма должна быть положительной")
	ErrInvalidPaymentMethod = errors.New("метод оплаты должен быть: card, cash или online_wallet")
	ErrAmountOverLimit      = errors.New("сумма не может превышать итоговую стоимость заказа")
	ErrPaymentsNotFound     = errors.New("вы еще не оплачивали ни один заказ")
	ErrOnePaymentNotFound   = errors.New("платёж не найден")
)

var errsMap = map[error]Err{
	ErrOnePaymentNotFound: {
		StatusCode: http.StatusNotFound,
		Msg: ErrOnePaymentNotFound.Error(),
	},
	ErrUserNotFound: {
		StatusCode: http.StatusNotFound,
		Msg:        ErrUserNotFound.Error(),
	},
	ErrInvalidStatusTransition: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrInvalidStatusTransition.Error(),
	},
	ErrInvalidOrderStatus: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrInvalidOrderStatus.Error(),
	},
	ErrMedicineNotFound: {
		StatusCode: http.StatusNotFound,
		Msg:        ErrMedicineNotFound.Error(),
	},
	ErrItemNotFound: {
		StatusCode: http.StatusNotFound,
		Msg:        ErrItemNotFound.Error(),
	},
	ErrOrdersNotFound: {
		StatusCode: http.StatusNotFound,
		Msg:        ErrOrdersNotFound.Error(),
	},
	ErrPaymentsNotFound: {
		StatusCode: http.StatusOK,
		Msg:        ErrPaymentsNotFound.Error(),
	},

	ErrMedicineOutOfStock: {
		StatusCode: http.StatusConflict,
		Msg:        ErrMedicineOutOfStock.Error(),
	},
	ErrInsufficientStock: {
		StatusCode: http.StatusConflict,
		Msg:        ErrInsufficientStock.Error(),
	},
	ErrItemAlreadyInCart: {
		StatusCode: http.StatusConflict,
		Msg:        ErrItemAlreadyInCart.Error(),
	},

	ErrInvalidQuantity: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrInvalidQuantity.Error(),
	},
	ErrUserIDMismatch: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrUserIDMismatch.Error(),
	},
	ErrEmptyRequiredFields: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrEmptyRequiredFields.Error(),
	},
	ErrAddressLengthInvalid: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrAddressLengthInvalid.Error(),
	},
	ErrCommentLengthInvalid: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrCommentLengthInvalid.Error(),
	},
	ErrPromocodeInactive: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrPromocodeInactive.Error(),
	},
	ErrPromocodeExpired: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrPromocodeExpired.Error(),
	},
	ErrDiscountTooHigh: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrDiscountTooHigh.Error(),
	},
	ErrAmountMustBePositive: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrAmountMustBePositive.Error(),
	},
	ErrInvalidPaymentMethod: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrInvalidPaymentMethod.Error(),
	},
	ErrAmountOverLimit: {
		StatusCode: http.StatusBadRequest,
		Msg:        ErrAmountOverLimit.Error(),
	},

	ErrCartEmpty: {
		StatusCode: http.StatusOK,
		Msg:        ErrCartEmpty.Error(),
	},

	ErrPromocodeNotFound: {
		StatusCode: http.StatusNotFound,
		Msg:        ErrPromocodeNotFound.Error(),
	},
	ErrPromoUsageLimit: {
		StatusCode: http.StatusConflict,
		Msg:        ErrPromoUsageLimit.Error(),
	},
	ErrPromoUserLimit: {
		StatusCode: http.StatusConflict,
		Msg:        ErrPromoUserLimit.Error(),
	},
}

func Get(err error) Err {
	if e, ok := errsMap[err]; ok {
		return e
	}

	return Err{
		StatusCode: http.StatusInternalServerError,
		Msg:        "внутренняя ошибка сервера",
	}
}
