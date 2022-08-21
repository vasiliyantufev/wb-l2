package main

import "fmt"

func main() {
	mobile := NewMobileAlert()

	fmt.Println(mobile.Alert())

	mobile.SetState(&MobileAlertSong{})

	fmt.Println(mobile.Alert())
}

// MobileAlertStater обеспечивает общий интерфейс для различных состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует предупреждение в зависимости от его состояния.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку предупреждения
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState изменяет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert это MobileAlert конструктор.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует виброзвонок
type MobileAlertVibration struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong реализует звуковой сигнал
type MobileAlertSong struct {
}

// Alert возвращает строку предупреждения
func (a *MobileAlertSong) Alert() string {
	return "один, два, три..."
}
