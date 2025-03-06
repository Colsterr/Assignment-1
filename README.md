# 🌍 Country Information API (Assignment-1)
This is a Go-based REST API that provides information about countries, including general details and historical population data.

## 🚀 Deployment
This service is deployed on Render and accessible at:
🔗 **Render URL**: [https://assignment-1-hian.onrender.com](https://assignment-1-hian.onrender.com)

---

## 📌 API Endpoints

### 1️⃣ **Country Info**
#### 📌 Fetch country details by ISO 3166-2 code
- **URL**: `/countryinfo/v1/info/{country_code}`
- **Example (Norway)**:  
  🔗 [Get Norway Info](https://assignment-1-hian.onrender.com/countryinfo/v1/info/no)

#### 📌 Fetch country details with city limit
- **URL**: `/countryinfo/v1/info/{country_code}?limit={n}`
- **Example (Limit to 10 cities for Norway)**:  
  🔗 [Get Norway Info (10 Cities)](https://assignment-1-hian.onrender.com/countryinfo/v1/info/no?limit=10)

---

### 2️⃣ **Country Population**
#### 📌 Fetch historical population data
- **URL**: `/countryinfo/v1/population/{country_code}`
- **Example (Norway)**:  
  🔗 [Get Norway Population](https://assignment-1-hian.onrender.com/countryinfo/v1/population/no)

#### 📌 Fetch population within a year range
- **URL**: `/countryinfo/v1/population/{country_code}?limit={start-end}`
- **Example (Norway from 2010 to 2015)**:  
  🔗 [Get Norway Population (2010-2015)](https://assignment-1-hian.onrender.com/countryinfo/v1/population/no?limit=2010-2015)

---

### 3️⃣ **Service Status**
#### 📌 Check API health & uptime
- **URL**: `/countryinfo/v1/status`
- **Example**:  
  🔗 [Get API Status](https://assignment-1-hian.onrender.com/countryinfo/v1/status)

---

## 🔧 Setup & Running Locally
### 1️⃣ Install Go
Ensure Go is installed and working (`go version` should return Go 1.20+).

### 2️⃣ Clone Repository
```sh
git clone https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2025-workspace/servank/assignment-1.git
cd assignment-1
