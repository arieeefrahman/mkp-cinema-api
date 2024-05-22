# Backend Test - Apps2pay
#### Name : Arief Rahman
#### Position Applied : Backend Engineer

## How To Run the Project

### Requirements
- **Golang**: go is installed and running.
- **PostgreSQL**: Ensure PostgreSQL is installed and running.
- **Redis**: Ensure Redis is installed and running.

### Environment Variables
Create a `.env` file in the root directory of the project with the following content:

```bash
APP_PORT=

DB_HOST=
DB_PORT=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=

REDIS_EXPIRED_DURATION=
REDIS_ADDRESS=
REDIS_SECRET_KEY=
```

Fill in the values.

### Installation and Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/arieeefrahman/mkp-cinema-api.git
   cd mkp-cinema-api
   ```
2. **Install Dependencies**
    ```bash
    go mod tidy
    ```
3. **Set up PostgreSQL**<br> 
    Create a database in PostgreSQL with the name specified in .env file under DB_NAME.
    
    Run the provided SQL script to set up the database schema:
    ```bash
    psql -h <DB_HOST> -U <DB_USERNAME> -d <DB_NAME> -f /mkp-cinema-api/apps2pay_script.sql
    ```

### Running the Project
1. Ensure PostgreSQL and Redis are running.
2. Verify that the environment variables are correctly set in the `.env` file.
3. Run the application:
```bash
go run main.go
```