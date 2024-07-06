package main

import (
    "encoding/json"
    "net/http"
    "log"
    "github.com/google/uuid"
    "strings"
)


// Global map to store receipt data in memory
var receipts = make(map[string]int)

// Handler for processing receipts
func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
    var receipt Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    id := uuid.NewString() // Generate a new unique ID
    points := CalculatePoints(receipt) // Calculate points
    receipts[id] = points // Store points with ID as key
    log.Printf("Receipt processed and stored with ID %s and Points %d", id, points)
    json.NewEncoder(w).Encode(map[string]string{"id": id}) // Respond with ID
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 4 || parts[3] != "points" {
        http.Error(w, "Invalid path", http.StatusBadRequest)
        return
    }
    id := parts[2]
    log.Printf("Attempting to retrieve points for ID: %s", id)
    if points, exists := receipts[id]; exists {
        json.NewEncoder(w).Encode(PointsResponse{Points: points})
    } else {
        log.Printf("Failed to find points for ID: %s", id)
        http.Error(w, "Receipt not found", http.StatusNotFound)
    }
}

func main() {
    http.HandleFunc("/receipts/process", processReceiptHandler)
    http.HandleFunc("/receipts/", getPointsHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}