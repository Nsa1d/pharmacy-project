package apperrors

import "errors"

var (
	ErrUserNotFound       = errors.New("пользователь не найден")
	ErrUserAlreadyExists  = errors.New("пользователь с таким email или телефоном уже существует")
	ErrMedicineNotFound   = errors.New("лекарство с таким ID не найдено")
	ErrItemNotFound       = errors.New("лекарства с таким номером нет в корзине")
	ErrMedicineOutOfStock = errors.New("лекарства нет в наличии")
	ErrInsufficientStock  = errors.New("недостаточно товара на складе")
	ErrInvalidQuantity    = errors.New("количество товара должно быть положительным")
	ErrUserIDMismatch     = errors.New("user_id в URL и в теле запроса не совпадают")
	ErrCartEmpty          = errors.New("вы еще ничего не добавили в корзину")
	ErrItemAlreadyInCart  = errors.New("товар уже в корзине, измените количество")
)
