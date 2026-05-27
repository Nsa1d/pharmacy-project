package apperrors

import "errors"

var (
	ErrUserNotFound       = errors.New("пользователь с таким ID не найден") 
	ErrMedicineNotFound   = errors.New("лекарство с таким ID не найдено") 
	ErrMedicineOutOfStock = errors.New("лекарства нет в наличии") 
	ErrInsufficientStock  = errors.New("недостаточно товара на складе") 
	ErrInvalidQuantity = errors.New("количество товара должно быть положительным") 
	ErrUserIDMismatch     = errors.New("user_id в URL и в теле запроса не совпадают")
	ErrCartEmpty = errors.New("вы еще ничего не добавили в корзину")
)
