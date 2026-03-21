gopackage repositories

import "lottery-system/models"

type LotteryTicketRepository struct{}

// param: ticketNumber string
// return: *models.LotteryTicket (nil ถ้าไม่เจอ), error
func (r *LotteryTicketRepository) FindByTicketNumber(ticketNumber string) (*models.LotteryTicket, error) {
}

// param: order models.Order
// return: models.Order (saved), error
func (r *LotteryTicketRepository) SaveOrder(order models.Order) (models.Order, error) {
}

// param: userId string
// return: []models.LotteryTicket, error
func (r *LotteryTicketRepository) FindTicketsByUserId(userId string) ([]models.LotteryTicket, error) {
}
