import pandas as pd
import matplotlib.pyplot as plt
import os

# Load log data
log_data = []

# Use the correct path to your log file
log_file_path = '/Users/kasrahmi/Desktop/intern-project/new/intern-project/loadgen/loadgen_log_rps.txt'

# Read the log file
with open(log_file_path, 'r') as log_file:
    current_rps = None
    for line in log_file:
        if "Starting test with" in line:
            # Extract RPS from log line
            current_rps = int(line.split(" ")[3])
        elif "Request" in line and "Latency" in line:
            parts = line.split(", ")
            timestamp = parts[0].split(": ")[1]
            request_name = parts[1].split(": ")[1]
            response = parts[2].split(": ")[1]
            latency = int(parts[3].split(": ")[1].split(" ")[0])
            worker_id = parts[4].split(": ")[1].strip()
            log_data.append([current_rps, timestamp, request_name, response, latency, worker_id])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=["RPS", "Timestamp", "Request", "Response", "Latency", "WorkerID"])

# Convert latency to numeric
df['Latency'] = pd.to_numeric(df['Latency'])

# Group by RPS
grouped = df.groupby('RPS')

# Initialize lists for plotting
rps_levels = []
avg_latencies = []
success_rates = []

# Calculate metrics for each RPS level
for rps, group in grouped:
    avg_latency = group['Latency'].mean()
    success_rate = (group['Response'] != "Error").mean() * 100
    rps_levels.append(rps)
    avg_latencies.append(avg_latency)
    success_rates.append(success_rate)
    print(f"RPS {rps}: Average Latency = {avg_latency:.2f} ms, Success Rate = {success_rate:.2f}%")

# Plot average latency vs RPS
plt.figure(figsize=(10, 6))
plt.plot(rps_levels, avg_latencies, marker='o')
plt.title('Average Latency vs RPS')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Average Latency (ms)')
plt.grid(True)
plt.show()

# Plot success rate vs RPS
plt.figure(figsize=(10, 6))
plt.plot(rps_levels, success_rates, marker='o', color='green')
plt.title('Success Rate vs RPS')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Success Rate (%)')
plt.grid(True)
plt.show()
