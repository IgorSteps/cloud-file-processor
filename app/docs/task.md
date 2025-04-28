# Cloud Software Engineer Take-Home Task (Mock)

## Event-Driven File Processing Pipeline

### Background

You're joining a company that needs to process user-uploaded CSV files containing transaction records. When a user uploads a file, the system should:

1. Store the file securely.
2. Trigger background processing (e.g. parse, validate, summarize).
3. Return a processing status to the user.
4. Notify the user when processing is complete.

### Requirements

#### Functional

- Provide a **REST API** with the following endpoints:
  - `POST /upload`: Accepts a CSV file upload.
  - `GET /status/{fileId}`: Returns the processing status of a file.
  - `GET /summary/{fileId}`: Returns the processed summary (e.g., total transactions, sum of amounts).

#### Processing

- The backend should:
  - Store the raw file in **cloud object storage** (e.g. AWS S3, GCS, or MinIO if you're local).
  - Trigger a **background worker** to process the file (can be async function, queue + worker, or serverless).
  - Store results in a **cloud-native database** (e.g. DynamoDB, Firestore, or PostgreSQL on RDS).

#### Infra

- Use **Infrastructure-as-Code** (Terraform, Pulumi, CDK, or similar) to define:
  - API deployment
  - Storage and database resources
  - Background processing infra (queue, lambda, container, etc.)

#### Optional (Bonus)

- Add **authentication** (API key or OAuth)
- Include **metrics** and **logging**
- Add **unit/integration tests**
- Deploy using **CI/CD** (e.g. GitHub Actions, Terraform Cloud)

### 📦 Deliverables

1. Link to GitHub repo (or a zip file with your project)
2. A **README** that includes:
   - Setup instructions
   - Your architecture explanation
   - How to run and test locally
3. Any **design decisions** or tradeoffs
