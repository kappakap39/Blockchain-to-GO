package routers

import (
	"blockchain_to_go/controllers"
	"encoding/json"
	"net/http"
)

// GetBlockchainHandler จัดการการร้องขอเพื่อดึงข้อมูลบล็อกเชนทั้งหมด
func GetBlockchainHandler(chain *controllers.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chain.Blocks)
	}
}

// AddBlockHandler จัดการการร้องขอเพื่อเพิ่มบล็อกใหม่
func AddBlockHandler(chain *controllers.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data struct {
			Data string `json:"data"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		chain.AddBlock(data.Data)
		w.WriteHeader(http.StatusCreated)
	}
}
