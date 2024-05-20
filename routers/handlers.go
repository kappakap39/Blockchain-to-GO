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
			Data []string `json:"data"` // เปลี่ยนจาก string เป็น []string เพื่อรองรับอาร์เรย์ของสตริง
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// เรียกใช้ฟังก์ชัน AddBlock และรับผลลัพธ์จากการเพิ่มบล็อก
		newBlocks := chain.AddBlock(data.Data)

		// สร้าง response สำหรับแสดงว่าบล็อกถูกเพิ่มเข้าไปในบล็อกเชนแล้ว
		res := map[string]interface{}{
			"message": "บล็อกถูกเพิ่มเข้าไปในบล็อกเชนแล้ว",
			"data":    newBlocks,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
