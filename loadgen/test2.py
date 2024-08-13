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
            duration = int(parts[3].split(": ")[1].replace(" ms", ""))  # Remove " ms" and convert to int
            request_name = parts[4].split(": ")[1]
            response = parts[5].split(": ")[1]
            latency = int(parts[6].split(": ")[1].split(" ")[0])  # Extract latency as an integer
            worker_id = parts[7].split(": ")[1].strip()
            log_data.append([timestamp, rps, distribution, duration, request_name, response, latency, worker_id])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=["Timestamp", "RPS", "Distribution", "Duration", "Request", "Response", "Latency", "WorkerID"])

# Convert Timestamp to datetime and filter out the first 4 minutes
df['Timestamp'] = pd.to_datetime(df['Timestamp'])
start_time = df['Timestamp'].min() + pd.Timedelta(minutes=4)
filtered_df = df[df['Timestamp'] >= start_time]

# Verify the filtering
print(f"Data starts from: {filtered_df['Timestamp'].min()}")

# Calculate average latency per combination of RPS, Distribution, and Duration
e2e_slowdown = filtered_df.groupby(['RPS', 'Distribution', 'Duration']).agg({
    'Latency': 'mean'
}).reset_index()

# Rename the column for clarity
e2e_slowdown.columns = ['RPS', 'Distribution', 'Duration', 'Avg Latency']

print(e2e_slowdown)

import seaborn as sns

# Set up the matplotlib figure
plt.figure(figsize=(14, 8))

# Plotting average latency against RPS, Distribution, and Duration
sns.lineplot(data=e2e_slowdown, x='RPS', y='Avg Latency', hue='Distribution', style='Duration', markers=True)

plt.title('Average Latency by RPS, Distribution, and Duration')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Average Latency (ms)')
plt.grid(True)
plt.legend(title='Distribution & Duration', loc='upper left')
plt.show()

import matplotlib.pyplot as plt
import seaborn as sns

# Plot average latency
plt.figure(figsize=(14, 8))
sns.lineplot(data=e2e_slowdown, x='RPS', y='Avg Latency', hue='Distribution', style='Duration', markers=True)

plt.title('Average Latency by RPS, Distribution, and Duration')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Average Latency (ms)')
plt.grid(True)
plt.legend(title='Distribution & Duration', loc='upper left')
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

plt.figure(figsize=(10, 6))
plt.hist(filtered_df['Latency'], bins=20, color='blue', edgecolor='black')
plt.title('Histogram of Latencies')
plt.xlabel('Latency (ms)')
plt.ylabel('Frequency')
plt.show()
