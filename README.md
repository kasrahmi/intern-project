# Load Testing Framework for Knative Autoscaling

## Project Overview

This project involves developing a load testing framework to evaluate the performance and scalability of distributed systems deployed on Knative. The framework uses a gRPC-based architecture for load testing, enabling parameterized testing to assess the impact of various configurations on performance metrics such as latency and success rate.

## Objectives

- **Develop a Load Testing Framework**
- **Deploy and Evaluate Knative's Autoscaling**
- **Conduct Parameterized Load Testing**
- **Collect and Analyze Performance Data**
- **Visualize and Interpret Results**

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (1.16+)
- [Docker](https://www.docker.com/get-started)
- [Knative](https://knative.dev/docs/install/)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/)

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/kasrahmi/intern-project.git
   cd intern-project
2. **Build the Load Generator**:
   ```bash
   cd loadgen
   go build -o loadgen main.go
3. **Build the Worker**:
   ```bash
   cd ../worker
   go build -o worker main.go
#### Or you can use the docker images :
1.  **Pull images**:
   ```bash
   docker pull docker.io/kasrahmi/loadgen-intern
   docker pull docker.io/kasrahmi/worker-intern
   ```
### Deployment

1. **Set Up a Single-Node Knative Cluster**: Follow the [vHive quickstart](https://github.com/vhive-serverless/vHive/blob/main/docs/quickstart_guide.md#1-deploy-functions) guide to set up the cluster.

2. **Deploy Load Generator**:
    ```bash
    kubectl apply -f kubernetes/loadgen-service.yaml
    kubectl apply -f kubernetes/loadgen-deployment.yaml
    ```
    You can find these two yaml files in /loadgen/yamls directory.

3. **Deploy Worker Functions**:
    ```bash
    kn service create worker-service \
    	--image=docker.io/kasrahmi/worker-intern:latest \
    	--port=50052
    ```

4. **Check Deployment Status**:
    ```bash
    kubectl get pods
    kubectl logs <loadgen-pod-name>
    ```
### Local Usage

1. **Run the Load Generator Locally**:

    ```bash
    ./loadgen --rps=10 --duration=10s --distribution=uniform
    ```
2.  **Run workers Locally** :
   	```bash
    ./worker --port=50052
    ```
    ```bash
    ./worker --port=50053
    ```

4. **Access Logs**: Logs are stored in the `logs/` directory for post-run analysis.

5. **Visualize Data**: Use the `data_analyzis.py` script to generate plots:

    ```bash
    python data_analyzis.py
    ```
### Experimentation

1. **Parameter Sweep**:
    - Vary RPS from 5 to 50 (steps of 5).
    - Test invocation distribution (Uniform vs. Poisson).
    - Adjust CPU-spin duration from 100 ms to 1 second (steps of 100 ms).
    - Run a grid search of these parameters.

2. **Data Collection**:
    - Collect performance data for 12 minutes, excluding the first 4 minutes for warm-up.
    - Refine metrics to calculate end-to-end (E2E) slowdown.

3. **Data Visualization**:
    - Analyze metrics using Python scripts.
    - Visualize data with graphs showing latency, success rate, and E2E slowdown.

### Results and Analysis

1. **Average Latency**:
    - **Impact of Invocation Distribution**: Poisson distribution generally results in higher average latency compared to Uniform distribution.
    - **Effect of RPS**: Average latency increases with higher RPS, more pronounced with Poisson distribution.
    - **Influence of CPU-spin Duration**: Longer CPU-spin durations lead to higher latency, especially under Poisson distribution.

2. **E2E Slowdown**:
    - **Impact of Invocation Distribution**: Poisson distribution shows a significantly higher E2E slowdown compared to Uniform distribution.
    - **Effect of RPS**: Increased RPS results in greater E2E slowdown, more noticeable under Poisson distribution.
    - **Influence of CPU-spin Duration**: Longer CPU-spin durations exacerbate E2E slowdown, particularly with Poisson-distributed requests.

## Acknowledgments

- [Knative](https://knative.dev/docs/)
- gRPC
- Docker
- Kubernetes
