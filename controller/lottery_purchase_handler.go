package controller

import "net/http"

type LotteryPurchaseHandler struct{}

// POST /api/lottery/purchase
// param: req.body = { userId: string, ticketNumber: string }
// return: { orderId: string, status: string, paymentDeadline: time.Time }
func (h *LotteryPurchaseHandler) PurchaseLottery(w http.ResponseWriter, r *http.Request) {
}

// GET /api/lottery/inventory/:userId
// param: userId string (path param)
// return: []LotteryTicket
func (h *LotteryPurchaseHandler) GetMyInventory(w http.ResponseWriter, r *http.Request) {
} 
