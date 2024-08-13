import pandas as pd
import matplotlib.pyplot as plt

# Load log data
log_data = []

# Use the correct path to your log file
log_file_path = '/Users/kasrahmi/Desktop/intern-project/new/intern-project/loadgen/loadgen_log.txt'

# Read the log file
with open(log_file_path, 'r') as log_file:
    for line in log_file:
        if "Request" in line and "Latency" in line:
            parts = line.split(", ")
            timestamp = parts[0].split(": ")[1]
            rps = int(parts[1].split(": ")[1])
            distribution = parts[2].split(": ")[1]
            duration = int(parts[3].split(": ")[1].replace(
                " ms", ""))  # Remove " ms" and convert to int
            request_name = parts[4].split(": ")[1]
            response = parts[5].split(": ")[1]
            # Extract latency as an integer
            latency = int(parts[6].split(": ")[1].split(" ")[0])
            worker_id = parts[7].split(": ")[1].strip()
            log_data.append([timestamp, rps, distribution, duration,
                            request_name, response, latency, worker_id])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=[
                  "Timestamp", "RPS", "Distribution", "Duration", "Request", "Response", "Latency", "WorkerID"])

# Convert latency to numeric
df['Latency'] = pd.to_numeric(df['Latency'])

# Calculate average latency and success rate per RPS level
avg_latency = df['Latency'].mean()
success_rate = (df['Response'] != "Error").mean()

print(f"Average Latency: {avg_latency:.2f} ms")
print(f"Success Rate: {success_rate * 100:.2f}%")

# Plot latency histogram
plt.figure(figsize=(10, 6))
plt.hist(df['Latency'], bins=20, color='blue', edgecolor='black')
plt.title('Histogram of Latencies')
plt.xlabel('Latency (ms)')
plt.ylabel('Frequency')
plt.show()

# Group by worker ID to see distribution of requests
worker_distribution = df['WorkerID'].value_counts()

# Plot worker distribution
plt.figure(figsize=(10, 6))
worker_distribution.plot(kind='bar')
plt.title('Request Distribution by Worker ID')
plt.xlabel('Worker ID')
plt.ylabel('Number of Requests')
plt.show()

# Analyze results for each RPS level
rps_grouped = df.groupby('RPS').agg(
    {'Latency': 'mean', 'Response': lambda x: (x != "Error").mean()})
rps_grouped.columns = ['Avg Latency', 'Success Rate']

# Plot average latency per RPS level
plt.figure(figsize=(10, 6))
rps_grouped['Avg Latency'].plot(kind='line', marker='o')
plt.title('Average Latency per RPS Level')
plt.xlabel('RPS')
plt.ylabel('Average Latency (ms)')
plt.grid(True)
plt.show()

# Plot success rate per RPS level
plt.figure(figsize=(10, 6))
rps_grouped['Success Rate'].plot(kind='line', marker='o', color='green')
plt.title('Success Rate per RPS Level')
plt.xlabel('RPS')
plt.ylabel('Success Rate')
plt.grid(True)
plt.show()
