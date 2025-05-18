# SalesTracker
A backend system built with Golang and MariaDB to track and analyze sales data. 

## Prerequisites 

- **Golang:** v1.24+
- **MariaDb:** v11.5+
- **Git**
- **Postman**

## Getting Started

#### 1. Clone a Repository

```bash
git@github.com:Partheeban-R/SalesTracker.git
```

#### 1. Configure The Database

```toml
#configure the database with your credientials
MariaDBServer    = "localhost" # database hosted ip
MariaDBPort      = 3306  # hosted port
MariaDBUser      = "root"
MariaDBPassword  = "root"
MariaDBDatabase  ="YourDatabaseName"
MariaDBDBType    ="mysql"
```


#### 3. Initialize & Run

```bash
go mod tidy
go run main.go
```

Server start on : `http://localhost:29091`

#### 4.CSV File Upload

Upoad the Csv file under `SalesData/` directory the program will automatically fetch the data on the desire automated time or through api trigger.

#### 4.Log File

To track the program execution and error debugging the log file is created on when application start or restart.
Auto Sync Log maintained in the `SyncTracking` table 
```sql
select * from SyncTracking 
order by id desc;
```


### Api Reference

| Endpoint                         | Method | Params                        |  Header                                   | Description                                                      
|----------------------------------|--------|-------------------------------|-------------------------------------------|-------------------------------------------                       
| `/SyncData`                      | GET    | None                          |   None                                    | Data Syncing through api trigger.                                  
| `/getRevenue/total`              | GET    | None                          |   FROMDT = YYYY-MM-DD ,TODT = YYYY-MM-DD  | Total Revenue for give date range.                                                
| `/getRevenue/product`            | GET    | None                          |   FROMDT = YYYY-MM-DD ,TODT = YYYY-MM-DD  | Product wise revenue obtained for give date range.                                            
| `/getRevenue/category`           | GET    | None                          |   FROMDT = YYYY-MM-DD ,TODT = YYYY-MM-DD  | Category wise revenue obtained for give date range.                                                
| `/getRevenue/region`             | GET    | None                          |   FROMDT = YYYY-MM-DD ,TODT = YYYY-MM-DD  | Region wise revenue obtained for give date range.                                                


## Sample API Requests 
1. **SyncData**: Api Trigger for Data syncing. 
Api Url: <http://localhost:29091/SyncData>
Method: GET
Header : None

**Sample Response** 
```json
{
    "status": "S",
    "errMsg": ""
}
```
2. **Total Revenue**: Total Revenue for give date range. 
Api Url: <http://localhost:29091/getRevenue/total>
Method: GET
Header : ` FROMDT = 2025-01-01` , `TODT = 2025-02-02`

**Sample Response** 
```json
{
    "TotalRevenue": 5036.46,
    "status": "S",
    "errMsg": ""
}
```
3. **Product Revenue**: Product based Revenue for give date range. 
Api Url: <http://localhost:29091/getRevenue/product>
Method: GET
Header : ` FROMDT = 2025-01-01` , `TODT = 2025-02-02`

**Sample Response** 
```json
{
    "ProductRevenue": [
        {
            "Product": "UltraBoost Running Shoes",
            "Revenue": 557.9
        },
        {
            "Product": "Sony WH-1000XM5 Headphones",
            "Revenue": 361.84
        },
        {
            "Product": "iPhone 15 Pro",
            "Revenue": 3931.95
        },
        {
            "Product": "Levi's 501 Jeans",
            "Revenue": 184.77
        }
    ],
    "status": "S",
    "errMsg": ""
}
```
4. **Category Revenue**: Category based Revenue for give date range. 
Api Url: <http://localhost:29091/getRevenue/category>
Method: GET
Header : ` FROMDT = 2025-01-01` , `TODT = 2025-02-02`

**Sample Response** 
```json
{
    "CategoryRevenue": [
        {
            "Revenue": 184.77,
            "Category": "Clothing"
        },
        {
            "Revenue": 4293.79,
            "Category": "Electronics"
        },
        {
            "Revenue": 557.9,
            "Category": "Shoes"
        }
    ],
    "status": "S",
    "errMsg": ""
}
```
5. **Region Revenue**: Region based Revenue for give date range. 
Api Url: <http://localhost:29091/getRevenue/region>
Method: GET
Header : ` FROMDT = 2025-01-01` , `TODT = 2025-02-02`

**Sample Response** 
```json
{
    "RegionRevenue": [
        {
            "Revenue": 1675.84,
            "Region": "Europe"
        },
        {
            "Revenue": 554.67,
            "Region": "North America"
        },
        {
            "Revenue": 2805.95,
            "Region": "South America"
        }
    ],
    "status": "S",
    "errMsg": ""
}
```
