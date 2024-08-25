import pandas as pd
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
import numpy as np
import os

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

# # Plot 3D Success Rate
# plot_3d_metric(success_rate, '3D Success Rate', 'Success Rate', 'viridis')

# # Plot 3D Average Latency
# plot_3d_metric(average_latency, '3D Average Latency', 'Latency (ms)', 'plasma')

# # Plot 3D E2E Slowdown
# plot_3d_metric(e2e_slowdown, '3D E2E Slowdown', 'Slowdown (relative)', 'inferno')

# Discuss behaviors across different RPS, invocation distribution, and CPU spin duration
for rps, dist, dur in grouped.groups:
    num_requests = grouped.get_group((rps, dist, dur)).shape[0]
    num_errors = grouped.get_group((rps, dist, dur))['Response'].value_counts().get('Error', 0)
    error_rate = num_errors / num_requests
    if error_rate > 0.1:
        print(f"Configuration with RPS={rps}, Distribution={dist}, Duration={dur} ms is unable to run due to high error rate ({error_rate:.2%}).")


def plot_latency_cdf_and_e2e_slowdown(warmup_filtered_df, output_file_latency, output_file_e2e_slowdown, image_dir):
    distributions = ['Uniform', 'Poisson']
    results_latency = []
    results_e2e_slowdown = []

    if not os.path.exists(image_dir):
        os.makedirs(image_dir)

    for dist in distributions:
        # Data for the given distribution
        dist_data = warmup_filtered_df[warmup_filtered_df['Distribution'] == dist]

        # Calculate E2E Slowdown base
        min_latency = dist_data['Latency'].min()

        # Plot CDF of Latency
        fig, ax = plt.subplots(figsize=(10, 8))
        for rps, group_data in dist_data.groupby('RPS'):
            latencies = group_data['Latency'].dropna().values
            latencies.sort()
            cdf = np.arange(len(latencies)) / float(len(latencies))

            # Plot the CDF
            label = f'RPS={rps}'
            line, = ax.plot(latencies, cdf, label=label)

            # Calculate metrics for Latency
            avg_latency = np.mean(latencies)
            median = np.percentile(latencies, 50)
            p90 = np.percentile(latencies, 90)
            p99 = np.percentile(latencies, 99)

            # Calculate E2E Slowdown
            e2e_slowdown = avg_latency / min_latency

            # Record results
            results_latency.append({
                'Distribution': dist,
                'RPS': rps,
                'Average Latency (ms)': avg_latency,
                'Median Latency (ms)': median,
                '90th Percentile (ms)': p90,
                '99th Percentile (ms)': p99
            })

            results_e2e_slowdown.append({
                'Distribution': dist,
                'RPS': rps,
                'E2E Slowdown': e2e_slowdown
            })

            # Plot median and P99 lines for Latency
            ax.axvline(median, color=line.get_color(), linestyle='--', alpha=0.7, label=f'{label} Median')
            ax.axvline(p99, color=line.get_color(), linestyle=':', alpha=0.7, label=f'{label} P99')

        ax.set_xlabel('Latency (ms)')
        ax.set_ylabel('CDF')
        ax.set_title(f'CDF of Latency for {dist} Distribution')
        ax.legend(loc='best')
        plt.grid(True)
        
        # Save CDF of Latency Plot
        latency_plot_path = os.path.join(image_dir, f'latency_cdf_{dist}.png')
        plt.savefig(latency_plot_path)
        plt.close()
        print(f"Saved Latency CDF plot to {latency_plot_path}")

    # Plot E2E Slowdown for both distributions
    fig, ax = plt.subplots(figsize=(10, 8))
    
    for dist in distributions:
        e2e_slowdown_data = [entry for entry in results_e2e_slowdown if entry['Distribution'] == dist]
        rps_values = [entry['RPS'] for entry in e2e_slowdown_data]
        e2e_values = [entry['E2E Slowdown'] for entry in e2e_slowdown_data]

        ax.plot(rps_values, e2e_values, marker='o', linestyle='-', label=f'{dist} E2E Slowdown')

    ax.set_xlabel('RPS')
    ax.set_ylabel('E2E Slowdown')
    ax.set_title('E2E Slowdown for Uniform and Poisson Distributions')
    ax.legend(loc='best')
    plt.grid(True)
    
    # Save E2E Slowdown Plot
    e2e_slowdown_plot_path = os.path.join(image_dir, 'e2e_slowdown.png')
    plt.savefig(e2e_slowdown_plot_path)
    plt.close()
    print(f"Saved E2E Slowdown plot to {e2e_slowdown_plot_path}")

    # Save metrics to files
    results_latency_df = pd.DataFrame(results_latency)
    results_e2e_slowdown_df = pd.DataFrame(results_e2e_slowdown)

    results_latency_df.to_csv(output_file_latency, index=False)
    results_e2e_slowdown_df.to_csv(output_file_e2e_slowdown, index=False)

    print(f"Latency metrics saved to {output_file_latency}")
    print(f"E2E Slowdown metrics saved to {output_file_e2e_slowdown}")

# Define the paths to save the results files and images
output_file_latency = 'loadgen/results/latency_metrics.csv'  # Update this with your desired file path for latency metrics
output_file_e2e_slowdown = 'loadgen/results/e2e_slowdown_metrics.csv'  # Update this with your desired file path for E2E slowdown metrics
image_dir = 'analyze-images/'  # Directory to save the plots

# Plot CDF of Latency and E2E Slowdown, save metrics, and images
plot_latency_cdf_and_e2e_slowdown(warmup_filtered_df, output_file_latency, output_file_e2e_slowdown, image_dir)
