gopackage services

import "lottery-system/models"

type LotteryPurchaseService struct{}

// param: userId string, ticketNumber string
// return: models.Order, error
func (s *LotteryPurchaseService) CreateOrder(userId string, ticketNumber string) (models.Order, error) {
}

// param: orderId string
// return: bool (true = ยังอยู่ในเวลาที่กำหนด), error
func (s *LotteryPurchaseService) ValidateTimeout(orderId string) (bool, error) {
}

// param: userId string
// return: []models.LotteryTicket, error
func (s *LotteryPurchaseService) GetInventoryByUser(userId string) ([]models.LotteryTicket, error) {
} 
