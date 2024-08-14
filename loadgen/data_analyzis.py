import pandas as pd
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# Load log data
log_data = []

# Path to the log file
log_file_path = '/Users/kasrahmi/Desktop/intern-project/new/intern-project/loadgen/loadgen_log.txt'

# Read the log file
with open(log_file_path, 'r') as log_file:
    for line in log_file:
        if "Starting test" in line:
            continue  # Skip the start test lines
        if "Response" in line or "Error" in line:
            parts = line.split(", ")
            timestamp = parts[0].split(": ", 1)[1]
            rps = int(parts[1].split(": ")[1])
            distribution = parts[2].split(": ")[1]
            duration = int(parts[3].split(": ")[1].split(" ")[0])
            request_name = parts[4].split(": ")[1]
            response = "Success" if "Response" in line else "Error"
            latency = int(parts[6].split(": ")[1].split(" ")[0]) if "Response" in line else None
            worker_id = parts[7].split(": ")[1].strip() if "Response" in line else None

            log_data.append([timestamp, rps, distribution, duration, request_name, response, latency, worker_id])

# Create a DataFrame
df = pd.DataFrame(log_data, columns=["Timestamp", "RPS", "Distribution", "Duration", "Request", "Response", "Latency", "WorkerID"])

# Convert Timestamp to datetime
df['Timestamp'] = pd.to_datetime(df['Timestamp'])

# Filter out the warm-up period (first 4 minutes)
start_time = df['Timestamp'].min()
warmup_filtered_df = df[df['Timestamp'] >= start_time + pd.Timedelta(minutes=4)]

# Verify that data exists after filtering
if warmup_filtered_df.empty:
    raise ValueError("No data available after filtering out the warm-up period. Please check the log data.")

# Calculate success rate and average latency per configuration
grouped = warmup_filtered_df.groupby(['RPS', 'Distribution', 'Duration'])
success_rate = grouped['Response'].apply(lambda x: (x == "Success").mean())
average_latency = grouped['Latency'].mean()

# Calculate E2E slowdown (overhead)
if not average_latency.empty:
    min_latency = average_latency.min()
    e2e_slowdown = average_latency / min_latency
else:
    e2e_slowdown = pd.Series(dtype='float64')

# Map distribution to a numerical value for plotting
distribution_map = {'Uniform': 0, 'Poisson': 1}

# Prepare data for 3D plotting
def prepare_plot_data(metric_series):
    plot_data = []
    for (rps, dist, dur), value in metric_series.items():
        plot_data.append((rps, distribution_map[dist], dur, value))
    return pd.DataFrame(plot_data, columns=['RPS', 'Distribution', 'Duration', 'Metric'])

# Plot 3D metrics
def plot_3d_metric(metric_series, title, zlabel, color):
    data = prepare_plot_data(metric_series)
    fig = plt.figure(figsize=(10, 8))
    ax = fig.add_subplot(111, projection='3d')

    scatter = ax.scatter(data['RPS'], data['Distribution'], data['Duration'], c=data['Metric'], cmap=color, marker='o')
    ax.set_xlabel('RPS')
    ax.set_ylabel('Distribution (0=Uniform, 1=Poisson)')
    ax.set_zlabel('Duration (ms)')
    ax.set_title(title)

    # Add color bar to indicate metric value
    plt.colorbar(scatter, ax=ax, label=zlabel)

    plt.show()

# Plot 3D Success Rate
plot_3d_metric(success_rate, '3D Success Rate', 'Success Rate', 'viridis')

# Plot 3D Average Latency
plot_3d_metric(average_latency, '3D Average Latency', 'Latency (ms)', 'plasma')

# Plot 3D E2E Slowdown
plot_3d_metric(e2e_slowdown, '3D E2E Slowdown', 'Slowdown (relative)', 'inferno')

# Discuss behaviors across different RPS, invocation distribution, and CPU spin duration
for rps, dist, dur in grouped.groups:
    num_requests = grouped.get_group((rps, dist, dur)).shape[0]
    num_errors = grouped.get_group((rps, dist, dur))['Response'].value_counts().get('Error', 0)
    error_rate = num_errors / num_requests
    if error_rate > 0.1:
        print(f"Configuration with RPS={rps}, Distribution={dist}, Duration={dur} ms is unable to run due to high error rate ({error_rate:.2%}).")
