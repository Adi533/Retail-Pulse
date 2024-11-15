package job

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"retailpulse/utils"
	"strconv"
	"sync"
	"time"
)

// this file contains two functions which submit the job and check for each status
// in submitjob func i m submitting the data while also processing each image and if we encounter any error while this
// processing of image we put status failed in it and setup error.
// in getjobstatus i m using it to get the status of job

var (
	jobs   = make(map[int]Job) // Store jobs with their job IDs
	jobsMu sync.Mutex          // Mutex to handle concurrent access to the `jobs` map
)

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURL  []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type Job struct {
	JobID     int     `json:"job_id"`
	Visits    []Visit `json:"visits"`
	Status    string  `json:"status"`
	Error     []Error `json:"error"`
	Completed bool    `json:"completed"`
}

type Error struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Count  int     `json:"count"`
		Visits []Visit `json:"visits"`
	}

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || len(request.Visits) != request.Count {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{error:""}`))
		return
	}

	// Create a new job
	jobID := rand.Int()
	job := Job{
		JobID:  jobID,
		Visits: request.Visits,
		Status: "ongoing",
	}

	// Process each visit (download images, calculate perimeter)
	for _, visit := range job.Visits {
		for _, imageURL := range visit.ImageURL {
			err := utils.DownloadAndProcessImage(imageURL)
			if err != nil {
				job.Status = "failed"
				job.Error = append(job.Error, Error{StoreID: visit.StoreID, Error: err.Error()})
				break
			}
		}
		// Simulate random GPU processing delay
		time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	}

	// Set job as completed if no errors occurred
	if len(job.Error) == 0 {
		job.Status = "completed"
		job.Completed = true
	}

	// Store the job in the map
	jobsMu.Lock()
	jobs[jobID] = job
	jobsMu.Unlock()

	// Respond with the job ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": jobID})
}

func GetJobInfoHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "JobID is required", http.StatusBadRequest)
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(jobID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	// Fetch job data
	jobsMu.Lock()
	job, exists := jobs[id]
	jobsMu.Unlock()

	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	// Prepare the response based on job status
	var response interface{}
	switch job.Status {
	case "completed", "ongoing":
		response = struct {
			Status string `json:"status"`
			JobID  int    `json:"job_id"`
		}{
			Status: job.Status,
			JobID:  job.JobID,
		}

	case "failed":
		response = struct {
			Status string  `json:"status"`
			JobID  int     `json:"job_id"`
			Error  []Error `json:"error"`
		}{
			Status: job.Status,
			JobID:  job.JobID,
			Error:  job.Error,
		}

	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
