#!/bin/bash

# Color codes for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

API_URL="http://localhost:8080"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}    API Testing Script - Kasir API${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Test 1: Health Check
echo -e "${YELLOW}[Test 1] Health Check${NC}"
echo "GET /healthz"
curl -X GET "$API_URL/healthz"
echo -e "\n${GREEN}✓ Health check passed${NC}\n"

# Test 2: Get all products
echo -e "${YELLOW}[Test 2] Get All Products${NC}"
echo "GET /api/products"
curl -X GET "$API_URL/api/products"
echo -e "\n${GREEN}✓ Products retrieved${NC}\n"

# Test 3: Create a category
echo -e "${YELLOW}[Test 3] Create Category${NC}"
echo "POST /api/categories"
CATEGORY_RESPONSE=$(curl -X POST "$API_URL/api/categories" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Makanan Ringan"
  }')
echo "$CATEGORY_RESPONSE"
CATEGORY_ID=$(echo "$CATEGORY_RESPONSE" | grep -o '"id":[0-9]*' | head -1 | grep -o '[0-9]*')
echo -e "\n${GREEN}✓ Category created with ID: $CATEGORY_ID${NC}\n"

# Test 4: Get all categories
echo -e "${YELLOW}[Test 4] Get All Categories${NC}"
echo "GET /api/categories"
curl -X GET "$API_URL/api/categories"
echo -e "\n${GREEN}✓ Categories retrieved${NC}\n"

# Test 5: Create a product
echo -e "${YELLOW}[Test 5] Create Product${NC}"
echo "POST /api/products"
PRODUCT_RESPONSE=$(curl -X POST "$API_URL/api/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng",
    "price": 5000,
    "stock": 50,
    "category_id": 1
  }')
echo "$PRODUCT_RESPONSE"
echo -e "\n${GREEN}✓ Product created${NC}\n"

# Test 6: Get product by ID
echo -e "${YELLOW}[Test 6] Get Product by ID${NC}"
echo "GET /api/products/1"
curl -X GET "$API_URL/api/products/1"
echo -e "\n${GREEN}✓ Product retrieved${NC}\n"

# Test 7: Update product
echo -e "${YELLOW}[Test 7] Update Product${NC}"
echo "PUT /api/products/1"
curl -X PUT "$API_URL/api/products/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Indomie Goreng Premium",
    "price": 6000,
    "stock": 60,
    "category_id": 1
  }'
echo -e "\n${GREEN}✓ Product updated${NC}\n"

# Test 8: Create checkout (transaction)
echo -e "${YELLOW}[Test 8] Create Checkout/Transaction${NC}"
echo "POST /api/checkout"
CHECKOUT_RESPONSE=$(curl -X POST "$API_URL/api/checkout" \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {
        "product_id": 1,
        "quantity": 3
      }
    ]
  }')
echo "$CHECKOUT_RESPONSE"
echo -e "\n${GREEN}✓ Transaction created${NC}\n"

# Test 9: Get daily sales report
echo -e "${YELLOW}[Test 9] Get Daily Sales Report (Today)${NC}"
echo "GET /api/report/hari-ini"
curl -X GET "$API_URL/api/report/hari-ini"
echo -e "\n${GREEN}✓ Daily report retrieved${NC}\n"

# Test 10: Get sales report by date range
echo -e "${YELLOW}[Test 10] Get Sales Report by Date Range${NC}"
echo "GET /api/report?start_date=2026-01-01&end_date=2026-02-28"
curl -X GET "$API_URL/api/report?start_date=2026-01-01&end_date=2026-02-28"
echo -e "\n${GREEN}✓ Range report retrieved${NC}\n"

# Test 11: Delete category
echo -e "${YELLOW}[Test 11] Delete Category${NC}"
echo "DELETE /api/categories/5"
curl -X DELETE "$API_URL/api/categories/5"
echo -e "\n${GREEN}✓ Category deleted${NC}\n"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}    All Tests Completed!${NC}"
echo -e "${BLUE}========================================${NC}"
