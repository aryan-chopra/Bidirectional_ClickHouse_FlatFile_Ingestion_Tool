# Bidirectional ClickHouse ↔ FlatFile Ingestion Tool

A full-stack data ingestion tool built with **React** and **Go (Gin)** that supports uploading CSV files to a **ClickHouse** database and downloading data from ClickHouse as CSV. Designed for high-speed, large-volume data pipelines with support for column selection, chunked processing, and real-time feedback.

---

##  Features

-  **Bidirectional ingestion** between flat CSV files and ClickHouse.
-  **Upload FlatFile** to ClickHouse with selectable columns.
-  **Streamed upload** in chunks for large datasets.
-  **Automatic type inference** to create ClickHouse table schema.
-  **Download ClickHouse tables** as FlatFile in batches.
-  **User-friendly frontend** with live status and progress.

---

##  Tech Stack

- **Frontend**: React, TailwindCSS, PapaParse
- **Backend**: Go (Gin), ClickHouse Go Client
- **Database**: ClickHouse (Cloud or Local)

---

##  Usage Instructions

###  Uploading a CSV to ClickHouse

1. **Choose source as File:**
   - Click the **"Upload"** tab.
   - Select a file from your computer.
   - The file will be parsed using PapaParse in the browser.
2. **Select Columns:**
   - Choose the columns you want to include in the upload.
3. **Enter ClickHouse Details:**
   - Host: `localhost` or your ClickHouse cloud URL
   - Port: `9000` (default for native), or `8123` (HTTP)
   - Username & Password
   - Database name
   - Table name
4. **Start Upload:**
   - Click the **"Upload"** button.
   - The frontend sends data in chunks.
   - A status window shows real-time progress.

> On success, a new table will be visible in your ClickHouse database.

---

###  Downloading Data from ClickHouse

1. **Choose source as ClickHouse"** tab.
2. **Enter ClickHouse Credentials:**
   - Fill in your database connection details.
3. **Fetch Tables:**
   - Click **"Fetch Tables"** to get a list of tables in the selected database.
4. **Select a Table:**
   - Choose the table you want to export.
5. **Download:**
   - Click **"Download"**.
   - The data is fetched from ClickHouse in chunks and compiled into a single file.
   - The file will be automatically downloaded when complete.
  
##  Setup Instructions

### Clone the Repository

```bash
git clone https://github.com/aryan-chopra/Bidirectional_ClickHouse_FlatFile_Ingestion_Tool.git
cd Bidirectional_ClickHouse_FlatFile_Ingestion_Tool
```

### Backend Setup (Go + Gin)
1. **Install Go (if not already)**

2. **Navigate to the backend folder:**
```bash
cd backend
```
3. **Install dependencies:**

```bash
go mod tidy
```

4. **Run the backend server:**
```bash
go run .
```
Server will run on http://localhost:8080

### Frontend Setup (React + Vite)
1. **Navigate to the frontend folder:**
```bash
cd ../frontend
```

2. **Install dependencies:**
```bash
npm install
```

3. **Run the development server:**

```bash
npm run dev
```
App will be available at http://localhost:5173
