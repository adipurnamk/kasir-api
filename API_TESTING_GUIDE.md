# API Testing Guide for Kasir API

## Prerequisites
- Make sure the server is running on `http://localhost:8080`
- Have PostgreSQL database running with proper schema
- All tables should be created: `products`, `categories`, `transactions`, `transaction_details`

## Quick Test using curl

### 1. Health Check
```bash
curl -X GET http://localhost:8080/healthz
```

### 2. Create a Category
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Makanan Ringan"
  }'
```

### 3. Get All Categories
```bash
curl -X GET http://localhost:8080/api/categories
```

### 4. Create a Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng",
    "price": 5000,
    "stock": 50,
    "category_id": 1
  }'
```

### 5. Get All Products
```bash
curl -X GET http://localhost:8080/api/products
```

### 6. Get Product by ID
```bash
curl -X GET http://localhost:8080/api/products/1
```

### 7. Update Product
```bash
curl -X PUT http://localhost:8080/api/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng Premium",
    "price": 6000,
    "stock": 60,
    "category_id": 1
  }'
```

### 8. Delete Product (if supported)
```bash
curl -X DELETE http://localhost:8080/api/products/1
```

### 9. Create a Transaction (Checkout)
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "product_id": 1,
        "quantity": 3
      },
      {
        "product_id": 2,
        "quantity": 2
      }
    ]
  }'
```

### 10. Get Today's Sales Report
```bash
curl -X GET http://localhost:8080/api/report/hari-ini
```

**Response:**
```json
{
  "total_revenue": 45000,
  "total_transaksi": 5,
  "produk_terlaris": {
    "nama": "Indomie Goreng",
    "qty_terjual": 12
  }
}
```

### 11. Get Sales Report by Date Range
```bash
curl -X GET "http://localhost:8080/api/report?start_date=2026-01-01&end_date=2026-02-28"
```

**Response:**
```json
[
  {
    "tanggal": "2026-02-08",
    "total_revenue": 45000,
    "total_transaksi": 5,
    "produk_terlaris": {
      "nama": "Indomie Goreng",
      "qty_terjual": 12
    }
  }
]
```

## Run Automated Tests

Make the test script executable and run it:

```bash
chmod +x test-api.sh
./test-api.sh
```

This will run all API tests in sequence with formatted output.

## Alternative: Using Postman

1. Import the API endpoints into Postman
2. Create a collection with the following requests:
   - Health Check (GET)
   - Create Category (POST)
   - Get All Categories (GET)
   - Create Product (POST)
   - Get All Products (GET)
   - Get Product by ID (GET)
   - Update Product (PUT)
   - Create Transaction (POST)
   - Get Daily Sales Report (GET)
   - Get Sales Report by Range (GET)

## Notes

- Ensure database tables have `created_at` column with default `CURRENT_TIMESTAMP` for transactions
- Product stock is automatically decremented during checkout
- The best-selling product is calculated based on total quantity sold
- Date format for reports should be `YYYY-MM-DD`
