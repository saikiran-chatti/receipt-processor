
# Receipt Processor

## Overview
This Receipt Processor is a web service that calculates and manages points for retail receipts. It provides an API for processing receipts and retrieving point totals based on specific rules.

## Features
- Process receipts and calculate points based on predefined rules
- Generate unique IDs for each processed receipt
- Retrieve point totals for processed receipts

## API Endpoints

### Process Receipt
- **URL:** `/receipts/process`
- **Method:** `POST`
- **Request Body:** JSON receipt data
- **Response:** JSON object containing the generated receipt ID

Example response:
```json
{
  "id": "7c32965b-fda3-4396-b784-28ccfa5a736b"
}
```

### Get Points
- **URL:** `/receipts/{id}/points`
- **Method:** `GET`
- **Response:** JSON object with the points awarded for the receipt

Example response:
```json
{
  "points": 32
}
```

## Setup and Running

### Prerequisites
- Go 1.16 or later
- Docker (optional)

### Running Locally
1. Clone the repository:
   ```
   git clone https://github.com/yourusername/receipt-processor.git
   cd receipt-processor
   ```

2. Build and run the application:
   ```
   go build
   ./receipt-processor
   ```

The service will start on `localhost:8080`.

### Running with Docker
1. Build the Docker image:
   ```
   docker build -t receipt-processor .
   ```

2. Run the container:
   ```
   docker run -p 8080:8080 --name receipt-processor-container receipt-processor
   ```

## Usage Examples

### Processing a Receipt
```bash
curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      }
    ],
    "total": "6.49"
  }'
```

### Retrieving Points
```bash
curl http://localhost:8080/receipts/7c32965b-fda3-4396-b784-28ccfa5a736b/points
```
