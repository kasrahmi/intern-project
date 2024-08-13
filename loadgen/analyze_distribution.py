import pandas as pd
import matplotlib.pyplot as plt

# Load log data
log_data = []

# Use the correct path to your log file
log_file_path = '/Users/kasrahmi/Desktop/intern-project/new/intern-project/loadgen/loadgen_log_distribution.txt'

# Read the log file
with open(log_file_path, 'r') as log_file:
    current_distribution = ""
    current_rps = 0
    for line in log_file:
        if "Starting test with" in line:
            # Parse distribution and RPS
            parts = line.split(", ")
            current_rps = int(parts[0].split(" ")[3])
            current_distribution = parts[1].split(": ")[1].strip()
        elif "Request" in line and "Latency" in line:
            parts = line.split(", ")
            timestamp = parts[0].split(": ")[1]
            request_name = parts[1].split(": ")[1]
            response = parts[2].split(": ")[1]
            latency = int(parts[3].split(": ")[1].split(" ")[0])
            worker_id = parts[4].split(": ")[1].strip()
            log_data.append([timestamp, request_name, response, latency, worker_id, current_rps, current_distribution])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=["Timestamp", "Request", "Response", "Latency", "WorkerID", "RPS", "Distribution"])

# Convert latency to numeric
df['Latency'] = pd.to_numeric(df['Latency'])

# Calculate average latency and success rate per RPS level and distribution
results = df.groupby(['RPS', 'Distribution']).agg({
    'Latency': 'mean',
    'Response': lambda x: (x != "Error").mean() * 100  # Success rate as percentage
}).rename(columns={'Latency': 'AvgLatency', 'Response': 'SuccessRate'}).reset_index()

# Print results
print(results)

# Plot latency for each distribution
plt.figure(figsize=(12, 6))
for distribution in df['Distribution'].unique():
    subset = df[df['Distribution'] == distribution]
    plt.hist(subset['Latency'], bins=20, alpha=0.5, label=f'{distribution} distribution')

plt.title('Latency Distribution')
plt.xlabel('Latency (ms)')
plt.ylabel('Frequency')
plt.legend()
plt.show()

# Plot average latency vs RPS
plt.figure(figsize=(12, 6))
for distribution in results['Distribution'].unique():
    subset = results[results['Distribution'] == distribution]
    plt.plot(subset['RPS'], subset['AvgLatency'], marker='o', label=f'{distribution} distribution')

plt.title('Average Latency vs RPS')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Average Latency (ms)')
plt.legend()
plt.grid(True)
plt.show()

# Plot success rate vs RPS
plt.figure(figsize=(12, 6))
for distribution in results['Distribution'].unique():
    subset = results[results['Distribution'] == distribution]
    plt.plot(subset['RPS'], subset['SuccessRate'], marker='o', label=f'{distribution} distribution')

plt.title('Success Rate vs RPS')
plt.xlabel('Requests Per Second (RPS)')
plt.ylabel('Success Rate (%)')
plt.legend()
plt.grid(True)
plt.show()

# Plot worker distribution
worker_distribution = df.groupby(['WorkerID', 'Distribution']).size().unstack().fillna(0)
worker_distribution.plot(kind='bar', stacked=True, figsize=(12, 6))
plt.title('Request Distribution by Worker ID and Distribution Type')
plt.xlabel('Worker ID')
plt.ylabel('Number of Requests')
plt.legend(title='Distribution')
plt.show()
