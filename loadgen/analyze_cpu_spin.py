import pandas as pd
import matplotlib.pyplot as plt

# Load log data
log_data = []

# Use the correct path to your log file
log_file_path = '/Users/kasrahmi/Desktop/intern-project/new/intern-project/loadgen/loadgen_log.txt'

# Read the log file
with open(log_file_path, 'r') as log_file:
    current_duration = 0
    for line in log_file:
        if "Starting test with CPU-spin duration" in line:
            current_duration = int(line.split(" ")[5])
        elif "Request" in line and "Latency" in line:
            parts = line.split(", ")
            timestamp = parts[0].split(": ")[1]
            request_name = parts[1].split(": ")[1]
            response = parts[2].split(": ")[1]
            latency = int(parts[3].split(": ")[1].split(" ")[0])
            worker_id = parts[4].split(": ")[1].strip()
            log_data.append([timestamp, request_name, response, latency, worker_id, current_duration])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=["Timestamp", "Request", "Response", "Latency", "WorkerID", "Duration"])

# Convert latency to numeric
df['Latency'] = pd.to_numeric(df['Latency'])

# Calculate average latency and success rate per duration level
results = df.groupby('Duration').agg({
    'Latency': 'mean',
    'Response': lambda x: (x != "Error").mean() * 100  # Success rate as percentage
}).rename(columns={'Latency': 'AvgLatency', 'Response': 'SuccessRate'}).reset_index()

# Print results
print(results)

# Plot average latency vs CPU-spin duration
plt.figure(figsize=(12, 6))
plt.plot(results['Duration'], results['AvgLatency'], marker='o')
plt.title('Average Latency vs CPU-spin Duration')
plt.xlabel('CPU-spin Duration (ms)')
plt.ylabel('Average Latency (ms)')
plt.grid(True)
plt.show()

# Plot success rate vs CPU-spin duration
plt.figure(figsize=(12, 6))
plt.plot(results['Duration'], results['SuccessRate'], marker='o')
plt.title('Success Rate vs CPU-spin Duration')
plt.xlabel('CPU-spin Duration (ms)')
plt.ylabel('Success Rate (%)')
plt.grid(True)
plt.show()
