#!/bin/bash

output_file="index.md"
source_folder="slides"
front_matter_flag=0

rm -f "$output_file"

function process_files {
  local folder="$1"
  for item in "$folder"/*; do
    if [ -f "$item" ]; then
      if [ $front_matter_flag -eq 0 ]; then
        cat "$item" >> "$output_file"
        front_matter_flag=1
      else
        printf "\n---\n" >> "$output_file"
        sed '1{/^---$/!q;};1,/^---$/d' "$item" >> "$output_file"
      fi
    elif [ -d "$item" ]; then
      process_files "$item"
    fi
  done
}

process_files "$source_folder"

echo "All .md files from folder $source_folder have been appended to $output_file."

# Fix assets path for GH Pages
sed -i 's/\/assets/\.\/assets/g' "$output_file"