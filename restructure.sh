#!/usr/bin/env bash

# Define an associative array with source files and their respective target directories
declare -A files_to_directories=(
    ["4-First_Goroutine/NotUsingGoroutines.go"]="4-First_Goroutine/not_using_goroutines/main.go"
    ["4-First_Goroutine/UsingGoroutines.go"]="4-First_Goroutine/using_goroutines/main.go"
    ["7-Channels/2-UsingChannels.go"]="7-Channels/using_channels/main.go"
    ["7-Channels/3-Channels_Sync_1.go"]="7-Channels/channels_sync_1/main.go"
    ["7-Channels/3-Channels_Sync_2.go"]="7-Channels/channels_sync_2/main.go"
    ["7-Channels/3-Channels_Sync_3.go"]="7-Channels/channels_sync_3/main.go"
    ["7-Channels/3-Channels_Sync_4.go"]="7-Channels/channels_sync_4/main.go"
    ["7-Channels/3-bufferedChannels.go"]="7-Channels/buffered_channels/main.go"
    ["7-Channels/4-multipleChannels_Select.go"]="7-Channels/multiple_channels_select/main.go"
    ["8-IO_Bound_vs_CPU_Bound/CPU_1-sequential.go"]="8-IO_Bound_vs_CPU_Bound/cpu_1_sequential/main.go"
    ["8-IO_Bound_vs_CPU_Bound/CPU_2-waitgroup.go"]="8-IO_Bound_vs_CPU_Bound/cpu_2_waitgroup/main.go"
    ["8-IO_Bound_vs_CPU_Bound/CPU_3-channels.go"]="8-IO_Bound_vs_CPU_Bound/cpu_3_channels/main.go"
    ["8-IO_Bound_vs_CPU_Bound/IO_1-web_status_sequential.go"]="8-IO_Bound_vs_CPU_Bound/io_1_web_status_sequential/main.go"
    ["8-IO_Bound_vs_CPU_Bound/IO_2-web_status_waitgroup.go"]="8-IO_Bound_vs_CPU_Bound/io_2_web_status_waitgroup/main.go"
    ["8-IO_Bound_vs_CPU_Bound/IO_3-web_status_buff_channels.go"]="8-IO_Bound_vs_CPU_Bound/io_3_web_status_buff_channels/main.go"
)

# Iterate over the array and move files to their new locations
for src in "${!files_to_directories[@]}"; do
    dest="${files_to_directories[$src]}"
    
    # Check if the source file exists
    if [ -e "$src" ]; then
        # Create the target directory if it doesn't exist
        mkdir -p "$(dirname "$dest")"
        
        # Move the file to the new location
        mv "$src" "$dest"
    else
        echo "Source file $src does not exist."
    fi
done

echo "Files have been moved to their respective directories."