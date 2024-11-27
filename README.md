
# Image-Based Expense Tracker

A Go-based backend service that processes receipt images to extract expense details like vendor, total amount, and date using OCR. The application stores the extracted data, provides APIs for retrieval, and generates expense reports in CSV or PDF format.

---

## Features

- **Upload Receipts:** Accepts receipt images for processing.
- **OCR Integration:** Uses OCR to extract text from images.
- **Data Parsing:** Extracts structured data (vendor, amount, date) from OCR results.
- **Expense Storage:** Stores expenses in a PostgreSQL database.
- **Report Generation:** Generates expense reports in CSV and PDF formats.
- **Scalable Architecture:** Modular and extensible design.

---

## Project Layout

```plaintext
image-expense-tracker/
├── cmd/
│   └── app/
│       └── main.go                 # Main entry point for the application
├── config/
│   └── config.go                   # Configuration builder and related logic
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── upload_handler.go   # Handles receipt upload
│   │   │   ├── expense_handler.go  # Handles expense retrieval
│   │   │   ├── report_handler.go   # Handles report generation
│   │   └── router.go               # Sets up routes and middleware
│   ├── core/
│   │   ├── expense/
│   │   │   ├── models.go           # Expense data model
│   │   │   ├── repository.go       # Repository interface and implementations
│   │   │   ├── service.go          # Business logic for expense operations
│   │   ├── ocr/
│   │   │   ├── processor.go        # OCRProcessor interface and strategies
│   │   │   ├── tesseract.go        # Tesseract implementation
│   │   │   ├── google_vision.go    # Google Vision implementation
│   │   └── parser/
│   │       └── expense_parser.go   # Facade for parsing OCR results into structured data
│   └── report/
│       ├── csv_report.go           # CSV report generator
│       ├── pdf_report.go           # PDF report generator
│       └── generator.go            # Template Method for report generation
├── pkg/
│   └── utils/
│       ├── logger.go               # Centralized logger
│       ├── validation.go           # Input validation utilities
│       ├── storage.go              # File storage utilities
│       └── response.go             # Response helper functions
├── test/
│   └── integration/
│       ├── expense_test.go         # Integration tests for expense operations
│       └── upload_test.go          # Integration tests for receipt uploads
├── Dockerfile                      # Dockerfile for containerizing the application
├── go.mod                          # Go module dependencies
├── go.sum                          # Go module dependency checksums
├── README.md                       # Project documentation
└── .env                            # Environment variables for local dev
```

---

## Getting Started

### Prerequisites

- Go 1.18+ installed
- PostgreSQL database
- Tesseract OCR (or Google Vision API key for OCR)

### Running the Project

1. Clone the repository:
   ```bash
   git clone https://github.com/Rashad-j/image-based-expense-tracker.git
   cd image-expense-tracker
   ```

2. Set up environment variables:
   - Export your `.env` variables 

3. Build and run the application:
   ```bash
   go run ./cmd/app/main.go
   ```

---

## License

This project is licensed under the MIT License.

---

## Contributing

Feel free to fork this repository, open issues, or submit pull requests. Contributions are welcome!
