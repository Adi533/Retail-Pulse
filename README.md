# Retail Pulse Documentation

## 1. **Description**

This project is a Go-based web service designed to handle jobs related to store visits. It provides functionality for submitting jobs and retrieving their status. The service processes each job by downloading images associated with store visits and reports any errors encountered during processing. The service tracks the status of each job, including whether it was completed successfully or failed due to issues such as missing store data or failed image downloads.

### Key Features:
- **Submit Job**: Accepts a set of store visits with associated image URLs and starts processing the job.
- **Track Job Status**: Allows users to query the status of a job using its job ID. Returns whether the job is ongoing, completed, or failed.
- **Error Reporting**: Provides detailed error information if any part of the job fails, including missing store data or issues with image downloads.

---

## 2. **Assumptions**

- The system assumes that each store visit contains a valid `store_id` and at least one image URL. The `store_id` must be a valid entry in the predefined store data.
- The `gorilla/mux` package is used for routing and handling HTTP requests.
- The Go environment is set up correctly, and the system uses Go modules for dependency management.
- The project is expected to run in a typical development or production environment where the required ports are open for communication.

---

## 3. **Installing (Setup) and Testing Instruction**

### **Installation**

1. **Clone the Repository**:
     ```bash
     git clone https://github.com/Adi533/Retail-pulse.git
     ```

2. **Navigate to the Project Folder**:
   ```bash
   cd retailpulse

3. **Install Dependencies:**
   ```bash
     go mod tidy
     ```
4. **Build the Project:**

    ```bash
      go build -o main .
    ```
5. **Run the Application:**

    ```bash
      ./main
    ```
The server will be running on http://localhost:8080.

### **Testing the API**
1. **Submit a Job:**

- To submit a job, send a POST request to /submit-job with the following JSON body:
```json
{
  "count": 2,
  "visits": [
    {
      "store_id": "S00339218",
      "image_url": ["http://example.com/image1.jpg"],
      "visit_time": "2024-11-15T12:00:00Z"
    },
    {
      "store_id": "S01408764",
      "image_url": ["http://example.com/image2.jpg"],
      "visit_time": "2024-11-15T12:30:00Z"
    }
  ]
}
```
2. **Get Job Status:**

- To get the status of a job, send a GET request to /get-job-info?jobid={job_id}, replacing {job_id} with the actual job ID returned when you submitted the job.

## 4. **Brief description of the work environment used to run this project**
1. **Computer/Operating System:**
- OS: Windows (can also work on mac or linux).
- CPU: Amd ryzen 7.
- RAM: 16 GB.
2. **Text Editor/IDE:**
- Visual Studio Code: A lightweight and powerful code editor with extensions for Go development.
3. **Libraries/Packages:**
- Go Version: 1.21
- gorilla/mux: A routing library for Go.
- Go modules: Dependency management for Go projects.
- Alpine Linux: A minimal Docker base image used for the containerized environment.

## 5. **Improvements for the Future**
If given more time, the following improvements could be made:

1. **Error Handling and Logging:**
Improve error handling by adding more detailed logs for debugging purposes, especially when dealing with network or image download failures.
Implement centralized logging with tools like logrus or zap for better log management and analysis.
2. **Multi-Threaded Processing:**
Introduce concurrency for handling multiple store visits in parallel. This will improve the performance, especially when there are large numbers of images to process.
3. **API Enhancements:**
Implement pagination for job retrieval, especially when the number of jobs grows large.
Add authentication and authorization for securing the API, ensuring that only authorized users can submit jobs or check job statuses.
4. **Testing and CI/CD:**
Implement automated tests (unit and integration tests) to ensure code quality and reduce the likelihood of regressions.
Set up a Continuous Integration/Continuous Deployment (CI/CD) pipeline using GitHub Actions or similar tools for automated testing, building, and deployment.
5. **Frontend Interface:**
Develop a frontend web interface to allow users to submit jobs and track their statuses via a user-friendly interface, instead of just using the raw API endpoints.
6. **Dockerization:**
Although the Dockerfile is present, further improve the Docker setup by using a multi-stage build to reduce the size of the final image.
7. **Error Handling for Missing Store IDs:**
Extend the error handling for missing store IDs or other incomplete data to make the system more fault-tolerant.

