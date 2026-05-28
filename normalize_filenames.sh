#!/bin/bash

# Script to rename all files in web/photos/ to sequential numbers (1.ext, 2.ext, etc.)

set -e

PHOTOS_DIR="web/photos"
TEMP_PREFIX="temp_rename_"

# Check if directory exists
if [ ! -d "$PHOTOS_DIR" ]; then
    echo "Error: Directory $PHOTOS_DIR does not exist"
    exit 1
fi

# Change to photos directory
cd "$PHOTOS_DIR" || exit 1

echo "Starting filename normalization in $PHOTOS_DIR..."
echo ""

# Collect all image files (excluding json files)
declare -a files
for file in *; do
    # Skip if not a file
    [ ! -f "$file" ] && continue

    # Skip json files
    [[ "$file" == *.json ]] && continue

    # Skip txt files
    [[ "$file" == *.txt ]] && continue

    # Add to array
    files+=("$file")
done

# Sort files alphabetically
IFS=$'\n' sorted_files=($(sort <<<"${files[*]}"))
unset IFS

# First pass: rename to temporary names to avoid conflicts
counter=1
declare -a temp_mappings

for file in "${sorted_files[@]}"; do
    extension="${file##*.}"
    # Convert extension to lowercase
    extension=$(echo "$extension" | tr '[:upper:]' '[:lower:]')
    # Convert .jpeg to .jpg
    [ "$extension" = "jpeg" ] && extension="jpg"

    temp_name="${TEMP_PREFIX}${counter}.${extension}"
    mv "$file" "$temp_name"
    temp_mappings+=("$temp_name|$counter.$extension|$file")
    counter=$((counter + 1))
done

# Second pass: rename from temp names to final names
for mapping in "${temp_mappings[@]}"; do
    IFS='|' read -r temp_name final_name original_name <<< "$mapping"

    echo "Renaming: $original_name -> $final_name"

    mv "$temp_name" "$final_name"
done

echo ""
echo "====================================="
echo "Normalization complete!"
echo "Files renamed: ${#sorted_files[@]}"
echo "====================================="

# Return to original directory
cd - > /dev/null
